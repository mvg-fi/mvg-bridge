package constants

var (
	MTGMembers   = []string{"aca77da7-450c-4e34-867d-92ee07c1cfee", "3fb68263-4f06-476e-83db-503d25d56b93", "51186d7e-d488-417d-a031-b4e34f4fdf86", "1e92e114-bfa9-4989-b8a8-8e728bf432ef"}
	MTGThreshold = uint8(3)
)

/*
Get Mixpay payeeId

curl -i -X POST https://api.mixpay.me/v1/multisig \
  -d "receivers[]"="aca77da7-450c-4e34-867d-92ee07c1cfee" \
  -d "receivers[]"="3fb68263-4f06-476e-83db-503d25d56b93" \
  -d "receivers[]"="51186d7e-d488-417d-a031-b4e34f4fdf86" \
  -d "receivers[]"="1e92e114-bfa9-4989-b8a8-8e728bf432ef" \
  -d "threshold"=3
*/
