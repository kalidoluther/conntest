# AT Commands

## Serial asynchronous automatic dialling and control

### 5.2.1 Command line general format

> A command line is made up of three elements: the prefix, the body, and the termination character.  
The command line prefix consists of the characters "AT" (IA5 4/1, 5/4) or "at" (IA5 6/1, 7/4), or, to repeat the execution of the previous command line, the characters "A/" (IA5 4/1, 2/15) or "a/" (IA5 6/1, 2/15).  
The body is made up of individual commands as specified later in this Recommendation.  
Space characters (IA5 2/0) are ignored and may be used freely for formatting purposes, unless they are embedded in numeric or string constants (5.4.2.1 or 5.4.2.2).  
The termination character may not appear in the body. The DCE shall be capable of accepting at least 40 characters in the body.  
The termination character may be selected by a user option (parameter S3), the default being CR (IA5 0/13).  

## Comands

### AT+CSQ (Get Signal Strength)

Description
```
AT+CSQ AT command returns the signal strength of the device.
Possible values are,
0 113 dBm or less
1 111 dBm
2...30 109... 53 dBm
31 51 dBm or greater
99 not known or not detectable
```

Usage
```
Command             Possible Response(s)
+CSQ                +CSQ: +CME, ERROR:
+CSQ=?              +CSQ: (list of supporteds), (list of supporteds)
```

## Examples


```

AT+CGMI    // Request manufacturer identification
AT+CGMM    // Request model identification
AT+CGSN    // Request product serial number identification
AT+CSUB    // Request the module version and chip

AT+CPIN?     // Request the state of the SIM card
AT+CICCID    // Read ICCID from SIM card
AT+CNUM      // Request the subscriber number
AT+CNMP?     // Preferred mode selection
AT+COPS?     // Check the current network operator

AT+IPREX?    // Check local baud rate
AT+CRESET    // Reset the module

AT+CGPS=1      // Start GPS session
AT+CGPSINFO    // Get GPS fixed position information
AT+CGPS=0      // Stop GPS session

Send and receive SMS
AT+CSCA="+27841000000"       // Set the SMS service centre address for CellC
AT+CMGF=1              // Select SMS message format
AT+CMGS="xxxxxx"       // Send message to "xxxxxx"(the receiver number).

AT+CMGR=3        // Read message
AT+CMGD=3        // Delete message

Make a call
 AT+CSDVC         // Switch voice channel device
    AT+CSDVC=1       // 1-Handset, 3-Speaker phone
    AT+CLVL=2        // Set loudspeaker volume level to 2, the level range is 0 to 5
ATDxxxxx;        // Call to xxxxx
    AT+CHUP          // Hang up the call
    AT+CLIP=1     // Calling line identification presentation
    ATA         // Call answer

HTTP test
You can test the SIM7600 LTE HTTP request using following commands.
AT+HTTPINIT         // Initialize and start the HTTP
    AT+HTTPPARA="URL","http://www.makerfabs.com"  // Set the URL
    AT+HTTPACTION=0    // Connect the HTTP. (0-get, 1-post, 2-head)
    AT+HTTPHEAD        // Read the response's header.
    AT+HTTPREAD=0,3    // Read the content (“3” means the number of the reading data)

Test the SD card for SIM7600
AT+FSCD=D:                    // Select SD card directory as current directory
    AT+FSLS                       // List directories/files in current directory
    AT+CFTRANRX="D: TEST.txt",10  // Transfer a file to EFS
    AT+CFTRANTX="D: TEST.txt"     // Transfer a file from EFS to host
```
