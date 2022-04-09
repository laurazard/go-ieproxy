//go:build !windows && !darwin
// +build !windows,!darwin

package ieproxy

func getConf() ProxyConf {
	return ProxyConf{}
}

func reloadConf() ProxyConf {
	return getConf()
}

func overrideEnvWithStaticProxy(pc ProxyConf, setenv envSetter) {
}
