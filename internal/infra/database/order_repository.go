package database

import (
	"cep-gin-clean-arch/models"
	"encoding/json"
	"errors"
	"os"

	"github.com/supabase-community/supabase-go"
)

type CEPRepository struct{}

func NewCEPRepository() *CEPRepository {
	return &CEPRepository{}
}

type Address struct {
	CEP    string `json:"CEP"`
	Estado string `json:"Estado"`
	Cidade string `json:"Cidade"`
	Bairro string `json:"Bairro"`
	Rua    string `json:"Rua"`
}

const jsonDados = `
[
	{"CEP": "90010000", "Estado": "RS", "Cidade": "Porto Alegre", "Bairro": "Centro Histórico", "Rua": "Rua da Praia"},
    {"CEP": "90010100", "Estado": "RS", "Cidade": "Porto Alegre", "Bairro": "Centro Histórico", "Rua": "Rua dos Andradas"},
    {"CEP": "90020000", "Estado": "RS", "Cidade": "Porto Alegre", "Bairro": "Centro", "Rua": "Avenida Borges de Medeiros"},
    {"CEP": "90030000", "Estado": "RS", "Cidade": "Porto Alegre", "Bairro": "Moinhos de Vento", "Rua": "Rua Mostardeiro"},
    {"CEP": "91010000", "Estado": "RS", "Cidade": "Porto Alegre", "Bairro": "Cristo Redentor", "Rua": "Avenida Assis Brasil"},
    {"CEP": "91020000", "Estado": "RS", "Cidade": "Porto Alegre", "Bairro": "Passo D'Areia", "Rua": "Avenida João Wallig"},
    {"CEP": "91030000", "Estado": "RS", "Cidade": "Porto Alegre", "Bairro": "Sarandi", "Rua": "Avenida Sertório"},
    {"CEP": "92010000", "Estado": "RS", "Cidade": "Canoas", "Bairro": "Centro", "Rua": "Rua Tiradentes"},
    {"CEP": "92020000", "Estado": "RS", "Cidade": "Canoas", "Bairro": "Niterói", "Rua": "Avenida Santos Ferreira"},
    {"CEP": "93010000", "Estado": "RS", "Cidade": "São Leopoldo", "Bairro": "Centro", "Rua": "Rua Independência"},
    {"CEP": "93020000", "Estado": "RS", "Cidade": "São Leopoldo", "Bairro": "Scharlau", "Rua": "Rua São João"},
    {"CEP": "94010000", "Estado": "RS", "Cidade": "Gravataí", "Bairro": "Centro", "Rua": "Rua Anápio Gomes"},
    {"CEP": "94020000", "Estado": "RS", "Cidade": "Gravataí", "Bairro": "Vera Cruz", "Rua": "Avenida Dorival Cândido Luz de Oliveira"},
    {"CEP": "95010000", "Estado": "RS", "Cidade": "Caxias do Sul", "Bairro": "Centro", "Rua": "Rua Sinimbu"},
    {"CEP": "95020000", "Estado": "RS", "Cidade": "Caxias do Sul", "Bairro": "São Pelegrino", "Rua": "Rua Os Dezoito do Forte"},
    {"CEP": "01001000", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Sé", "Rua": "Praça da Sé"},
    {"CEP": "01001001", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Sé", "Rua": "Praça da Sé"},
    {"CEP": "01153000", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Barra Funda", "Rua": "Rua Vitorino Carmilo"},
    {"CEP": "03223050", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Jardim Independência", "Rua": "Rua Dario Meira"},
    {"CEP": "03255000", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Vila Tolstoi", "Rua": "Rua José Antônio Fontes"},
    {"CEP": "03513050", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Vila Matilde", "Rua": "Rua Antônio Juvenal"},
    {"CEP": "03683010", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Vila União (Zona Leste)", "Rua": "Rua Antônio Olímpio"},
    {"CEP": "04295001", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Vila Vera", "Rua": "Avenida Coronel José Pires de Andrade"},
    {"CEP": "05425070", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Pinheiros", "Rua": "Avenida Doutora Ruth Cardoso"},
    {"CEP": "05434020", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Vila Madalena", "Rua": "Rua Morás"},
    {"CEP": "08090284", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Jardim Helena", "Rua": "Rua 03 de Outubro"},
    {"CEP": "08410500", "Estado": "SP", "Cidade": "São Paulo", "Bairro": "Vila Santa Cruz (Zona Leste)", "Rua": "Rua Doutor Pedro Batista"},
    {"CEP": "09925000", "Estado": "SP", "Cidade": "Diadema", "Bairro": "Campanário", "Rua": "Rodovia dos Imigrantes"},
    {"CEP": "11441710", "Estado": "SP", "Cidade": "Guarujá", "Bairro": "Enseada", "Rua": "Rua Projetada Nove"},
    {"CEP": "12209530", "Estado": "SP", "Cidade": "São José dos Campos", "Bairro": "Centro", "Rua": "Rua José de Alencar"},
    {"CEP": "13327220", "Estado": "SP", "Cidade": "Salto", "Bairro": "Jardim Saltense", "Rua": "Rua Acácio Rodrigues de Moraes"},
    {"CEP": "13401130", "Estado": "SP", "Cidade": "Piracicaba", "Bairro": "Paulista", "Rua": "Rua Sud Mennucci"},
    {"CEP": "17022113", "Estado": "SP", "Cidade": "Bauru", "Bairro": "Vila São Paulo", "Rua": "Rua Gaudêncio Piola"},
    {"CEP": "20020030", "Estado": "RJ", "Cidade": "Rio de Janeiro", "Bairro": "Centro", "Rua": "Praça Academia"}
]
`

