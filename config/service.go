package config

var (
    ServiceDomain = "https://127.0.0.1" //注意必须要加协议
    ServicePort   = "8090"
    ServiceEnv    = "local" //local,production,uat,test
)

func IsLocal() bool {
    return Environment("local")
}

func Environment(env ...string) bool {
    for _, s := range env {
        if ServiceEnv == s {
            return true
        }
    }
    return false
}
