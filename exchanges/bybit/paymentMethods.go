package bybit

var bybitPaymentMethodsMap = map[string]string{
	"1":   "A-Bank",
	"2":   "Alipay",
	"3":   "Wechat",
	"4":   "7-Eleven",
	"5":   "Advcash",
	"6":   "Airtel Money",
	"7":   "AirTM",
	"9":   "Alior Bank",
	"10":  "Apple Pay",
	"11":  "Bank of Georgia",
	"12":  "Bank Jago",
	"13":  "Bank Pivdenny",
	"14":  "Bank Transfer",
	"15":  "Bank Transfer(pero)",
	"16":  "BBVA",
	"17":  "BCA Mobile",
	"18":  "Cash Deposit to Bank",
	"19":  "CIMB Niaga",
	"20":  "Coin.ph",
	"21":  "Coinpay",
	"22":  "Credit Agricole",
	"23":  "DANA",
	"24":  "DSK Bank",
	"25":  "EasyPay",
	"26":  "Foward Bank",
	"27":  "FPS",
	"28":  "Gcash",
	"29":  "Google Pay",
	"30":  "GoPay",
	"31":  "Idea Bank",
	"32":  "IMPS",
	"33":  "JKOPAY",
	"34":  "KredoBank",
	"35":  "LandBank of the Philippines",
	"36":  "Line Pay",
	"37":  "Maybank",
	"38":  "Metropolitan Bank of the Philippines",
	"40":  "Mobile Top-up",
	"41":  "MoMo",
	"42":  "Monikwik",
	"43":  "Monobank",
	"44":  "MTS-Bank",
	"45":  "NEO",
	"46":  "Oschadbank",
	"48":  "OTP",
	"49":  "OTP Bank",
	"50":  "Payall",
	"51":  "Payeer",
	"52":  "Paymaya",
	"54":  "PayPal",
	"55":  "Paytm",
	"56":  "Perfect Money",
	"57":  "PhonePe",
	"58":  "PNB",
	"59":  "Post Bank",
	"60":  "Privat Bank",
	"61":  "PUMB",
	"62":  "QIWI",
	"63":  "Raiffeisen Bank Aval",
	"64":  "Raiffeisenbank",
	"65":  "Revolut",
	"66":  "Rosselkhozbank",
	"69":  "Shopee",
	"70":  "SmartPay",
	"71":  "Sovkimbank",
	"72":  "Sportbank",
	"73":  "SWIFT",
	"74":  "Tascombank",
	"75":  "Tinkoff",
	"76":  "Touch n Go",
	"77":  "Transfers with specific bank",
	"78":  "Wise",
	"79":  "True Money",
	"80":  "Ukrsibbank",
	"81":  "UnionBank of the Philippines",
	"82":  "UPI",
	"83":  "ViettelPay",
	"84":  "VNPay",
	"85":  "VNPT Pay",
	"87":  "Western Union",
	"88":  "Yandex.Money",
	"89":  "ZaloPay",
	"90":  "Cash in Person",
	"91":  "Grab Pay",
	"93":  "Bank of America",
	"95":  "Bundle",
	"96":  "Cash app",
	"97":  "Chipper Cash",
	"98":  "Dash-App",
	"99":  "Flip",
	"100": "Garanti",
	"101": "GoMoney",
	"102": "Home Credit Bank(Russia)",
	"103": "Inecobank",
	"104": "Interbank",
	"105": "Itaú Brazil",
	"106": "Jenius PayMe",
	"107": "K Bank",
	"108": "LinkAja",
	"109": "M-pesa Paybill",
	"110": "Mandiri Pay",
	"111": "NETELLER",
	"112": "Orange Money-OM",
	"113": "OVO",
	"114": "Papara",
	"115": "Permata Me",
	"116": "Pipol Pay",
	"117": "SEA Bank",
	"118": "SEPA",
	"119": "ShopeePay-SEA",
	"120": "Yap! (BNI)",
	"121": "Zelle",
	"122": "Ziraat",
	"123": "Banco Brubank",
	"124": "Banco del Sol",
	"125": "Banco Rio",
	"126": "BancolombiaS.A",
	"127": "Bank RBK",
	"128": "Lemon Cash",
	"129": "Mercadopago",
	"130": "Pago Movil",
	"131": "Prex",
	"132": "Reba",
	"133": "Santander Poland",
	"134": "Uala",
	"135": "Wilobank",
	"136": "ABN AMRO",
	"137": "Banesco",
	"138": "Bizum",
	"139": "BNP Paribas",
	"142": "Davivienda S.A",
	"143": "Easypaisa-PK Only",
	"144": "ForteBank",
	"145": "iDEAL",
	"146": "Idram",
	"147": "ING",
	"148": "Jazzcash",
	"149": "Jysan Bank",
	"150": "Kaspi Bank",
	"151": "Kuveyt Turk",
	"152": "La Banque postale",
	"153": "MICB",
	"154": "Millenium",
	"155": "Moneygram",
	"156": "N26",
	"157": "Paysend.com",
	"158": "Paysera",
	"159": "PKO Bank",
	"160": "Postepay",
	"162": "Skrill",
	"163": "Societe Generale",
	"164": "Sofort",
	"165": "TBC Bank",
	"166": "UniCredit",
	"167": "Uphold",
	"168": "Victoriabank",
	"169": "Vodafone cash",
	"170": "WorldRemit",
	"171": "Yape",
	"172": "Upaisa",
	"173": "Absolut Bank",
	"174": "BOS Bank",
	"175": "stc pay",
	"176": "M-Pesa Kenya(Safaricom)",
	"177": "M-Pesa(Vodafone)",
	"178": "MTN Mobile Money",
	"179": "Afriex",
	"180": "EQ Bank",
	"182": "Arca",
	"183": "Tigo Pesa",
	"184": "Tigo Money",
	"185": "Rosbank",
	"186": "Banco Pichinca",
	"187": "Banesco Panama",
	"188": "Banco Guayaquil",
	"189": "Zinli",
	"190": "Produbanco",
	"191": "Banco general panama",
	"192": "Credit Bank of Peru",
	"193": "Mercantil Bank Panama",
	"194": "ABA",
	"195": "Mony",
	"196": "Bank Transfer (Argentina)",
	"197": "Banco del pacifico",
	"198": "Banistmo Panama",
	"199": "Banco Bolivariano",
	"200": "Banco de Credito",
	"201": "ScotiaBank Peru",
	"202": "MAIB",
	"203": "Halyk Bank",
	"204": "Bank transfer (Cambodia)",
	"205": "Itaú Uruguay",
	"206": "Facebank International",
	"207": "Venmo",
	"208": "Bank of the Republic of Uruguay",
	"209": "Nequi",
	"210": "BAC Credomatic",
	"211": "CenterCredit Bank",
	"212": "KapitalBank",
	"213": "Ameriabank",
	"214": "Banco Santander Uruguay",
	"215": "Bank of the philippine Islands (BPI)",
	"216": "BAC Costa Rica",
	"217": "Bakong",
	"218": "OCA Blue",
	"219": "Bank Transfer (Turkey)",
	"220": "Bank transfer (India)",
	"221": "Bank transfer (Costa Rica)",
	"222": "Banco Agricola SV",
	"223": "Banco Estado",
	"224": "FinComBank",
	"225": "Interac e-transfer",
	"226": "Lloyds Bank",
	"227": "Scotiabank Panama",
	"228": "Transferencia ACH (Panamá)",
	"229": "WingMoney",
	"230": "Banco de Chile",
	"231": "CitiBank (Russia)",
	"232": "Monzo",
	"233": "BCR Bank",
	"234": "Banco BAC Credomatic SV",
	"235": "Banks transfer (Vietnam)",
	"236": "Daviplata",
	"237": "Red Pagos",
	"238": "UniBank",
	"239": "Banco de Costa Rica",
	"241": "Multibank Panama",
	"242": "BBVA Uruguay",
	"244": "TD Bank",
	"245": "BCEL",
	"246": "Banco Hipotecario SV",
	"247": "Banco Lafise Nicaragua",
	"248": "Banco Promerica SV",
	"249": "Bank Transfer (Middle East)",
	"250": "Cashpack",
	"251": "Central Bank of Uruguay",
	"252": "Pix",
	"253": "BNC Banco Nacional de Credito",
	"254": "Banco Fucerep",
	"255": "Banco Union",
	"256": "Bandes Uruguay",
	"257": "Blik",
	"258": "CIB Bank",
	"259": "CashU",
	"260": "DayhanBank",
	"261": "Dashen",
	"262": "Eurasian Bank",
	"263": "Home Credit Kazakhstan",
	"264": "Nagad",
	"265": "Pasha Bank",
	"266": "Recargas prepago",
	"267": "Rocket",
	"268": "Scotiabank Colpatria",
	"269": "Scotiabank Uruguay",
	"270": "Turon Bank",
	"271": "Uralsib Bank",
	"272": "Vostochny Bank",
	"273": "Xalq Bank",
	"274": "ЮMoney",
	"275": "bKash",
	"276": "Banco de Bogota",
	"277": "Movii",
	"278": "Banco Falabella",
	"279": "Efecty",
	"280": "Altyn Bank",
	"281": "Nurbank",
	"282": "Humo",
	"283": "Uzcard",
	"284": "Paynet",
	"285": "InfinBank",
	"286": "Uzpromstroybank",
	"287": "Asaka Bank",
	"288": "Ipoteka Bank",
	"289": "Agrobank",
	"290": "Uzbek National Bank",
	"291": "Optima Bank",
	"292": "mBank",
	"293": "Commercial bank Kyrgyzstan",
	"294": "Elcart",
	"295": "DemirBank",
	"296": "O!Money",
	"297": "RSK Bank",
	"298": "KICB",
	"299": "STP",
	"300": "Citibanamex",
	"301": "OXXO",
	"302": "Bank Transfer (Venezuela)",
	"303": "SEPA Instant",
	"304": "CIB",
	"305": "National Bank of Egypt (NBE)",
	"306": "Banque Misr",
	"307": "Alex Bank",
	"308": "QNB",
	"309": "Qatar National Bank QNB",
	"310": "Ahlibank",
	"311": "HSBC Uruguay",
	"312": "Banque du Caire",
	"313": "Alinma Bank",
	"314": "National Bank of Kuwait (K.S.C) (NBK)",
	"315": "urpay",
	"316": "Mercantil",
	"317": "Banco de Venezuela",
	"318": "Provincial",
	"319": "Western Discount Bank",
	"320": "Bancaribe",
	"321": "Bancamiga",
	"322": "Banplus",
	"323": "Ubii Pagos",
	"324": "Banco Activo",
	"325": "Banco Plaza",
	"326": "Recarga Pines",
	"327": "Tu Dinero Ya (Banco Plaza)",
	"328": "ZEN",
	"329": "Starling Bank",
	"330": "Bunq",
	"331": "HSBC Bank Middle East Limited - Bahrain",
	"332": "MTBank",
	"334": "PriorBank",
	"335": "Statusbank",
	"336": "Technobank",
	"337": "TK Bank",
	"338": "RBC Royal Bank",
	"339": "CIBC",
	"340": "BMO",
	"341": "Tangerine Bank",
	"342": "GEO Pay",
	"343": "Al Rajhi Bank",
	"344": "Saudi National Bank (AlAhli Bank)",
	"345": "Riyad Bank",
	"346": "The Saudi British Bank (SABB)",
	"347": "Arab Bank",
	"348": "Liv. KSA",
	"349": "ABB Bank",
	"350": "Bank Transfer (Taiwan)",
	"351": "WavePay",
	"352": "FNB-ewallet",
	"353": "PicPay",
	"354": "PagSeguro",
	"355": "Bank Pekao",
	"356": "Getin Noble Bank",
	"361": "CIH Bank",
	"362": "Attijari Bank",
	"363": "Attijariwafa National Bank",
	"364": "BMCE Bank",
	"365": "BRD Bank",
	"367": "Crédit Banque Populaire du Maroc",
	"368": "Ligo",
	"369": "BanBif",
	"370": "Plin",
	"371": "ArmEconomBank",
	"372": "ArmBusinessBank",
	"373": "IDBank",
	"374": "Bank BRI",
	"375": "Swish",
	"376": "OCBC NISP",
	"377": "Sberbank",
	"378": "Gazprombank",
	"379": "Alfa Bank",
	"380": "Otkritie Bank ",
	"381": "VTB Bank",
	"382": "SBP",
	"383": "MIR",
	"384": "CEC",
	"385": "Bank Transilvania",
	"386": "IME Pay",
	"387": "Khalti",
	"388": "Esewa",
}