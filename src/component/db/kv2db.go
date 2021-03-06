package db

import (
	gp "github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"common/logger"
	"protobuf"
	rpcplus "component/rpc"
	"fmt"
)

func KVQuery(db *rpcplus.Client, table, uid string, value interface{}) (exist bool, err error) {
	//ts("KVQuery", table, uid)
	//defer te("KVQuery", table, uid)

	var reply protobuf.DBQueryResult

	err = db.Call(protobuf.Network_Protocol(protobuf.DB_Protocol_eDB_Query), protobuf.DBQuery{table, uid}, &reply)

	if err != nil {
		logger.Error("KVQuery Error On Query %s : %s (%s)", table, uid, err.Error())
		return
	}

	switch reply.Code {
	case protobuf.Ok:

		var dst []byte

		dst, err = snappy.Decode(nil, reply.Value)

		if err != nil {
			logger.Error("KVQuery Unmarshal Error On snappy.Decode %s : %s (%s)", table, uid, err.Error())
			return
		}


		switch value.(type) {
			case gp.Message:

				err = gp.Unmarshal(dst, value.(gp.Message))

				if err != nil {
					logger.Error("KVQuery Unmarshal Error On Query %s : %s (%s)", table, uid, err.Error())
					return
				}
			case *[]byte:

				*(value.(*[]byte)) = dst

			default:
				logger.Error("KVQuery args type error %v", value)
				return
		}

		exist = true
		return

	case protobuf.NoExist:
		return
	}

	logger.Error("KVQuery Unknow DBReturn %d", reply.Code)

	return false, fmt.Errorf("KVQuery Unknow DBReturn %d", reply.Code)
}

func KVWrite(db *rpcplus.Client, table, uid string, value interface{}) (result bool, err error) {
	//ts("KVWrite", table, uid)
	//defer te("KVWrite", table, uid)

	var buf []byte

	switch value.(type) {
		case gp.Message:

			buf, err = gp.Marshal(value.(gp.Message))
			if err != nil {
				logger.Error("KVWrite Error On Marshal %s : %s (%s)", table, uid, err.Error())
				return
			}

		case []byte:

			buf = value.([]byte)

		default:
			logger.Error("KVWrite args type error %v", value)
		return
	}

	dst := snappy.Encode(nil, buf)

	var reply protobuf.DBWriteResult
	err = db.Call(protobuf.Network_Protocol(protobuf.DB_Protocol_eDB_Write), protobuf.DBWrite{table, uid, dst}, &reply)

	if err != nil {
		logger.Error("KVWrite Error On Create %s: %s (%s)", table, uid, err.Error())
		return
	}

	if reply.Code != protobuf.Ok {
		logger.Error("KVWrite Error On Create %s: %s code (%d)", table, uid, reply.Code)
		return
	}

	result = true
	return
}

func KVDelete(db *rpcplus.Client, table, uid string) (result bool, err error) {
	//ts("KVDelete", table, uid)
	//defer te("KVDelete", table, uid)

	var reply protobuf.DBDelResult
	err = db.Call(protobuf.Network_Protocol(protobuf.DB_Protocol_eDB_Delete), protobuf.DBDel{table, uid}, &reply)

	if err != nil {
		logger.Error("KVDelete Error On %s: %s (%s)", table, uid, err.Error())
		return
	}

	if reply.Code != protobuf.Ok {
		logger.Error("KVDelete Error On %s: %s code (%d)", table, uid, reply.Code)
		return
	}

	result = true
	return
}
