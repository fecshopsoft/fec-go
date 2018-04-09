package security

import(
    "github.com/dgrijalva/jwt-go"
    "time"
    "fmt"
    "errors"
)

var hmacSampleSecret = []byte("fdasfdsafdsfdsf534re#$ed");
// token 加密
func JwtSignToken(data interface{}) (string, error) {
    nowTime := time.Now().Unix()
    expired := int64(nowTime + 86400)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "data": data,
        "logined": 1,
        "expired": expired,
    })
    // Sign and get the complete encoded token as a string using the secret
    tokenString, err := token.SignedString(hmacSampleSecret)
    return tokenString, err
}

func JwtParse(tokenString string) (interface{}, int, int64, error) {
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
        return nil, 0, 0, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        lg, ok := claims["logined"].(float64)
        if !ok {
            return nil, 0, 0, errors.New("JwtParse logined convert float64 fail")
        }
        logined := int(lg)
        ex, ok := claims["expired"].(float64)
        if !ok {
            return nil, 0, 0, errors.New("JwtParse expired convert float64 fail")
        }
        expired := int64(ex)
        return claims["data"], logined, expired, nil
    } else {
        return nil, 0, 0, errors.New("JwtParse token fail")
    }

}



// 通过siteUid ，得到加密的access_token
func JwtSignAccessToken(siteUid string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "website_uid": siteUid,
    })
    // Sign and get the complete encoded token as a string using the secret
    tokenString, err := token.SignedString(hmacSampleSecret)
    return tokenString, err
}
//  通过加密的access_token，通过siteUid
func JwtParseAccessToken(tokenString string) (string, error) {
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
        return "", err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        website_uid, ok := claims["website_uid"].(string)
        if !ok {
            return "", errors.New("JwtParseAccessToken website_uid convert string fail")
        }
        return website_uid, nil
    } else {
        return "", errors.New("JwtParseAccessToken token fail")
    }

}




