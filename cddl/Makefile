.DEFAULT_GOAL := check

SHELL := /bin/bash

%.cbor: %.diag ; diag2cbor.rb $< > $@

CDDL_FILE := cmw.cddl

DIAG_FILES := $(wildcard example*.diag)

CBOR_FILES := $(DIAG_FILES:.diag=.cbor)
PRETTY_FILES := $(CBOR_FILES:.cbor=.pretty)

CLEANFILES += $(CBOR_FILES)
CLEANFILES += $(PRETTY_FILES)

check: $(CDDL_FILE) $(CBOR_FILES)
	@for f in $(CBOR_FILES) ; do \
		echo ">> validating $$f" ; \
		cddl $(CDDL_FILE) validate $$f || exit 1; \
		echo ">> prettyfying $$f" ; \
		cbor2pretty.rb $$f > $${f%.cbor}.pretty ; \
	done

clean: ; $(RM) $(CLEANFILES)
