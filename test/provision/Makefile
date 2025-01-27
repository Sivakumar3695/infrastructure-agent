# Include Ansible dependencies common installation strategy
include $(CURDIR)/test/ansible/Ansible.common

TERRAFORM_DIR := $(CURDIR)/test/provision/terraform
ANSIBLE_FOLDER := $(CURDIR)/test/provision/terraform

.DEFAULT_GOAL := provision

.PHONY: terraform/backend
terraform/backend:
ifndef TERRAFORM_STATE_KEY
	$(error TERRAFORM_STATE_KEY is undefined)
endif
	sed "s/TERRAFORM_STATE_KEY/${TERRAFORM_STATE_KEY}/g" "$(TERRAFORM_DIR)/terraform.backend.tf.dist" > "$(TERRAFORM_DIR)/terraform.backend.tf"
ifndef TAG_OR_UNIQUE_NAME
	$(error TAG_OR_UNIQUE_NAME is undefined)
endif
	cp "$(TERRAFORM_DIR)/caos.auto.tfvars.dist" "$(TERRAFORM_DIR)/caos.auto.tfvars"
	sed -i -e "s/PREFIX/$(PREFIX)/g" "$(TERRAFORM_DIR)/caos.auto.tfvars"
	sed -i -e "s/TAG_OR_UNIQUE_NAME/$(TAG_OR_UNIQUE_NAME)/g" "$(TERRAFORM_DIR)/caos.auto.tfvars"

.PHONY: test/provision
test/provision: terraform/backend
ifndef PLATFORM
	$(error PLATFORM is undefined)
endif
ifndef ANSIBLE_PASSWORD_WINDOWS
	$(error ANSIBLE_PASSWORD_WINDOWS is undefined)
endif
ifndef PREFIX
	$(error PREFIX is undefined)
endif
ifndef NR_LICENSE_KEY
	$(error NR_LICENSE_KEY is undefined)
endif
ifndef INVENTORY_OUTPUT
	$(error INVENTORY_OUTPUT is undefined)
endif
	sed -i "s/PREFIX:/${PREFIX}:/g" "$(TERRAFORM_DIR)/caos.auto.tfvars"
	$(MAKE) ansible/dependencies ANSIBLE_FOLDER="$(ANSIBLE_FOLDER)"
	terraform -chdir=$(TERRAFORM_DIR) init -reconfigure && \
	TF_VAR_platform="$(PLATFORM)" \
	TF_VAR_windows_password="$(ANSIBLE_PASSWORD_WINDOWS)" \
	TF_VAR_nr_license_key="$(NR_LICENSE_KEY)" \
	TF_VAR_inventory_output="$(INVENTORY_OUTPUT)" \
	terraform -chdir=$(TERRAFORM_DIR) apply -auto-approve

.PHONY: test/provision/clean
test/provision/clean: terraform/backend ansible/clean
	terraform -chdir=$(TERRAFORM_DIR) init -reconfigure && \
	TF_VAR_nr_license_key="$(NR_LICENSE_KEY)" \
	TF_VAR_platform="$(PLATFORM)" \
	TF_VAR_windows_password="$(ANSIBLE_PASSWORD_WINDOWS)" \
	TF_VAR_inventory_output="$(INVENTORY_OUTPUT)" \
    terraform -chdir=$(TERRAFORM_DIR) destroy -auto-approve && \
	rm "$(TERRAFORM_DIR)/terraform.backend.tf" "$(TERRAFORM_DIR)/caos.auto.tfvars"
	$(MAKE) ansible/clean ANSIBLE_FOLDER=$(ANSIBLE_FOLDER)
