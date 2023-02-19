package voice

var dataMap map[VOICE][]byte

func init() {
	dataMap = map[VOICE][]byte{
		VOICE_A:       a,
		VOICE_A1:      a1,
		VOICE_AI1:     ai1,
		VOICE_AI2:     ai2,
		VOICE_AI3:     ai3,
		VOICE_AI4:     ai4,
		VOICE_AN1:     an1,
		VOICE_AN3:     an3,
		VOICE_AN4:     an4,
		VOICE_ANG1:    ang1,
		VOICE_ANG2:    ang2,
		VOICE_AO1:     ao1,
		VOICE_AO2:     ao2,
		VOICE_AO3:     ao3,
		VOICE_AO4:     ao4,
		VOICE_BA:      ba,
		VOICE_BA1:     ba1,
		VOICE_BA2:     ba2,
		VOICE_BA3:     ba3,
		VOICE_BA4:     ba4,
		VOICE_BAI1:    bai1,
		VOICE_BAI2:    bai2,
		VOICE_BAI3:    bai3,
		VOICE_BAI4:    bai4,
		VOICE_BAN1:    ban1,
		VOICE_BAN3:    ban3,
		VOICE_BAN4:    ban4,
		VOICE_BANG1:   bang1,
		VOICE_BANG3:   bang3,
		VOICE_BANG4:   bang4,
		VOICE_BAO1:    bao1,
		VOICE_BAO2:    bao2,
		VOICE_BAO3:    bao3,
		VOICE_BAO4:    bao4,
		VOICE_BEI1:    bei1,
		VOICE_BEI3:    bei3,
		VOICE_BEI4:    bei4,
		VOICE_BEN1:    ben1,
		VOICE_BEN3:    ben3,
		VOICE_BEN4:    ben4,
		VOICE_BENG1:   beng1,
		VOICE_BENG4:   beng4,
		VOICE_BI1:     bi1,
		VOICE_BI2:     bi2,
		VOICE_BI3:     bi3,
		VOICE_BI4:     bi4,
		VOICE_BIAN1:   bian1,
		VOICE_BIAN3:   bian3,
		VOICE_BIAN4:   bian4,
		VOICE_BIAO1:   biao1,
		VOICE_BIAO3:   biao3,
		VOICE_BIE1:    bie1,
		VOICE_BIE2:    bie2,
		VOICE_BIE3:    bie3,
		VOICE_BIN1:    bin1,
		VOICE_BIN4:    bin4,
		VOICE_BING1:   bing1,
		VOICE_BING3:   bing3,
		VOICE_BING4:   bing4,
		VOICE_BO:      bo,
		VOICE_BO1:     bo1,
		VOICE_BO2:     bo2,
		VOICE_BO3:     bo3,
		VOICE_BU3:     bu3,
		VOICE_BU4:     bu4,
		VOICE_CA1:     ca1,
		VOICE_CAI1:    cai1,
		VOICE_CAI2:    cai2,
		VOICE_CAI3:    cai3,
		VOICE_CAI4:    cai4,
		VOICE_CAN1:    can1,
		VOICE_CAN2:    can2,
		VOICE_CAN3:    can3,
		VOICE_CAN4:    can4,
		VOICE_CANG1:   cang1,
		VOICE_CANG2:   cang2,
		VOICE_CAO1:    cao1,
		VOICE_CAO2:    cao2,
		VOICE_CAO3:    cao3,
		VOICE_CE4:     ce4,
		VOICE_CENG2:   ceng2,
		VOICE_CENG4:   ceng4,
		VOICE_CHA1:    cha1,
		VOICE_CHA2:    cha2,
		VOICE_CHA4:    cha4,
		VOICE_CHAI1:   chai1,
		VOICE_CHAI2:   chai2,
		VOICE_CHAN1:   chan1,
		VOICE_CHAN2:   chan2,
		VOICE_CHAN3:   chan3,
		VOICE_CHAN4:   chan4,
		VOICE_CHANG1:  chang1,
		VOICE_CHANG2:  chang2,
		VOICE_CHANG3:  chang3,
		VOICE_CHANG4:  chang4,
		VOICE_CHAO1:   chao1,
		VOICE_CHAO2:   chao2,
		VOICE_CHAO3:   chao3,
		VOICE_CHE1:    che1,
		VOICE_CHE3:    che3,
		VOICE_CHE4:    che4,
		VOICE_CHEN2:   chen2,
		VOICE_CHEN4:   chen4,
		VOICE_CHENG1:  cheng1,
		VOICE_CHENG2:  cheng2,
		VOICE_CHENG3:  cheng3,
		VOICE_CHENG4:  cheng4,
		VOICE_CHI1:    chi1,
		VOICE_CHI2:    chi2,
		VOICE_CHI3:    chi3,
		VOICE_CHI4:    chi4,
		VOICE_CHONG1:  chong1,
		VOICE_CHONG2:  chong2,
		VOICE_CHONG3:  chong3,
		VOICE_CHOU1:   chou1,
		VOICE_CHOU2:   chou2,
		VOICE_CHOU3:   chou3,
		VOICE_CHOU4:   chou4,
		VOICE_CHU1:    chu1,
		VOICE_CHU2:    chu2,
		VOICE_CHU3:    chu3,
		VOICE_CHU4:    chu4,
		VOICE_CHUAI1:  chuai1,
		VOICE_CHUAN1:  chuan1,
		VOICE_CHUAN2:  chuan2,
		VOICE_CHUAN3:  chuan3,
		VOICE_CHUAN4:  chuan4,
		VOICE_CHUANG1: chuang1,
		VOICE_CHUANG2: chuang2,
		VOICE_CHUANG3: chuang3,
		VOICE_CHUANG4: chuang4,
		VOICE_CHUI1:   chui1,
		VOICE_CHUI2:   chui2,
		VOICE_CHUN1:   chun1,
		VOICE_CHUN2:   chun2,
		VOICE_CHUN3:   chun3,
		VOICE_CHUO1:   chuo1,
		VOICE_CHUO4:   chuo4,
		VOICE_CI2:     ci2,
		VOICE_CI3:     ci3,
		VOICE_CI4:     ci4,
		VOICE_CONG1:   cong1,
		VOICE_CONG2:   cong2,
		VOICE_COU4:    cou4,
		VOICE_CU1:     cu1,
		VOICE_CU4:     cu4,
		VOICE_CUAN4:   cuan4,
		VOICE_CUI1:    cui1,
		VOICE_CUI4:    cui4,
		VOICE_CUN1:    cun1,
		VOICE_CUN2:    cun2,
		VOICE_CUN4:    cun4,
		VOICE_CUO1:    cuo1,
		VOICE_CUO4:    cuo4,
		VOICE_DA:      da,
		VOICE_DA1:     da1,
		VOICE_DA2:     da2,
		VOICE_DA3:     da3,
		VOICE_DA4:     da4,
		VOICE_DAI1:    dai1,
		VOICE_DAI3:    dai3,
		VOICE_DAI4:    dai4,
		VOICE_DAN1:    dan1,
		VOICE_DAN3:    dan3,
		VOICE_DAN4:    dan4,
		VOICE_DANG1:   dang1,
		VOICE_DANG3:   dang3,
		VOICE_DANG4:   dang4,
		VOICE_DAO1:    dao1,
		VOICE_DAO3:    dao3,
		VOICE_DAO4:    dao4,
		VOICE_DE:      de,
		VOICE_DE2:     de2,
		VOICE_DENG1:   deng1,
		VOICE_DENG3:   deng3,
		VOICE_DENG4:   deng4,
		VOICE_DI1:     di1,
		VOICE_DI2:     di2,
		VOICE_DI3:     di3,
		VOICE_DI4:     di4,
		VOICE_DIAN1:   dian1,
		VOICE_DIAN3:   dian3,
		VOICE_DIAN4:   dian4,
		VOICE_DIAO1:   diao1,
		VOICE_DIAO4:   diao4,
		VOICE_DIE1:    die1,
		VOICE_DIE2:    die2,
		VOICE_DING1:   ding1,
		VOICE_DING3:   ding3,
		VOICE_DING4:   ding4,
		VOICE_DIU1:    diu1,
		VOICE_DONG1:   dong1,
		VOICE_DONG3:   dong3,
		VOICE_DONG4:   dong4,
		VOICE_DOU1:    dou1,
		VOICE_DOU3:    dou3,
		VOICE_DOU4:    dou4,
		VOICE_DU1:     du1,
		VOICE_DU2:     du2,
		VOICE_DU3:     du3,
		VOICE_DU4:     du4,
		VOICE_DUAN1:   duan1,
		VOICE_DUAN3:   duan3,
		VOICE_DUAN4:   duan4,
		VOICE_DUI1:    dui1,
		VOICE_DUI4:    dui4,
		VOICE_DUN1:    dun1,
		VOICE_DUN3:    dun3,
		VOICE_DUN4:    dun4,
		VOICE_DUO1:    duo1,
		VOICE_DUO2:    duo2,
		VOICE_DUO3:    duo3,
		VOICE_DUO4:    duo4,
		VOICE_E2:      e2,
		VOICE_E4:      e4,
		VOICE_EN1:     en1,
		VOICE_ER2:     er2,
		VOICE_ER3:     er3,
		VOICE_ER4:     er4,
		VOICE_FA1:     fa1,
		VOICE_FA2:     fa2,
		VOICE_FA3:     fa3,
		VOICE_FAN1:    fan1,
		VOICE_FAN2:    fan2,
		VOICE_FAN3:    fan3,
		VOICE_FAN4:    fan4,
		VOICE_FANG1:   fang1,
		VOICE_FANG2:   fang2,
		VOICE_FANG3:   fang3,
		VOICE_FANG4:   fang4,
		VOICE_FEI1:    fei1,
		VOICE_FEI2:    fei2,
		VOICE_FEI3:    fei3,
		VOICE_FEI4:    fei4,
		VOICE_FEN1:    fen1,
		VOICE_FEN2:    fen2,
		VOICE_FEN3:    fen3,
		VOICE_FEN4:    fen4,
		VOICE_FENG1:   feng1,
		VOICE_FENG2:   feng2,
		VOICE_FENG3:   feng3,
		VOICE_FENG4:   feng4,
		VOICE_FOU3:    fou3,
		VOICE_FU1:     fu1,
		VOICE_FU2:     fu2,
		VOICE_FU3:     fu3,
		VOICE_FU4:     fu4,
		VOICE_GA4:     ga4,
		VOICE_GAI1:    gai1,
		VOICE_GAI3:    gai3,
		VOICE_GAI4:    gai4,
		VOICE_GAN1:    gan1,
		VOICE_GAN3:    gan3,
		VOICE_GAN4:    gan4,
		VOICE_GANG1:   gang1,
		VOICE_GANG3:   gang3,
		VOICE_GAO1:    gao1,
		VOICE_GAO3:    gao3,
		VOICE_GAO4:    gao4,
		VOICE_GE1:     ge1,
		VOICE_GE2:     ge2,
		VOICE_GE4:     ge4,
		VOICE_GEI3:    gei3,
		VOICE_GEN1:    gen1,
		VOICE_GENG1:   geng1,
		VOICE_GENG3:   geng3,
		VOICE_GENG4:   geng4,
		VOICE_GONG1:   gong1,
		VOICE_GONG3:   gong3,
		VOICE_GONG4:   gong4,
		VOICE_GOU1:    gou1,
		VOICE_GOU3:    gou3,
		VOICE_GOU4:    gou4,
		VOICE_GU1:     gu1,
		VOICE_GU3:     gu3,
		VOICE_GU4:     gu4,
		VOICE_GUA1:    gua1,
		VOICE_GUA3:    gua3,
		VOICE_GUA4:    gua4,
		VOICE_GUAI1:   guai1,
		VOICE_GUAI3:   guai3,
		VOICE_GUAI4:   guai4,
		VOICE_GUAN1:   guan1,
		VOICE_GUAN3:   guan3,
		VOICE_GUAN4:   guan4,
		VOICE_GUANG1:  guang1,
		VOICE_GUANG3:  guang3,
		VOICE_GUANG4:  guang4,
		VOICE_GUI1:    gui1,
		VOICE_GUI3:    gui3,
		VOICE_GUI4:    gui4,
		VOICE_GUN3:    gun3,
		VOICE_GUN4:    gun4,
		VOICE_GUO1:    guo1,
		VOICE_GUO2:    guo2,
		VOICE_GUO3:    guo3,
		VOICE_GUO4:    guo4,
		VOICE_HA1:     ha1,
		VOICE_HA2:     ha2,
		VOICE_HAI1:    hai1,
		VOICE_HAI2:    hai2,
		VOICE_HAI3:    hai3,
		VOICE_HAI4:    hai4,
		VOICE_HAN1:    han1,
		VOICE_HAN2:    han2,
		VOICE_HAN3:    han3,
		VOICE_HAN4:    han4,
		VOICE_HANG1:   hang1,
		VOICE_HANG2:   hang2,
		VOICE_HAO2:    hao2,
		VOICE_HAO3:    hao3,
		VOICE_HAO4:    hao4,
		VOICE_HE1:     he1,
		VOICE_HE2:     he2,
		VOICE_HE4:     he4,
		VOICE_HEI1:    hei1,
		VOICE_HEN2:    hen2,
		VOICE_HEN3:    hen3,
		VOICE_HEN4:    hen4,
		VOICE_HENG1:   heng1,
		VOICE_HENG2:   heng2,
		VOICE_HONG1:   hong1,
		VOICE_HONG2:   hong2,
		VOICE_HOU2:    hou2,
		VOICE_HOU3:    hou3,
		VOICE_HOU4:    hou4,
		VOICE_HU1:     hu1,
		VOICE_HU2:     hu2,
		VOICE_HU3:     hu3,
		VOICE_HU4:     hu4,
		VOICE_HUA1:    hua1,
		VOICE_HUA2:    hua2,
		VOICE_HUA4:    hua4,
		VOICE_HUAI2:   huai2,
		VOICE_HUAI4:   huai4,
		VOICE_HUAN1:   huan1,
		VOICE_HUAN2:   huan2,
		VOICE_HUAN3:   huan3,
		VOICE_HUAN4:   huan4,
		VOICE_HUANG1:  huang1,
		VOICE_HUANG2:  huang2,
		VOICE_HUANG3:  huang3,
		VOICE_HUI1:    hui1,
		VOICE_HUI2:    hui2,
		VOICE_HUI3:    hui3,
		VOICE_HUI4:    hui4,
		VOICE_HUN1:    hun1,
		VOICE_HUN2:    hun2,
		VOICE_HUN4:    hun4,
		VOICE_HUO1:    huo1,
		VOICE_HUO2:    huo2,
		VOICE_HUO3:    huo3,
		VOICE_HUO4:    huo4,
		VOICE_JI1:     ji1,
		VOICE_JI2:     ji2,
		VOICE_JI3:     ji3,
		VOICE_JI4:     ji4,
		VOICE_JIA1:    jia1,
		VOICE_JIA2:    jia2,
		VOICE_JIA3:    jia3,
		VOICE_JIA4:    jia4,
		VOICE_JIAN1:   jian1,
		VOICE_JIAN3:   jian3,
		VOICE_JIAN4:   jian4,
		VOICE_JIANG1:  jiang1,
		VOICE_JIANG3:  jiang3,
		VOICE_JIANG4:  jiang4,
		VOICE_JIAO1:   jiao1,
		VOICE_JIAO3:   jiao3,
		VOICE_JIAO4:   jiao4,
		VOICE_JIE1:    jie1,
		VOICE_JIE2:    jie2,
		VOICE_JIE3:    jie3,
		VOICE_JIE4:    jie4,
		VOICE_JIN1:    jin1,
		VOICE_JIN3:    jin3,
		VOICE_JIN4:    jin4,
		VOICE_JING1:   jing1,
		VOICE_JING3:   jing3,
		VOICE_JING4:   jing4,
		VOICE_JIONG3:  jiong3,
		VOICE_JIU1:    jiu1,
		VOICE_JIU3:    jiu3,
		VOICE_JIU4:    jiu4,
		VOICE_JU1:     ju1,
		VOICE_JU2:     ju2,
		VOICE_JU3:     ju3,
		VOICE_JU4:     ju4,
		VOICE_JUAN1:   juan1,
		VOICE_JUAN3:   juan3,
		VOICE_JUAN4:   juan4,
		VOICE_JUE2:    jue2,
		VOICE_JUN1:    jun1,
		VOICE_JUN4:    jun4,
		VOICE_KA1:     ka1,
		VOICE_KA3:     ka3,
		VOICE_KAI1:    kai1,
		VOICE_KAI3:    kai3,
		VOICE_KAN1:    kan1,
		VOICE_KAN3:    kan3,
		VOICE_KAN4:    kan4,
		VOICE_KANG1:   kang1,
		VOICE_KANG2:   kang2,
		VOICE_KANG4:   kang4,
		VOICE_KAO3:    kao3,
		VOICE_KAO4:    kao4,
		VOICE_KE1:     ke1,
		VOICE_KE2:     ke2,
		VOICE_KE3:     ke3,
		VOICE_KE4:     ke4,
		VOICE_KEN3:    ken3,
		VOICE_KENG1:   keng1,
		VOICE_KONG1:   kong1,
		VOICE_KONG3:   kong3,
		VOICE_KONG4:   kong4,
		VOICE_KOU1:    kou1,
		VOICE_KOU3:    kou3,
		VOICE_KOU4:    kou4,
		VOICE_KU1:     ku1,
		VOICE_KU3:     ku3,
		VOICE_KU4:     ku4,
		VOICE_KUA1:    kua1,
		VOICE_KUA3:    kua3,
		VOICE_KUA4:    kua4,
		VOICE_KUAI4:   kuai4,
		VOICE_KUAN1:   kuan1,
		VOICE_KUAN3:   kuan3,
		VOICE_KUANG1:  kuang1,
		VOICE_KUANG2:  kuang2,
		VOICE_KUANG4:  kuang4,
		VOICE_KUI1:    kui1,
		VOICE_KUI2:    kui2,
		VOICE_KUI4:    kui4,
		VOICE_KUN1:    kun1,
		VOICE_KUN3:    kun3,
		VOICE_KUN4:    kun4,
		VOICE_KUO4:    kuo4,
		VOICE_LA:      la,
		VOICE_LA1:     la1,
		VOICE_LA3:     la3,
		VOICE_LA4:     la4,
		VOICE_LAI2:    lai2,
		VOICE_LAI4:    lai4,
		VOICE_LAN2:    lan2,
		VOICE_LAN3:    lan3,
		VOICE_LAN4:    lan4,
		VOICE_LANG2:   lang2,
		VOICE_LANG3:   lang3,
		VOICE_LANG4:   lang4,
		VOICE_LAO1:    lao1,
		VOICE_LAO2:    lao2,
		VOICE_LAO3:    lao3,
		VOICE_LAO4:    lao4,
		VOICE_LE:      le,
		VOICE_LE1:     le1,
		VOICE_LE4:     le4,
		VOICE_LEI1:    lei1,
		VOICE_LEI2:    lei2,
		VOICE_LEI3:    lei3,
		VOICE_LEI4:    lei4,
		VOICE_LENG2:   leng2,
		VOICE_LENG3:   leng3,
		VOICE_LENG4:   leng4,
		VOICE_LI1:     li1,
		VOICE_LI2:     li2,
		VOICE_LI3:     li3,
		VOICE_LI4:     li4,
		VOICE_LIA3:    lia3,
		VOICE_LIAN2:   lian2,
		VOICE_LIAN3:   lian3,
		VOICE_LIAN4:   lian4,
		VOICE_LIANG2:  liang2,
		VOICE_LIANG3:  liang3,
		VOICE_LIANG4:  liang4,
		VOICE_LIAO1:   liao1,
		VOICE_LIAO2:   liao2,
		VOICE_LIAO3:   liao3,
		VOICE_LIAO4:   liao4,
		VOICE_LIE3:    lie3,
		VOICE_LIE4:    lie4,
		VOICE_LIN1:    lin1,
		VOICE_LIN2:    lin2,
		VOICE_LIN3:    lin3,
		VOICE_LIN4:    lin4,
		VOICE_LING2:   ling2,
		VOICE_LING3:   ling3,
		VOICE_LING4:   ling4,
		VOICE_LIU1:    liu1,
		VOICE_LIU2:    liu2,
		VOICE_LIU3:    liu3,
		VOICE_LIU4:    liu4,
		VOICE_LONG2:   long2,
		VOICE_LONG3:   long3,
		VOICE_LOU2:    lou2,
		VOICE_LOU3:    lou3,
		VOICE_LOU4:    lou4,
		VOICE_LU2:     lu2,
		VOICE_LU3:     lu3,
		VOICE_LU4:     lu4,
		VOICE_LUAN2:   luan2,
		VOICE_LUAN3:   luan3,
		VOICE_LUAN4:   luan4,
		VOICE_LUN1:    lun1,
		VOICE_LUN2:    lun2,
		VOICE_LUN4:    lun4,
		VOICE_LUO1:    luo1,
		VOICE_LUO2:    luo2,
		VOICE_LUO3:    luo3,
		VOICE_LUO4:    luo4,
		VOICE_LV2:     lv2,
		VOICE_LV3:     lv3,
		VOICE_LV4:     lv4,
		VOICE_LVE4:    lve4,
		VOICE_MA:      ma,
		VOICE_MA1:     ma1,
		VOICE_MA2:     ma2,
		VOICE_MA3:     ma3,
		VOICE_MA4:     ma4,
		VOICE_MAI2:    mai2,
		VOICE_MAI3:    mai3,
		VOICE_MAI4:    mai4,
		VOICE_MAN2:    man2,
		VOICE_MAN3:    man3,
		VOICE_MAN4:    man4,
		VOICE_MANG2:   mang2,
		VOICE_MANG3:   mang3,
		VOICE_MAO1:    mao1,
		VOICE_MAO2:    mao2,
		VOICE_MAO3:    mao3,
		VOICE_MAO4:    mao4,
		VOICE_ME:      me,
		VOICE_MEI2:    mei2,
		VOICE_MEI3:    mei3,
		VOICE_MEI4:    mei4,
		VOICE_MEN:     men,
		VOICE_MEN2:    men2,
		VOICE_MEN4:    men4,
		VOICE_MENG2:   meng2,
		VOICE_MENG3:   meng3,
		VOICE_MENG4:   meng4,
		VOICE_MI1:     mi1,
		VOICE_MI2:     mi2,
		VOICE_MI3:     mi3,
		VOICE_MI4:     mi4,
		VOICE_MIAN2:   mian2,
		VOICE_MIAN3:   mian3,
		VOICE_MIAN4:   mian4,
		VOICE_MIAO2:   miao2,
		VOICE_MIAO3:   miao3,
		VOICE_MIAO4:   miao4,
		VOICE_MIE4:    mie4,
		VOICE_MIN2:    min2,
		VOICE_MIN3:    min3,
		VOICE_MING2:   ming2,
		VOICE_MING4:   ming4,
		VOICE_MIU4:    miu4,
		VOICE_MO1:     mo1,
		VOICE_MO2:     mo2,
		VOICE_MO3:     mo3,
		VOICE_MO4:     mo4,
		VOICE_MOU2:    mou2,
		VOICE_MOU3:    mou3,
		VOICE_MU3:     mu3,
		VOICE_MU4:     mu4,
		VOICE_NA2:     na2,
		VOICE_NA3:     na3,
		VOICE_NA4:     na4,
		VOICE_NAI3:    nai3,
		VOICE_NAI4:    nai4,
		VOICE_NAN2:    nan2,
		VOICE_NANG2:   nang2,
		VOICE_NAO2:    nao2,
		VOICE_NAO3:    nao3,
		VOICE_NAO4:    nao4,
		VOICE_NE:      ne,
		VOICE_NEI3:    nei3,
		VOICE_NEI4:    nei4,
		VOICE_NEN4:    nen4,
		VOICE_NENG2:   neng2,
		VOICE_NI1:     ni1,
		VOICE_NI2:     ni2,
		VOICE_NI3:     ni3,
		VOICE_NI4:     ni4,
		VOICE_NIAN2:   nian2,
		VOICE_NIAN3:   nian3,
		VOICE_NIAN4:   nian4,
		VOICE_NIANG2:  niang2,
		VOICE_NIANG4:  niang4,
		VOICE_NIAO3:   niao3,
		VOICE_NIAO4:   niao4,
		VOICE_NIE1:    nie1,
		VOICE_NIE4:    nie4,
		VOICE_NIN2:    nin2,
		VOICE_NING2:   ning2,
		VOICE_NING4:   ning4,
		VOICE_NIU2:    niu2,
		VOICE_NIU3:    niu3,
		VOICE_NONG2:   nong2,
		VOICE_NONG4:   nong4,
		VOICE_NU2:     nu2,
		VOICE_NU3:     nu3,
		VOICE_NU4:     nu4,
		VOICE_NUAN3:   nuan3,
		VOICE_NUO2:    nuo2,
		VOICE_NUO4:    nuo4,
		VOICE_NV3:     nv3,
		VOICE_NVE4:    nve4,
		VOICE_O2:      o2,
		VOICE_OU1:     ou1,
		VOICE_OU3:     ou3,
		VOICE_PA1:     pa1,
		VOICE_PA2:     pa2,
		VOICE_PA4:     pa4,
		VOICE_PAI1:    pai1,
		VOICE_PAI2:    pai2,
		VOICE_PAI4:    pai4,
		VOICE_PAN1:    pan1,
		VOICE_PAN2:    pan2,
		VOICE_PAN4:    pan4,
		VOICE_PANG1:   pang1,
		VOICE_PANG2:   pang2,
		VOICE_PANG4:   pang4,
		VOICE_PAO1:    pao1,
		VOICE_PAO2:    pao2,
		VOICE_PAO3:    pao3,
		VOICE_PAO4:    pao4,
		VOICE_PEI1:    pei1,
		VOICE_PEI2:    pei2,
		VOICE_PEI4:    pei4,
		VOICE_PEN1:    pen1,
		VOICE_PEN2:    pen2,
		VOICE_PENG1:   peng1,
		VOICE_PENG2:   peng2,
		VOICE_PENG3:   peng3,
		VOICE_PENG4:   peng4,
		VOICE_PI1:     pi1,
		VOICE_PI2:     pi2,
		VOICE_PI3:     pi3,
		VOICE_PI4:     pi4,
		VOICE_PIAN1:   pian1,
		VOICE_PIAN4:   pian4,
		VOICE_PIAO1:   piao1,
		VOICE_PIAO2:   piao2,
		VOICE_PIAO4:   piao4,
		VOICE_PIE1:    pie1,
		VOICE_PIN1:    pin1,
		VOICE_PIN2:    pin2,
		VOICE_PIN3:    pin3,
		VOICE_PIN4:    pin4,
		VOICE_PING1:   ping1,
		VOICE_PING2:   ping2,
		VOICE_PO1:     po1,
		VOICE_PO2:     po2,
		VOICE_PO3:     po3,
		VOICE_PO4:     po4,
		VOICE_POU1:    pou1,
		VOICE_PU1:     pu1,
		VOICE_PU2:     pu2,
		VOICE_PU3:     pu3,
		VOICE_PU4:     pu4,
		VOICE_QI1:     qi1,
		VOICE_QI2:     qi2,
		VOICE_QI3:     qi3,
		VOICE_QI4:     qi4,
		VOICE_QIA1:    qia1,
		VOICE_QIA4:    qia4,
		VOICE_QIAN1:   qian1,
		VOICE_QIAN2:   qian2,
		VOICE_QIAN3:   qian3,
		VOICE_QIAN4:   qian4,
		VOICE_QIANG1:  qiang1,
		VOICE_QIANG2:  qiang2,
		VOICE_QIANG3:  qiang3,
		VOICE_QIAO1:   qiao1,
		VOICE_QIAO2:   qiao2,
		VOICE_QIAO3:   qiao3,
		VOICE_QIAO4:   qiao4,
		VOICE_QIE3:    qie3,
		VOICE_QIE4:    qie4,
		VOICE_QIN1:    qin1,
		VOICE_QIN2:    qin2,
		VOICE_QIN3:    qin3,
		VOICE_QIN4:    qin4,
		VOICE_QING1:   qing1,
		VOICE_QING2:   qing2,
		VOICE_QING3:   qing3,
		VOICE_QING4:   qing4,
		VOICE_QIONG2:  qiong2,
		VOICE_QIU1:    qiu1,
		VOICE_QIU2:    qiu2,
		VOICE_QU1:     qu1,
		VOICE_QU2:     qu2,
		VOICE_QU3:     qu3,
		VOICE_QU4:     qu4,
		VOICE_QUAN1:   quan1,
		VOICE_QUAN2:   quan2,
		VOICE_QUAN3:   quan3,
		VOICE_QUAN4:   quan4,
		VOICE_QUE1:    que1,
		VOICE_QUE2:    que2,
		VOICE_QUE4:    que4,
		VOICE_QUN2:    qun2,
		VOICE_RAN2:    ran2,
		VOICE_RAN3:    ran3,
		VOICE_RANG2:   rang2,
		VOICE_RANG3:   rang3,
		VOICE_RANG4:   rang4,
		VOICE_RAO2:    rao2,
		VOICE_RAO3:    rao3,
		VOICE_RAO4:    rao4,
		VOICE_RE3:     re3,
		VOICE_RE4:     re4,
		VOICE_REN2:    ren2,
		VOICE_REN3:    ren3,
		VOICE_REN4:    ren4,
		VOICE_RENG1:   reng1,
		VOICE_RENG2:   reng2,
		VOICE_RI4:     ri4,
		VOICE_RONG1:   rong1,
		VOICE_RONG2:   rong2,
		VOICE_RONG3:   rong3,
		VOICE_ROU2:    rou2,
		VOICE_ROU4:    rou4,
		VOICE_RU2:     ru2,
		VOICE_RU3:     ru3,
		VOICE_RU4:     ru4,
		VOICE_RUAN3:   ruan3,
		VOICE_RUI3:    rui3,
		VOICE_RUI4:    rui4,
		VOICE_RUN4:    run4,
		VOICE_RUO4:    ruo4,
		VOICE_SA1:     sa1,
		VOICE_SA3:     sa3,
		VOICE_SA4:     sa4,
		VOICE_SAI1:    sai1,
		VOICE_SAI4:    sai4,
		VOICE_SAN1:    san1,
		VOICE_SAN3:    san3,
		VOICE_SAN4:    san4,
		VOICE_SANG1:   sang1,
		VOICE_SANG3:   sang3,
		VOICE_SANG4:   sang4,
		VOICE_SAO1:    sao1,
		VOICE_SAO3:    sao3,
		VOICE_SE4:     se4,
		VOICE_SEN1:    sen1,
		VOICE_SENG1:   seng1,
		VOICE_SHA1:    sha1,
		VOICE_SHA2:    sha2,
		VOICE_SHA3:    sha3,
		VOICE_SHA4:    sha4,
		VOICE_SHAI1:   shai1,
		VOICE_SHAI4:   shai4,
		VOICE_SHAN1:   shan1,
		VOICE_SHAN3:   shan3,
		VOICE_SHAN4:   shan4,
		VOICE_SHANG:   shang,
		VOICE_SHANG1:  shang1,
		VOICE_SHANG3:  shang3,
		VOICE_SHANG4:  shang4,
		VOICE_SHAO1:   shao1,
		VOICE_SHAO2:   shao2,
		VOICE_SHAO3:   shao3,
		VOICE_SHAO4:   shao4,
		VOICE_SHE1:    she1,
		VOICE_SHE2:    she2,
		VOICE_SHE3:    she3,
		VOICE_SHE4:    she4,
		VOICE_SHEI2:   shei2,
		VOICE_SHEN1:   shen1,
		VOICE_SHEN2:   shen2,
		VOICE_SHEN3:   shen3,
		VOICE_SHEN4:   shen4,
		VOICE_SHENG1:  sheng1,
		VOICE_SHENG2:  sheng2,
		VOICE_SHENG3:  sheng3,
		VOICE_SHENG4:  sheng4,
		VOICE_SHI:     shi,
		VOICE_SHI1:    shi1,
		VOICE_SHI2:    shi2,
		VOICE_SHI3:    shi3,
		VOICE_SHI4:    shi4,
		VOICE_SHOU1:   shou1,
		VOICE_SHOU3:   shou3,
		VOICE_SHOU4:   shou4,
		VOICE_SHU1:    shu1,
		VOICE_SHU2:    shu2,
		VOICE_SHU3:    shu3,
		VOICE_SHU4:    shu4,
		VOICE_SHUA1:   shua1,
		VOICE_SHUA3:   shua3,
		VOICE_SHUAI1:  shuai1,
		VOICE_SHUAI3:  shuai3,
		VOICE_SHUAI4:  shuai4,
		VOICE_SHUAN1:  shuan1,
		VOICE_SHUAN4:  shuan4,
		VOICE_SHUANG1: shuang1,
		VOICE_SHUANG3: shuang3,
		VOICE_SHUI3:   shui3,
		VOICE_SHUI4:   shui4,
		VOICE_SHUN3:   shun3,
		VOICE_SHUN4:   shun4,
		VOICE_SHUO1:   shuo1,
		VOICE_SHUO4:   shuo4,
		VOICE_SI1:     si1,
		VOICE_SI3:     si3,
		VOICE_SI4:     si4,
		VOICE_SONG1:   song1,
		VOICE_SONG3:   song3,
		VOICE_SONG4:   song4,
		VOICE_SOU1:    sou1,
		VOICE_SOU4:    sou4,
		VOICE_SU1:     su1,
		VOICE_SU2:     su2,
		VOICE_SU4:     su4,
		VOICE_SUAN1:   suan1,
		VOICE_SUAN4:   suan4,
		VOICE_SUI1:    sui1,
		VOICE_SUI2:    sui2,
		VOICE_SUI3:    sui3,
		VOICE_SUI4:    sui4,
		VOICE_SUN1:    sun1,
		VOICE_SUN3:    sun3,
		VOICE_SUO:     suo,
		VOICE_SUO1:    suo1,
		VOICE_SUO3:    suo3,
		VOICE_TA1:     ta1,
		VOICE_TA3:     ta3,
		VOICE_TA4:     ta4,
		VOICE_TAI1:    tai1,
		VOICE_TAI2:    tai2,
		VOICE_TAI4:    tai4,
		VOICE_TAN1:    tan1,
		VOICE_TAN2:    tan2,
		VOICE_TAN3:    tan3,
		VOICE_TAN4:    tan4,
		VOICE_TANG1:   tang1,
		VOICE_TANG2:   tang2,
		VOICE_TANG3:   tang3,
		VOICE_TANG4:   tang4,
		VOICE_TAO1:    tao1,
		VOICE_TAO2:    tao2,
		VOICE_TAO3:    tao3,
		VOICE_TAO4:    tao4,
		VOICE_TE4:     te4,
		VOICE_TENG2:   teng2,
		VOICE_TI1:     ti1,
		VOICE_TI2:     ti2,
		VOICE_TI3:     ti3,
		VOICE_TI4:     ti4,
		VOICE_TIAN1:   tian1,
		VOICE_TIAN2:   tian2,
		VOICE_TIAN3:   tian3,
		VOICE_TIAO1:   tiao1,
		VOICE_TIAO2:   tiao2,
		VOICE_TIAO4:   tiao4,
		VOICE_TIE1:    tie1,
		VOICE_TIE3:    tie3,
		VOICE_TING1:   ting1,
		VOICE_TING2:   ting2,
		VOICE_TING3:   ting3,
		VOICE_TONG1:   tong1,
		VOICE_TONG2:   tong2,
		VOICE_TONG3:   tong3,
		VOICE_TONG4:   tong4,
		VOICE_TOU1:    tou1,
		VOICE_TOU2:    tou2,
		VOICE_TOU4:    tou4,
		VOICE_TU1:     tu1,
		VOICE_TU2:     tu2,
		VOICE_TU3:     tu3,
		VOICE_TU4:     tu4,
		VOICE_TUAN2:   tuan2,
		VOICE_TUI1:    tui1,
		VOICE_TUI2:    tui2,
		VOICE_TUI3:    tui3,
		VOICE_TUI4:    tui4,
		VOICE_TUN1:    tun1,
		VOICE_TUN2:    tun2,
		VOICE_TUO1:    tuo1,
		VOICE_TUO2:    tuo2,
		VOICE_TUO3:    tuo3,
		VOICE_TUO4:    tuo4,
		VOICE_WA:      wa,
		VOICE_WA1:     wa1,
		VOICE_WA2:     wa2,
		VOICE_WA3:     wa3,
		VOICE_WA4:     wa4,
		VOICE_WAI1:    wai1,
		VOICE_WAI4:    wai4,
		VOICE_WAN1:    wan1,
		VOICE_WAN2:    wan2,
		VOICE_WAN3:    wan3,
		VOICE_WAN4:    wan4,
		VOICE_WANG1:   wang1,
		VOICE_WANG2:   wang2,
		VOICE_WANG3:   wang3,
		VOICE_WANG4:   wang4,
		VOICE_WEI1:    wei1,
		VOICE_WEI2:    wei2,
		VOICE_WEI3:    wei3,
		VOICE_WEI4:    wei4,
		VOICE_WEN1:    wen1,
		VOICE_WEN2:    wen2,
		VOICE_WEN3:    wen3,
		VOICE_WEN4:    wen4,
		VOICE_WENG1:   weng1,
		VOICE_WO1:     wo1,
		VOICE_WO3:     wo3,
		VOICE_WO4:     wo4,
		VOICE_WU1:     wu1,
		VOICE_WU2:     wu2,
		VOICE_WU3:     wu3,
		VOICE_WU4:     wu4,
		VOICE_XI1:     xi1,
		VOICE_XI2:     xi2,
		VOICE_XI3:     xi3,
		VOICE_XI4:     xi4,
		VOICE_XIA1:    xia1,
		VOICE_XIA2:    xia2,
		VOICE_XIA4:    xia4,
		VOICE_XIAN1:   xian1,
		VOICE_XIAN2:   xian2,
		VOICE_XIAN3:   xian3,
		VOICE_XIAN4:   xian4,
		VOICE_XIANG1:  xiang1,
		VOICE_XIANG2:  xiang2,
		VOICE_XIANG3:  xiang3,
		VOICE_XIANG4:  xiang4,
		VOICE_XIAO1:   xiao1,
		VOICE_XIAO2:   xiao2,
		VOICE_XIAO3:   xiao3,
		VOICE_XIAO4:   xiao4,
		VOICE_XIE1:    xie1,
		VOICE_XIE2:    xie2,
		VOICE_XIE3:    xie3,
		VOICE_XIE4:    xie4,
		VOICE_XIN1:    xin1,
		VOICE_XIN4:    xin4,
		VOICE_XING1:   xing1,
		VOICE_XING2:   xing2,
		VOICE_XING3:   xing3,
		VOICE_XING4:   xing4,
		VOICE_XIONG1:  xiong1,
		VOICE_XIONG2:  xiong2,
		VOICE_XIU1:    xiu1,
		VOICE_XIU3:    xiu3,
		VOICE_XIU4:    xiu4,
		VOICE_XU1:     xu1,
		VOICE_XU2:     xu2,
		VOICE_XU3:     xu3,
		VOICE_XU4:     xu4,
		VOICE_XUAN1:   xuan1,
		VOICE_XUAN2:   xuan2,
		VOICE_XUAN3:   xuan3,
		VOICE_XUAN4:   xuan4,
		VOICE_XUE1:    xue1,
		VOICE_XUE2:    xue2,
		VOICE_XUE3:    xue3,
		VOICE_XUE4:    xue4,
		VOICE_XUN1:    xun1,
		VOICE_XUN2:    xun2,
		VOICE_XUN4:    xun4,
		VOICE_YA:      ya,
		VOICE_YA1:     ya1,
		VOICE_YA2:     ya2,
		VOICE_YA3:     ya3,
		VOICE_YA4:     ya4,
		VOICE_YAN1:    yan1,
		VOICE_YAN2:    yan2,
		VOICE_YAN3:    yan3,
		VOICE_YAN4:    yan4,
		VOICE_YANG1:   yang1,
		VOICE_YANG2:   yang2,
		VOICE_YANG3:   yang3,
		VOICE_YANG4:   yang4,
		VOICE_YAO1:    yao1,
		VOICE_YAO2:    yao2,
		VOICE_YAO3:    yao3,
		VOICE_YAO4:    yao4,
		VOICE_YE1:     ye1,
		VOICE_YE2:     ye2,
		VOICE_YE3:     ye3,
		VOICE_YE4:     ye4,
		VOICE_YI1:     yi1,
		VOICE_YI2:     yi2,
		VOICE_YI3:     yi3,
		VOICE_YI4:     yi4,
		VOICE_YIN1:    yin1,
		VOICE_YIN2:    yin2,
		VOICE_YIN3:    yin3,
		VOICE_YIN4:    yin4,
		VOICE_YING1:   ying1,
		VOICE_YING2:   ying2,
		VOICE_YING3:   ying3,
		VOICE_YING4:   ying4,
		VOICE_YO1:     yo1,
		VOICE_YONG1:   yong1,
		VOICE_YONG3:   yong3,
		VOICE_YONG4:   yong4,
		VOICE_YOU1:    you1,
		VOICE_YOU2:    you2,
		VOICE_YOU3:    you3,
		VOICE_YOU4:    you4,
		VOICE_YU1:     yu1,
		VOICE_YU2:     yu2,
		VOICE_YU3:     yu3,
		VOICE_YU4:     yu4,
		VOICE_YUAN1:   yuan1,
		VOICE_YUAN2:   yuan2,
		VOICE_YUAN3:   yuan3,
		VOICE_YUAN4:   yuan4,
		VOICE_YUE1:    yue1,
		VOICE_YUE4:    yue4,
		VOICE_YUN1:    yun1,
		VOICE_YUN2:    yun2,
		VOICE_YUN3:    yun3,
		VOICE_YUN4:    yun4,
		VOICE_ZA2:     za2,
		VOICE_ZA3:     za3,
		VOICE_ZAI1:    zai1,
		VOICE_ZAI3:    zai3,
		VOICE_ZAI4:    zai4,
		VOICE_ZAN2:    zan2,
		VOICE_ZAN4:    zan4,
		VOICE_ZANG1:   zang1,
		VOICE_ZANG4:   zang4,
		VOICE_ZAO1:    zao1,
		VOICE_ZAO2:    zao2,
		VOICE_ZAO3:    zao3,
		VOICE_ZAO4:    zao4,
		VOICE_ZE2:     ze2,
		VOICE_ZEI2:    zei2,
		VOICE_ZEN3:    zen3,
		VOICE_ZENG1:   zeng1,
		VOICE_ZENG4:   zeng4,
		VOICE_ZHA1:    zha1,
		VOICE_ZHA2:    zha2,
		VOICE_ZHA3:    zha3,
		VOICE_ZHA4:    zha4,
		VOICE_ZHAI1:   zhai1,
		VOICE_ZHAI2:   zhai2,
		VOICE_ZHAI3:   zhai3,
		VOICE_ZHAI4:   zhai4,
		VOICE_ZHAN1:   zhan1,
		VOICE_ZHAN3:   zhan3,
		VOICE_ZHAN4:   zhan4,
		VOICE_ZHANG1:  zhang1,
		VOICE_ZHANG3:  zhang3,
		VOICE_ZHANG4:  zhang4,
		VOICE_ZHAO1:   zhao1,
		VOICE_ZHAO3:   zhao3,
		VOICE_ZHAO4:   zhao4,
		VOICE_ZHE:     zhe,
		VOICE_ZHE1:    zhe1,
		VOICE_ZHE2:    zhe2,
		VOICE_ZHE3:    zhe3,
		VOICE_ZHE4:    zhe4,
		VOICE_ZHEN1:   zhen1,
		VOICE_ZHEN3:   zhen3,
		VOICE_ZHEN4:   zhen4,
		VOICE_ZHENG1:  zheng1,
		VOICE_ZHENG3:  zheng3,
		VOICE_ZHENG4:  zheng4,
		VOICE_ZHI1:    zhi1,
		VOICE_ZHI2:    zhi2,
		VOICE_ZHI3:    zhi3,
		VOICE_ZHI4:    zhi4,
		VOICE_ZHONG1:  zhong1,
		VOICE_ZHONG3:  zhong3,
		VOICE_ZHONG4:  zhong4,
		VOICE_ZHOU1:   zhou1,
		VOICE_ZHOU2:   zhou2,
		VOICE_ZHOU3:   zhou3,
		VOICE_ZHOU4:   zhou4,
		VOICE_ZHU1:    zhu1,
		VOICE_ZHU2:    zhu2,
		VOICE_ZHU3:    zhu3,
		VOICE_ZHU4:    zhu4,
		VOICE_ZHUA1:   zhua1,
		VOICE_ZHUAI1:  zhuai1,
		VOICE_ZHUAN1:  zhuan1,
		VOICE_ZHUAN3:  zhuan3,
		VOICE_ZHUAN4:  zhuan4,
		VOICE_ZHUANG1: zhuang1,
		VOICE_ZHUANG4: zhuang4,
		VOICE_ZHUI1:   zhui1,
		VOICE_ZHUI4:   zhui4,
		VOICE_ZHUN1:   zhun1,
		VOICE_ZHUN3:   zhun3,
		VOICE_ZHUO1:   zhuo1,
		VOICE_ZHUO2:   zhuo2,
		VOICE_ZI:      zi,
		VOICE_ZI1:     zi1,
		VOICE_ZI3:     zi3,
		VOICE_ZI4:     zi4,
		VOICE_ZONG1:   zong1,
		VOICE_ZONG3:   zong3,
		VOICE_ZONG4:   zong4,
		VOICE_ZOU3:    zou3,
		VOICE_ZOU4:    zou4,
		VOICE_ZU1:     zu1,
		VOICE_ZU2:     zu2,
		VOICE_ZU3:     zu3,
		VOICE_ZUAN1:   zuan1,
		VOICE_ZUI3:    zui3,
		VOICE_ZUI4:    zui4,
		VOICE_ZUN1:    zun1,
		VOICE_ZUO2:    zuo2,
		VOICE_ZUO3:    zuo3,
		VOICE_ZUO4:    zuo4,
	}
}

func GetVoice(v VOICE) []byte {
	return dataMap[v]
}
