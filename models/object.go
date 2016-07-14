package models

type Object struct {

		TableName struct {}																				`sql:"dm_informatieportal,alias:object" json:"-"`

		Aangewezigheid_correcte_sleutel string										`json:"aangewezigheid_correcte_sleutel"`
		Aantal_gevaren string																			`json:"aantal_gevaren"`
		Aanwezigheid_bluswater string															`json:"aanwezigheid_bluswater"`
		Aanwezigheid_sleutelkluis string													`json:"aanwezigheid_sleutelkluis"`
		Aanwezigheid_voedingspunt_droge_blusleiding string				`json:"aanwezigheid_voedingspunt_droge_blusleiding"`
		Adres_bouwjaar int64																			`json:"adres_bouwjaar"`
		Adres_gemeentenaam string																	`json:"adres_gemeentenaam"`
		Adres_huisletter string																		`json:"adres_huisletter"`
		Adres_huisnummer int64																		`json:"adres_huisnummer"`
		Adres_huisnummertoevoeging string													`json:"adres_huisnummertoevoeging"`
		Adres_postcode string																			`json:"adres_postcode"`
		Adres_straatnaam string																		`json:"adres_straatnaam"`
		Adres_woonplaats string																		`json:"adres_woonplaats"`
		Alternatieve_toegang_geregeld string											`json:"alternatieve_toegang_geregeld"`
		Alternatieve_toegang_omschrijving string									`json:"alternatieve_toegang_omschrijving"`
		Bag_id string																							`json:"bag_id"`
		Bijzondere_gevaren string																	`json:"bijzondere_gevaren"`
		Date_updated	string																			`json:"-"`
		Foto_brandmeldpaneel string																`json:"foto_brandmeldpaneel"`
		Foto_brandweer_hoofdingang string													`json:"foto_brandweer_hoofdingang"`
		Foto_neveningang string																		`json:"foto_neveningang"`
		Foto_per_gevaar string																		`json:"foto_per_gevaar"`
		Geometry_center_geo string																`json:"geometry_center_geo"`
		Geometry_center_geo_1 string															`json:"-"`
		Geometry_omtrek_geo string																`json:"-"`
		Geometry_rd_x string																			`json:"geometry_rd_x"`
		Geometry_rd_y string																			`json:"geometry_rd_y"`
		Hub_object_bk int64																				`json:"hub_object_bk"`
		Hub_object_id int64																				`json:"hub_object_id"`
		Hub_object_naam string																		`json:"hub_object_naam"`
		Object_functie string																			`json:"object_functie"`
		Object_functie_1 string																		`json:"-"`
		Orienteren_check string																		`json:"orienteren_check"`
		Overige_gevaren string																		`json:"overige_gevaren"`
		Risicovol_categorie string																`json:"risicovol_categorie"`
		Risicovol_naam string																			`json:"risicovol_naam"`
		Risicovol_soort string																		`json:"risicovol_soort"`
		Risicovol_type string																			`json:"risicovol_type"`
		Voldoende_opstelplaats_bij_ingang string									`json:"voldoende_opstelplaats_bij_ingang"`
		Zuurstof_opslag string																		`json:"zuurstof_opslag"`
}