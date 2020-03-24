package drops

import "github.com/Viva-con-Agua/echo-pool/nats"

type StringList []string

var LogoutList StringList

// add string to list, used for uuid
func (l StringList) Add(uuid string) StringList {
	return append(l, uuid)
}

// delete string from list, used for uuid
func (l StringList) Delete(uuid string) StringList {
	// search position
	for i, v := range l {
		if v == uuid {
			//copy last element to position
			l[i] = l[len(l)-1]
			// make last element empty
			l[len(l)-1] = ""
			// return reduced list
			return l[:len(l)-1]
		}
	}
	return l
}

func (l StringList) Contains(uuid string) bool {
	if len(l) == 0 {
		return false
	}
	for _, v := range l {
		if v == uuid {
			return true
		}
	}
	return false
}

func NatsLogout() {
	nats.Nats.Subscribe("LOGOUT", func(s string){
		LogoutList = LogoutList.Add(s)
	})
}
