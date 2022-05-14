# Syniti Take-Home Assignment

## Questions:

### What’s your proudest achievement? It can be a personal project or something you’ve worked on professionally. Just a short paragraph is fine, but I’d love to know why you’re proud of it.

The achievement I am most proud of is founding [my own company](https://knowtworthy.com/) and [taking that company through a startup incubator](https://news.engineering.utoronto.ca/startups-to-watch-from-hatchery-demo-day-2018/). Along with the founding team, I worked to build a business case for our meeting minutes software solution and pitch it to investors. Furthermore, I built our Minimum Viable  Product to give to users and scaled that MVP into the complete product that we were able to sell to customers. What makes me proud about this achievement is that I dove into the deep end by starting a company from scratch and managed to build it into a functioning business. This experience taught me how to architect, build, and manage a complex software system from the ground up.

### What's a personal project you're currently working on? This could be a coding side project, hobby, or otherwise real world project you're working on.

[My current personal project](https://github.com/adambcomer/WiscKey) is a database storage engine based on the [WiscKey paper by Lu et al. 2017](https://www.usenix.org/system/files/conference/fast16/fast16-papers-lu.pdf). This database storage engine is based on a Log-Structured Merge-Tree storage system, but it separates the keys and values into different logs. By splitting the keys and values, the storage engine reduces the amount of space and write amplification on the hard disks. My project is to implement this database storage engine in C. In university, I took a course on C, but I never did anything more than a few assignments in the language. The goal is to hone my computer science skills by building a database storage engine that is close to the metal.

### Tell us about a technical book or article you read recently, why you liked it, and why we should read it.

One of the technical resources I recently read was the [Readings in Database Systems, 5th Edition](http://www.redbook.io/index.html). This resource is a primer for someone interested in database systems and the internal systems within a database. One aspect I like a lot is this resource contains a lot of history to explain to a younger reader, such as myself, the origins of many modern database systems. I enjoyed reading this primer because it gave me a starting place to investigate specific database system components that I would love to implement in the future. I would recommend this primer to anyone who works with databases. Many engineers treat databases as black boxes that hold data and return it upon issuing a query. Understanding how an RDMS optimizes a query or the isolation levels within a database gives a software engineer a better understanding of the guarantees and limitations of their databases. In turn, engineers that have a deeper understanding of database systems can write better, more efficient code.

### Tell us about one of your favorite products (physical or software) and one specific aspect that makes it truly great.

My favorite product is Google Search. As a product, it is extremely simple, a single text box. Behind that single search box, it ranks millions of pages for relevancy based on a user's query. Google's ability to extract the latent structure of the web and parse the ambiguity of the web is what makes it so powerful. When I am building something, I use this product as my model for an ideal software system. I try to make my projects simple to use and capable of inferring the details that a non-expert wouldn't know to set upfront. And, if possible, I try to use the latent structure within a project to design a cleaner solution for end users.

## Code:

I wrote the assignment in Golang. Inside `main.go`, you will find the application entrypoint to read the records from the file and filter out the invalid/duplicate records. In `record.go`, you will find the record struct and a method to determine if a record is independently valid.

### Run Application

```shell
go run .
```

### Run Tests

```shell
go test . -v
```