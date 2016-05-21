package idgen

import "launchpad.net/gozk"
import "time"
import "fmt"
import "strings"
import "strconv"

type ZkGen struct {
	conn     *zookeeper.Conn
	duration time.Duration
	host     string
}

func NewZkGen(host string, duration time.Duration) *ZkGen {

	ins := new(ZkGen)
	ins.duration = duration
	ins.host = host

	zk, session, err := zookeeper.Dial(host, duration)

	if err != nil {
		fmt.Println("Can't connect:%v", err)
		return nil
	}

	event := <-session

	if event.State != zookeeper.STATE_CONNECTED {

		fmt.Println("Can't connect:%v", event)

		return nil
	}

	ins.conn = zk

	return ins
}

func (this *ZkGen) Get() uint16 {

	new_path, err := this.conn.Create("/counter-", "0", zookeeper.SEQUENCE, zookeeper.WorldACL(zookeeper.PERM_ALL))

	if err != nil {

		fmt.Println("Can't create counter: %v", err)
		return 0
	}

	parts := strings.SplitN(new_path, "-", 2)

	if len(parts) < 2 {

		fmt.Println("not valid new_path")

		return 0
	}

	sSeqNum := parts[1]

	iSeqNum, errConv := strconv.ParseUint(sSeqNum, 10, 64)

	if errConv != nil {

		fmt.Println("not valid str")

		return 0
	}

	return uint16(iSeqNum & 0x03ff)
}
