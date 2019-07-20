##{{ProjectName}}


**Run**

`make run`

**Run test**

Theses tests should be run without any external resource. So all dependencies must be mocked.


Run all tests:
`make test`

Run specific test:
`make test method=TestExtractToken`

**Run integration tests**

The integration are tests that needs access to test resource like postgres, cassandra and kafka

Run all tests:
`make integration`

Run specific test:
`make integration method=TestFirebaseHttp`

**Release**

`make release`

**Deploy**

- Increase version number in version.sh file

- `make deploy`

**Send Raw Push**

 
 `
 curl --header "Content-Type: application/json;charset=UTF-8" -X POST -d '
  {
     "user_ids" : [2058851620],
     "data" : {}
  }
  ' http://test:test@192.168.10.152:5050/admin/push/send
  `
  
  `
  curl --header "Content-Type: application/json;charset=UTF-8" -X POST -d '
   {
      "user_ids" : [2058851620],
      "title": "salam",
      "text" : "body",
      "image" : "https://res.cloudinary.com/demo/image/upload/w_250,h_250,c_mfit/w_700/sample.jpg",
      "icon"  : ""
      "tag" : "1"
      "data" : {}
   }
   ' http://test:test@192.168.10.152:5050/admin/push/send
   `
   
**Send Push To All Users**   
  `
  curl --header "Content-Type: application/json;charset=UTF-8" -X POST -d '
   {
      "title": "salam",
      "text" : "body",
      "image" : "https://res.cloudinary.com/demo/image/upload/w_250,h_250,c_mfit/w_700/sample.jpg",
      "icon"  : ""
      "tag" : "1"
      "data" : {}
   }
   ' http://test:test@192.168.10.152:5050/admin/push/sendtoall
   `
   
**Get result of push**      
   
`
curl http://test:test@192.168.10.152:5050/admin/push/result
`