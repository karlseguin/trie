package trie

import (
	"math/rand"
	"strings"
)

// generate words for testing

//took these from https://github.com/weaver/moniker
var (
	words = []string{"aback","abaft","abandoned","abashed","aberrant","abhorrent","abide","abiding","abject","ablaze","able","able","abnormal","aboard","aboriginal","abortive","abounding","abrasive","abrupt","absent","absorbed","absorbing","abstracted","absurd","abundant","abusive","accelerate","accept","acceptable","accessible","accidental","accomplish","account","accurate","achieve","achieve","acid","acidic","acoustic","acoustics","acquire","acrid","act","acted","action","activate","activity","actor","actually","ad hoc","adamant","adapt","adaptable","add","addicted","addition","address","adhesive","adjoining","adjustment","administer","admire","admit","adopt","adorable","adventurous","advertisement","advice","advise","afford","afraid","aftermath","afternoon","afterthought","aggressive","agonizing","agree","agreeable","agreement","ahead","air","airplane","airport","ajar","alarm","alcoholic","alert","alert","alight","alike","alive","alleged","alley","allow","alluring","aloof","altered","amazing","ambiguous","ambitious","amount","amuck","amuse","amused","amusement","amusing","analyze","ancient","anger","angle","angry","animal","animated","announce","annoy","annoyed","annoying","answer","answer","ant","anticipate","ants","anxious","apathetic","apologize","apparatus","apparel","appear","applaud","apple","apples","appliance","applied","appoint","appraise","appreciate","approval","approve","aquatic","arbitrate","arch","argue","argument","arise","arithmetic","arm","army","aromatic","arrange","arrest","arrive","arrogant","art","ascertain","ashamed","ask","aspiring","assemble","assess","assist","assorted","assure","astonishing","attach","attack","attack","attain","attempt","attempt","attend","attention","attract","attraction","attractive","audited","aunt","auspicious","authority","automatic","available","average","avoid","awake","awake","aware","awesome","awful","axiomatic",
										"babies","baby","back","back","bad","badge","bag","bait","bake","balance","balance","ball","balloon","balls","ban","banana","band","bang","barbarous","bare","base","baseball","bashful","basin","basket","basketball","bat","bat","bath","bathe","battle","battle","bawdy","be","bead","beam","beam","bean","bear","bear","bears","beast","beat","beautiful","become","bed","bedroom","beds","bee","beef","beetle","befitting","beg","beggar","begin","beginner","behave","behavior","behold","belief","believe","bell","belligerent","bells","belong","bend","beneficial","bent","berry","berserk","beset","best","bet","better","bewildered","bid","big","bike","bikes","billowy","bind","bird","birds","birth","birthday","bit","bite","bite","bite-sized","bitter","bizarre","black","black-and-white","blade","bleach","bleed","bless","blind","blink","blood","bloody","blot","blow","blow","blue","blue-eyed","blush","blushing","board","boast","boat","boats","body","boil","boiling","bolt","bomb","bomb","bone","book","book","books","boorish","boot","border","bore","bored","boring","borrow","bottle","bounce","bouncy","boundary","boundless","bow","box","box","boy","boys","brain","brainy","brake","brake","branch","branch","brash","brass","brave","brawny","bread","break","breakable","breakfast","breath","breathe","breed","breezy","brick","bridge","brief","brief","bright","bright","bring","broad","broadcast","broken","brother","brothers","brown","bruise","brush","brush","bubble","bubble","bucket","budget","build","building","bulb","bump","bumpy","bun","burly","burn","burn","burst","burst","bury","bushes","business","bust","bustling","busy","butter","button","buy","buzz",
										"cabbage","cable","cactus","cagey","cake","cakes","calculate","calculating","calculator","calendar","call","callous","calm","camera","camp","camp","can","cannon","canvas","cap","capable","capricious","caption","car","card","care","care","careful","careless","caring","carpenter","carriage","carry","cars","cart","carve","cast","cast","cat","catalog","catch","cats","cattle","cause","cause","cautious","cave","ceaseless","celery","cellar","cemetery","cent","certain","chain","chair","chairs","chalk","challenge","chance","change","change","changeable","channel","charge","charming","chart","chase","cheap","cheat","check","cheer","cheerful","cheese","chemical","cherries","cherry","chess","chew","chicken","chickens","chief","childlike","children","chilly","chin","chivalrous","choke","choose","chop","chubby","chunky","church","circle","claim","clam","clammy","clap","clarify","class","classify","classy","clean","clean","clear","clear","clever","cling","clip","clock","clocks","cloistered","close","closed","cloth","clothe","cloud","clouds","cloudy","clover","club","clumsy","cluttered","coach","coach","coal","coast","coat","cobweb","coherent","coil","coil","cold","collar","collect","color","color","colorful","colossal","comb","comb","combative","come","comfort","comfortable","command","committee","common","communicate","company","compare","comparison","compete","competition","compile","complain","complete","complete","complex","compose","compute","conceive","concentrate","conceptualize","concern","concerned","conclude","condemned","condition","conduct","confess","confront","confuse","confused","connect","connection","conscious","conserve","consider","consist","consolidate","construct","consult","contain","continue","contract","control","control","convert","cooing","cook","cool","cooperative","coordinate","coordinated","copper","copy","copy","cord","cork","corn","correct","correlate","cost","cough","cough","counsel","count","country","courageous","cover","cover","cow","cowardly","cows","crabby","crack","crack","cracker","crash","crate","craven","crawl","crayon","crazy","cream","create","creator","creature","credit","creep","creepy","crib","crime","critique","crook","crooked","cross","crow","crowd","crowded","crown","cruel","crush","crush","cry","cry","cub","cuddly","cultured","cumbersome","cup","cure","curious","curl","curly","current","curtain","curve","curve","curved","curvy","cushion","cut","cut","cute","cute","cycle","cynical",
										"dad","daffy","daily","dam","damage","damaged","damaging","damp","dance","dangerous","dapper","dare","dark","dashing","daughter","day","dazzling","dead","deadpan","deafening","deal","dear","death","debonair","debt","decay","deceive","decide","decision","decisive","decorate","decorous","deep","deeply","deer","defeated","defective","defiant","define","degree","delay","delegate","delicate","delicious","delight","delightful","delirious","deliver","demonic","demonstrate","depend","dependent","depressed","deranged","describe","descriptive","desert","deserted","deserve","design","design","desire","desk","destroy","destruction","detail","detail","detailed","detect","determine","determined","develop","development","devilish","devise","diagnose","didactic","different","difficult","dig","digestion","diligent","dime","dinner","dinosaurs","direct","direction","direful","dirt","dirty","disagree","disagreeable","disappear","disapprove","disarm","disastrous","discover","discovery","discreet","discussion","disease","disgust","disgusted","disgusting","disillusioned","dislike","dispensable","dispense","display","disprove","dissect","distance","distinct","distribute","distribution","disturbed","dive","divergent","divert","divide","division","dizzy","do","dock","doctor","dog","dogs","doll","dolls","domineering","donkey","door","double","doubt","doubtful","downtown","drab","draconian","draft","drag","drain","drain","dramatic","dramatize","draw","drawer","dream","dreary","dress","dress","drink","drink","drip","drive","driving","drop","drop","drown","drug","drum","drum","drunk","dry","dry","duck","ducks","dull","dust","dust","dusty","dusty","dwell","dynamic","dysfunctional",
										"eager","ear","early","earn","earsplitting","earth","earthquake","earthy","easy","eat","eatable","economic","edge","edited","educate","educated","education","effect","efficacious","efficient","egg","eggnog","eggs","eight","elastic","elated","elbow","elderly","electric","elegant","elfin","eliminate","elite","embarrass","embarrassed","eminent","employ","empty","empty","enacted","enchanted","enchanting","encourage","encouraging","end","end","endurable","endure","energetic","enforce","engine","engineer","enhance","enjoy","enlist","enormous","ensure","enter","entertain","entertaining","enthusiastic","envious","equable","equal","erect","erratic","error","escape","establish","estimate","ethereal","evaluate","evanescent","evasive","even","event","examine","example","exceed","excellent","exchange","excite","excited","exciting","exclusive","excuse","execute","exercise","exhibit","exist","existence","exotic","expand","expansion","expect","expedite","expensive","experience","experiment","expert","explain","explode","express","extend","extra-large","extra-small","extract","exuberant","exultant","eye","eyes",
										"fabulous","face","face","facilitate","fact","fade","faded","fail","faint","fair","fairies","faithful","fall","fallacious","false","familiar","family","famous","fan","fanatical","fancy","fancy","fang","fantastic","far","far-flung","farm","farmer","fascinated","fast","fasten","fat","father","father","faucet","faulty","fax","fear","fear","fearful","fearless","feast","feather","feeble","feed","feel","feeling","feet","feigned","female","fence","fertile","festive","fetch","few","fiction","field","fierce","fifth","fight","fight","file","fill","film","filthy","finalize","finance","find","fine","finger","finger","finicky","fire","fire","fireman","first","fish","fit","five","fix","fixed","flag","flagrant","flaky","flame","flap","flash","flashy","flat","flavor","flawless","flee","flesh","flight","flimsy","fling","flippant","float","flock","flood","floor","flow","flower","flower","flowers","flowery","fluffy","fluttering","fly","fly","foamy","fog","fold","fold","follow","food","fool","foolish","foot","forbid","force","force","forecast","forego","foregoing","foresee","foretell","forget","forgetful","forgive","fork","form","form","formulate","forsake","fortunate","four","fowl","fragile","frail","frame","frame","frantic","free","freeze","freezing","frequent","fresh","fretful","friction","friend","friendly","friends","frighten","frightened","frightening","frog","frogs","front","fruit","fry","fuel","full","fumbling","functional","funny","furniture","furry","furtive","future","futuristic","fuzzy",
										"gabby","gainful","game","gamy","gaping","garden","garrulous","gate","gather","gaudy","gaze","geese","general","generate","gentle","get","ghost","giant","giants","giddy","gifted","gigantic","giraffe","girl","girls","give","glamorous","glass","gleaming","glib","glistening","glorious","glossy","glove","glow","glue","glue","go","goat","godly","gold","goldfish","good","good-bye","goofy","goose","gorgeous","govern","government","governor","grab","graceful","grade","graduate","grain","grandfather","grandiose","grandmother","grape","grass","grate","grateful","gratis","gray","grease","greasy","great","greedy","green","greet","grey","grieving","grin","grind","grip","grip","groan","groovy","grotesque","grouchy","ground","group","grow","growth","grubby","gruesome","grumpy","guarantee","guard","guarded","guess","guide","guide","guiltless","guitar","gullible","gun","gusty","guttural",
										"habitual","hair","haircut","half","hall","hallowed","halting","hammer","hammer","hand","hand","handle","hands","handsome","handsomely","handwrite","handy","hang","hanging","hapless","happen","happy","harass","harbor","hard","hard-to-find","harm","harmonious","harmony","harsh","hat","hate","hate","hateful","haunt","head","head","heady","heal","health","healthy","heap","hear","hearing","heart","heartbreaking","heat","heat","heavenly","heavy","hellish","help","help","helpful","helpless","hen","hesitant","hide","hideous","high","high-pitched","highfalutin","hilarious","hill","hissing","historical","history","hit","hobbies","hold","hole","holiday","holistic","hollow","home","homeless","homely","honey","honorable","hook","hook","hop","hope","hope","horn","horrible","horse","horses","hose","hospitable","hospital","hot","hot","hour","house","houses","hover","hug","huge","hulking","hum","humdrum","humor","humorous","hungry","hunt","hurried","hurry","hurt","hurt","hushed","husky","hydrant","hypnotic","hypothesize","hysterical",
										"ice","icicle","icky","icy","idea","identify","idiotic","ignorant","ignore","ill","ill-fated","ill-informed","illegal","illustrate","illustrious","imaginary","imagine","immense","imminent","impartial","imperfect","implement","impolite","important","imported","impossible","impress","improve","improvise","impulse","incandescent","include","income","incompetent","inconclusive","increase","increase","incredible","induce","industrious","industry","inexpensive","infamous","influence","inform","initiate","inject","injure","ink","inlay","innate","innocent","innovate","input","inquisitive","insect","insidious","inspect","inspire","install","instinctive","institute","instruct","instrument","insurance","insure","integrate","intelligent","intend","intensify","interest","interest","interesting","interfere","interlay","internal","interpret","interrupt","interview","introduce","invent","invention","inventory","investigate","invincible","invite","irate","iron","irritate","irritating","island","itch","itchy",
										"jaded","jagged","jail","jail","jam","jam","jar","jazzy","jealous","jeans","jelly","jellyfish","jewel","jittery","jobless","jog","join","join","joke","joke","jolly","journey","joyous","judge","judge","judicious","juggle","juice","juicy","jumbled","jump","jump","jumpy","justify","juvenile",
										"kaput","keen","keep","kept","kettle","key","kick","kick","kill","kind","kindhearted","kindly","kiss","kiss","kite","kitten","kittens","kitty","knee","kneel","knife","knit","knock","knot","knot","knotty","know","knowing","knowledge","knowledgeable","known",
										"label","labored","laborer","lace","lackadaisical","lacking","ladybug","lake","lame","lamentable","lamp","land","land","language","languid","large","last","last","late","laugh","laugh","laughable","launch","lavish","lawyer","lay","lazy","lead","lead","leaf","lean","lean","leap","learn","learned","learning","leather","leave","lecture","led","left","leg","legal","legs","lend","let","lethal","letter","letters","lettuce","level","level","level","lewd","library","license","lick","lie","lift","lifted","light","light","light","lighten","like","like","likeable","limit","limping","line","linen","lip","liquid","list","list","listen","literate","little","live","lively","lively","living","lizards","load","loaf","locate","lock","lock","locket","log","lonely","long","long","long-term","longing","look","look","loose","lopsided","lose","loss","loud","loutish","love","love","lovely","loving","low","low","lowly","lucky","ludicrous","lumber","lumpy","lunch","lunchroom","lush","luxuriant","lying","lyrical",
										"macabre","machine","macho","maddening","madly","magenta","magic","magical","magnificent","maid","mailbox","maintain","majestic","make","makeshift","male","malicious","mammoth","man","man","manage","manager","maniacal","manipulate","manufacture","many","map","map","marble","march","mark","mark","marked","market","market","married","marry","marvelous","mask","mass","massive","match","match","mate","material","materialistic","matter","mature","meal","mean","mean","measly","measure","measure","meat","meaty","meddle","mediate","medical","meek","meet","meeting","mellow","melodic","melt","melt","melted","memorize","memory","men","mend","mentor","merciful","mere","messy","metal","mice","middle","mighty","military","milk","milk","milky","mind","mindless","mine","mine","miniature","minister","minor","mint","minute","miscreant","mislead","miss","misspell","mist","mistake","misty","misunderstand","mitten","mix","mixed","moan","moaning","model","modern","modify","moldy","mom","momentous","money","monitor","monkey","month","moon","moor","morning","mother","motion","motionless","motivate","mountain","mountainous","mourn","mouth","move","move","mow","muddle","muddled","mug","multiply","mundane","murder","murky","muscle","mushy","music","mute","mysterious",
										"nail","nail","naive","name","name","nappy","narrow","nasty","nation","natural","naughty","nauseating","navigate","near","neat","nebulous","necessary","neck","need","need","needle","needless","needy","negotiate","neighborly","nerve","nervous","nest","nest","net","new","news","next","nice","nifty","night","nimble","nine","nippy","nod","noise","noiseless","noisy","nominate","nonchalant","nondescript","nonstop","normal","normalize","north","nose","nostalgic","nosy","note","note","notebook","notice","noxious","null","number","number","numberless","numerous","nut","nutritious","nutty",
										"oafish","oatmeal","obedient","obeisant","obese","obey","object","obnoxious","obscene","obsequious","observant","observation","observe","obsolete","obtain","obtainable","occur","ocean","oceanic","odd","offbeat","offend","offer","offer","office","officiate","oil","old","old-fashioned","omniscient","one","onerous","open","open","operate","operation","opinion","opposite","optimal","orange","orange","oranges","order","order","ordinary","organic","organization","organize","oriented","originate","ornament","ossified","outgoing","outrageous","outstanding","oval","oven","overcome","overconfident","overdo","overdraw","overflow","overhear","overjoyed","overrated","overt","overtake","overthrow","overwrought","owe","owl","own","owner",
										"pack","paddle","page","pail","pain","painful","painstaking","paint","paint","pale","paltry","pan","pancake","panicky","panoramic","paper","parallel","parcel","parched","parent","park","park","parsimonious","part","part","participate","partner","party","pass","passenger","past","paste","paste","pastoral","pat","patch","pathetic","pause","pay","payment","peace","peaceful","pear","peck","pedal","peel","peep","pen","pencil","penitent","perceive","perfect","perfect","perform","periodic","permissible","permit","perpetual","person","persuade","pest","pet","petite","petite","pets","phobic","phone","photograph","physical","picayune","pick","pickle","picture","pie","pies","pig","pigs","pilot","pin","pinch","pine","pink","pinpoint","pioneer","pipe","piquant","pizzas","place","place","placid","plain","plan","plane","planes","plant","plant","plant","plantation","plants","plastic","plastic","plate","plausible","play","play","playground","plead","pleasant","please","pleasure","plot","plough","plucky","plug","pocket","point","point","pointless","poised","poison","poke","police","polish","polish","polite","political","pollution","poor","pop","popcorn","porter","position","possess","possessive","possible","post","pot","potato","pour","powder","power","powerful","practice","praised","pray","preach","precede","precious","predict","prefer","premium","prepare","prescribe","present","present","preserve","preset","preside","press","pretend","pretty","prevent","previous","price","pricey","prick","prickly","print","print","prison","private","probable","process","process","procure","produce","produce","productive","profess","profit","profuse","program","progress","project","promise","promote","proofread","property","propose","prose","protect","protective","protest","proud","prove","provide","psychedelic","psychotic","public","publicize","puffy","pull","pull","pump","pump","pumped","punch","puncture","punish","punishment","puny","purchase","purple","purpose","purring","push","push","pushy","put","puzzled","puzzling",
										"quack","quaint","qualify","quarrelsome","quarter","quartz","queen","question","question","questionable","queue","quick","quickest","quicksand","quiet","quiet","quill","quilt","quince","quirky","quit","quiver","quixotic","quizzical",
										"rabbit","rabbits","rabid","race","racial","radiate","ragged","rail","railway","rain","rain","rainstorm","rainy","raise","rake","rambunctious","rampant","range","rank","rapid","rare","raspy","rat","rate","rate","ratty","ray","reach","reaction","read","reading","ready","real","realign","realize","reason","reason","rebel","receipt","receive","receptive","recess","recognize","recommend","reconcile","recondite","record","record","recruit","red","reduce","redundant","refer","reflect","reflective","refuse","regret","regret","regular","regulate","rehabilitate","reign","reinforce","reject","rejoice","relate","relation","relax","release","relieved","religion","rely","remain","remarkable","remember","remind","reminiscent","remove","render","reorganize","repair","repeat","replace","reply","report","represent","representative","reproduce","repulsive","request","request","rescue","research","resolute","resolve","resonant","respect","respond","responsible","rest","restored","restructure","retire","retrieve","return","review","revise","reward","rhetorical","rhyme","rhythm","rice","rich","rid","riddle","ride","rifle","right","righteous","rightful","rigid","ring","ring","rings","rinse","ripe","rise","risk","ritzy","river","road","roasted","rob","robin","robust","rock","rock","rod","roll","roll","romantic","roof","room","roomy","root","rose","rot","rotten","rough","round","route","royal","rub","rub","ruddy","rude","ruin","rule","rule","run","run","rural","rush","rustic","ruthless",
										"sable","sack","sack","sad","safe","sail","sail","salt","salty","same","sand","sassy","satisfy","satisfying","save","savory","saw","say","scale","scandalous","scarce","scare","scarecrow","scared","scarf","scary","scatter","scattered","scene","scent","schedule","school","science","scientific","scintillating","scissors","scold","scorch","scrape","scratch","scrawny","scream","screeching","screw","screw","scribble","scrub","sea","seal","search","seashore","seat","second","second-hand","secret","secretary","secretive","secure","sedate","see","seed","seek","seemly","select","selection","selective","self","selfish","sell","send","sense","sense","separate","separate","serious","servant","serve","service","set","settle","sew","shade","shade","shaggy","shake","shake","shaky","shallow","shame","shape","shape","share","sharp","shave","shear","shed","sheep","sheet","shelf","shelter","shine","shiny","ship","shirt","shiver","shivering","shock","shock","shocking","shoe","shoe","shoes","shoot","shop","shop","short","show","show","shrill","shrink","shrug","shut","shut","shy","sick","side","sidewalk","sigh","sign","sign","signal","silent","silent","silk","silky","silly","silver","simple","simplify","simplistic","sin","sincere","sing","sink","sink","sip","sister","sisters","sit","six","size","skate","sketch","ski","skillful","skin","skinny","skip","skirt","sky","slap","slave","slay","sleep","sleep","sleepy","sleet","slide","slim","slimy","sling","slink","slip","slip","slippery","slit","slope","sloppy","slow","slow","small","smart","smash","smash","smell","smell","smelly","smile","smile","smiling","smite","smoggy","smoke","smoke","smooth","snail","snails","snake","snakes","snatch","sneak","sneaky","sneeze","sneeze","sniff","snobbish","snore","snotty","snow","snow","soak","soap","society","sock","soda","sofa","soft","soggy","solid","solve","somber","son","song","songs","soothe","soothsay","sophisticated","sordid","sore","sore","sort","sort","sound","sound","soup","sour","sow","space","spade","spare","spark","spark","sparkle","sparkling","speak","special","specify","spectacular","speed","spell","spend","spicy","spiders","spiffy","spiky","spill","spin","spiritual","spit","spiteful","splendid","split","spoil","sponge","spooky","spoon","spot","spot","spotless","spotted","spotty","spray","spread","spring","spring","sprout","spurious","spy","squalid","square","square","squash","squeak","squeal","squealing","squeamish","squeeze","squirrel","stage","stain","staking","stale","stamp","stamp","stand","standing","star","stare","start","start","statement","station","statuesque","stay","steadfast","steady","steal","steam","steel","steep","steer","stem","step","step","stereotyped","stew","stick","stick","sticks","sticky","stiff","stimulate","stimulating","sting","stingy","stink","stir","stitch","stitch","stocking","stomach","stone","stop","stop","store","store","stormy","story","stove","straight","strange","stranger","strap","straw","stream","streamline","street","strengthen","stretch","stretch","stride","strike","string","string","strip","striped","strive","stroke","strong","structure","structure","study","stuff","stupendous","stupid","sturdy","subdued","sublet","subsequent","substance","substantial","subtract","succeed","successful","succinct","suck","sudden","suffer","sugar","suggest","suggestion","suit","suit","sulky","summarize","summer","sun","super","superb","superficial","supervise","supply","support","support","suppose","supreme","surprise","surprise","surround","suspect","suspend","swanky","swear","sweat","sweater","sweep","sweet","swell","sweltering","swift","swim","swim","swing","swing","switch","symbolize","symptomatic","synonymous","synthesize","system","systemize",
										"table","taboo","tabulate","tacit","tacky","tail","take","talented","talk","talk","tall","tame","tame","tan","tangible","tangy","tank","tap","target","tart","taste","taste","tasteful","tasteless","tasty","tawdry","tax","teach","teaching","team","tear","tearful","tease","tedious","teeny","teeny-tiny","teeth","telephone","tell","telling","temper","temporary","tempt","ten","tendency","tender","tense","tense","tent","tenuous","terrible","terrific","terrify","territory","test","test","tested","testy","texture","thank","thankful","thaw","theory","therapeutic","thick","thin","thing","things","think","thinkable","third","thirsty","thirsty","thought","thoughtful","thoughtless","thread","threatening","three","thrill","thrive","throat","throne","throw","thrust","thumb","thunder","thundering","tick","ticket","tickle","tidy","tie","tiger","tight","tightfisted","time","time","tin","tiny","tip","tire","tired","tiresome","title","toad","toe","toes","tomatoes","tongue","tooth","toothbrush","toothpaste","toothsome","top","torpid","touch","touch","tough","tour","tow","towering","town","toy","toys","trace","trade","trade","trail","train","train","trains","tramp","tranquil","transcribe","transfer","transform","translate","transport","transport","trap","trashy","travel","tray","tread","treat","treatment","tree","trees","tremble","tremendous","trick","trick","tricky","trip","trip","trite","trot","trouble","trouble","troubled","troubleshoot","trousers","truck","trucks","truculent","true","trust","truthful","try","tub","tug","tumble","turkey","turn","turn","tutor","twig","twist","twist","two","type","typical",
										"ubiquitous","ugliest","ugly","ultra","umbrella","unable","unaccountable","unadvised","unarmed","unbecoming","unbiased","uncle","uncovered","undergo","understand","understood","undertake","underwear","undesirable","undress","unequal","unequaled","uneven","unfasten","unhealthy","unify","uninterested","unique","unit","unite","unkempt","unknown","unlock","unnatural","unpack","unruly","unsightly","unsuitable","untidy","untidy","unused","unusual","unwieldy","unwritten","upbeat","update","upgrade","uphold","uppity","upset","upset","uptight","use","use","used","useful","useless","utilize","utopian","utter","uttermost",
										"vacation","vacuous","vagabond","vague","valuable","value","van","vanish","various","vase","vast","vegetable","veil","vein","vengeful","venomous","verbalize","verdant","verify","verse","versed","vessel","vest","vex","victorious","view","vigorous","violent","violet","visit","visitor","vivacious","voice","voiceless","volatile","volcano","volleyball","voracious","voyage","vulgar",
										"wacky","waggish","wail","wait","waiting","wake","wakeful","walk","walk","wall","wander","wandering","want","wanting","war","warlike","warm","warm","warn","wary","wash","wash","waste","waste","wasteful","watch","watch","water","water","watery","wave","wave","waves","wax","way","weak","wealth","wealthy","wear","weary","weather","weave","wed","week","weep","weigh","weight","welcome","well-groomed","well-made","well-off","well-to-do","wend","wet","wet","wheel","whimsical","whine","whip","whip","whirl","whisper","whispering","whistle","whistle","white","whole","wholesale","wicked","wide","wide-eyed","wiggly","wild","wilderness","willing","win","wind","wind","window","windy","wine","wing","wink","winter","wipe","wire","wiry","wise","wish","wish","wistful","withdraw","withhold","withstand","witty","wobble","woebegone","woman","womanly","women","wonder","wonderful","wood","wooden","wool","woozy","word","work","work","workable","worm","worried","worry","worthless","wound","wrap","wrathful","wreck","wren","wrench","wrestle","wretched","wriggle","wring","wrist","write","writer","writing","wrong","wry",
										"x-ray",
										"yak","yam","yard","yarn","yawn","year","yell","yellow","yielding","yoke","young","youthful","yummy",
										"zany","zealous","zebra","zephyr","zesty","zinc","zip","zipper","zippy","zonked","zoo","zoom"}
	scrtch = make([]string, 10)
	wordCount = len(words)
)

func randomWord() string {
	count := rand.Intn(3)+1
	for i := 0; i < count; i++ {
		scrtch[i] = words[rand.Intn(wordCount)]
	}
	return strings.Join(scrtch[:count], " ")
}

func randomPrefix() string {
	word := words[rand.Intn(wordCount)]
	return word[:rand.Intn(len(word))]
}