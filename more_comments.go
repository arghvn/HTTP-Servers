package main

// In written code, Resp is a structure.
// This stratum has a field called body,
// for example, when a Google search result shows us an image as a result,
//  that image is related to this field of the stratum.
// body is a type of interface.

// For example in this part of the code:
// Func Get (url string) (resp *response , err error )

// *response
// Is a struct with its fields :
// Namly fileds = status (string)
// Statuscode (int)
// Proto(string)
// Protomajor(int)
// Protominor(int)

// Body in terms of typing
// Io.readcloser is a popular interface.

// Anything that wants to run the body must have the Read and Close methods.
// We conclude that Resp.body definitely has these two methods.
// Because if they did not, they would definitely not be returned.
