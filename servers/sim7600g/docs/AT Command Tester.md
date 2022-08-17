## AT Command Teser

### Output Sample
```

Found ports :COM8


Sending AT query..
AT

OK
Successfull response for AT query..

Enabling echo and verbose mode
ATE1V1

OK
AT+CGMM

SIMCOM_SIM7600G-H

OK
Model Number : SIMCOM_SIM7600G-H
AT+CGMI

SIMCOM INCORPORATED

OK
Manufacturer : SIMCOM INCORPORATED

Please try   for features specific to Simcom modules.

AT+CNUM

OK
Command error..

AT+CSQ

+CSQ: 16,99

OK
Signal level is -81 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

AT+CSQ

+CSQ: 16,99

OK
Signal level is -81 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

AT+CSQ

+CSQ: 16,99

OK
Signal level is -81 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).


+CMTI: "SM",2
Getting Alert Mode..
AT+CALM?

ERROR
Command error..

Setting Alert Mode to Normal.
AT+CALM=0

ERROR
Command error..

Getting Alert Mode..
AT+CALM?

ERROR
Command error..

Getting ringer sound level..
AT+CRSL?

ERROR
Command error..

Getting loud speaker level.
AT+CLVL?


+CLVL: 4

OK
Disconnecting port COM8.


Sending AT query..
AT

OK
Successfull response for AT query..

Enabling echo and verbose mode
ATE1V1

OK
AT+CGMM

SIMCOM_SIM7600G-H

OK
Model Number : SIMCOM_SIM7600G-H
AT+CGMI

SIMCOM INCORPORATED

OK
Manufacturer : SIMCOM INCORPORATED

Please try   for features specific to Simcom modules.

Getting Alert Mode..
AT+CALM?

ERROR
Command error..

AT+CSQ

+CSQ: 6,99

OK
Signal level is -101 dbm. Signal condition is marginal.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

AT+CSQ

+CSQ: 6,99

OK
Signal level is -101 dbm. Signal condition is marginal.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).



Setting Alert Mode to Normal.
Getting ringer sound level..
Disconnecting port COM8.


Sending AT query..
AT

OK
Successfull response for AT query..

Enabling echo and verbose mode
ATE1V1

OK
AT+CGMM

SIMCOM_SIM7600G-H

OK
Model Number : SIMCOM_SIM7600G-H
AT+CGMI

SIMCOM INCORPORATED

OK
Manufacturer : SIMCOM INCORPORATED

Please try   for features specific to Simcom modules.

AT+CGMR

+CGMR: LE20B03SIM7600M22

OK
Revision : +CGMR: LE20B03SIM7600M22

AT+CGSN

862636050029311

OK
The IMEI(International Mobile Equipment IdentityI is 862636050029311
AT+CIMI

655070216918286

OK
Phone Number : 655070216918286
AT+CCLK?

+CCLK: "80/01/06,07:09:04"

OK
Current time is"80/01/06,07:09:04"
AT+CPIN?

+CPIN: READY

OK
SIM is ready
AT+CGDCONT?

+CGDCONT: 1,"IP","CMNET","0.0.0.0",0,0,0,0
+CGDCONT: 2,"IPV4V6","","0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0",0,0,0,0
+CGDCONT: 3,"IPV4V6","","0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0",0,0,0,0
+CGDCONT: 4,"IPV4V6","","0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0",0,0,0,0
+CGDCONT: 5,"IPV4V6","","0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0",0,0,0,0
+CGDCONT: 6,"IPV4V6","","0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0",0,0,0,0

OK
Following connection profiles are available,

CID-> 1
PDP Type->IP
APN->CMNET
PDP Address->0.0.0.0
Data Compression->0
Header Compression->0

CID-> 2
PDP Type->IPV4V6
APN->
PDP Address->0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0
Data Compression->0
Header Compression->0

CID-> 3
PDP Type->IPV4V6
APN->
PDP Address->0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0
Data Compression->0
Header Compression->0

CID-> 4
PDP Type->IPV4V6
APN->
PDP Address->0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0
Data Compression->0
Header Compression->0

CID-> 5
PDP Type->IPV4V6
APN->
PDP Address->0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0
Data Compression->0
Header Compression->0

CID-> 6
PDP Type->IPV4V6
APN->
PDP Address->0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0
Data Compression->0
Header Compression->0
AT+CSQ

+CSQ: 16,99

OK
Signal level is -81 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

AT+IPR?

+IPR: 115200

OK
Serial port speed is set to 115200.
AT+COPS?

+COPS: 1,0,"Cell C  Cell C",7

OK
Device is currently on "Cell C  Cell C" network.

AT+CREG?

+CREG: 0,1

OK
The device is registered in home network.

AT+CGACT?

+CGACT: 1,0
+CGACT: 2,0
+CGACT: 3,0
+CGACT: 4,0
+CGACT: 5,0
+CGACT: 6,1

OK
AT+CGPADDR= 6

+CGPADDR: 6,10.29.127.164

OK
IP Address of the connected profile is 10.29.127.164

AT+CGATT?

+CGATT: 1

OK
Device is attached to the network
AT+CFUN?

+CFUN: 1

OK
Device has Full functionality.
AT+CGCLASS?

+CGCLASS: "A"

OK
The device supports station class "A"
AT+CSCA?

+CSCA: "+27841000000",145

OK
SMS service center address is +27841000000
AT+CMGF?

+CMGF: 1

OK
SMS message for is configured for Text mode
Checking registration status...

AT+CREG?

+CREG: 0,1

OK
The device is registered in home network.

AT+CGREG?

+CGREG: 0,1

OK
The device is registered in home network.

Dialing number 0839776374

ATD0839776374;

OK
Voice call successfull


VOICE CALL: BEGIN
Hanging up the call..

ATH

OK
Call sucessfully dis-connected..

No call is connected..


VOICE CALL: END: 000034


NO CARRIER

Disconnecting port COM8.

average dbm is -81
AT+CSQ

+CSQ: 18,99

OK
Signal level is -77 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

AT+CSQ

+CSQ: 18,99

OK
Signal level is -77 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).


Sending AT query..
AT

OK
Successfull response for AT query..

AT+CSQ

+CSQ: 18,99

OK
Signal level is -77 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

Enabling echo and verbose mode
ATE1V1

OK
AT+CSQ

+CSQ: 18,99

OK
Signal level is -77 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

AT+CGMM

SIMCOM_SIM7600G-H

OK
Model Number : SIMCOM_SIM7600G-H
AT+CSQ

+CSQ: 18,99

OK
Signal level is -77 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

AT+CGMI

SIMCOM INCORPORATED

OK
Manufacturer : SIMCOM INCORPORATED
average dbm is -77
AT+CSQ

+CSQ: 18,99

OK
Signal level is -77 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).


Please try   for features specific to Simcom modules.

AT+CSQ

+CSQ: 18,99

OK
Signal level is -77 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

AT+CSQ

+CSQ: 18,99

OK
Signal level is -77 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

AT+CSQ

+CSQ: 18,99

OK
Signal level is -77 dbm. Signal condition is good.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

AT+CSQ

+CSQ: 14,99

OK
Signal level is -85 dbm. Signal condition is OK.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

average dbm is -78
AT+CSQ

+CSQ: 14,99

OK
Signal level is -85 dbm. Signal condition is OK.The signal strength range is -53 dbm (Excellent) to -109 dbm (Marginal).

Reading SMS Messages..

AT+CMGL="ALL"

+CMGL: 0,"REC READ","+27840001656","","22/08/08,12:01:40+08"
You have less than 79390 MB Anytime data remaining. Check your balance on the Cell C app, at www.cellc.co.za or by dialing *147# Cell C
+CMGL: 1,"REC READ","+2784000136800051","","22/08/08,12:03:29+08"
Afrihost OTP: 3778
+CMGL: 2,"REC UNREAD","+2787085101217044","","22/08/08,12:27:00+08"
Lays Chips 120g any 2 for R34! SPAR Rewards Customers only. Ends 21/8 For more offers visit https://go.spar.co.za/5LBFvD or *120*7070#. Stop to opt out

OK
AT+CMGD= 0

OK
SMS Delete successfull

AT+CMGL="ALL"

+CMGL: 1,"REC READ","+2784000136800051","","22/08/08,12:03:29+08"
Afrihost OTP: 3778
+CMGL: 2,"REC READ","+2787085101217044","","22/08/08,12:27:00+08"
Lays Chips 120g any 2 for R34! SPAR Rewards Customers only. Ends 21/8 For more offers visit https://go.spar.co.za/5LBFvD or *120*7070#. Stop to opt out

OK
AT+CMGD= 1

OK
SMS Delete successfull

AT+CMGL="ALL"

+CMGL: 2,"REC READ","+2787085101217044","","22/08/08,12:27:00+08"
Lays Chips 120g any 2 for R34! SPAR Rewards Customers only. Ends 21/8 For more offers visit https://go.spar.co.za/5LBFvD or *120*7070#. Stop to opt out

OK
AT+CMGD= 2

OK
SMS Delete successfull

AT+CMGL="ALL"

OK
Checking registration status...

AT+CREG?

+CREG: 0,1

OK
The device is registered in home network.

AT+CGREG?

+CGREG: 0,1

OK
The device is registered in home network.

Dialing number 0839776374

ATD0839776374;

OK
Voice call successfull


VOICE CALL: BEGIN


+RXDTMF: 1


VOICE CALL: END: 000352


NO CARRIER

Setting Alert Mode to Normal.
Disconnecting port COM8.


MISSED_CALL: 08:48AM +27839776374

Sending AT query..
AT

OK
Successfull response for AT query..

Enabling echo and verbose mode
ATE1V1

OK
AT+CGMM

SIMCOM_SIM7600G-H

OK
Model Number : SIMCOM_SIM7600G-H
AT+CGMI

SIMCOM INCORPORATED

OK
Manufacturer : SIMCOM INCORPORATED

Please try   for features specific to Simcom modules.

Setting Alert Mode to Normal.
AT+CALM=0

ERROR
Command error..

Checking registration status...

AT+CREG?

+CREG: 0,1

OK
The device is registered in home network.

AT+CGREG?

+CGREG: 0,1

OK
The device is registered in home network.

Dialing number 0839776374

ATD0839776374;

OK
Voice call successfull


VOICE CALL: BEGIN
Hanging up the call..

ATH

OK
Call sucessfully dis-connected..


VOICE CALL: END: 000033


NO CARRIER

Reading SMS Messages..

Port not connected.

Port not connected.

Disconnecting port COM8.


Sending AT query..
AT

OK
Successfull response for AT query..

Enabling echo and verbose mode
ATE1V1

OK
AT+CGMM

SIMCOM_SIM7600G-H

OK
Model Number : SIMCOM_SIM7600G-H
Reading SMS Messages..

AT+CMGL="ALL"

+CMGL: 0,"REC UNREAD","2784000162816067","","22/08/08,16:32:40+08"
SPECIAL WOMENS DAY WHATSAPP STICKERS! Dial *147*6*2# NOW to download and use them to appreciate all the women in your life!R3/day. T&Cs.CellC. To stop reply:out
+CMGL: 1,"REC UNREAD","+2784000125811001","","22/08/08,16:50:50+08"
FNB :-) Transfer funds between accounts for free on the FNB App, FNB a div of FRB Ltd Auth FSP & NCRCP20
+CMGL: 2,"REC UNREAD","14154888273","","22/08/09,11:33:49+00"
198601 is your GitHub authentication code.

@github.com #198601

OK

Please try   for features specific to Simcom modules.

Checking registration status...

Cannot make voice call without device being registered

Disconnecting port COM8.

Found ports :COM8


Sending AT query..
AT

OK
Successfull response for AT query..

Enabling echo and verbose mode
ATE1V1

OK
AT+CGMM

SIMCOM_SIM7600G-H

OK
Model Number : SIMCOM_SIM7600G-H
AT+CGMI

SIMCOM INCORPORATED

OK
Manufacturer : SIMCOM INCORPORATED

Please try   for features specific to Simcom modules.

Checking registration status...

AT+CREG?

+CREG: 0,1

OK
The device is registered in home network.

AT+CGREG?

+CGREG: 0,1

OK
The device is registered in home network.

Dialing number 0839776374

ATD0839776374;

OK
Voice call successfull


VOICE CALL: BEGIN


+RXDTMF: 1


VOICE CALL: END: 001404


NO CARRIER

```