package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hewan struct {
	ID           	primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Jenis 	        string             	`bson:"jenis,omitempty" json:"jenis,omitempty" example:"Anjing"`
	Umur 	  		string 			 	`bson:"umur,omitempty" json:"umur,omitempty" example:"8"`
	Ras 			string 			 	`bson:"ras,omitempty" json:"ras,omitempty" example:"Pitbull"`
}

type MakananHewan struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Hewan           Hewan             	`bson:"hewan,omitempty" json:"hewan,omitempty"`
	JenisMakanan    string             	`bson:"jenismakanan,omitempty" json:"jenismakanan,omitempty" example:"Basah"`
	Bahan         	string             	`bson:"bahan,omitempty" json:"bahan,omitempty" example:"Daging Ikan"`
	Berat           string             	`bson:"berat,omitempty" json:"berat,omitempty" example:"4"`
	Rasa         	string             	`bson:"rasa,omitempty" json:"rasa,omitempty" example:"Tuna"`
	Merk		    string             	`bson:"merk,omitempty" json:"merk,omitempty" example:"Whiskas"`
	Harga		    string             	`bson:"harga,omitempty" json:"harga,omitempty" example:"500000"`
	Tanggal  		primitive.DateTime 	`bson:"tanggal,omitempty" json:"tanggal,omitempty" swaggertype:"string" example:"2024-09-01T00:00:00Z" format:"date-time"`
}

type ReqHewan struct {
	Jenis 	        string             	`bson:"jenis,omitempty" json:"jenis,omitempty" example:"Anjing"`
	Umur 	  		string 			 	`bson:"umur,omitempty" json:"umur,omitempty" example:"8"`
	Ras 			string 			 	`bson:"ras,omitempty" json:"ras,omitempty" example:"Pitbull"`
}

type ReqMakananHewan struct {
	Hewan           ReqHewan             	`bson:"hewan,omitempty" json:"hewan,omitempty"`
	JenisMakanan    string             	`bson:"jenismakanan,omitempty" json:"jenismakanan,omitempty" example:"Basah"`
	Bahan         	string             	`bson:"bahan,omitempty" json:"bahan,omitempty" example:"Daging Ikan"`
	Berat           string             	`bson:"berat,omitempty" json:"berat,omitempty" example:"4"`
	Rasa         	string             	`bson:"rasa,omitempty" json:"rasa,omitempty" example:"Tuna"`
	Merk		    string             	`bson:"merk,omitempty" json:"merk,omitempty" example:"Whiskas"`
	Harga		    string             	`bson:"harga,omitempty" json:"harga,omitempty" example:"500000"`
	Tanggal  		primitive.DateTime 	`bson:"tanggal,omitempty" json:"tanggal,omitempty" swaggertype:"string" example:"2024-09-01T00:00:00Z" format:"date-time"`
}