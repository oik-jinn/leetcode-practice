package main

import "fmt"
import "os"

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
	os.Exit(1)
}

func restoreIpAddresses(s string) []string {
    
    arr := cutString("", s, 4)

    for i := 0; i < len(arr); i++ {
        arr[i] = arr[i][1:]
    }
    return arr
}

func cutString(prefix, str string, num int) []string {
	if len(str) < num || len(str) > 3*num {
		return nil
	}

	if num == 1 {
		if len(str) > 1 && str[0]== '0' {
			return nil
		}

		if (len(str)==3 && str <= "255") || len(str)<3 {
			return []string{prefix+"."+str}
		} else {
			return nil
		}
	}

	result := make([]string, 0)
	for i := 1; i< 4 && i < len(str); i++ {

		if len(str[0:i]) > 1 && str[0] == '0' {
            continue
        }
        if len(str[0:i]) == 3 && str[0:i] > "255" {
            continue
        }

        arr := cutString(prefix+"."+str[:i], str[i:], num-1)
        if nil == arr {
        	continue
        }

        result = append(result, arr...)
	}

	return result
}    
