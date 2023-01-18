package repository_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/entities"
)

func TestConsistency(t *testing.T) {
	sut := repository.NewInMem()

	existingEmps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.NotNil(t, existingEmps)

	existingEmpCount := len(existingEmps)

	assert.Equal(t, 3, existingEmpCount)

	var wg sync.WaitGroup

	noOfEmpsToCreate := 100

	for i := 0; i < noOfEmpsToCreate; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newEmp := entities.Employee{Name: "Gaurav", Department: "LnD"}
			savedEmp, err := sut.Save(newEmp)

			assert.Nil(t, err)
			assert.NotNil(t, savedEmp)

			allEmps, err := sut.ListAll()

			assert.Nil(t, err)
			assert.NotNil(t, allEmps)
		}()
	}

	wg.Wait()

	allEmps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.NotNil(t, allEmps)

	assert.Equal(t, existingEmpCount+noOfEmpsToCreate, len(allEmps))
}
