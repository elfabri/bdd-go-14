package main

type perm string

const (
    Read  perm = "read"
    Write perm = "write"
    Exec  perm = "execute"
)

var Admin = "admin"
var User = perm("user")

func checkPermission(p perm) {
    // check the permission
}

func main() {
    // 1. In which case will the go compiler
    // rise an error?
    checkPermission(Write)
    checkPermission(User)
    checkPermission(Read)
    // checkPermission(Admin)  // <- error cannot use Admin (t string) as perm

    // 2. Will 'checkPermission(perm(Admin))'
    // raise an error from the compiler? no, it won't
    checkPermission(perm(Admin))
}
