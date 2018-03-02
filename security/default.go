package security

import(
    "github.com/dgrijalva/jwt-go"
    "time"
    "fmt"
)

var hmacSampleSecret = []byte("fdasfdsafdsfdsf534re#$ed");



func JwtSignToken(data interface{}) (string, error) {
    nowTime := time.Now().Unix()
    expired := int64(nowTime + 86400)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "data": data,
        "expired": expired,
    })
    
    // Sign and get the complete encoded token as a string using the secret
    tokenString, err := token.SignedString(hmacSampleSecret)
    return tokenString, err
}


func JwtParse(tokenString string) (interface{}, interface{}, error) {
    
    // Parse takes the token string and a function for looking up the key. The latter is especially
    // useful if you use multiple keys for your application.  The standard is to use 'kid' in the
    // head of the token to identify which key to use, but the parsed token (head and claims) is provided
    // to the callback, providing flexibility.
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Don't forget to validate the alg is what you expect:
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
        return hmacSampleSecret, nil
    })
    
    if err != nil {
        return nil, 0, nil
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims["data"], claims["expired"], nil
    } else {
        return nil, 0, nil
    }

}









