package cluster

import (
	"time"

	"github.com/AsynkronIT/gonet"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/log"
	"github.com/AsynkronIT/protoactor-go/remote"
)

var cfg *ClusterConfig

func Start(clusterName, address string, provider ClusterProvider) {
	StartWithConfig(NewClusterConfig(clusterName, address, provider))
}

func StartWithConfig(config *ClusterConfig) {
	cfg = config

	//TODO: make it possible to become a cluster even if remoting is already started
	remote.Start(cfg.Address, cfg.RemotingOption...)

	address := actor.ProcessRegistry.Address
	h, p := gonet.GetAddress(address)
	plog.Info("Starting Proto.Actor cluster", log.String("address", address))
	kinds := remote.GetKnownKinds()

	//for each known kind, spin up a partition-kind actor to handle all requests for that kind
	setupPartition(kinds)
	setupPidCache()
	setupMemberList()

	cfg.ClusterProvider.RegisterMember(cfg.Name, h, p, kinds, cfg.InitialMemberStatusValue, cfg.MemberStatusValueSerializer)
	cfg.ClusterProvider.MonitorMemberStatusChanges()
}

func Shutdown(graceful bool) {
	if graceful {
		cfg.ClusterProvider.Shutdown()
		//This is to wait ownership transferring complete.
		time.Sleep(time.Millisecond * 2000)
		stopMemberList()
		stopPidCache()
		stopPartition()
	}

	remote.Shutdown(graceful)

	address := actor.ProcessRegistry.Address
	plog.Info("Stopped Proto.Actor cluster", log.String("address", address))
}

//Get a PID to a virtual actor
func Get(name string, kind string) (*actor.PID, remote.ResponseStatusCode) {
	//Check Cache
	if pid, ok := pidCache.getCache(name); ok {
		return pid, remote.ResponseStatusCodeOK
	}

	//Get Pid
	address := memberList.getPartitionMember(name, kind)
	if address == "" {
		//No available member found
		return nil, remote.ResponseStatusCodeUNAVAILABLE
	}

	//package the request as a remote.ActorPidRequest
	req := &remote.ActorPidRequest{
		Kind: kind,
		Name: name,
	}

	//ask the DHT partition for this name to give us a PID
	remotePartition := partition.partitionForKind(address, kind)
	r, err := remotePartition.RequestFuture(req, cfg.TimeoutTime).Result()
	if err == actor.ErrTimeout {
		plog.Error("PidCache Pid request timeout")
		return nil, remote.ResponseStatusCodeTIMEOUT
	} else if err != nil {
		plog.Error("PidCache Pid request error", log.Error(err))
		return nil, remote.ResponseStatusCodeERROR
	}

	response, ok := r.(*remote.ActorPidResponse)
	if !ok {
		return nil, remote.ResponseStatusCodeERROR
	}

	statusCode := remote.ResponseStatusCode(response.StatusCode)
	switch statusCode {
	case remote.ResponseStatusCodeOK:
		//save cache
		pidCache.addCache(name, response.Pid)
		//tell the original requester that we have a response
		return response.Pid, statusCode
	default:
		//forward to requester
		return response.Pid, statusCode
	}
}

//RemoveCache at PidCache
func RemoveCache(name string) {
	pidCache.removeCacheByName(name)
}
