package server

/*
весь этот файл реализация шаблона для установки дефолтных значений
*/

type Options struct {
	Port string
	Scan Scan
}

type Option func(opt *Options)

func WithPort(port string) Option {
	return func(o *Options) {
		if port != "" {
			o.Port = port
		}
	}
}

func WithCrawler(scan Scan) Option {
	return func(opt *Options) {
		opt.Scan = scan
	}
}

func newOptions(opts ...Option) *Options {
	o := Options{Port: "8012"} //установка дефолтного порта

	for _, opt := range opts {
		opt(&o) //приминение новых параметров. Здесь же устанавливается недефолтный порт.
	}

	return &o

}
