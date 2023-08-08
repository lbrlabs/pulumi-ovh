// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.Me.outputs.GetMeCurrency;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetMeResult {
    /**
     * @return Postal address of the account
     * 
     */
    private String address;
    /**
     * @return Area of the account
     * 
     */
    private String area;
    /**
     * @return City of birth
     * 
     */
    private String birthCity;
    /**
     * @return Birth date
     * 
     */
    private String birthDay;
    /**
     * @return City of the account
     * 
     */
    private String city;
    /**
     * @return This is the national identification number of the company that possess this account
     * 
     */
    private String companyNationalIdentificationNumber;
    /**
     * @return Type of corporation
     * 
     */
    private String corporationType;
    /**
     * @return Country of the account
     * 
     */
    private String country;
    private List<GetMeCurrency> currencies;
    /**
     * @return The customer code of this account (a numerical value used for identification when contacting support via phone call)
     * 
     */
    private String customerCode;
    /**
     * @return Email address
     * 
     */
    private String email;
    /**
     * @return Fax number
     * 
     */
    private String fax;
    /**
     * @return First name
     * 
     */
    private String firstname;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Italian SDI
     * 
     */
    private String italianSdi;
    /**
     * @return Preferred language for this account
     * 
     */
    private String language;
    /**
     * @return Legal form of the account
     * 
     */
    private String legalform;
    /**
     * @return Name of the account holder
     * 
     */
    private String name;
    /**
     * @return National Identification Number of this account
     * 
     */
    private String nationalIdentificationNumber;
    /**
     * @return Nic handle / customer identifier
     * 
     */
    private String nichandle;
    /**
     * @return Name of the organisation for this account
     * 
     */
    private String organisation;
    /**
     * @return OVHcloud subsidiary
     * 
     */
    private String ovhCompany;
    /**
     * @return OVHcloud subsidiary
     * 
     */
    private String ovhSubsidiary;
    /**
     * @return Phone number
     * 
     */
    private String phone;
    /**
     * @return Country code of the phone number
     * 
     */
    private String phoneCountry;
    /**
     * @return Gender of the account holder
     * 
     */
    private String sex;
    /**
     * @return Backup email address
     * 
     */
    private String spareEmail;
    /**
     * @return State of the postal address
     * 
     */
    private String state;
    /**
     * @return The resource URN of the account, to be used when writing IAM policies
     * 
     */
    private String urn;
    /**
     * @return VAT number
     * 
     */
    private String vat;
    /**
     * @return Zipcode of the address
     * 
     */
    private String zip;

    private GetMeResult() {}
    /**
     * @return Postal address of the account
     * 
     */
    public String address() {
        return this.address;
    }
    /**
     * @return Area of the account
     * 
     */
    public String area() {
        return this.area;
    }
    /**
     * @return City of birth
     * 
     */
    public String birthCity() {
        return this.birthCity;
    }
    /**
     * @return Birth date
     * 
     */
    public String birthDay() {
        return this.birthDay;
    }
    /**
     * @return City of the account
     * 
     */
    public String city() {
        return this.city;
    }
    /**
     * @return This is the national identification number of the company that possess this account
     * 
     */
    public String companyNationalIdentificationNumber() {
        return this.companyNationalIdentificationNumber;
    }
    /**
     * @return Type of corporation
     * 
     */
    public String corporationType() {
        return this.corporationType;
    }
    /**
     * @return Country of the account
     * 
     */
    public String country() {
        return this.country;
    }
    public List<GetMeCurrency> currencies() {
        return this.currencies;
    }
    /**
     * @return The customer code of this account (a numerical value used for identification when contacting support via phone call)
     * 
     */
    public String customerCode() {
        return this.customerCode;
    }
    /**
     * @return Email address
     * 
     */
    public String email() {
        return this.email;
    }
    /**
     * @return Fax number
     * 
     */
    public String fax() {
        return this.fax;
    }
    /**
     * @return First name
     * 
     */
    public String firstname() {
        return this.firstname;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Italian SDI
     * 
     */
    public String italianSdi() {
        return this.italianSdi;
    }
    /**
     * @return Preferred language for this account
     * 
     */
    public String language() {
        return this.language;
    }
    /**
     * @return Legal form of the account
     * 
     */
    public String legalform() {
        return this.legalform;
    }
    /**
     * @return Name of the account holder
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return National Identification Number of this account
     * 
     */
    public String nationalIdentificationNumber() {
        return this.nationalIdentificationNumber;
    }
    /**
     * @return Nic handle / customer identifier
     * 
     */
    public String nichandle() {
        return this.nichandle;
    }
    /**
     * @return Name of the organisation for this account
     * 
     */
    public String organisation() {
        return this.organisation;
    }
    /**
     * @return OVHcloud subsidiary
     * 
     */
    public String ovhCompany() {
        return this.ovhCompany;
    }
    /**
     * @return OVHcloud subsidiary
     * 
     */
    public String ovhSubsidiary() {
        return this.ovhSubsidiary;
    }
    /**
     * @return Phone number
     * 
     */
    public String phone() {
        return this.phone;
    }
    /**
     * @return Country code of the phone number
     * 
     */
    public String phoneCountry() {
        return this.phoneCountry;
    }
    /**
     * @return Gender of the account holder
     * 
     */
    public String sex() {
        return this.sex;
    }
    /**
     * @return Backup email address
     * 
     */
    public String spareEmail() {
        return this.spareEmail;
    }
    /**
     * @return State of the postal address
     * 
     */
    public String state() {
        return this.state;
    }
    /**
     * @return The resource URN of the account, to be used when writing IAM policies
     * 
     */
    public String urn() {
        return this.urn;
    }
    /**
     * @return VAT number
     * 
     */
    public String vat() {
        return this.vat;
    }
    /**
     * @return Zipcode of the address
     * 
     */
    public String zip() {
        return this.zip;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String address;
        private String area;
        private String birthCity;
        private String birthDay;
        private String city;
        private String companyNationalIdentificationNumber;
        private String corporationType;
        private String country;
        private List<GetMeCurrency> currencies;
        private String customerCode;
        private String email;
        private String fax;
        private String firstname;
        private String id;
        private String italianSdi;
        private String language;
        private String legalform;
        private String name;
        private String nationalIdentificationNumber;
        private String nichandle;
        private String organisation;
        private String ovhCompany;
        private String ovhSubsidiary;
        private String phone;
        private String phoneCountry;
        private String sex;
        private String spareEmail;
        private String state;
        private String urn;
        private String vat;
        private String zip;
        public Builder() {}
        public Builder(GetMeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.area = defaults.area;
    	      this.birthCity = defaults.birthCity;
    	      this.birthDay = defaults.birthDay;
    	      this.city = defaults.city;
    	      this.companyNationalIdentificationNumber = defaults.companyNationalIdentificationNumber;
    	      this.corporationType = defaults.corporationType;
    	      this.country = defaults.country;
    	      this.currencies = defaults.currencies;
    	      this.customerCode = defaults.customerCode;
    	      this.email = defaults.email;
    	      this.fax = defaults.fax;
    	      this.firstname = defaults.firstname;
    	      this.id = defaults.id;
    	      this.italianSdi = defaults.italianSdi;
    	      this.language = defaults.language;
    	      this.legalform = defaults.legalform;
    	      this.name = defaults.name;
    	      this.nationalIdentificationNumber = defaults.nationalIdentificationNumber;
    	      this.nichandle = defaults.nichandle;
    	      this.organisation = defaults.organisation;
    	      this.ovhCompany = defaults.ovhCompany;
    	      this.ovhSubsidiary = defaults.ovhSubsidiary;
    	      this.phone = defaults.phone;
    	      this.phoneCountry = defaults.phoneCountry;
    	      this.sex = defaults.sex;
    	      this.spareEmail = defaults.spareEmail;
    	      this.state = defaults.state;
    	      this.urn = defaults.urn;
    	      this.vat = defaults.vat;
    	      this.zip = defaults.zip;
        }

        @CustomType.Setter
        public Builder address(String address) {
            this.address = Objects.requireNonNull(address);
            return this;
        }
        @CustomType.Setter
        public Builder area(String area) {
            this.area = Objects.requireNonNull(area);
            return this;
        }
        @CustomType.Setter
        public Builder birthCity(String birthCity) {
            this.birthCity = Objects.requireNonNull(birthCity);
            return this;
        }
        @CustomType.Setter
        public Builder birthDay(String birthDay) {
            this.birthDay = Objects.requireNonNull(birthDay);
            return this;
        }
        @CustomType.Setter
        public Builder city(String city) {
            this.city = Objects.requireNonNull(city);
            return this;
        }
        @CustomType.Setter
        public Builder companyNationalIdentificationNumber(String companyNationalIdentificationNumber) {
            this.companyNationalIdentificationNumber = Objects.requireNonNull(companyNationalIdentificationNumber);
            return this;
        }
        @CustomType.Setter
        public Builder corporationType(String corporationType) {
            this.corporationType = Objects.requireNonNull(corporationType);
            return this;
        }
        @CustomType.Setter
        public Builder country(String country) {
            this.country = Objects.requireNonNull(country);
            return this;
        }
        @CustomType.Setter
        public Builder currencies(List<GetMeCurrency> currencies) {
            this.currencies = Objects.requireNonNull(currencies);
            return this;
        }
        public Builder currencies(GetMeCurrency... currencies) {
            return currencies(List.of(currencies));
        }
        @CustomType.Setter
        public Builder customerCode(String customerCode) {
            this.customerCode = Objects.requireNonNull(customerCode);
            return this;
        }
        @CustomType.Setter
        public Builder email(String email) {
            this.email = Objects.requireNonNull(email);
            return this;
        }
        @CustomType.Setter
        public Builder fax(String fax) {
            this.fax = Objects.requireNonNull(fax);
            return this;
        }
        @CustomType.Setter
        public Builder firstname(String firstname) {
            this.firstname = Objects.requireNonNull(firstname);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder italianSdi(String italianSdi) {
            this.italianSdi = Objects.requireNonNull(italianSdi);
            return this;
        }
        @CustomType.Setter
        public Builder language(String language) {
            this.language = Objects.requireNonNull(language);
            return this;
        }
        @CustomType.Setter
        public Builder legalform(String legalform) {
            this.legalform = Objects.requireNonNull(legalform);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder nationalIdentificationNumber(String nationalIdentificationNumber) {
            this.nationalIdentificationNumber = Objects.requireNonNull(nationalIdentificationNumber);
            return this;
        }
        @CustomType.Setter
        public Builder nichandle(String nichandle) {
            this.nichandle = Objects.requireNonNull(nichandle);
            return this;
        }
        @CustomType.Setter
        public Builder organisation(String organisation) {
            this.organisation = Objects.requireNonNull(organisation);
            return this;
        }
        @CustomType.Setter
        public Builder ovhCompany(String ovhCompany) {
            this.ovhCompany = Objects.requireNonNull(ovhCompany);
            return this;
        }
        @CustomType.Setter
        public Builder ovhSubsidiary(String ovhSubsidiary) {
            this.ovhSubsidiary = Objects.requireNonNull(ovhSubsidiary);
            return this;
        }
        @CustomType.Setter
        public Builder phone(String phone) {
            this.phone = Objects.requireNonNull(phone);
            return this;
        }
        @CustomType.Setter
        public Builder phoneCountry(String phoneCountry) {
            this.phoneCountry = Objects.requireNonNull(phoneCountry);
            return this;
        }
        @CustomType.Setter
        public Builder sex(String sex) {
            this.sex = Objects.requireNonNull(sex);
            return this;
        }
        @CustomType.Setter
        public Builder spareEmail(String spareEmail) {
            this.spareEmail = Objects.requireNonNull(spareEmail);
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder urn(String urn) {
            this.urn = Objects.requireNonNull(urn);
            return this;
        }
        @CustomType.Setter
        public Builder vat(String vat) {
            this.vat = Objects.requireNonNull(vat);
            return this;
        }
        @CustomType.Setter
        public Builder zip(String zip) {
            this.zip = Objects.requireNonNull(zip);
            return this;
        }
        public GetMeResult build() {
            final var o = new GetMeResult();
            o.address = address;
            o.area = area;
            o.birthCity = birthCity;
            o.birthDay = birthDay;
            o.city = city;
            o.companyNationalIdentificationNumber = companyNationalIdentificationNumber;
            o.corporationType = corporationType;
            o.country = country;
            o.currencies = currencies;
            o.customerCode = customerCode;
            o.email = email;
            o.fax = fax;
            o.firstname = firstname;
            o.id = id;
            o.italianSdi = italianSdi;
            o.language = language;
            o.legalform = legalform;
            o.name = name;
            o.nationalIdentificationNumber = nationalIdentificationNumber;
            o.nichandle = nichandle;
            o.organisation = organisation;
            o.ovhCompany = ovhCompany;
            o.ovhSubsidiary = ovhSubsidiary;
            o.phone = phone;
            o.phoneCountry = phoneCountry;
            o.sex = sex;
            o.spareEmail = spareEmail;
            o.state = state;
            o.urn = urn;
            o.vat = vat;
            o.zip = zip;
            return o;
        }
    }
}
