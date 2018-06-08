# go-clockout
If we don't know the time of clocking in today, we can use this tool to calculate the time to clock out.

Everyday I go to work, and press button to turn on my computer. So I use this start time to calculate the time that I can clock out.

On Windows, we can easyily use the event of ID 12 to get the start time.

### Windows command and result
```
wevtutil qe system /q:Event[EventData[Data[@Name='StartTime']^>'2018-06-08T00:00:00']]
```
```
<Event xmlns='http://schemas.microsoft.com/win/2004/08/events/event'>
<System>
<Provider Name='Microsoft-Windows-Kernel-General' Guid='{A68CA8B7-004F-D7B6-A698-07E2DE0F1F5D}'/>
<EventID>12</EventID>
<Version>0</Version>
<Level>4</Level>
<Task>0</Task>
<Opcode>0</Opcode>
<Keywords>0x8000000000000000</Keywords>
<TimeCreated SystemTime='2018-06-08T02:03:00.734000400Z'/>
<EventRecordID>302816</EventRecordID>
<Correlation/>
<Execution ProcessID='4' ThreadID='8'/>
<Channel>System</Channel>
<Computer> xxxxx </Computer>
<Security UserID='S-1-5-18'/>
</System>
<EventData>
<Data Name='MajorVersion'>6</Data>
<Data Name='MinorVersion'>1</Data>
<Data Name='BuildVersion'>7601</Data>
<Data Name='QfeVersion'>23807</Data>
<Data Name='ServiceVersion'>1</Data>
<Data Name='BootMode'>0</Data>
<Data Name='StartTime'>2018-06-08T02:03:00.109999300Z</Data>
</EventData>
</Event>
```

### Tool execution result
```
clock in  : 2018-06-08 10:03
clock out : 2018-06-08 19:18
```