func (r *CEPRepository) Buscar(cep string) (models.CEPResponse, error) {

	// Primeiro busca os dados "em memória"
	dados := []byte(jsonDados)

	var addresses []Address

	// Unmarshal JSON data
	err := json.Unmarshal(dados, &addresses)
	if err != nil {
		return models.CEPResponse{}, errors.New("Erro ao acessar dados de CEP")
	}

	// Create a map for quick lookup
	addressMap := make(map[string]Address)
	for _, address := range addresses {
		addressMap[address.CEP] = address
	}
	// Se o CEP estiver na memória, retorna os dados
	desiredAddress, found := addressMap[cep]
	if found {
		return models.CEPResponse{
			Estado: desiredAddress.Estado,
			Cidade: desiredAddress.Cidade,
			Bairro: desiredAddress.Bairro,
			Rua:    desiredAddress.Rua,
		}, nil
	}

	// Caso nao encontre, busca na tabela Supabase
	// Se der erro ao buscar na tabela Supabase, retorna um objeto vazio e o erro, que sera logado no sentry
	// Exemplo: CEP 99150000 nao se encontra no json mas está na tabela Supabase
	address, err := buscarDadosNaTabelaSupabase(cep)
	if err != nil {
		return models.CEPResponse{}, errors.New("Supabase: Erro ao acessar dados de CEP: " + err.Error())
	}

	// Se encontrou na tabela Supabase, retorna os dados
	return models.CEPResponse{
		Estado: address.Estado,
		Cidade: address.Cidade,
		Bairro: address.Bairro,
		Rua:    address.Rua,
	}, nil

}

func buscarDadosNaTabelaSupabase(cep string) (Address, error) {
	// Cliente do Supabase
	client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), nil)
	if err != nil {
		return Address{}, err
	}

	data, _, err := client.From("cep").
		Select("rua, bairro, cidade, estado", "exact", false).
		Eq("cep", cep).
		Execute()
	if err != nil {
		return Address{}, err
	}

	var addresses []Address
	err = json.Unmarshal([]byte(data), &addresses)
	if err != nil {
		return Address{}, err
	}

	// Caso retorne um array vazio, retorna um erro
	if len(addresses) == 0 {
		return Address{}, errors.New("CEP não encontrado")
	}

	// se o len de addresses for maior que 0, retorna o primeiro endereco
	return addresses[0], nil
}
