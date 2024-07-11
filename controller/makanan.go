package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alfianbr16/backend-tb/config"
	inimodel "github.com/alfianbr16/package-tb/model"
	cek "github.com/alfianbr16/package-tb/module"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetMakanan godoc
// @Summary Get All Data Makanan.
// @Description Mengambil semua data makanan.
// @Tags Makanan
// @Accept json
// @Produce json
// @Success 200 {object} MakananHewan
// @Router /makanan [get]
func GetMakanan(c *fiber.Ctx) error {
	ps := cek.GetAllMakanan(config.Ulbimongoconn, "makananhewan")
	return c.JSON(ps)
}

// GetMakananID godoc
// @Summary Get By ID Data Makanan.
// @Description Ambil per ID data makanan.
// @Tags Makanan
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} MakananHewan
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /makanan/{id} [get]
func GetMakananID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := cek.GetMakananFromID(objID, config.Ulbimongoconn, "makananhewan")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

// InsertDataMakanan godoc
// @Summary Insert data makanan.
// @Description Input data makanan.
// @Tags Makanan
// @Accept json
// @Produce json
// @Param request body ReqMakananHewan true "Payload Body [RAW]"
// @Success 200 {object} MakananHewan
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertDataMakanan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var makan inimodel.MakananHewan

	// Parse body request ke struct makan
	if err := c.BodyParser(&makan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Lakukan operasi insert ke database
	insertedID, err := cek.InsertMakanan(db, "makananhewan",
		makan.Hewan,
		makan.JenisMakanan,
		makan.Bahan,
		makan.Berat,
		makan.Rasa,
		makan.Merk,
		makan.Harga)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Berhasil menyimpan data
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// UpdateData godoc
// @Summary Update data makanan.
// @Description Ubah data makanan.
// @Tags Makanan
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqMakananHewan true "Payload Body [RAW]"
// @Success 200 {object} MakananHewan
// @Failure 400
// @Failure 500
// @Router /update/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var makan inimodel.MakananHewan
	if err := c.BodyParser(&makan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdateMakanan function with the parsed ID and the MakananHewan object
	err = cek.UpdateMakanan(db, "makananhewan",
		objectID,
		makan.Hewan,
		makan.JenisMakanan,
		makan.Bahan,
		makan.Berat,
		makan.Rasa,
		makan.Merk,
		makan.Harga)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// DeleteMakananByID godoc
// @Summary Delete data makanan.
// @Description Hapus data makanan.
// @Tags Makanan
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeleteMakananByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = cek.DeleteMakananByID(objID, config.Ulbimongoconn, "makananhewan")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}