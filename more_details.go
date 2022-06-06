package main

// RESP is a Go library that provides a reader, writer, and server implementation for the Redis RESP Protocol.

// RESP is short for REdis Serialization Protocol. While the protocol was designed specifically for Redis,
//  it can be used for other client-server software projects.
// Reader and Writer types for streaming RESP values from files, networks, or byte streams.
// Server Implementation for creating your own RESP server. Clients use the same tools and libraries as Redis.
// Append-only File type for persisting RESP values to disk.

// package resp provides a reader, writer, and server implementation for the RESP protocol. http://redis.io/topics/protocol
// RESP is short for "REdis Serialization Protocol". While the protocol was designed specifically
//  for Redis, it can be used for other client-server software projects.
// RESP has the advantages of being human readable and with performance of a binary protocol.
// Tag literally means tag and etiquette, but in the world of programming, including web pages
//  and HTML coding, tags mean predefined grammatical symbols of this language, which are usually
//  divided into two parts. The first part indicates the beginning of the tag and the second part indicates the end.
// The body defines the document, also known as the site body tag, and contains all the contents
//  of an HTML document, such as text, links, images, tables, lists, and so on.
// That means we have to create every visual item on our website inside the <body> tag.
// In other words, the visual framework of your site contains many tags.
// All of these visual tags are inside the <body> tag.
// We can specify the font size, font color and other items seen on web pages in this tag.
// The <body> tag includes an opening tag and a closing tag.
// The closed body tag is the end of the site's visual body code.
// This tag opens after the head tag closes and closes before the html tag closes.
// How to use the body tag:
// This tag introduces the main body of the HTML page and contains all the contents that
// can be displayed on the page such as posts, photos, links and..
// Only the contents of this section are displayed by the browser as output, and the output
// content codes are defined between the open and closed <body> tags.
// This tag can also contain graphic details of the page.
// This tag contains the body and the main and visual content of web pages such as text, links, tables, images and..
//  The body element is the backbone of a web page, and no page without this tag can display its other elements.
