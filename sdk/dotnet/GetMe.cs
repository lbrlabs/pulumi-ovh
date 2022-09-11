// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Ovh
{
    public static class GetMe
    {
        /// <summary>
        /// Use this data source to get information about the current OVH account.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myaccount = Ovh.GetMe.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMeResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMeResult>("ovh:index/getMe:getMe", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetMeResult
    {
        public readonly string Address;
        public readonly string Area;
        public readonly string BirthCity;
        public readonly string BirthDay;
        public readonly string City;
        public readonly string CompanyNationalIdentificationNumber;
        public readonly string CorporationType;
        public readonly string Country;
        public readonly ImmutableArray<Outputs.GetMeCurrencyResult> Currencies;
        public readonly string CustomerCode;
        public readonly string Email;
        public readonly string Fax;
        public readonly string Firstname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ItalianSdi;
        public readonly string Language;
        public readonly string Legalform;
        public readonly string Name;
        public readonly string NationalIdentificationNumber;
        public readonly string Nichandle;
        public readonly string Organisation;
        public readonly string OvhCompany;
        public readonly string OvhSubsidiary;
        public readonly string Phone;
        public readonly string PhoneCountry;
        public readonly string Sex;
        public readonly string SpareEmail;
        public readonly string State;
        public readonly string Vat;
        public readonly string Zip;

        [OutputConstructor]
        private GetMeResult(
            string address,

            string area,

            string birthCity,

            string birthDay,

            string city,

            string companyNationalIdentificationNumber,

            string corporationType,

            string country,

            ImmutableArray<Outputs.GetMeCurrencyResult> currencies,

            string customerCode,

            string email,

            string fax,

            string firstname,

            string id,

            string italianSdi,

            string language,

            string legalform,

            string name,

            string nationalIdentificationNumber,

            string nichandle,

            string organisation,

            string ovhCompany,

            string ovhSubsidiary,

            string phone,

            string phoneCountry,

            string sex,

            string spareEmail,

            string state,

            string vat,

            string zip)
        {
            Address = address;
            Area = area;
            BirthCity = birthCity;
            BirthDay = birthDay;
            City = city;
            CompanyNationalIdentificationNumber = companyNationalIdentificationNumber;
            CorporationType = corporationType;
            Country = country;
            Currencies = currencies;
            CustomerCode = customerCode;
            Email = email;
            Fax = fax;
            Firstname = firstname;
            Id = id;
            ItalianSdi = italianSdi;
            Language = language;
            Legalform = legalform;
            Name = name;
            NationalIdentificationNumber = nationalIdentificationNumber;
            Nichandle = nichandle;
            Organisation = organisation;
            OvhCompany = ovhCompany;
            OvhSubsidiary = ovhSubsidiary;
            Phone = phone;
            PhoneCountry = phoneCountry;
            Sex = sex;
            SpareEmail = spareEmail;
            State = state;
            Vat = vat;
            Zip = zip;
        }
    }
}
