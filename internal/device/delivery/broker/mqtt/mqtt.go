package mqtt

const QOS_LEVEL = 1

type HandlerFunc func(msg []byte) error
