package main

import (
  "encoding/json"
  "errors"
  "fmt"
//  "log"
  "net/http"
  "os"
  "strconv"
  "github.com/akash1808/golang-gin/mysql"
  //"github.com/akash1808/golang-gin/redisdb" 
  //jwtmiddleware "github.com/auth0/go-jwt-middleware"
  //jwt "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
)

type Response struct {
  Message string `json:"message"`
}

type Jwks struct {
  Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
  Kty string   `json:"kty"`
  Kid string   `json:"kid"`
  Use string   `json:"use"`
  N   string   `json:"n"`
  E   string   `json:"e"`
  X5c []string `json:"x5c"`
}

type Joke struct {
  ID    int    `json:"id" binding:"required"`
  Likes int    `json:"likes"`
  Joke  string `json:"joke" binding:"required"`
}

/** we'll create a list of jokes */
var jokes = []Joke{
  Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
  Joke{2, 0, "What do you call a fake noodle? An Impasta."},
  Joke{3, 0, "How many apples grow on a tree? All of them."},
  Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
  Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
  Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
  Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
}

//var jwtMiddleWare *jwtmiddleware.JWTMiddleware

func main() {
	mysqlobject:=mysql.GetMySQL()
	if (mysqlobject==nil) {
	  fmt.Println("Failed to Initialize mYsql")
	  return
  }
//  redisobject:= redisdb.Get()
//  if (redisobject==nil) {
//          fmt.Println("Failed to Initialize mYsql")
//          return
//  }
  //jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
  //ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
  //    aud := os.Getenv("AUTH0_API_AUDIENCE")
  //    checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
  //    if !checkAudience {
  //      return token, errors.New("Invalid audience.")
  ///    }
      // verify iss claim
   //   iss := os.Getenv("AUTH0_DOMAIN")
   //   checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
   //   if !checkIss {
   //     return token, errors.New("Invalid issuer.")
   //   }
      
    //  cert, err := getPemCert(token)
    //  if err != nil {
    //    log.Fatalf("could not get cert: %+v", err)
    //  }
      
    //  result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
    //  return result, nil
    //},  
   // SigningMethod: jwt.SigningMethodRS256,
 // })
  
//  jwtMiddleWare = jwtMiddleware
  // Set the router as the default one shipped with Gin
  router := gin.Default()
  
  // Serve the frontend
  
  api := router.Group("/api")
  {
    api.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
        "message": "pong",
      })
    })
    api.GET("/jokes",  JokeHandler)
  }
  // Start the app
  router.Run("0.0.0.0:8080")
}

func getPemCert(token string) (string, error) {
  cert := ""
  resp, err := http.Get(os.Getenv("AUTH0_DOMAIN") + ".well-known/jwks.json")
  if err != nil {
    return cert, err
  }
  defer resp.Body.Close()
  
  var jwks = Jwks{}
  err = json.NewDecoder(resp.Body).Decode(&jwks)
  
  if err != nil {
    return cert, err
  }
  
 // x5c := jwks.Keys[0].X5c
 // for k, v := range x5c {
 //   if token.Header["kid"] == jwks.Keys[k].Kid {
 //     cert = "-----BEGIN CERTIFICATE-----\n" + v + "\n-----END CERTIFICATE-----"
 //   }
//  }
  
  if cert == "" {
    return cert, errors.New("unable to find appropriate key")
  }
  
  return cert, nil
}

// authMiddleware intercepts the requests, and check for a valid jwt token
func authMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    // Get the client secret key
  //  _ = jwtMiddleWare.CheckJWT(c.Writer, c.Request)
  //  if err != nil {
      // Token not found
  //    fmt.Println(err)
  //    c.Abort()
  //    c.Writer.WriteHeader(http.StatusUnauthorized)
  //    c.Writer.Write([]byte("Unauthorized"))
  //    return
  //  }
  }
}

// JokeHandler returns a list of jokes available (in memory)
func JokeHandler(c *gin.Context) {
  c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
  c.Header("Content-Type", "application/json")
  
  c.JSON(http.StatusOK, jokes)
}

