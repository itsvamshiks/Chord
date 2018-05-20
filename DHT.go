package main

/**
- Phani Teja Kesha (KP38691)
- Seshagopalan Narasimhan (XP27536)
- Siddhant Goenka (VI91418)
- Vamshi Krishna Sai Nagabandi (SO03137)
*/
type DHT []string

func (d DHT) get(key int) string	{
	if key < len(d)	{
		return d[key]
	}else	{
		panic("Key not in Ring range");
	}
}//end of method


func (d DHT) put(key int, value string)	 bool {
	if key < len(d)	{
		d[key] = value
		return true
	}else	{
		return false
	}
}//end of method


func (d DHT) remove(key int) bool	{
	if key < len(d)	{
		var empty string
		d[key] = empty
		return true
	}else	{
		return false
	}
}//end of method
