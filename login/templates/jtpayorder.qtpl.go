// This file is automatically generated by qtc from "jtpayorder.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:1
package templates

//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:2
type JtpayOrder struct {
	P1_yingyongnum string
	P2_ordernumber string
	P3_money       string
	P6_ordertime   string
	P7_productcode string
	P8_sign        string
	P9_signtype    string
	P25_terminal   string
}

//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:14
func StreamJtpayOrderTemplate(qw422016 *qt422016.Writer, p *JtpayOrder) {
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:14
	qw422016.N().S(`
<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<title>竣付通</title>
</head>
<!--支付宝IOSwap支付请求提交页-->
<body onLoad="document.yeepay.submit();">
	<form name='yeepay' action='https://order.z.jtpay.com/jh-web-order/order/receiveOrder' method='post'>
	<input type='hidden' name='p1_yingyongnum' value='`)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:23
	qw422016.E().S(p.P1_yingyongnum)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:23
	qw422016.N().S(`'>
	<input type='hidden' name='p2_ordernumber' value='`)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:24
	qw422016.E().S(p.P2_ordernumber)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:24
	qw422016.N().S(`'>
	<input type='hidden' name='p3_money' value='`)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:25
	qw422016.E().S(p.P3_money)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:25
	qw422016.N().S(`'>
	<input type='hidden' name='p6_ordertime' value='`)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:26
	qw422016.E().S(p.P6_ordertime)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:26
	qw422016.N().S(`'>
	<input type='hidden' name='p7_productcode' value='`)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:27
	qw422016.E().S(p.P7_productcode)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:27
	qw422016.N().S(`'>
	<input type='hidden' name='p8_sign' value='`)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:28
	qw422016.E().S(p.P8_sign)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:28
	qw422016.N().S(`'>
	<input type='hidden' name='p9_signtype' value='`)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:29
	qw422016.E().S(p.P9_signtype)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:29
	qw422016.N().S(`'>
	<input type='hidden' name='p25_terminal' value='`)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:30
	qw422016.E().S(p.P25_terminal)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:30
	qw422016.N().S(`'>
	</form>
</body>
</html>
`)
//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
}

//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
func WriteJtpayOrderTemplate(qq422016 qtio422016.Writer, p *JtpayOrder) {
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
	StreamJtpayOrderTemplate(qw422016, p)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
	qt422016.ReleaseWriter(qw422016)
//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
}

//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
func JtpayOrderTemplate(p *JtpayOrder) string {
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
	qb422016 := qt422016.AcquireByteBuffer()
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
	WriteJtpayOrderTemplate(qb422016, p)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
	qs422016 := string(qb422016.B)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
	qt422016.ReleaseByteBuffer(qb422016)
	//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
	return qs422016
//line ../src/github.com/valyala/quicktemplate/examples/basicserver/templates/jtpayorder.qtpl:34
}
