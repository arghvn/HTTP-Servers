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

// Example :

// Resp,  _  :=  http.Get(http://google.com/)
// Body := make([ ]byte, 99999)
// Resp.body.Read(body)
// Fmt.println(string(body))

// It can be said that the resp structure is body and the body is interface type.
// Both resp and body are a kind of structure.
// Of course, this point is true when we are going to receive these two from the browser,
// if we receive these two items, we have textured them.
// But these structures actually set up an interface type.
// But while defining each of these, we typed them into the interface.
// It was actually an interface and then turned into a structure.
// This point is due to the rules of the package inside the said package,
// when the structure is created, we have to type it in the interface so
//  that the structure can be assigned to these two items.

// Inside the http package, the body is made of structure
// type and in order to be restored (works)
// Must have a series of methods.
// If it does not have these methods,
// it will not be restored and the program will give an error.
// With Intface, it is clear what forced us to have these methods.
// In fact, when we equate typing with a specific interface,
//  we mean that we want the body to have the required methods.

// There is also no need for the body to be a structure,
//  even if the body type is a String, Int, or whatever.
// But it has the methods mentioned, it can be used.
// So whenever we see that a function returns an interface,
//  that is, anything from the interface type,
// It can be concluded that it actually restores another type.
// We may not know that type and we do not know what that type is.
// It can only be said with certainty that the type being restored
// has the methods that the interface needs.

// So in this example we know that Resp is a structure that has a body.
// In the third line, Resp.body is called.
// We know that any type of body (it does not matter what type it is) only
// this type must have the Read method.
// In the next step, we call the relevant Read method.
// How the Read method works is slightly different from other methods and functions.
// This method works by taking an input that must be of reference type
// (Pass by reference type).
// That is, take a reference.
// This method automatically fills in the received reference and returns whatever is needed.
// In the second line of this example we have created a body and in the third line
//  we have filled it and in the fourth line we have called it.
// We do not check what is called in Read.
// Body has changed here.
// If we print after the second body line, what it returns is empty.
// And if we print after the third line, it will return the answer it received from Google.
// In the second line we have created an array of bytes, the byte (a kind of code)
// is easily stringed and vice versa
// The array length we created is too long (99999).
// If the length of the answer array that Google gives us is longer than this defined length,
//  it will cause a problem.
// This answer is cut and only 999999 bytes are filled and converted to string to be displayed.
// To prevent this problem, we must use the For loop.
// Or even better than the loop, you can use Ioutil.readall,
//  which reads and returns the length of the response, whatever it is.

// So the read method takes input and fills it but does not return it.
