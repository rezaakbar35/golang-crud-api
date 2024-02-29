package photoController

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/rezaakbar35/golang-crud-api/model"
    "gorm.io/gorm"
)

func GetAllPhoto(c *gin.Context){
    var photos []model.Photo

    if err := model.DB.Find(&photos).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"photo": photos})
}

func GetPhotoById(c *gin.Context){
    var photo model.Photo
    id := c.Param("id")

    if err := model.DB.First(&photo, id).Error; err != nil {
        switch err {
        case gorm.ErrRecordNotFound:
            c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Photo not found!"})
            return
        default:
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"photo": photo})
}

func CreatePhoto(c *gin.Context){
    var photo model.Photo

    if err := c.ShouldBindJSON(&photo); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    model.DB.Create(&photo)
    c.JSON(http.StatusOK, gin.H{"message": "Create Success!","photo": photo})
}

func UpdatePhoto(c *gin.Context){
    var photo model.Photo
    id := c.Param("id")

    if err := model.DB.First(&photo, id).Error; err != nil {
        switch err {
        case gorm.ErrRecordNotFound:
            c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Photo not found!"})
            return
        default:
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    if err := c.ShouldBindJSON(&photo); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := model.DB.Save(&photo).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Update Success!", "photo": photo})
}

func DeletePhoto(c *gin.Context){
    var photo model.Photo
    id := c.Param("id")

    if err := model.DB.First(&photo, id).Error; err != nil {
        switch err {
        case gorm.ErrRecordNotFound:
            c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Photo not found!"})
            return
        default:
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    if err := model.DB.Delete(&photo, id).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Delete Success!", "photo": photo})
}

