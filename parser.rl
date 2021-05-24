package uri

import (
	"fmt"
)

%% machine uri;
%% write data;

func RagelParse(str string) (*URI, error) {
	uri := &URI{}
	data := str
	cs := 0
	limit := len(data)
	p := 0 // data pointer
	m := 0 // marker for matching start position
	pe := limit // data end pointer
	eof := limit // End of data
%%{
	action sm   { m = p }
	action sip  { uri.scheme   = SIP      }
	action sips { uri.scheme   = SIPS     }
	action usrp { uri.userinfo = str[m:p] }
	action hstp { uri.hostport = str[m:p] }
	action prms { uri.params   = str[m:p] }
	action hdrs { uri.headers  = str[m:p] }

  unreserved      = alnum | [\-_.!~*'()];
  escaped         = "%" xdigit xdigit;
  reserved        = [;/?:@&=+$,];
  user_unreserved = [&=+$,;?/];
	hexseq          = xdigit{1,4} ( ":" xdigit{1,4} )*;
	hexpart         = hexseq | hexseq "::" hexseq? | "::" hexseq?;
	paramchar       = [[\]/:&+$] | unreserved | escaped;
	uriparam        = paramchar+ ("=" paramchar+)?;
	hdrchar         = [[\]/?:+$] | unreserved | escaped;
	header          = hdrchar+ "=" hdrchar*;

	domainlabel     = alnum | alnum ( alnum | "-" )* alnum;
	toplabel        = alpha | alpha ( alnum | "-" )* alnum;
	hostname        = ( domainlabel "." )* toplabel "."?;

	IPv4address     = digit{1,3} "." digit{1,3} "." digit{1,3} "." digit{1,3};
	IPv6address     = hexpart ( ":" IPv4address )?;
	IPv6reference   = "[" IPv6address "]";

	user            = ( unreserved | escaped | user_unreserved )+;
  password        = ( unreserved | escaped | [&=+$,] )*;
	host            = hostname | IPv4address | IPv6reference;
  port            = digit{1,5};
	
	scheme   = ("sip" %sip | "sips" %sips) ":";
	userinfo = user >sm (":" password )? %usrp "@";
	hostport = host >sm (":" port)? %hstp;
	params   = (";" uriparam)* >sm %prms;
	headers  = "?" (header ("&" header)*) >sm %hdrs;

	uri := scheme userinfo? hostport params headers?;
}%%
  %% write init;
	%% write exec;

	if cs >= uri_first_final {
		return uri, nil
	}
	return nil, fmt.Errorf("Invalid URI '%s'.", str)
}

/* vim: set filetype=go : */
