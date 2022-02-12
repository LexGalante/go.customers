package infrastructures

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lexgalante/go.customers/api/src/models"
)

//GetAddresFromViaCep -> search address in viacep
func GetAddresFromViaCep(cep string) (models.Address, error) {
	var address models.Address

	resp, err := http.Get(fmt.Sprintf("%s/%s/json/", os.Getenv("VIA_CEP_URL"), cep))
	if err != nil {
		return address, err
	}

	if resp.StatusCode == 500 || resp.StatusCode == 400 {
		return address, fmt.Errorf("error while trying to search address at viacep: %s", resp.Status)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return address, err
	}

	err = json.Unmarshal(body, &address)
	if err != nil {
		return address, err
	}

	return address, nil
}
