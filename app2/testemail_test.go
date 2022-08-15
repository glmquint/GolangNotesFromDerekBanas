package app2
import("testing")

func TestIsEmail(t *testing.T){
    _, err := IsEmail("hello")
    if err == nil {
        t.Error("hello is not an email") // PASS
    }
    _, err = IsEmail("glmquint@gmail.com")
    if err != nil {
        t.Error("glmquint@gmail.com is an email") // FAIL
    }
    _, err = IsEmail("glmquint@gmail")
    if err == nil {
        t.Error("glmquint@gmail is not an email") // FAIL
    }
}
