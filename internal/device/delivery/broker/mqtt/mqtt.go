package mqtt

const DEFAULT_QOS_LEVEL = 1

type HandlerFunc func(msg []byte) error
