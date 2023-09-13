# xiphias-model-common

This repository contains generated Go code for the Xiphias Model Common protocol buffers used by Kik Messenger.

### Installation

```bash
go get github.com/sokkit-io/xiphias-model-common@latest
```

### Usage

You can use the generated code to marshal and unmarshal protocol buffer messages, typically from a Kik XMPP stanza.

For example:

```go
package main

import (
	"encoding/base64"
	"github.com/golang/protobuf/proto"
	"github.com/sokkit-io/xiphias-model-common/generated/go"
	"strings"
)

func main() {
    // Let's say you've got a base64-encoded response from Kik that you want to decode. This happens to be a XiBareUserJid body.
    xiBareUserJidB64 := "Cgd0ZWRfdzZ3"
    
    // Remove trailing padding characters
    xiBareUserJidB64 = strings.TrimRight(xiBareUserJidB64, "=")
    
    // Decode the base64-encoded response
    xiBareUserJidBytes, _ := base64.RawURLEncoding.DecodeString(xiBareUserJidB64)
    
    // Prepare a XiBareUserJid struct to unmarshal the response into
    xiBareUserJid := &common.XiBareUserJid{}
    
    // Unmarshal the response into the XiBareUserJid struct
    _ = proto.Unmarshal(xiBareUserJidBytes, xiBareUserJid)
    
    // Print the JID
    println(xiBareUserJid.GetLocalPart())
}
```

### Notes

Protocol buffer definitions were extracted from the [Kik Messenger](https://kik.com) Android application
version **15.51.1.28280** using [pbtk](https://github.com/marin-m/pbtk).

Definitions will be updated as breaking changes are introduced.

### Acknowledgements

- [kikinteractive](https://github.com/kikinteractive) for creating [Kik Messenger](https://kik.com)
- [marin-m](https://github.com/marin-m) for creating [pbtk](https://github.com/marin-m/pbtk)
- [golang/protobuf](https://github.com/golang/protobuf) contributors
- Developers and skids of the Kik modding community