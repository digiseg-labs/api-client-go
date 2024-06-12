/*
Digiseg API

### Digiseg API documentation  # Introduction  This API let you harness the power of Digisegs powerful and tracking-free segmentation engine.  Audiences by Digiseg are available in 50+ countries, probablistically mapping neighborhood characteristics to the IP addresses observed on the internet - Household targeting & measurement for the post-cookie world.  ## Developer SDKs  In addition to using these APIs directly through any HTTP client, we provide a set of API client SDKs for popular programming languages:  <div class=\"api-clients\">   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-python\">     <i class=\"api-client-sdk-logo devicon-python-plain\"></i>     <p>API client for Python</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-ts\">     <i class=\"api-client-sdk-logo devicon-typescript-plain\"></i>     <p>API client for TypeScript</p>   </a>   <a class=\"api-client-box\" href=\"https://github.com/digiseg-labs/api-client-go\">     <i class=\"api-client-sdk-logo devicon-go-original-wordmark\"></i>     <p>API client for Go</p>   </a> </div> <div class=\"api-clients-breaker\" />  ## Audience taxonomy  Digiseg audiences are grouped into private and business audiences. In each group there are categories that then contain the audiences. The API endpoints that communicate audiences and household characteristics, audience codes are being used.  The following table can be used as a reference for audience codes. Note that Digiseg will at times update names of audiences for purposes of internationalization, clarity or other such purposes - but the codes will remain as-is and should be considered a stable point of reference for the audience.  | Group | Category | Audience Code | Audience Name | |-------|----------|---------------|---------------| | private | home_type | a1 | Apartment | |  |  | a2 | House | |  | savings | b1 | No Savings | |  |  | b2 | Smaller Savings | |  |  | b3 | Larger Savings | |  | lifecycle | c1 | Young couples and singles | |  |  | c2 | Early family life | |  |  | c3 | Middle-aged families | |  |  | c4 | Mature families | |  |  | c5 | Pensioners / Retirees | |  | cars | d1 | No cars | |  |  | d2 | 1 car | |  |  | d3 | 2 or more cars | |  | children | e1 | No children | |  |  | e2 | 1 child | |  |  | e3 | 2 or more children | |  | education | f1 | Basic | |  |  | f2 | Medium | |  |  | f3 | Higher | |  | neighbourhood_type | g1 | Countryside | |  |  | g2 | Village | |  |  | g3 | Suburban | |  |  | g4 | City | |  | income | h1 | Lowest 20% | |  |  | h2 | Lowest 20-40% | |  |  | h3 | Middle 40-60% | |  |  | h4 | Highest 60-80% | |  |  | h5 | Top 20% | |  | home_ownership | j1 | Rent | |  |  | j2 | Own | |  | building_age | k1 | Pre 1945 | |  |  | k2 | 1945-1989 | |  |  | k3 | 1990 until today | |  | living_space | l1 | Small | |  |  | l2 | Medium | |  |  | l3 | Large | |  | tech_level | n1 | Basic | |  |  | n2 | Medium | |  |  | n3 | High | | business | size | ba1 | Small Business | |  |  | ba2 | Medium Business | |  |  | ba3 | Larger Business |  There is also an interactive [Audience builder](https://digiseg.io/cookieless-audience-builder/) which lets you discover the targeting reach and power of combining various household characteristics into composite audiences. 

API version: 1.0.0
Contact: support@digiseg.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DayOfMonthStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DayOfMonthStats{}

// DayOfMonthStats Contains statistics about the day of the month that study activity has been measured. The time zone used to record these measurements is the time zone of the measured user, or UTC if the user's location cannot be resolved. 
type DayOfMonthStats struct {
	Var0 *MeasurementsContainer `json:"0,omitempty"`
	Var1 *MeasurementsContainer `json:"1,omitempty"`
	Var2 *MeasurementsContainer `json:"2,omitempty"`
	Var3 *MeasurementsContainer `json:"3,omitempty"`
	Var4 *MeasurementsContainer `json:"4,omitempty"`
	Var5 *MeasurementsContainer `json:"5,omitempty"`
	Var6 *MeasurementsContainer `json:"6,omitempty"`
	Var7 *MeasurementsContainer `json:"7,omitempty"`
	Var8 *MeasurementsContainer `json:"8,omitempty"`
	Var9 *MeasurementsContainer `json:"9,omitempty"`
	Var10 *MeasurementsContainer `json:"10,omitempty"`
	Var11 *MeasurementsContainer `json:"11,omitempty"`
	Var12 *MeasurementsContainer `json:"12,omitempty"`
	Var13 *MeasurementsContainer `json:"13,omitempty"`
	Var14 *MeasurementsContainer `json:"14,omitempty"`
	Var15 *MeasurementsContainer `json:"15,omitempty"`
	Var16 *MeasurementsContainer `json:"16,omitempty"`
	Var17 *MeasurementsContainer `json:"17,omitempty"`
	Var18 *MeasurementsContainer `json:"18,omitempty"`
	Var19 *MeasurementsContainer `json:"19,omitempty"`
	Var20 *MeasurementsContainer `json:"20,omitempty"`
	Var21 *MeasurementsContainer `json:"21,omitempty"`
	Var22 *MeasurementsContainer `json:"22,omitempty"`
	Var23 *MeasurementsContainer `json:"23,omitempty"`
	Var24 *MeasurementsContainer `json:"24,omitempty"`
	Var25 *MeasurementsContainer `json:"25,omitempty"`
	Var26 *MeasurementsContainer `json:"26,omitempty"`
	Var27 *MeasurementsContainer `json:"27,omitempty"`
	Var28 *MeasurementsContainer `json:"28,omitempty"`
	Var29 *MeasurementsContainer `json:"29,omitempty"`
	Var30 *MeasurementsContainer `json:"30,omitempty"`
	Var31 *MeasurementsContainer `json:"31,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DayOfMonthStats DayOfMonthStats

// NewDayOfMonthStats instantiates a new DayOfMonthStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDayOfMonthStats() *DayOfMonthStats {
	this := DayOfMonthStats{}
	return &this
}

// NewDayOfMonthStatsWithDefaults instantiates a new DayOfMonthStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDayOfMonthStatsWithDefaults() *DayOfMonthStats {
	this := DayOfMonthStats{}
	return &this
}

// GetVar0 returns the Var0 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar0() MeasurementsContainer {
	if o == nil || IsNil(o.Var0) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var0
}

// GetVar0Ok returns a tuple with the Var0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar0Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var0) {
		return nil, false
	}
	return o.Var0, true
}

// HasVar0 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar0() bool {
	if o != nil && !IsNil(o.Var0) {
		return true
	}

	return false
}

// SetVar0 gets a reference to the given MeasurementsContainer and assigns it to the Var0 field.
func (o *DayOfMonthStats) SetVar0(v MeasurementsContainer) {
	o.Var0 = &v
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar1() MeasurementsContainer {
	if o == nil || IsNil(o.Var1) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar1Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var1) {
		return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar1() bool {
	if o != nil && !IsNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given MeasurementsContainer and assigns it to the Var1 field.
func (o *DayOfMonthStats) SetVar1(v MeasurementsContainer) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar2() MeasurementsContainer {
	if o == nil || IsNil(o.Var2) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar2Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var2) {
		return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar2() bool {
	if o != nil && !IsNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given MeasurementsContainer and assigns it to the Var2 field.
func (o *DayOfMonthStats) SetVar2(v MeasurementsContainer) {
	o.Var2 = &v
}

// GetVar3 returns the Var3 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar3() MeasurementsContainer {
	if o == nil || IsNil(o.Var3) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar3Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var3) {
		return nil, false
	}
	return o.Var3, true
}

// HasVar3 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar3() bool {
	if o != nil && !IsNil(o.Var3) {
		return true
	}

	return false
}

// SetVar3 gets a reference to the given MeasurementsContainer and assigns it to the Var3 field.
func (o *DayOfMonthStats) SetVar3(v MeasurementsContainer) {
	o.Var3 = &v
}

// GetVar4 returns the Var4 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar4() MeasurementsContainer {
	if o == nil || IsNil(o.Var4) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var4
}

// GetVar4Ok returns a tuple with the Var4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar4Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var4) {
		return nil, false
	}
	return o.Var4, true
}

// HasVar4 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar4() bool {
	if o != nil && !IsNil(o.Var4) {
		return true
	}

	return false
}

// SetVar4 gets a reference to the given MeasurementsContainer and assigns it to the Var4 field.
func (o *DayOfMonthStats) SetVar4(v MeasurementsContainer) {
	o.Var4 = &v
}

// GetVar5 returns the Var5 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar5() MeasurementsContainer {
	if o == nil || IsNil(o.Var5) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var5
}

// GetVar5Ok returns a tuple with the Var5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar5Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var5) {
		return nil, false
	}
	return o.Var5, true
}

// HasVar5 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar5() bool {
	if o != nil && !IsNil(o.Var5) {
		return true
	}

	return false
}

// SetVar5 gets a reference to the given MeasurementsContainer and assigns it to the Var5 field.
func (o *DayOfMonthStats) SetVar5(v MeasurementsContainer) {
	o.Var5 = &v
}

// GetVar6 returns the Var6 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar6() MeasurementsContainer {
	if o == nil || IsNil(o.Var6) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var6
}

// GetVar6Ok returns a tuple with the Var6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar6Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var6) {
		return nil, false
	}
	return o.Var6, true
}

// HasVar6 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar6() bool {
	if o != nil && !IsNil(o.Var6) {
		return true
	}

	return false
}

// SetVar6 gets a reference to the given MeasurementsContainer and assigns it to the Var6 field.
func (o *DayOfMonthStats) SetVar6(v MeasurementsContainer) {
	o.Var6 = &v
}

// GetVar7 returns the Var7 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar7() MeasurementsContainer {
	if o == nil || IsNil(o.Var7) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var7
}

// GetVar7Ok returns a tuple with the Var7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar7Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var7) {
		return nil, false
	}
	return o.Var7, true
}

// HasVar7 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar7() bool {
	if o != nil && !IsNil(o.Var7) {
		return true
	}

	return false
}

// SetVar7 gets a reference to the given MeasurementsContainer and assigns it to the Var7 field.
func (o *DayOfMonthStats) SetVar7(v MeasurementsContainer) {
	o.Var7 = &v
}

// GetVar8 returns the Var8 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar8() MeasurementsContainer {
	if o == nil || IsNil(o.Var8) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var8
}

// GetVar8Ok returns a tuple with the Var8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar8Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var8) {
		return nil, false
	}
	return o.Var8, true
}

// HasVar8 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar8() bool {
	if o != nil && !IsNil(o.Var8) {
		return true
	}

	return false
}

// SetVar8 gets a reference to the given MeasurementsContainer and assigns it to the Var8 field.
func (o *DayOfMonthStats) SetVar8(v MeasurementsContainer) {
	o.Var8 = &v
}

// GetVar9 returns the Var9 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar9() MeasurementsContainer {
	if o == nil || IsNil(o.Var9) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var9
}

// GetVar9Ok returns a tuple with the Var9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar9Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var9) {
		return nil, false
	}
	return o.Var9, true
}

// HasVar9 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar9() bool {
	if o != nil && !IsNil(o.Var9) {
		return true
	}

	return false
}

// SetVar9 gets a reference to the given MeasurementsContainer and assigns it to the Var9 field.
func (o *DayOfMonthStats) SetVar9(v MeasurementsContainer) {
	o.Var9 = &v
}

// GetVar10 returns the Var10 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar10() MeasurementsContainer {
	if o == nil || IsNil(o.Var10) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var10
}

// GetVar10Ok returns a tuple with the Var10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar10Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var10) {
		return nil, false
	}
	return o.Var10, true
}

// HasVar10 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar10() bool {
	if o != nil && !IsNil(o.Var10) {
		return true
	}

	return false
}

// SetVar10 gets a reference to the given MeasurementsContainer and assigns it to the Var10 field.
func (o *DayOfMonthStats) SetVar10(v MeasurementsContainer) {
	o.Var10 = &v
}

// GetVar11 returns the Var11 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar11() MeasurementsContainer {
	if o == nil || IsNil(o.Var11) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var11
}

// GetVar11Ok returns a tuple with the Var11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar11Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var11) {
		return nil, false
	}
	return o.Var11, true
}

// HasVar11 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar11() bool {
	if o != nil && !IsNil(o.Var11) {
		return true
	}

	return false
}

// SetVar11 gets a reference to the given MeasurementsContainer and assigns it to the Var11 field.
func (o *DayOfMonthStats) SetVar11(v MeasurementsContainer) {
	o.Var11 = &v
}

// GetVar12 returns the Var12 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar12() MeasurementsContainer {
	if o == nil || IsNil(o.Var12) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var12
}

// GetVar12Ok returns a tuple with the Var12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar12Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var12) {
		return nil, false
	}
	return o.Var12, true
}

// HasVar12 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar12() bool {
	if o != nil && !IsNil(o.Var12) {
		return true
	}

	return false
}

// SetVar12 gets a reference to the given MeasurementsContainer and assigns it to the Var12 field.
func (o *DayOfMonthStats) SetVar12(v MeasurementsContainer) {
	o.Var12 = &v
}

// GetVar13 returns the Var13 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar13() MeasurementsContainer {
	if o == nil || IsNil(o.Var13) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var13
}

// GetVar13Ok returns a tuple with the Var13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar13Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var13) {
		return nil, false
	}
	return o.Var13, true
}

// HasVar13 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar13() bool {
	if o != nil && !IsNil(o.Var13) {
		return true
	}

	return false
}

// SetVar13 gets a reference to the given MeasurementsContainer and assigns it to the Var13 field.
func (o *DayOfMonthStats) SetVar13(v MeasurementsContainer) {
	o.Var13 = &v
}

// GetVar14 returns the Var14 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar14() MeasurementsContainer {
	if o == nil || IsNil(o.Var14) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var14
}

// GetVar14Ok returns a tuple with the Var14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar14Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var14) {
		return nil, false
	}
	return o.Var14, true
}

// HasVar14 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar14() bool {
	if o != nil && !IsNil(o.Var14) {
		return true
	}

	return false
}

// SetVar14 gets a reference to the given MeasurementsContainer and assigns it to the Var14 field.
func (o *DayOfMonthStats) SetVar14(v MeasurementsContainer) {
	o.Var14 = &v
}

// GetVar15 returns the Var15 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar15() MeasurementsContainer {
	if o == nil || IsNil(o.Var15) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var15
}

// GetVar15Ok returns a tuple with the Var15 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar15Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var15) {
		return nil, false
	}
	return o.Var15, true
}

// HasVar15 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar15() bool {
	if o != nil && !IsNil(o.Var15) {
		return true
	}

	return false
}

// SetVar15 gets a reference to the given MeasurementsContainer and assigns it to the Var15 field.
func (o *DayOfMonthStats) SetVar15(v MeasurementsContainer) {
	o.Var15 = &v
}

// GetVar16 returns the Var16 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar16() MeasurementsContainer {
	if o == nil || IsNil(o.Var16) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var16
}

// GetVar16Ok returns a tuple with the Var16 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar16Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var16) {
		return nil, false
	}
	return o.Var16, true
}

// HasVar16 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar16() bool {
	if o != nil && !IsNil(o.Var16) {
		return true
	}

	return false
}

// SetVar16 gets a reference to the given MeasurementsContainer and assigns it to the Var16 field.
func (o *DayOfMonthStats) SetVar16(v MeasurementsContainer) {
	o.Var16 = &v
}

// GetVar17 returns the Var17 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar17() MeasurementsContainer {
	if o == nil || IsNil(o.Var17) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var17
}

// GetVar17Ok returns a tuple with the Var17 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar17Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var17) {
		return nil, false
	}
	return o.Var17, true
}

// HasVar17 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar17() bool {
	if o != nil && !IsNil(o.Var17) {
		return true
	}

	return false
}

// SetVar17 gets a reference to the given MeasurementsContainer and assigns it to the Var17 field.
func (o *DayOfMonthStats) SetVar17(v MeasurementsContainer) {
	o.Var17 = &v
}

// GetVar18 returns the Var18 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar18() MeasurementsContainer {
	if o == nil || IsNil(o.Var18) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var18
}

// GetVar18Ok returns a tuple with the Var18 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar18Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var18) {
		return nil, false
	}
	return o.Var18, true
}

// HasVar18 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar18() bool {
	if o != nil && !IsNil(o.Var18) {
		return true
	}

	return false
}

// SetVar18 gets a reference to the given MeasurementsContainer and assigns it to the Var18 field.
func (o *DayOfMonthStats) SetVar18(v MeasurementsContainer) {
	o.Var18 = &v
}

// GetVar19 returns the Var19 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar19() MeasurementsContainer {
	if o == nil || IsNil(o.Var19) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var19
}

// GetVar19Ok returns a tuple with the Var19 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar19Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var19) {
		return nil, false
	}
	return o.Var19, true
}

// HasVar19 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar19() bool {
	if o != nil && !IsNil(o.Var19) {
		return true
	}

	return false
}

// SetVar19 gets a reference to the given MeasurementsContainer and assigns it to the Var19 field.
func (o *DayOfMonthStats) SetVar19(v MeasurementsContainer) {
	o.Var19 = &v
}

// GetVar20 returns the Var20 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar20() MeasurementsContainer {
	if o == nil || IsNil(o.Var20) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var20
}

// GetVar20Ok returns a tuple with the Var20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar20Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var20) {
		return nil, false
	}
	return o.Var20, true
}

// HasVar20 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar20() bool {
	if o != nil && !IsNil(o.Var20) {
		return true
	}

	return false
}

// SetVar20 gets a reference to the given MeasurementsContainer and assigns it to the Var20 field.
func (o *DayOfMonthStats) SetVar20(v MeasurementsContainer) {
	o.Var20 = &v
}

// GetVar21 returns the Var21 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar21() MeasurementsContainer {
	if o == nil || IsNil(o.Var21) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var21
}

// GetVar21Ok returns a tuple with the Var21 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar21Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var21) {
		return nil, false
	}
	return o.Var21, true
}

// HasVar21 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar21() bool {
	if o != nil && !IsNil(o.Var21) {
		return true
	}

	return false
}

// SetVar21 gets a reference to the given MeasurementsContainer and assigns it to the Var21 field.
func (o *DayOfMonthStats) SetVar21(v MeasurementsContainer) {
	o.Var21 = &v
}

// GetVar22 returns the Var22 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar22() MeasurementsContainer {
	if o == nil || IsNil(o.Var22) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var22
}

// GetVar22Ok returns a tuple with the Var22 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar22Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var22) {
		return nil, false
	}
	return o.Var22, true
}

// HasVar22 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar22() bool {
	if o != nil && !IsNil(o.Var22) {
		return true
	}

	return false
}

// SetVar22 gets a reference to the given MeasurementsContainer and assigns it to the Var22 field.
func (o *DayOfMonthStats) SetVar22(v MeasurementsContainer) {
	o.Var22 = &v
}

// GetVar23 returns the Var23 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar23() MeasurementsContainer {
	if o == nil || IsNil(o.Var23) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var23
}

// GetVar23Ok returns a tuple with the Var23 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar23Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var23) {
		return nil, false
	}
	return o.Var23, true
}

// HasVar23 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar23() bool {
	if o != nil && !IsNil(o.Var23) {
		return true
	}

	return false
}

// SetVar23 gets a reference to the given MeasurementsContainer and assigns it to the Var23 field.
func (o *DayOfMonthStats) SetVar23(v MeasurementsContainer) {
	o.Var23 = &v
}

// GetVar24 returns the Var24 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar24() MeasurementsContainer {
	if o == nil || IsNil(o.Var24) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var24
}

// GetVar24Ok returns a tuple with the Var24 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar24Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var24) {
		return nil, false
	}
	return o.Var24, true
}

// HasVar24 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar24() bool {
	if o != nil && !IsNil(o.Var24) {
		return true
	}

	return false
}

// SetVar24 gets a reference to the given MeasurementsContainer and assigns it to the Var24 field.
func (o *DayOfMonthStats) SetVar24(v MeasurementsContainer) {
	o.Var24 = &v
}

// GetVar25 returns the Var25 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar25() MeasurementsContainer {
	if o == nil || IsNil(o.Var25) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var25
}

// GetVar25Ok returns a tuple with the Var25 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar25Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var25) {
		return nil, false
	}
	return o.Var25, true
}

// HasVar25 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar25() bool {
	if o != nil && !IsNil(o.Var25) {
		return true
	}

	return false
}

// SetVar25 gets a reference to the given MeasurementsContainer and assigns it to the Var25 field.
func (o *DayOfMonthStats) SetVar25(v MeasurementsContainer) {
	o.Var25 = &v
}

// GetVar26 returns the Var26 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar26() MeasurementsContainer {
	if o == nil || IsNil(o.Var26) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var26
}

// GetVar26Ok returns a tuple with the Var26 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar26Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var26) {
		return nil, false
	}
	return o.Var26, true
}

// HasVar26 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar26() bool {
	if o != nil && !IsNil(o.Var26) {
		return true
	}

	return false
}

// SetVar26 gets a reference to the given MeasurementsContainer and assigns it to the Var26 field.
func (o *DayOfMonthStats) SetVar26(v MeasurementsContainer) {
	o.Var26 = &v
}

// GetVar27 returns the Var27 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar27() MeasurementsContainer {
	if o == nil || IsNil(o.Var27) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var27
}

// GetVar27Ok returns a tuple with the Var27 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar27Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var27) {
		return nil, false
	}
	return o.Var27, true
}

// HasVar27 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar27() bool {
	if o != nil && !IsNil(o.Var27) {
		return true
	}

	return false
}

// SetVar27 gets a reference to the given MeasurementsContainer and assigns it to the Var27 field.
func (o *DayOfMonthStats) SetVar27(v MeasurementsContainer) {
	o.Var27 = &v
}

// GetVar28 returns the Var28 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar28() MeasurementsContainer {
	if o == nil || IsNil(o.Var28) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var28
}

// GetVar28Ok returns a tuple with the Var28 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar28Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var28) {
		return nil, false
	}
	return o.Var28, true
}

// HasVar28 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar28() bool {
	if o != nil && !IsNil(o.Var28) {
		return true
	}

	return false
}

// SetVar28 gets a reference to the given MeasurementsContainer and assigns it to the Var28 field.
func (o *DayOfMonthStats) SetVar28(v MeasurementsContainer) {
	o.Var28 = &v
}

// GetVar29 returns the Var29 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar29() MeasurementsContainer {
	if o == nil || IsNil(o.Var29) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var29
}

// GetVar29Ok returns a tuple with the Var29 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar29Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var29) {
		return nil, false
	}
	return o.Var29, true
}

// HasVar29 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar29() bool {
	if o != nil && !IsNil(o.Var29) {
		return true
	}

	return false
}

// SetVar29 gets a reference to the given MeasurementsContainer and assigns it to the Var29 field.
func (o *DayOfMonthStats) SetVar29(v MeasurementsContainer) {
	o.Var29 = &v
}

// GetVar30 returns the Var30 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar30() MeasurementsContainer {
	if o == nil || IsNil(o.Var30) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var30
}

// GetVar30Ok returns a tuple with the Var30 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar30Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var30) {
		return nil, false
	}
	return o.Var30, true
}

// HasVar30 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar30() bool {
	if o != nil && !IsNil(o.Var30) {
		return true
	}

	return false
}

// SetVar30 gets a reference to the given MeasurementsContainer and assigns it to the Var30 field.
func (o *DayOfMonthStats) SetVar30(v MeasurementsContainer) {
	o.Var30 = &v
}

// GetVar31 returns the Var31 field value if set, zero value otherwise.
func (o *DayOfMonthStats) GetVar31() MeasurementsContainer {
	if o == nil || IsNil(o.Var31) {
		var ret MeasurementsContainer
		return ret
	}
	return *o.Var31
}

// GetVar31Ok returns a tuple with the Var31 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayOfMonthStats) GetVar31Ok() (*MeasurementsContainer, bool) {
	if o == nil || IsNil(o.Var31) {
		return nil, false
	}
	return o.Var31, true
}

// HasVar31 returns a boolean if a field has been set.
func (o *DayOfMonthStats) HasVar31() bool {
	if o != nil && !IsNil(o.Var31) {
		return true
	}

	return false
}

// SetVar31 gets a reference to the given MeasurementsContainer and assigns it to the Var31 field.
func (o *DayOfMonthStats) SetVar31(v MeasurementsContainer) {
	o.Var31 = &v
}

func (o DayOfMonthStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DayOfMonthStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var0) {
		toSerialize["0"] = o.Var0
	}
	if !IsNil(o.Var1) {
		toSerialize["1"] = o.Var1
	}
	if !IsNil(o.Var2) {
		toSerialize["2"] = o.Var2
	}
	if !IsNil(o.Var3) {
		toSerialize["3"] = o.Var3
	}
	if !IsNil(o.Var4) {
		toSerialize["4"] = o.Var4
	}
	if !IsNil(o.Var5) {
		toSerialize["5"] = o.Var5
	}
	if !IsNil(o.Var6) {
		toSerialize["6"] = o.Var6
	}
	if !IsNil(o.Var7) {
		toSerialize["7"] = o.Var7
	}
	if !IsNil(o.Var8) {
		toSerialize["8"] = o.Var8
	}
	if !IsNil(o.Var9) {
		toSerialize["9"] = o.Var9
	}
	if !IsNil(o.Var10) {
		toSerialize["10"] = o.Var10
	}
	if !IsNil(o.Var11) {
		toSerialize["11"] = o.Var11
	}
	if !IsNil(o.Var12) {
		toSerialize["12"] = o.Var12
	}
	if !IsNil(o.Var13) {
		toSerialize["13"] = o.Var13
	}
	if !IsNil(o.Var14) {
		toSerialize["14"] = o.Var14
	}
	if !IsNil(o.Var15) {
		toSerialize["15"] = o.Var15
	}
	if !IsNil(o.Var16) {
		toSerialize["16"] = o.Var16
	}
	if !IsNil(o.Var17) {
		toSerialize["17"] = o.Var17
	}
	if !IsNil(o.Var18) {
		toSerialize["18"] = o.Var18
	}
	if !IsNil(o.Var19) {
		toSerialize["19"] = o.Var19
	}
	if !IsNil(o.Var20) {
		toSerialize["20"] = o.Var20
	}
	if !IsNil(o.Var21) {
		toSerialize["21"] = o.Var21
	}
	if !IsNil(o.Var22) {
		toSerialize["22"] = o.Var22
	}
	if !IsNil(o.Var23) {
		toSerialize["23"] = o.Var23
	}
	if !IsNil(o.Var24) {
		toSerialize["24"] = o.Var24
	}
	if !IsNil(o.Var25) {
		toSerialize["25"] = o.Var25
	}
	if !IsNil(o.Var26) {
		toSerialize["26"] = o.Var26
	}
	if !IsNil(o.Var27) {
		toSerialize["27"] = o.Var27
	}
	if !IsNil(o.Var28) {
		toSerialize["28"] = o.Var28
	}
	if !IsNil(o.Var29) {
		toSerialize["29"] = o.Var29
	}
	if !IsNil(o.Var30) {
		toSerialize["30"] = o.Var30
	}
	if !IsNil(o.Var31) {
		toSerialize["31"] = o.Var31
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DayOfMonthStats) UnmarshalJSON(data []byte) (err error) {
	varDayOfMonthStats := _DayOfMonthStats{}

	err = json.Unmarshal(data, &varDayOfMonthStats)

	if err != nil {
		return err
	}

	*o = DayOfMonthStats(varDayOfMonthStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "0")
		delete(additionalProperties, "1")
		delete(additionalProperties, "2")
		delete(additionalProperties, "3")
		delete(additionalProperties, "4")
		delete(additionalProperties, "5")
		delete(additionalProperties, "6")
		delete(additionalProperties, "7")
		delete(additionalProperties, "8")
		delete(additionalProperties, "9")
		delete(additionalProperties, "10")
		delete(additionalProperties, "11")
		delete(additionalProperties, "12")
		delete(additionalProperties, "13")
		delete(additionalProperties, "14")
		delete(additionalProperties, "15")
		delete(additionalProperties, "16")
		delete(additionalProperties, "17")
		delete(additionalProperties, "18")
		delete(additionalProperties, "19")
		delete(additionalProperties, "20")
		delete(additionalProperties, "21")
		delete(additionalProperties, "22")
		delete(additionalProperties, "23")
		delete(additionalProperties, "24")
		delete(additionalProperties, "25")
		delete(additionalProperties, "26")
		delete(additionalProperties, "27")
		delete(additionalProperties, "28")
		delete(additionalProperties, "29")
		delete(additionalProperties, "30")
		delete(additionalProperties, "31")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDayOfMonthStats struct {
	value *DayOfMonthStats
	isSet bool
}

func (v NullableDayOfMonthStats) Get() *DayOfMonthStats {
	return v.value
}

func (v *NullableDayOfMonthStats) Set(val *DayOfMonthStats) {
	v.value = val
	v.isSet = true
}

func (v NullableDayOfMonthStats) IsSet() bool {
	return v.isSet
}

func (v *NullableDayOfMonthStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDayOfMonthStats(val *DayOfMonthStats) *NullableDayOfMonthStats {
	return &NullableDayOfMonthStats{value: val, isSet: true}
}

func (v NullableDayOfMonthStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDayOfMonthStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


