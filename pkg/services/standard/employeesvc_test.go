package standard_test

import (
	"context"

	errpkg "ashish.com/m/internal/errors"
	"ashish.com/m/mocks"
	"ashish.com/m/pkg/models"
	"ashish.com/m/pkg/services"
	stdsvc "ashish.com/m/pkg/services/standard"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("employee services", func() {
	var (
		ctrl              *gomock.Controller
		ctx               context.Context
		mockEmployeeStore *mocks.MockEmployeeStore
		employeeService   services.EmployeeService
		employee          models.Employee
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockEmployeeStore = mocks.NewMockEmployeeStore(ctrl)
		ctx = context.Background()
		employeeService = stdsvc.NewEmployeeService(mockEmployeeStore)
	})
	employee = models.Employee{
		Name:   "Test",
		Salary: "898989",
	}
	const validEmployeeID = "8r077e3f-76t4-76y4-er43-456yt229t886"
	Describe("get employee", func() {
		Context("get employee", func() {
			It("should return employee from db", func() {
				mockEmployeeStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&employee, nil)
				fetched, err := employeeService.GetEmployee(ctx, validEmployeeID)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(fetched).ShouldNot(BeNil())
			})
			It("should return error if employee id is empty", func() {
				emptyError := errpkg.NewEmptyError("employee id")
				fetched, err := employeeService.GetEmployee(ctx, "")

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError(emptyError))
				Expect(fetched).Should(BeNil())
			})
			It("should return error if employee is not available", func() {
				notFoundError := errpkg.NewRecordNotFoundError("employee", "search.criteria")
				mockEmployeeStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&models.Employee{}, notFoundError)

				fetched, err := employeeService.GetEmployee(ctx, validEmployeeID)
				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError(errpkg.NewResourceNotFoundError("employee", validEmployeeID)))
				Expect(fetched).Should(BeNil())
			})
		})
	})
})
