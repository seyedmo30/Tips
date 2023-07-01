موارد مصرف اینترفیس :

# Embedding interfaces in structs

این روش برای وابسته کردن یک کلاس ( استراکت ) به اینترفیس است ، در حقیقت در کانستراکتور این کلاس باید شی ساخته شود که تمام رسیور های اینترفیس را ایمپلمنت کرده باشد 


    type controller struct {
    	dataStore  protocol.DataStore
    }
    
    func NewController(
    	dataStore protocol.DataStore
    	) *controller {
    	return &controller{
    		dataStore:  dataStore
    	}
    }

در حقیقت مجبوریم شی بسازیم که dataStore را ایمپلمنت کرده باشد


 # ddd

 ss
