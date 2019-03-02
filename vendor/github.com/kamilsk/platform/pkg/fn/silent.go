package fn

// DoSilent accepts a result of fmt.Fprint* function family
// or Write method of the io.Writer interface and allow you to ignore it.
//
//  DoSilent(fmt.Fprintln(writer, "ignore the result"))
//
func DoSilent(int, error) {}

// DoSilent64 accepts a result of io.Copy* function family
// and allow you to ignore it.
//
//  DoSilent64(io.Copy(to, from))
//
func DoSilent64(int64, error) {}
