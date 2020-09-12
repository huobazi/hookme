# voiceover

```go

voiceover.Say("Hello...")
voiceover.Sayf("Error: %v\n", err)

m := make(map[string]interface{})
m["hello"] = "world";
m["haha"] = 1024;

voiceover.WithField("key1","value1").WithFields(m).Say("Bye...")

```