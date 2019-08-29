package config

type metadataYaml struct {
	Name string `yaml:"name"`
}

type RouterConfigYaml struct {
	APIVersion string              `yaml:"apiVersion"` // InfraService/v1
	Kind       string              `yaml:"kind"`       // CtxRouterConfig
	Metadata   metadataYaml        `yaml:"metadata"`
	Spec       configCtxRouterSpec `yaml:"spec,omitempty"`
}
type configCtxRouterSpec struct {
	Hosts     []string    `yaml:"hosts"`
	Protocols []string    `yaml:"protocols"`
	Routes    []routeYaml `yaml:"routes"`
}

type routeYaml struct {
	Name          string            `yaml:"name"`
	Path          string            `yaml:"path"`
	RequestHeader requestHeaderYaml `yaml:"requestHeader,omitempty"`
	CookieValues  cookieValuesYaml  `yaml:"cookieValues,omitempty"`
	FormValues    formValuesYaml    `yaml:"formValues,omitempty"`
	SessionValues sessionValuesYaml `yaml:"sessionValues,omitempty"`
	CustomLogic   string            `yaml:"customLogic,omitempty"`
	Target        targetYaml        `yaml:"target"`
}

type requestHeaderYaml struct {
	Require   []valuesYaml `yaml:"require,omitempty"`
	Forbidden []valuesYaml `yaml:"require,omitempty"`
}
type cookieValuesYaml struct {
	Require   []valuesYaml `yaml:"require,omitempty"`
	Forbidden []valuesYaml `yaml:"require,omitempty"`
}
type formValuesYaml struct {
	Require   []valuesYaml `yaml:"require,omitempty"`
	Forbidden []valuesYaml `yaml:"require,omitempty"`
}
type sessionValuesYaml struct {
	Require   []valuesYaml `yaml:"require,omitempty"`
	Forbidden []valuesYaml `yaml:"require,omitempty"`
}
type valuesYaml struct {
	Name   string   `yaml:"name"`
	Values []string `yaml:"values,omitempty"`
	Regex  string   `yaml:"regex,omitempty"`
}

type targetYaml struct {
	Addresses  []string `yaml:"addresses"`
	Middleware []string `yaml:"middleware,omitempty"`
}
