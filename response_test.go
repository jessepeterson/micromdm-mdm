package mdm

import (
	"github.com/groob/plist"
	"testing"
)

var (
	macOSElCapQueryResponses = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CommandUUID</key>
	<string>bb8b2033-b81a-4927-8aa1-993d4ba2ade9</string>
	<key>QueryResponses</key>
	<dict>
		<key>ActiveManagedUsers</key>
		<array>
			<string>00000000-1111-2222-3333-444455556666</string>
		</array>
		<key>AvailableDeviceCapacity</key>
		<real>4.5032081604003906</real>
		<key>AwaitingConfiguration</key>
		<false/>
		<key>BluetoothMAC</key>
		<string>de-ad-be-ef-ca-fe</string>
		<key>BuildVersion</key>
		<string>15F34</string>
		<key>CurrentConsoleManagedUser</key>
		<string>00000000-1111-2222-3333-444455556666</string>
		<key>DeviceCapacity</key>
		<real>476.13906860351562</real>
		<key>DeviceName</key>
		<string>micromdm-macos</string>
		<key>HostName</key>
		<string>micromdm-macos.local</string>
		<key>Languages</key>
		<array>
			<string>en</string>
		</array>
		<key>LocalHostName</key>
		<string>micromdm-macos</string>
		<key>Locales</key>
		<array>
			<string>en_AU</string>
			<string>eu</string>
			<string>hr_BA</string>
			<string>en_CM</string>
			<string>rw_RW</string>
			<string>en_SZ</string>
			<string>tk_Latn</string>
			<string>he_IL</string>
			<string>ar</string>
			<string>uz_Arab</string>
			<string>en_PN</string>
			<string>as</string>
			<string>en_NF</string>
			<string>rwk_TZ</string>
			<string>zh_Hant_TW</string>
			<string>gsw_LI</string>
			<string>th_TH</string>
			<string>ta_IN</string>
			<string>es_EA</string>
			<string>fr_GF</string>
			<string>ar_001</string>
			<string>en_RW</string>
			<string>tr_TR</string>
			<string>de_CH</string>
			<string>ee_TG</string>
			<string>en_NG</string>
			<string>fr_TG</string>
			<string>az</string>
			<string>fr_SC</string>
			<string>es_HN</string>
			<string>en_AG</string>
			<string>ru_KZ</string>
			<string>gsw</string>
			<string>dyo</string>
			<string>so_ET</string>
			<string>zh_Hant_MO</string>
			<string>de_BE</string>
			<string>km_KH</string>
			<string>my_MM</string>
			<string>mgh_MZ</string>
			<string>ee_GH</string>
			<string>es_EC</string>
			<string>kw_GB</string>
			<string>rm_CH</string>
			<string>en_ME</string>
			<string>nyn</string>
			<string>mk_MK</string>
			<string>bs_Cyrl_BA</string>
			<string>ar_MR</string>
			<string>en_BM</string>
			<string>ms_Arab</string>
			<string>en_AI</string>
			<string>gl_ES</string>
			<string>en_PR</string>
			<string>ha_Latn_GH</string>
			<string>ff_CM</string>
			<string>ne_IN</string>
			<string>or_IN</string>
			<string>khq_ML</string>
			<string>en_MG</string>
			<string>pt_TL</string>
			<string>en_LC</string>
			<string>ta_SG</string>
			<string>jmc_TZ</string>
			<string>om_ET</string>
			<string>lv_LV</string>
			<string>es_US</string>
			<string>en_PT</string>
			<string>vai_Latn_LR</string>
			<string>en_NL</string>
			<string>iu_Cans_CA</string>
			<string>cgg_UG</string>
			<string>ta</string>
			<string>en_MH</string>
			<string>to_TO</string>
			<string>zu_ZA</string>
			<string>shi_Latn_MA</string>
			<string>brx_IN</string>
			<string>ar_KM</string>
			<string>en_AL</string>
			<string>te</string>
			<string>chr_US</string>
			<string>yo_BJ</string>
			<string>fr_VU</string>
			<string>pa</string>
			<string>tg</string>
			<string>ks_Arab</string>
			<string>kea</string>
			<string>ksh_DE</string>
			<string>sw_CD</string>
			<string>th</string>
			<string>te_IN</string>
			<string>fr_RE</string>
			<string>ur_IN</string>
			<string>yo_NG</string>
			<string>ti</string>
			<string>guz_KE</string>
			<string>tk</string>
			<string>kl_GL</string>
			<string>ksf_CM</string>
			<string>mua_CM</string>
			<string>lag_TZ</string>
			<string>lb</string>
			<string>fr_TN</string>
			<string>es_PA</string>
			<string>pl_PL</string>
			<string>to</string>
			<string>hi_IN</string>
			<string>dje_NE</string>
			<string>es_GQ</string>
			<string>kok_IN</string>
			<string>pl</string>
			<string>fr_GN</string>
			<string>bem</string>
			<string>ha</string>
			<string>ckb</string>
			<string>lg</string>
			<string>tr</string>
			<string>en_PW</string>
			<string>en_NO</string>
			<string>nyn_UG</string>
			<string>sr_Latn_RS</string>
			<string>gsw_FR</string>
			<string>pa_Guru</string>
			<string>he</string>
			<string>sn_ZW</string>
			<string>qu_BO</string>
			<string>lu_CD</string>
			<string>mgo_CM</string>
			<string>ps_AF</string>
			<string>en_BS</string>
			<string>ug_Arab</string>
			<string>da</string>
			<string>ms_Latn_SG</string>
			<string>ps</string>
			<string>ln</string>
			<string>pt</string>
			<string>iu_Cans</string>
			<string>hi</string>
			<string>lo</string>
			<string>ebu</string>
			<string>de</string>
			<string>gu_IN</string>
			<string>seh</string>
			<string>en_CX</string>
			<string>en_ZM</string>
			<string>tzm_Latn_MA</string>
			<string>fr_HT</string>
			<string>fr_GP</string>
			<string>lt</string>
			<string>lu</string>
			<string>ln_CD</string>
			<string>vai_Latn</string>
			<string>el_GR</string>
			<string>lv</string>
			<string>en_KE</string>
			<string>sbp</string>
			<string>hr</string>
			<string>en_CY</string>
			<string>es_GT</string>
			<string>twq_NE</string>
			<string>zh_Hant_HK</string>
			<string>kln_KE</string>
			<string>fr_GQ</string>
			<string>chr</string>
			<string>hu</string>
			<string>es_UY</string>
			<string>fr_CA</string>
			<string>en_NR</string>
			<string>mer</string>
			<string>shi</string>
			<string>es_PE</string>
			<string>fr_SN</string>
			<string>bez</string>
			<string>sw_TZ</string>
			<string>wae_CH</string>
			<string>kkj</string>
			<string>hy</string>
			<string>kk_Cyrl_KZ</string>
			<string>en_CZ</string>
			<string>teo_KE</string>
			<string>teo</string>
			<string>dz_BT</string>
			<string>ar_JO</string>
			<string>mer_KE</string>
			<string>khq</string>
			<string>ln_CF</string>
			<string>nn_NO</string>
			<string>en_MO</string>
			<string>ar_TD</string>
			<string>dz</string>
			<string>ses</string>
			<string>en_BW</string>
			<string>en_AS</string>
			<string>ar_IL</string>
			<string>ms_Latn_BN</string>
			<string>bo_CN</string>
			<string>nnh</string>
			<string>teo_UG</string>
			<string>hy_AM</string>
			<string>ln_CG</string>
			<string>sr_Latn_BA</string>
			<string>en_MP</string>
			<string>ksb_TZ</string>
			<string>ar_SA</string>
			<string>smn_FI</string>
			<string>ar_LY</string>
			<string>en_AT</string>
			<string>so_KE</string>
			<string>fr_CD</string>
			<string>af_NA</string>
			<string>en_NU</string>
			<string>es_PH</string>
			<string>en_KI</string>
			<string>en_JE</string>
			<string>lkt</string>
			<string>fa_IR</string>
			<string>ky_Cyrl</string>
			<string>uz_Latn_UZ</string>
			<string>zh_Hans_CN</string>
			<string>ewo_CM</string>
			<string>fr_PF</string>
			<string>ca_IT</string>
			<string>en_BZ</string>
			<string>ar_KW</string>
			<string>pt_GW</string>
			<string>fr_FR</string>
			<string>am_ET</string>
			<string>en_VC</string>
			<string>fr_DJ</string>
			<string>fr_CF</string>
			<string>es_SV</string>
			<string>en_MS</string>
			<string>pt_ST</string>
			<string>ar_SD</string>
			<string>luy_KE</string>
			<string>gd_GB</string>
			<string>de_LI</string>
			<string>fr_CG</string>
			<string>ckb_IQ</string>
			<string>zh_Hans_SG</string>
			<string>en_MT</string>
			<string>ewo</string>
			<string>af_ZA</string>
			<string>os_GE</string>
			<string>om_KE</string>
			<string>nl_SR</string>
			<string>es_ES</string>
			<string>es_DO</string>
			<string>ar_IQ</string>
			<string>fr_CH</string>
			<string>nnh_CM</string>
			<string>es_419</string>
			<string>en_MU</string>
			<string>bm_Latn</string>
			<string>en_US_POSIX</string>
			<string>yav_CM</string>
			<string>luo_KE</string>
			<string>dua_CM</string>
			<string>et_EE</string>
			<string>en_IE</string>
			<string>ak_GH</string>
			<string>rwk</string>
			<string>es_CL</string>
			<string>kea_CV</string>
			<string>fr_CI</string>
			<string>ckb_IR</string>
			<string>fr_BE</string>
			<string>se</string>
			<string>en_NZ</string>
			<string>ky_Cyrl_KG</string>
			<string>en_LR</string>
			<string>en_KN</string>
			<string>nb_SJ</string>
			<string>sg</string>
			<string>sr_Cyrl_RS</string>
			<string>ru_RU</string>
			<string>en_ZW</string>
			<string>sv_AX</string>
			<string>si</string>
			<string>ga_IE</string>
			<string>en_VG</string>
			<string>ff_MR</string>
			<string>sk</string>
			<string>agq_CM</string>
			<string>fr_BF</string>
			<string>naq_NA</string>
			<string>sl</string>
			<string>en_MW</string>
			<string>mr_IN</string>
			<string>az_Latn</string>
			<string>en_LS</string>
			<string>de_AT</string>
			<string>ka</string>
			<string>sn</string>
			<string>sr_Latn_ME</string>
			<string>fr_NC</string>
			<string>so</string>
			<string>is_IS</string>
			<string>twq</string>
			<string>ig_NG</string>
			<string>sq</string>
			<string>fo_FO</string>
			<string>sr</string>
			<string>tzm</string>
			<string>ga</string>
			<string>om</string>
			<string>en_LT</string>
			<string>bas_CM</string>
			<string>se_NO</string>
			<string>ki</string>
			<string>nl_BE</string>
			<string>ar_QA</string>
			<string>gd</string>
			<string>sv</string>
			<string>kk</string>
			<string>sw</string>
			<string>es_CO</string>
			<string>az_Latn_AZ</string>
			<string>rn_BI</string>
			<string>or</string>
			<string>kl</string>
			<string>ca</string>
			<string>en_VI</string>
			<string>km</string>
			<string>os</string>
			<string>en_MY</string>
			<string>kn</string>
			<string>en_LU</string>
			<string>fr_SY</string>
			<string>ar_TN</string>
			<string>en_JM</string>
			<string>fr_PM</string>
			<string>ko</string>
			<string>fr_NE</string>
			<string>fr_MA</string>
			<string>gl</string>
			<string>ru_MD</string>
			<string>saq_KE</string>
			<string>ks</string>
			<string>fr_CM</string>
			<string>lb_LU</string>
			<string>gv_IM</string>
			<string>fr_BI</string>
			<string>en_LV</string>
			<string>ks_Arab_IN</string>
			<string>es_NI</string>
			<string>en_GB</string>
			<string>kw</string>
			<string>nl_SX</string>
			<string>dav_KE</string>
			<string>tr_CY</string>
			<string>ky</string>
			<string>en_UG</string>
			<string>nus_SD</string>
			<string>en_TC</string>
			<string>tzm_Latn</string>
			<string>ar_EG</string>
			<string>fr_BJ</string>
			<string>gu</string>
			<string>es_PR</string>
			<string>fr_RW</string>
			<string>sr_Cyrl_BA</string>
			<string>gv</string>
			<string>fr_MC</string>
			<string>cs</string>
			<string>bez_TZ</string>
			<string>es_CR</string>
			<string>asa_TZ</string>
			<string>ar_EH</string>
			<string>ms_Arab_BN</string>
			<string>mn_Cyrl</string>
			<string>sbp_TZ</string>
			<string>en_IL</string>
			<string>ha_Latn_NE</string>
			<string>lt_LT</string>
			<string>mfe</string>
			<string>en_GD</string>
			<string>cy</string>
			<string>ca_FR</string>
			<string>es_BO</string>
			<string>fr_BL</string>
			<string>bn_IN</string>
			<string>uz_Cyrl_UZ</string>
			<string>az_Cyrl</string>
			<string>en_IM</string>
			<string>sw_KE</string>
			<string>en_SB</string>
			<string>pa_Arab</string>
			<string>ur_PK</string>
			<string>haw_US</string>
			<string>ar_SO</string>
			<string>en_IN</string>
			<string>ha_Latn</string>
			<string>fil</string>
			<string>fr_MF</string>
			<string>en_WS</string>
			<string>es_CU</string>
			<string>ja_JP</string>
			<string>fy_NL</string>
			<string>en_SC</string>
			<string>en_IO</string>
			<string>pt_PT</string>
			<string>en_HK</string>
			<string>en_GG</string>
			<string>fr_MG</string>
			<string>de_LU</string>
			<string>ms_Latn_MY</string>
			<string>tg_Cyrl</string>
			<string>en_SD</string>
			<string>shi_Tfng</string>
			<string>ln_AO</string>
			<string>ug_Arab_CN</string>
			<string>as_IN</string>
			<string>en_GH</string>
			<string>ro_RO</string>
			<string>jgo_CM</string>
			<string>dua</string>
			<string>en_UM</string>
			<string>en_SE</string>
			<string>kn_IN</string>
			<string>en_KY</string>
			<string>vun_TZ</string>
			<string>kln</string>
			<string>en_GI</string>
			<string>ca_ES</string>
			<string>rof</string>
			<string>pt_CV</string>
			<string>kok</string>
			<string>pt_BR</string>
			<string>ar_DJ</string>
			<string>yi_001</string>
			<string>fi_FI</string>
			<string>tg_Cyrl_TJ</string>
			<string>zh</string>
			<string>es_PY</string>
			<string>ar_SS</string>
			<string>mua</string>
			<string>sr_Cyrl_ME</string>
			<string>vai_Vaii_LR</string>
			<string>en_001</string>
			<string>nl_NL</string>
			<string>en_TK</string>
			<string>si_LK</string>
			<string>en_SG</string>
			<string>sv_SE</string>
			<string>fr_DZ</string>
			<string>ca_AD</string>
			<string>pt_AO</string>
			<string>vi</string>
			<string>xog_UG</string>
			<string>xog</string>
			<string>en_IS</string>
			<string>nb</string>
			<string>seh_MZ</string>
			<string>es_AR</string>
			<string>sk_SK</string>
			<string>en_SH</string>
			<string>ti_ER</string>
			<string>nd</string>
			<string>az_Cyrl_AZ</string>
			<string>zu</string>
			<string>ne</string>
			<string>nd_ZW</string>
			<string>el_CY</string>
			<string>en_IT</string>
			<string>nl_BQ</string>
			<string>da_GL</string>
			<string>ja</string>
			<string>rm</string>
			<string>fr_ML</string>
			<string>rn</string>
			<string>en_VU</string>
			<string>rof_TZ</string>
			<string>ro</string>
			<string>ebu_KE</string>
			<string>ru_KG</string>
			<string>en_SI</string>
			<string>sg_CF</string>
			<string>mfe_MU</string>
			<string>nl</string>
			<string>brx</string>
			<string>bs_Latn</string>
			<string>fa</string>
			<string>zgh_MA</string>
			<string>en_GM</string>
			<string>shi_Latn</string>
			<string>en_FI</string>
			<string>nn</string>
			<string>en_EE</string>
			<string>ru</string>
			<string>kam_KE</string>
			<string>fur</string>
			<string>vai_Vaii</string>
			<string>ar_ER</string>
			<string>ti_ET</string>
			<string>rw</string>
			<string>ff</string>
			<string>luo</string>
			<string>fa_AF</string>
			<string>ha_Latn_NG</string>
			<string>nl_CW</string>
			<string>en_HR</string>
			<string>en_FJ</string>
			<string>fi</string>
			<string>pt_MO</string>
			<string>be</string>
			<string>en_US</string>
			<string>en_TO</string>
			<string>en_SK</string>
			<string>bg</string>
			<string>ru_BY</string>
			<string>it_IT</string>
			<string>ml_IN</string>
			<string>gsw_CH</string>
			<string>qu_EC</string>
			<string>fo</string>
			<string>sv_FI</string>
			<string>en_FK</string>
			<string>nus</string>
			<string>ta_LK</string>
			<string>vun</string>
			<string>sr_Latn</string>
			<string>fr</string>
			<string>en_SL</string>
			<string>bm</string>
			<string>ar_BH</string>
			<string>guz</string>
			<string>bn</string>
			<string>bo</string>
			<string>ar_SY</string>
			<string>lo_LA</string>
			<string>ne_NP</string>
			<string>uz_Latn</string>
			<string>be_BY</string>
			<string>es_IC</string>
			<string>sr_Latn_XK</string>
			<string>ar_MA</string>
			<string>pa_Guru_IN</string>
			<string>br</string>
			<string>luy</string>
			<string>kde_TZ</string>
			<string>bs</string>
			<string>fy</string>
			<string>fur_IT</string>
			<string>hu_HU</string>
			<string>ar_AE</string>
			<string>en_HU</string>
			<string>sah_RU</string>
			<string>zh_Hans</string>
			<string>en_FM</string>
			<string>sq_AL</string>
			<string>ko_KP</string>
			<string>en_150</string>
			<string>en_DE</string>
			<string>fr_MQ</string>
			<string>en_CA</string>
			<string>hsb_DE</string>
			<string>en_TR</string>
			<string>ro_MD</string>
			<string>es_VE</string>
			<string>fr_WF</string>
			<string>mt_MT</string>
			<string>kab</string>
			<string>nmg_CM</string>
			<string>en_GR</string>
			<string>ru_UA</string>
			<string>fr_MR</string>
			<string>tk_Latn_TM</string>
			<string>zh_Hans_MO</string>
			<string>mn_Cyrl_MN</string>
			<string>ff_GN</string>
			<string>bs_Cyrl</string>
			<string>sw_UG</string>
			<string>ko_KR</string>
			<string>en_DG</string>
			<string>bo_IN</string>
			<string>en_CC</string>
			<string>shi_Tfng_MA</string>
			<string>lag</string>
			<string>it_SM</string>
			<string>os_RU</string>
			<string>en_TT</string>
			<string>ms_Arab_MY</string>
			<string>sq_MK</string>
			<string>ms_Latn</string>
			<string>bem_ZM</string>
			<string>kde</string>
			<string>ar_OM</string>
			<string>cgg</string>
			<string>bas</string>
			<string>bm_Latn_ML</string>
			<string>kam</string>
			<string>es_MX</string>
			<string>sah</string>
			<string>wae</string>
			<string>en_GU</string>
			<string>zh_Hant</string>
			<string>fr_MU</string>
			<string>fr_KM</string>
			<string>ar_LB</string>
			<string>en_BA</string>
			<string>en_TV</string>
			<string>sr_Cyrl</string>
			<string>dje</string>
			<string>kab_DZ</string>
			<string>fil_PH</string>
			<string>se_SE</string>
			<string>vai</string>
			<string>hr_HR</string>
			<string>bs_Latn_BA</string>
			<string>nl_AW</string>
			<string>dav</string>
			<string>so_SO</string>
			<string>ar_PS</string>
			<string>en_FR</string>
			<string>uz_Cyrl</string>
			<string>ff_SN</string>
			<string>en_BB</string>
			<string>ki_KE</string>
			<string>naq</string>
			<string>en_SS</string>
			<string>mg_MG</string>
			<string>mas_KE</string>
			<string>en_RO</string>
			<string>en_PG</string>
			<string>mgh</string>
			<string>dyo_SN</string>
			<string>mas</string>
			<string>agq</string>
			<string>bn_BD</string>
			<string>haw</string>
			<string>yi</string>
			<string>nb_NO</string>
			<string>da_DK</string>
			<string>en_DK</string>
			<string>saq</string>
			<string>ug</string>
			<string>cy_GB</string>
			<string>fr_YT</string>
			<string>jmc</string>
			<string>ses_ML</string>
			<string>en_PH</string>
			<string>de_DE</string>
			<string>ar_YE</string>
			<string>yo</string>
			<string>lkt_US</string>
			<string>uz_Arab_AF</string>
			<string>jgo</string>
			<string>sl_SI</string>
			<string>uk</string>
			<string>en_CH</string>
			<string>asa</string>
			<string>lg_UG</string>
			<string>qu_PE</string>
			<string>mgo</string>
			<string>id_ID</string>
			<string>en_NA</string>
			<string>en_GY</string>
			<string>zgh</string>
			<string>pt_MZ</string>
			<string>fr_LU</string>
			<string>kk_Cyrl</string>
			<string>mas_TZ</string>
			<string>en_DM</string>
			<string>ta_MY</string>
			<string>dsb</string>
			<string>en_BE</string>
			<string>mg</string>
			<string>ur</string>
			<string>fr_GA</string>
			<string>ka_GE</string>
			<string>nmg</string>
			<string>en_TZ</string>
			<string>eu_ES</string>
			<string>ar_DZ</string>
			<string>id</string>
			<string>so_DJ</string>
			<string>hsb</string>
			<string>yav</string>
			<string>mk</string>
			<string>pa_Arab_PK</string>
			<string>ml</string>
			<string>en_ER</string>
			<string>ig</string>
			<string>se_FI</string>
			<string>mn</string>
			<string>ksb</string>
			<string>uz</string>
			<string>vi_VN</string>
			<string>ii</string>
			<string>qu</string>
			<string>en_PK</string>
			<string>ee</string>
			<string>mr</string>
			<string>ms</string>
			<string>en_ES</string>
			<string>sq_XK</string>
			<string>it_CH</string>
			<string>mt</string>
			<string>en_CK</string>
			<string>br_FR</string>
			<string>sr_Cyrl_XK</string>
			<string>ksf</string>
			<string>en_SX</string>
			<string>bg_BG</string>
			<string>en_PL</string>
			<string>af</string>
			<string>el</string>
			<string>cs_CZ</string>
			<string>fr_TD</string>
			<string>zh_Hans_HK</string>
			<string>is</string>
			<string>ksh</string>
			<string>my</string>
			<string>en</string>
			<string>it</string>
			<string>dsb_DE</string>
			<string>ii_CN</string>
			<string>smn</string>
			<string>iu</string>
			<string>eo</string>
			<string>en_ZA</string>
			<string>en_AD</string>
			<string>ak</string>
			<string>en_RU</string>
			<string>kkj_CM</string>
			<string>am</string>
			<string>es</string>
			<string>et</string>
			<string>uk_UA</string>
		</array>
		<key>Model</key>
		<string>iMac17,1</string>
		<key>ModelName</key>
		<string>iMac</string>
		<key>OSUpdateSettings</key>
		<dict>
			<key>AutoCheckEnabled</key>
			<false/>
			<key>AutomaticAppInstallationEnabled</key>
			<false/>
			<key>AutomaticOSInstallationEnabled</key>
			<false/>
			<key>AutomaticSecurityUpdatesEnabled</key>
			<true/>
			<key>BackgroundDownloadEnabled</key>
			<true/>
			<key>CatalogURL</key>
			<string>https://swscan.apple.com/content/catalogs/others/index-10.11-10.10-10.9-mountainlion-lion-snowleopard-leopard.merged-1.sucatalog.gz</string>
			<key>IsDefaultCatalog</key>
			<true/>
			<key>PerformPeriodicCheck</key>
			<true/>
			<key>PreviousScanDate</key>
			<date>2016-06-08T03:50:31Z</date>
			<key>PreviousScanResult</key>
			<integer>0</integer>
		</dict>
		<key>OSVersion</key>
		<string>10.11.5</string>
		<key>ProductName</key>
		<string>iMac17,1</string>
		<key>SerialNumber</key>
		<string>C12A1234567L</string>
		<key>UDID</key>
		<string>11111111-2222-3333-4444-555566667777</string>
		<key>WiFiMAC</key>
		<string>dd:ee:aa:dd:be:ef</string>
		<key>iTunesStoreAccountHash</key>
		<string>aaaaaaaaAAAAAAAAaaaaAAAAAAA=</string>
		<key>iTunesStoreAccountIsActive</key>
		<true/>
	</dict>
	<key>RequestType</key>
	<string>DeviceInformation</string>
	<key>Status</key>
	<string>Acknowledged</string>
	<key>UDID</key>
	<string>11111111-2222-3333-4444-555566667777</string>
</dict>
</plist>`

	// IOS 8 Fixture is here because it contains different keys to an IOS 9 device. NOTE: WiFi no cellular.
	ios8IpadQueryResponses = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CommandUUID</key>
	<string>f14b5d8b-e1ae-4a3a-b217-e060803db82b</string>
	<key>QueryResponses</key>
	<dict>
		<key>AvailableDeviceCapacity</key>
		<real>3.9073715209960938</real>
		<key>BatteryLevel</key>
		<real>1</real>
		<key>BluetoothMAC</key>
		<string>de-ad-be-ef-ca-fe</string>
		<key>BuildVersion</key>
		<string>12H143</string>
		<key>CellularTechnology</key>
		<integer>0</integer>
		<key>DeviceCapacity</key>
		<real>26.629673004150391</real>
		<key>DeviceName</key>
		<string>micromdm-ios-8-ipad</string>
		<key>EASDeviceIdentifier</key>
		<string>AAAAAAAAAAAAAAAAAAAAAAAAAA</string>
		<key>IsActivationLockEnabled</key>
		<false/>
		<key>IsCloudBackupEnabled</key>
		<false/>
		<key>IsDeviceLocatorServiceEnabled</key>
		<false/>
		<key>IsDoNotDisturbInEffect</key>
		<false/>
		<key>IsSupervised</key>
		<false/>
		<key>Model</key>
		<string>MD786X</string>
		<key>ModelName</key>
		<string>iPad</string>
		<key>OSVersion</key>
		<string>8.4</string>
		<key>ProductName</key>
		<string>iPad4,1</string>
		<key>SerialNumber</key>
		<string>ABCDABCDAB00</string>
		<key>UDID</key>
		<string>1111111111111111111111111111111111111111</string>
		<key>WiFiMAC</key>
		<string>dd:ee:aa:dd:be:e1</string>
		<key>iTunesStoreAccountHash</key>
		<string>aaaaaaaaAAAAAAAAaaaaAAAAAAA=</string>
		<key>iTunesStoreAccountIsActive</key>
		<true/>
	</dict>
	<key>Status</key>
	<string>Acknowledged</string>
	<key>UDID</key>
	<string>1111111111111111111111111111111111111111</string>
</dict>
</plist>`

	ios8IphoneQueryResponses = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CommandUUID</key>
	<string>c3911e20-e407-4ef8-b919-e9e3bce35ac6</string>
	<key>QueryResponses</key>
	<dict>
		<key>AvailableDeviceCapacity</key>
		<real>24.264816284179688</real>
		<key>BatteryLevel</key>
		<real>0.54000002145767212</real>
		<key>BluetoothMAC</key>
		<string>de-ad-be-ef-ca-fe</string>
		<key>BuildVersion</key>
		<string>12F70</string>
		<key>CarrierSettingsVersion</key>
		<string>19.0</string>
		<key>CellularTechnology</key>
		<integer>3</integer>
		<key>CurrentMCC</key>
		<string>504</string>
		<key>CurrentMNC</key>
		<string>02</string>
		<key>DataRoamingEnabled</key>
		<false/>
		<key>DeviceCapacity</key>
		<real>55.89947509765625</real>
		<key>DeviceName</key>
		<string>micromdm-iphone6-ios8</string>
		<key>EASDeviceIdentifier</key>
		<string>AAAAAAAAAAAAAAAAAAAAAAAAAA</string>
		<key>ICCID</key>
		<string>1111 2222 3333 4444 5555</string>
		<key>IMEI</key>
		<string>11 222222 333333 4</string>
		<key>IsActivationLockEnabled</key>
		<false/>
		<key>IsCloudBackupEnabled</key>
		<false/>
		<key>IsDeviceLocatorServiceEnabled</key>
		<false/>
		<key>IsDoNotDisturbInEffect</key>
		<false/>
		<key>IsRoaming</key>
		<false/>
		<key>IsSupervised</key>
		<false/>
		<key>MEID</key>
		<string>00000000000000</string>
		<key>Model</key>
		<string>MG4F2X</string>
		<key>ModelName</key>
		<string>iPhone</string>
		<key>ModemFirmwareVersion</key>
		<string>2.23.03</string>
		<key>OSVersion</key>
		<string>8.3</string>
		<key>PersonalHotspotEnabled</key>
		<false/>
		<key>PhoneNumber</key>
		<string>+00111111111</string>
		<key>ProductName</key>
		<string>iPhone7,2</string>
		<key>SIMCarrierNetwork</key>
		<string>Tincan</string>
		<key>SerialNumber</key>
		<string>000000000001</string>
		<key>SubscriberCarrierNetwork</key>
		<string>Tincan</string>
		<key>SubscriberMCC</key>
		<string>502</string>
		<key>SubscriberMNC</key>
		<string>02</string>
		<key>UDID</key>
		<string>1111111111111111111111111111111111111112</string>
		<key>WiFiMAC</key>
		<string>dd:ee:aa:dd:be:e2</string>
		<key>iTunesStoreAccountHash</key>
		<string>aaaaaaaaAAAAAAAAaaaaAAAAAAA=</string>
		<key>iTunesStoreAccountIsActive</key>
		<true/>
	</dict>
	<key>Status</key>
	<string>Acknowledged</string>
	<key>UDID</key>
	<string>1111111111111111111111111111111111111112</string>
</dict>
</plist>`
)

func TestQueryResponseMac(t *testing.T) {
	response := &Response{}
	plistBuf := []byte(macOSElCapQueryResponses)
	err := plist.Unmarshal(plistBuf, response)

	if err != nil {
		t.Fatal(err)
	}
}

func TestQueryResponseIpadIOS8(t *testing.T) {
	response := &Response{}
	plistBuf := []byte(ios8IpadQueryResponses)
	err := plist.Unmarshal(plistBuf, response)

	if err != nil {
		t.Fatal(err)
	}
}

func TestQueryResponseIphoneIOS8(t *testing.T) {
	response := &Response{}
	plistBuf := []byte(ios8IphoneQueryResponses)
	err := plist.Unmarshal(plistBuf, response)

	if err != nil {
		t.Fatal(err)
	}
}
