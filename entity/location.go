package entity

import (
	"time"
)

type Location struct {
	GPS_longitude_SiteDeVente           *float64   `json:"GPS_longitude_SiteDeVente,omitempty"`
	GPS_latitude_SiteDeVente            *float64   `json:"GPS_latitude_SiteDeVente,omitempty"`
	Code_SiteDeVente                    string     `json:"Code_SiteDeVente"`
	Code_SiteDeVente_new                string     `json:"Code_SiteDeVente_new"`
	Nom_SiteDeVente                     string     `json:"Nom_SiteDeVente"`
	Adr1_SiteDeVente                    *string    `json:"Adr1_SiteDeVente,omitempty"`
	Adr2_SiteDeVente                    *string    `json:"Adr2_SiteDeVente,omitempty"`
	Adr3_SiteDeVente                    string     `json:"Adr3_SiteDeVente"`
	CodePostal_SiteDeVente              string     `json:"CodePostal_SiteDeVente"`
	Ville_SiteDeVente                   string     `json:"Ville_SiteDeVente"`
	Dpt_SiteDeVente                     string     `json:"Dpt_SiteDeVente"`
	Tel_SiteDeVente                     string     `json:"Tel_SiteDeVente"`
	Date_Ouverture_SiteDeVente          *time.Time `json:"Date_Ouverture_SiteDeVente,omitempty"`
	Fax_SiteDeVente                     string     `json:"Fax_SiteDeVente"`
	Date_Fermeture_SiteDeVente          *time.Time `json:"Date_Fermeture_SiteDeVente,omitempty"`
	Rcs_SiteDeVente                     *string    `json:"rcs_SiteDeVente,omitempty"`
	Email_SiteDeVente                   *string    `json:"email_SiteDeVente,omitempty"`
	Id_type_activite_SiteDeVente        *uint64    `json:"id_type_activite_SiteDeVente,omitempty"`
	Type_activite_SiteDeVente           *string    `json:"type_activite_SiteDeVente,omitempty"`
	Site_Web_SiteDeVente                *string    `json:"Site_Web_SiteDeVente,omitempty"`
	Id_web_sitedevente                  *uint64    `json:"id_web_sitedevente,omitempty"`
	Descriptif_SiteDeVente              *string    `json:"descriptif_SiteDeVente,omitempty"`
	Infos_cplt_SiteDeVente              *string    `json:"infos_cplt_SiteDeVente,omitempty"`
	Id_Localite_SiteDeVente             *uint64    `json:"id_Localite_SiteDeVente,omitempty"`
	Localite_SiteDeVente                *string    `json:"Localite_SiteDeVente,omitempty"`
	Id_Type_Flux_VO_SiteDeVente         *uint64    `json:"id_Type_Flux_VO_SiteDeVente,omitempty"`
	Type_Flux_VO_SiteDeVente            *string    `json:"Type_Flux_VO_SiteDeVente,omitempty"`
	Id_Volume_Occasion_SiteDeVente      *uint64    `json:"id_Volume_Occasion_SiteDeVente,omitempty"`
	Volume_Occasion_SiteDeVente         *string    `json:"Volume_Occasion_SiteDeVente,omitempty"`
	Id_Type_DMS_SiteDeVente             *uint64    `json:"id_Type_DMS_SiteDeVente,omitempty"`
	Type_DMS_SiteDeVente                *string    `json:"Type_DMS_SiteDeVente,omitempty"`
	Capital_SiteDeVente                 *uint64    `json:"Capital_SiteDeVente,omitempty"`
	Id_Forme_Juridique_SiteDeVente      *uint64    `json:"id_Forme_Juridique_SiteDeVente,omitempty"`
	Forme_Juridique_SiteDeVente         *string    `json:"Forme_Juridique_SiteDeVente,omitempty"`
	Code_Concession                     *string    `json:"Code_Concession,omitempty"`
	Nom_Concession                      *string    `json:"Nom_Concession,omitempty"`
	Adr1_Concession                     *string    `json:"Adr1_Concession,omitempty"`
	Adr2_Concession                     *string    `json:"Adr2_Concession,omitempty"`
	Adr3_Concession                     *string    `json:"Adr3_Concession,omitempty"`
	CodePostal_Concession               *string    `json:"CodePostal_Concession,omitempty"`
	Ville_Concession                    *string    `json:"Ville_Concession,omitempty"`
	Rcs_Concession                      *string    `json:"rcs_Concession,omitempty"`
	Date_Ouverture_Concession           *time.Time `json:"Date_Ouverture_Concession,omitempty"`
	Date_Fermeture_Concession           *time.Time `json:"Date_Fermeture_Concession,omitempty"`
	Tel_Concession                      *string    `json:"Tel_Concession,omitempty"`
	Fax_Concession                      *string    `json:"Fax_Concession,omitempty"`
	Email_Concession                    *string    `json:"Email_Concession,omitempty"`
	Site_Web_Concession                 *string    `json:"Site_Web_Concession,omitempty"`
	Prenom_directeur_Concession         *string    `json:"Prenom_directeur_Concession,omitempty"`
	Nom_directeur_Concession            *string    `json:"Nom_directeur_Concession,omitempty"`
	Id_Groupe_investissement_Concession *uint64    `json:"id_Groupe_investissement_Concession,omitempty"`
	Groupe_investissement_Concession    *string    `json:"Groupe_investissement_Concession,omitempty"`
	Id_region_vente                     *uint64    `json:"id_region_vente,omitempty"`
	Libelle_region_vente                *string    `json:"libelle_region_vente,omitempty"`
	Id_region_SAV                       *uint64    `json:"id_region_SAV,omitempty"`
	Libelle_region_SAV                  *string    `json:"libelle_region_SAV,omitempty"`
	Id_region_PRA                       *uint64    `json:"id_region_PRA,omitempty"`
	Libelle_region_PRA                  *string    `json:"libelle_region_PRA,omitempty"`
	Code_departement                    *string    `json:"code_departement,omitempty"`
	Nom_departement                     *string    `json:"nom_departement,omitempty"`
	Code_region                         *string    `json:"code_region,omitempty"`
	Nom_region                          *string    `json:"nom_region,omitempty"`
	Champ_Support                       *string    `json:"champ_Support,omitempty"`
}
