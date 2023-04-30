

# Single Responsibility

این اصل میگه هر کلاس باید یک کار را انجام دهد به عبارت دیگر یک کلاس فقط باید به یک دلیل تغییر کنه

برای رعایت این اصل ، یکی از نکات این است که ریسیور های استراکچر را با دقت بنویسیم ، زیرا شاید یه استراکت بیش از چند ریسیور داشته باشد و در نتیجه بیش از یک کار انجام دهد


# Open/close

برای رعایت این اصل ، باید برنامه جوری نوشته شده باشه که در صورتی که بخوایم تغییری در آینده ایجاد کنیم ، تنها چیزی به برنامه اضافه کنیم و کد های قبلی دچار تغییر زیاد یا حذف نشن



# Liskov

میگه که در صورتی که از یک شی ، چند نمونه بسازیم ، باید برنامه جوری باشه که اگر نمونه ها را جا بجا کنیم ، برنامه به خطا نخوره


برای رعایت این اصل ، باید در ایجاد وراثت دقت کنیم ، مثلا اگر کلاس پدر یک متد دارد ، نباید کلاس فرزند همان متد را داشته باشد 


مثلا برای ساخت مسطیل نیاز به عرض و طول داریم در حالی که برای مربع نیازی به هر دو نداریم: 



    type ConvexQuadrilateral interface {
      GetArea() int
    }

    type Rectangle interface {
      ConvexQuadrilateral
      SetA(a int)
      SetB(b int)
    }

    type Oblong struct {
      Rectangle
      a int
      b int
    }

    func (o *Oblong) SetA(a int) {
      o.a = a
    }

    func (o *Oblong) SetB(b int) {
      o.b = b
    }

    func (o Oblong) GetArea() int {
      return o.a * o.b
    }

    type Square struct {
      Rectangle
      a int
    }

    func (o *Square) SetA(a int) {
      o.a = a
    }

    func (o Square) GetArea() int {
      return o.a * o.a
    }

    func (o *Square) SetB(b int) {
      //
      // should it be o.a = b ?
      // or should it be empty?
      //
    }

برای این کار می بایست وراثت را در خود اینترفیس ها پیاده کنیم

    type ConvexQuadrilateral interface {
      GetArea() int
    }

    type EquilateralQuadrilateral interface {
      ConvexQuadrilateral
      SetA(a int)
    }

    type NonEquilateralQuadrilateral interface {
      ConvexQuadrilateral
      SetA(a int)
      SetB(b int)
    }

    type NonEquiangularQuadrilateral interface {
      ConvexQuadrilateral
      SetAngle(angle float64)
    }

    type Oblong struct {
      NonEquilateralQuadrilateral
      a int
      b int
    }

    type Square struct {
      EquilateralQuadrilateral
      a int
    }

    type Parallelogram struct {
      NonEquilateralQuadrilateral
      NonEquiangularQuadrilateral
      a     int
      b     int
      angle float64
    }

    type Rhombus struct {
      EquilateralQuadrilateral
      NonEquiangularQuadrilateral
      a     int
      angle float64
    }
