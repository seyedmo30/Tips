# unit test

بخش های بسیار کوچک و فانکشن ها را با این روش تست می کنیم .

در این روش ، ورودی و خروجی و نتیجه ی فانشکن را از قبل مشخص می کنیم و تست را اجرا می کنیم .

# integration test

گاهی اوقات تست به پیش نیاز هایی وابسته است ، مثلا برای تست یک ریکویست ، با وظیفه ثبت در دیتابیس یا خطا ، ما نیاز داریم که علاوه بر مشخص کردن ورودی و خروجی ، کانکتور ها ( یا dependency ها ) را هم مشخص کنیم . در این صورت می توانیم ابتدا تمامی کلاس ها یا استراکت ها را معرفی کنیم . مثلا تنظیمات کانکشن به دیتا بیس یا تنظیمات روتر . 

 
      type ParserTestSuite struct {
        suite.Suite
        Parser protocol.Parser
      }

      func (s *ParserTestSuite) SetupTest() {
        parser := viper_parser.Parser{}
        s.Parser = &parser

      }
      func TestExampleTestSuite(t *testing.T) {
        suite.Run(t, &ParserTestSuite{})
      }

      func (s *ParserTestSuite) TestMigration() {
        conn := NewConnectionPostgres(s.Parser)
        err := conn.Migration(context.Background())
        s.Nil(err)

      }
