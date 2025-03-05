package ordersvr

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thannaton/interview-golang/internal/core/constant/enum"
)

type (
	GetOrderInput struct {
		No                int
		PlatformProductId string
		Qty               int
		UnitPrice         float64
		TotalPrice        float64
	}

	GetOrderOutput struct {
		No         int
		ProductId  string
		MaterialId string
		ModelId    string
		Qty        int
		UnitPrice  float64
		TotalPrice float64
	}
)

func (s *orderService) Get(c *gin.Context, input GetOrderInput) ([]GetOrderOutput, error) {
	var totalQty, lastNo int
	mapInfos := make(map[string]GetOrderOutput)
	mapFreeGift := make(map[string]GetOrderOutput)

	platfromProducts := strings.Split(input.PlatformProductId, "/")
	for i, product := range platfromProducts {
		lastNo = i + 1
		qty := 1
		id := string(regexp.MustCompile(`[A-Z0-9]+-[A-Z0-9]+-[A-Z0-9]+-[A-Z]+|[A-Z0-9]+-[A-Z0-9]+-[A-Z0-9]+`).Find([]byte(product)))
		matchQty := regexp.MustCompile(`\*(\d+)`).FindStringSubmatch(product)
		if len(matchQty) > 1 {
			qty, _ = strconv.Atoi(matchQty[1])
		}

		splitProduct := strings.Split(id, "-")
		filmTypeId, textureId, phoneModelId := splitProduct[0], splitProduct[1], splitProduct[2]
		if len(splitProduct) > 3 && splitProduct[3] != "" {
			phoneModelId = fmt.Sprintf("%v-%v", phoneModelId, splitProduct[3])
		}

		mapInfos[id] = GetOrderOutput{
			No:         lastNo,
			ProductId:  id,
			MaterialId: fmt.Sprintf("%v-%v", filmTypeId, textureId),
			ModelId:    phoneModelId,
			Qty:        qty,
		}

		if val, ok := mapFreeGift[textureId]; !ok {
			mapFreeGift[textureId] = GetOrderOutput{
				ProductId:  fmt.Sprintf("%v-CLEANNER", textureId),
				Qty:        qty,
				UnitPrice:  float64(0),
				TotalPrice: float64(0),
			}
		} else {
			val.Qty += qty
		}

		totalQty += qty
	}

	result := make([]GetOrderOutput, 0)

	unitPrice := float64(input.UnitPrice) / float64(totalQty)

	for _, value := range mapInfos {
		result = append(result, GetOrderOutput{
			No:         value.No,
			ProductId:  value.ProductId,
			MaterialId: value.MaterialId,
			ModelId:    value.ModelId,
			Qty:        value.Qty,
			UnitPrice:  unitPrice,
			TotalPrice: unitPrice * float64(value.Qty),
		})

	}

	result = append(result, addWipingCloth(lastNo, totalQty))
	lastNo++

	if val := addFreeGift(lastNo, enum.Clear.String(), mapFreeGift); val != (GetOrderOutput{}) {
		result = append(result, val)
		lastNo++
	}

	if val := addFreeGift(lastNo, enum.Matte.String(), mapFreeGift); val != (GetOrderOutput{}) {
		result = append(result, val)
		lastNo++
	}

	if val := addFreeGift(lastNo, enum.Privacy.String(), mapFreeGift); val != (GetOrderOutput{}) {
		result = append(result, val)
		lastNo++
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].No < result[j].No
	})

	return result, nil
}

func addWipingCloth(runningNo int, qty int) GetOrderOutput {
	return GetOrderOutput{
		No:         runningNo,
		ProductId:  enum.WipingCloth.String(),
		Qty:        qty,
		UnitPrice:  float64(0),
		TotalPrice: float64(0),
	}
}

func addFreeGift(runingNo int, textureId string, mapFreeGift map[string]GetOrderOutput) GetOrderOutput {
	value, ok := mapFreeGift[textureId]
	if !ok {
		return GetOrderOutput{}
	}

	value.No = runingNo
	return value
}
