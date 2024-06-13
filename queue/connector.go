package queue

// DispatcherOption dispatcher option
type DispatcherOption interface {
	Apply(Dispatcher)
}

type DispatcherOptionClause func(Dispatcher)

func (fn DispatcherOptionClause) Apply(dispatcher Dispatcher) {
	fn(dispatcher)
}

type DispatcherConnector interface {
	Connect(options ...DispatcherOption)
}

type DispatcherResolver interface {
	Resolve(config map[string]any) Dispatcher
}

type FailedJobProviderOption interface {
	Apply(FailedJobProvider)
}

type FailedJobProviderOptionClause func(FailedJobProvider)

func (oc FailedJobProviderOptionClause) Apply(provider FailedJobProvider) {
	oc(provider)
}

type FailedJobProviderConnector interface {
	Connect(options ...FailedJobProviderOption)
}

type FailedJobProviderResolver interface {
	Connect(config map[string]any) FailedJobProvider
}
