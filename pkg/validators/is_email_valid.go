/* AWS Lambda in GoLang — The Ultimate Guide
   https://www.softkraft.co/aws-lambda-in-golang/

   NOTE:  this function isn't actually checking to see if an email address exists, 
   only that the string is of the proper form that it might be an email address
*/
package validators

  import "regexp"

  func IsEmailValid(email string) bool {
    var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{1,64}@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

    if len(email) < 3 || len(email) > 254 || !rxEmail.MatchString(email) {
        return false
    }
    return true
  }
