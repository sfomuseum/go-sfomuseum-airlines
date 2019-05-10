package wikipedia

// if you're reading this it means the data below comes from:
// https://github.com/sfomuseum-data/sfomuseum-data-enterprise/blob/master/utils/python/wikipedia-airline-codes
// https://github.com/sfomuseum-data/sfomuseum-data-enterprise/blob/master/misc/wikipedia-airline-codes.csv
// that code should be moved to this package and the data below should be updated to reference
// SFO Museum WOF IDs in https://github.com/sfomuseum-data/sfomuseum-data-enterprise/tree/master/data
// (20190328/thisisaaronland)

const lookupTable string = `iata_code,icao_code,callsign,wikipedia_id,name
,EVY,,No._34_Squadron_RAAF,"34 Squadron, Royal Australian Air Force"
,GNL,GENERAL,135_Airways,135 Airways
1T,,,Hitit_Computer_Services,Hitit Computer Services
,WYT,WYTON,2_Sqn_No_1_Elementary_Flying_Training_School,2 Sqn No 1 Elementary Flying Training School
,TFU,THJY,213_Flight_Unit,213th Flight Unit
,CHD,CHKALOVSK-AVIA,223rd_Flight_Unit,223rd Flight Unit
,TTF,CARGO UNIT,224th_Flight_Unit,224th Flight Unit
,TWF,CLOUD RUNNER,247_Jet_Ltd,247 Jet Ltd
,SEC,SECUREX,3D_Aviation,3D Aviation
Q5,MLA,MILE-AIR,40-Mile_Air,40-Mile Air
,QRT,QUARTET,4D_Air,4D Air
,PIU,PRIMA,43_Air_School,43 Air School
4R,SEK,,Star_East_Airline,Star East Airline
4O,AIJ,ABC Aerolineas,Interjet,Interjet
7A,XRC,,Express_Air_Cargo,Express Air Cargo
7C,JJA,JEJU AIR,Jeju_Air,Jeju Air
,BRO,BROADSWORD,2Excel_Aviation,2Excel Aviation
P6,PSC,PASCAN,Pascan_Aviation,Pascan Aviation
,WSN,WINGSPAN,Advanced_Air,Advanced Air
,GBT,GLOBETROTTER,A-Jet_Aviation_Aircraft_Management,A-Jet Aviation Aircraft Management
,AJR,JET MONGOLIA,A-Jet_Aviation_Company,A-Jet Aviation Company
,SFM,AIR SAFAR,A-Safar_Air_Services,A-Safar Air Services
,AJJ,ATLANTIC JET,A2_Jet_Leasing,A2 Jet Leasing
,XXV,,AASANA,AASANA
,LEP,LECOSTA,Air_Costa,Air Costa
,NKP,ABAKAN AIR,Abakan_Air,Abakan Air
,ADD,,Advanced_Air,Advanced Air
,OWN,AERO OWEN,Aero_Owen,Aero Owen
,ASR,SOTRAVIA,Aero_Sotravia,Aero Sotravia
,ARH,AEROHELCA,Aerohelicopteros,Aerohelicopteros
,VMX,VENTA,Aeroventas_de_Mexico,Aeroventas de Mexico
,ACC,,Airspeed_Charter,Airspeed Charter
,FSY,FROSTY,Algonquin_Airlink,Algonquin Airlink
,TTX,TWISTER,Alliance_Air_Charters,Alliance Air Charters
,ALN,TOLEMAC,Alpha_Jet,Alpha Jet
,UJX,ATLAS UKRAINE,Atlas_Ukraine_Airlines,Atlas Ukraine Airlines
,RGR,AVIOR REGIONAL,Avior_Regional,Avior Regional
,NGF,ANGEL FLIGHT,Air_Charity_Network,Air Charity Network
,WFT,WORLD FLIGHT,Aircharters_Worldwide,Aircharters Worldwide
,ATT,ATTAWASOL AIR,Attawasol_Airlines,Attawasol Airlines
,EDY,STOBART,Apollo_Air_Service,Apollo Air Service
JY,IWY,ISLANDWAYS,InterCaribbean_Airways,InterCaribbean Airways
JU,ASL,AIR SERBIA,Air_Serbia,Air Serbia
QH,LYN,ALTYN AVIA,Air_Kyrgyzstan,Air Kyrgyzstan
XK,CCM,CORSICA,Air_Corsica,Air Corsica
,AHS,HIGH SKY,AHS_Air_International,AHS Air International
,ROH,,Aberdair_Aviation_Ghana,Aberdair Aviation Ghana
,ROO,AERO ROA,Aero_Roa,Aero Roa
,AWF,,Aeroforward,Aeroforward
,SUP,SUN SPEED,Aeronautical_Charters,Aeronautical Charters
,PSO,AEROPEGASO,Aerotaxis_Pegaso,Aerotaxis Pegaso
EI,EIN,SHAMROCK,Aer_Lingus,Aer Lingus
,VLB,VOLTA,Air_Volta,Air Volta
,WAS,,Air-Watania,Air-Watania
,FCJ,FRACJET,AirSprint_US,AirSprint US
,TEW,TEAMWORK,Airteam_Charter,Airteam Charter
,STT,STAR CHARTER,Alpha_Star_Charter,Alpha Star Charter
,LBZ,,Angkasa_Super_Service,Angkasa Super Service
,HEZ,ARROW,Arrow_Aviation,Arrow Aviation
,UAH,,Air_Experience_Flight,"Air Experience Flight, Cranwell"
,DRO,ADRO SERVICIOS,Adro_Servicios_Aereos,Adro Servicios Aereos
,RVQ,REVA AIR,Aero_Jet_International,Aero Jet International
,ASK,MULTISKY,Aerosky,Aerosky
,AEH,AEROCUTTER,Aero4m,Aero4m
,ERO,AEROECOM,Aeroecom,Aeroecom
A8,XAU,PEARL,Aerolink_Uganda,Aerolink Uganda
,NKY,AEROMON,Aeromonkey,Aeromonkey
,AWP,,Aeroworld_Pakistan,Aeroworld Pakistan
,AGA,GEOLINE,AG_Air,AG Air
,ABZ,ISLAND LIFEFLIGHT,Air_Ambulance_Services,Air Ambulance Services
RV,ROU,ROUGE,Air_Canada_Rouge,Air Canada Rouge
,CNM,MENGYUAN,Air_China_Inner_Mongolia,Air China Inner Mongolia
,VRE,COTE DIVORIE,Air_C%C3%B4te_d%27Ivoire,Air Côte d'Ivoire
,AWL,,Air_Walser,Air Walser
,AXY,LEGEND,Air_X_Charter,Air X Charter
,OES,ART AUSTRIA,ART_Aviation,ART Aviation
,ASF,AUSTRIAN AIRFORCE,Austrian_Air_Force,Austrian Air Force
,AVG,AVILEF,Aviation_Legacy,Aviation Legacy
7A,AZY,AZTEC WORLD,Aztec_Worldwide_Airlines,Aztec Worldwide Airlines
1B,,,Abacus_International,Abacus International
W9,AAB,ABG,Abelag_Aviation,Abelag Aviation
6U,ACX,LOADMASTER,Air_Cargo_Germany,Air Cargo Germany
,AAD,SUNRISE,Aero_Aviation_Centre_Ltd.,Aero Aviation Centre Ltd.
,SII,ASEISA,Aero_Servicios_Ejecutivos_Internacionales,Aero Servicios Ejecutivos Internacionales
,BZS,BINIZA,Aero_Biniza,Aero Biniza
,ACO,AERO COMONDU,Aero_Comondu,Aero Comondu
,AET,AERO PALMA,Aero-Palma,Aero-Palma
,ABM,ALBATROS ESPANA,Aero_Albatros_(Spanish_airline),Aero Albatros
E4,RSO,AERO ASIA,Aero_Asia_International,Aero Asia International
ZI,AAF,AIGLE AZUR,Aigle_Azur,Aigle Azur
,AAM,,Aim_Air,Aim Air
,AAO,ATLANTIS AIR,Atlantis_Airlines_(USA),Atlantis Airlines (USA)
,AAP,AEROVISTA GROUP,Aerovista_Airlines_(Kyrgyzstan),Aerovista Airlines
AE,AE,,Air_Ceylon,Air Ceylon
OZ,AAR,ASIANA,Asiana_Airlines,Asiana Airlines
4K*,AAS,AL-ASS,Askari_Aviation,Askari Aviation
,AAU,AUSTASIA,Australia_Asia_Airlines,Australia Asia Airlines
,AAV,ASTRO-PHIL,Astro_Air_International,Astro Air International
8U,AAW,AFRIQIYAH,Afriqiyah_Airways,Afriqiyah Airways
Q9,AFU,AFRINAT,Afrinat_International_Airlines,Afrinat International Airlines
,AAX,AFREX,Afric'air_Express,Afric'air Express
,BRL,BRASD'OR,Air_Brasd'or,Air Brasd'or
,AFH,FECTO,Air_Fecteau,Air Fecteau
,BRM,BOOMERANG,Air_500,Air 500
,AAG,ATLANTIC,Atlantic_Flight_Training,Atlantic Flight Training
KI,AAG,ATLANTIC,Air_Atlantique,Air Atlantique
QB,AAJ,AIR ALMA,Air_Alma,Air Alma
,ACS,,Air_Cess,Air Cess
,ADT,AIR DORVAL,Air_Dorval,Air Dorval
,AHN,AIR HUNGARIA,Air_Hungaria,Air Hungaria
,AHR,ADRIATIC,Air_Adriatic,Air Adriatic
LD,AHK,AIR HONG KONG,Air_Hong_Kong,Air Hong Kong
,AHS,AIRSAR,Air_Viggi_San_Raffaele,Air Viggi San Raffaele
,AEQ,LUNA,Air_Express,Air Express
,AAI,BOREALIS,Air_Aurora,Air Aurora
,ACU,AFRISPIRIT,Air_Cargo_Transportation_System,Air Cargo Transportation System
,ACV,,Air_Charter_Service,Air Charter Service
,ADC,ATLAN-DOMINICAN,Air_Atlantic_Dominicana,Air Atlantic Dominicana
,ADW,AIR ANDAMAN,Air_Andaman,Air Andaman
UX,AEA,EUROPA,Air_Europa,Air Europa
IG,AEY,AIR ITALY,Air_Italy_(2005–2018),Air Italy
IG,ISS,AIR ITALY,Air_Italy,Air Italy
,ASW,AIRSOUTHWEST,Air_Southwest_Ltd.,Air Southwest Ltd.
,ASX,AIRSPEC,Air_Special,Air Special
NX,AMU,AIR MACAO,Air_Macau,Air Macau
ZV,AMW,AIR MIDWEST,Air_Midwest,Air Midwest
HM,SEY,SEYCHELLES,Air_Seychelles,Air Seychelles
,SFB,AIR SOFIA,Air_Sofia,Air Sofia
,BRF,AIR BRAVO,Air_Bravo,Air Bravo
AF,AFR,AIRFRANS,Air_France,Air France
,ACG,AIR PARTNER,Air_Partner,Air Partner
SB,ACI,AIRCALIN,Air_Caledonie_International,Air Caledonie International
,VSG,VISIG,AirClass_Airways,AirClass Airways
EH,AKX,ALFA WING,Air_Nippon_Network_Co._Ltd.,Air Nippon Network Co. Ltd.
,ALM,ANTILLEAN,Air_ALM,Air ALM
,ALN,CHICAGO LINCOLN,Air_Lincoln,Air Lincoln
,ACM,WEST CAL,Air_Caledonia,Air Caledonia
,AGM,AIR GUAM,Air_Guam,Air Guam
ZW,AWI,WISCONSIN,Air_Wisconsin,Air Wisconsin
,ALU,LUXORJET,Air_Luxor_STP,Air Luxor STP
,RSI,AIR SUNSHINE,Air_Sunshine,Air Sunshine
GN,AGN,GOLF NOVEMBER,Air_Gabon,Air Gabon
,AFS,,Air_Data,Air Data
,AFV,AFRIQUE VACANCE,Air_Afrique_Vacancies,Air Afrique Vacancies
NQ,AJX,AIR JAPAN,Air_Japan,Air Japan
VD,LIB,AIR LIBERTE,Air_Libert%C3%A9,Air Liberté
TT,KLA,KAUNAS,Air_Lithuania,Air Lithuania
,AMG,AIR MINAS,Air_Minas_Linhas_A%C3%A9reas,Air Minas Linhas Aéreas
4N,ANT,AIR NORTH,Air_North_Charter,Air North Charter - Canada
,ANV,AIR NEVADA,Air_Nevada,Air Nevada
NZ,ANZ,NEW ZEALAND,New_Zealand_National_Airways_Corporation,New Zealand National Airways Corporation 
TE,ANZ,NEW ZEALAND,TEAL,Tasman Empire Airways Limited
,AOE,,Air_One_Executive,Air One Executive
,AMI,AIR MALDIVES,Air_Maldives,Air Maldives
QM,AML,MALAWI,Air_Malawi,Air Malawi
,AMN,MONTENEGRO,Air_Montenegro,Air Montenegro
ML,BIE,MEDITERRANEE,Air_Mediterranee,Air Mediterranee
P8,MKG,MEKONG,Air_Mekong,Air Mekong
,TAH,AIR MOOREA,Air_Moorea,Air Moorea
,AVZ,AIR VALENCIA,Air_Valencia,Air Valencia
,AMO,AIR MONTREAL,Air_Montreal_(Air_Holdings_Inc.),Air Montreal (Air Holdings Inc.)
BM,,,Air_Sicilia,Air Sicilia
,AMR,AIR AM,Air_Specialties_Corporation,Air Specialties Corporation
,AMS,AIR MUSKOKA,Air_Muskoka,Air Muskoka
,AOJ,ASTERIX,Avcon_Jet,Avcon Jet
,AJU,AIRJETSUL,Air_Jetsul,Air Jetsul
,AKA,,Air_Korea_Co._Ltd.,Air Korea Co. Ltd.
,LIV,LIVONIA,Air_Livonia,Air Livonia
ZX,ABL,AIRCOACH,Air_BC,Air BC
,ABN,,Air_Fret_Senegal,Air Fret Senegal
,LJA,AIR JAMAHIRIYA,Air_Jamahiriya_Company,Air Jamahiriya Company
G8,AGB,,Air_Service_Gabon,Air Service Gabon
7T,AGV,AIR GLACIERS,Air_Glaciers,Air Glaciers
,MVM,PEGASUS,Air_Cargo_America,Air Cargo America
,AMY,AIR AMBAR,Air_Ambar,Air Ambar
6V,VGA,AIR VEGAS,Air_Vegas,Air Vegas
,AOU,AIR TRACTOR,Air_Tractor,Air Tractor
,APA,CAN-AM,Air_Park_Aviation_Ltd.,Air Park Aviation Ltd.
,APG,AIR PEOPLE,Air_People_International,Air People International
NH,ANA,ALL NIPPON,All_Nippon_Airways,All Nippon Airways
,ANB,AIR NAV,Air_Navigation_And_Trading_Co._Ltd.,Air Navigation And Trading Co. Ltd.
,NGO,AIR ANGOL,Air-Angol,Air-Angol
TZ,TWG,TWINGOOSE,,air-taxi Europe
,NGP,REGAL EAGLE,Air_Nigeria,Air Nigeria
2Q,SNC,NIGHT CARGO,Air_Cargo_Carriers,Air Cargo Carriers
,SND,ARSAM,Air_Samarkand,Air Samarkand
V7,SNG,AIR SENEGAL,Air_Senegal_International,Air Senegal International
,SNY,AIR SANDY,Air_Sandy,Air Sandy
,AII,INTEGRA,Air_Integra,Air Integra
,BFF,AIR BAFFIN,Air_Baffin,Air Baffin
,BDM,BANDAMA,Air_Bandama,Air Bandama
AB,BER,AIR BERLIN,Air_Berlin,Air Berlin
,ABT,AIR BROUSSE,Air_Brousse,Air Brousse
,APV,AIR PLAN,Air_Plan_International,Air Plan International
,ARX,AIREX,"Air_Xpress,_Inc.","Air Xpress, Inc."
,HTT,HOTEL TANGO,Air_Tchad,Air Tchad
,ARZ,AIR RESORTS,Air_Resorts,Air Resorts
,ASB,AIR SPRAY,Air-Spray_1967_Ltd.,Air-Spray 1967 Ltd.
,ASC,AIR STAR,Air_Star_Corporation,Air Star Corporation
4D,ASD,AIR SINAI,Air_Sinai,Air Sinai
,AQN,BUSHAIR,Air_Queensland,Air Queensland
,ARC,,Air_Routing_International_Corp.,Air Routing International Corp.
QN,ARR,AIR ARMENIA,Air_Armenia,Air Armenia
,ABR,CONTRACT,ASL_Airlines_Ireland,ASL Airlines Ireland
,AIL,AIR ILLINOIS,Air_Illinois,Air Illinois
AI,AIC,AIRINDIA,Air_India,Air India Limited
,AIG,,Air_Inter_Gabon,Air Inter Gabon
PJ,SPM,SAINT-PIERRE,Air_Saint_Pierre,Air Saint Pierre
SZ,WOW,SWALLOW,Air_Southwest,Air Southwest
,ATJ,SNOOPY,Air_Traffic_GmbH,Air Traffic GmbH
8C,ATN,AIR TRANSPORT,Air_Transport_International,Air Transport International
,ATQ,MULTI,Air_Transport_Schiphol,Air Transport Schiphol
,ATS,,Air_Transport_Service,Air Transport Service
,AVG,DJIBOUTI FALCON,Air_Falcon,Air Falcon
,AUX,,Air_Uganda_International_Ltd.,Air Uganda International Ltd.
NF,AVN,AIR VAN,Air_Vanuatu,Air Vanuatu
ZB,BUB,BOURBON,Air_Bourbon,Air Bourbon
CC,ABD,ATLANTA,Air_Atlanta_Icelandic,Air Atlanta Icelandic
,AIE,AIR INUIT,Air_Inuit,Air Inuit
,AIS,SURESTE,Air_Sureste,Air Sureste
RB,SBK,Air Srpska,Air_Srpska,Air Srpska
TN,THT,TAHITI AIRLINES,Air_Tahiti_Nui,Air Tahiti Nui
SW,NMB,NAMIBIA,Air_Namibia,Air Namibia
,NSK,INTERSALONIKA,Air_Intersalonika,Air Intersalonika
,NTL,AIR ANATOLIA,Air_Anatolia,Air Anatolia
,SGA,AIR SAIGON,Air_Saigon,Air Saigon
,AFW,AFRAIR,Afrique_Regional_Airways,Afrique Regional Airways
AW,AFW,BLACKSTAR,Africa_World_Airlines,Africa World Airlines
,ACX,PARAIR,Air_Charters,Air Charters
PE,AEL,AIR EUROPE,Air_Europe_Italy,Air Europe Italy
JM,AJM,JAMAICA,Air_Jamaica,Air Jamaica
,AWN,AIR NIAMEY,Air_Niamey,Air Niamey
,AWT,AIR WEST,Air_West,Air West
6G,AWW,RED DRAGON,Air_Wales,Air Wales
TX,FWI,FRENCH WEST,Air_Cara%C3%AFbes,Air Caraïbes
IX,AXB,EXPRESS INDIA,Air_India_Express,Air India Express
,AXD,AIR SUDEX,Air_Express,Air Express
,BSB,ARBAS,Air_Wings,Air Wings
BT,BTI,AIRBALTIC,Air_Baltic,Air Baltic
,ANI,NIGALANTIC,Air_Atlantic_(Nig)_Limited,Air Atlantic (Nig) Limited
EL,ANK,ANK AIR,Air_Nippon,Air Nippon
YW,ANE,AIR NOSTRUM,Air_Nostrum,Air Nostrum
PX,ANG,NIUGINI,Air_Niugini,Air Niugini
G9,ABY,ARABIA,Air_Arabia,Air Arabia
AC,ACA,AIR CANADA,Air_Canada,Air Canada
AP,LAV,ALBASTAR,AlbaStar,AlbaStar
,MHS,AIR MEMPHIS,Air_Memphis,Air Memphis
XT,AXL,EXEL COMMUTER,Air_Exel,Air Exel
,AZF,AIR ZERMATT,Air_Zermatt_AG,Air Zermatt AG
UM,AZW,AIR ZIMBABWE,Air_Zimbabwe,Air Zimbabwe
,MHU,MEPHIS UGANDA,Air_Memphis_(Uganda),Air Memphis
,MKH,AIR MARRAKECH,Air_Marrakech_Service,Air Marrakech Service
,AZX,AZIMA,Air_Max_Africa,Air Max Africa
S2,RSH,SAHARA,Air_Sahara,Air Sahara
,ATA,,Air_Transport_Association,Air Transport Association
TC,ATC,TANZANIA,Air_Tanzania,Air Tanzania
,XAC,,Air_Charter_World,Air Charter World
2J,VBW,BURKINA,Air_Burkina,Air Burkina
,ATH,AIR TRAVEL,Air_Travel_Corp.,Air Travel Corp.
KM,AMC,AIR MALTA,Air_Malta,Air Malta
YT,TGA,AIR TOGO,Air_Togo,Air Togo
,ASJ,SATELLITE,Air_Satellite,Air Satellite
,ASN,,Air_and_Sea_Transport,Air and Sea Transport
,ASS,AIR CLASS,"Air_Class,_S.A._de_C.V.","Air Class, S.A. de C.V."
,NPL,AIR NEPAL,Air_Nepal_International,Air Nepal International
,NPR,,Air_Napier,Air Napier
,WAM,TAXI CARGO,Air_Taxi_&_Cargo,Air Taxi & Cargo
,RSM,AIR SOMALIA,Air_Somalia,Air Somalia
,AWZ,,AirWest,AirWest
G4,AAY,ALLEGIANT,Allegiant_Air,Allegiant Air
,AAZ,ANGUS,Angus_Aviation,Angus Aviation
,ABA,ARTEM-AVIA,Artem-Avia,Artem-Avia
,ABB,AFRICAN BUSINESS,African_Business_and_Transportations,African Business and Transportations
,ABE,ABAN,Aban_Air,Aban Air
,ABF,SKYWINGS,Aerial_Oy,Aerial Oy
,ABG,ABAKAN-AVIA,Abakan-Avia,Abakan-Avia
M3,TUS,Turismo,ABSA_Cargo_Airline,ABSA Cargo
,ABJ,,Abaet%C3%A9_Linhas_A%C3%A9reas,Abaeté Linhas Aéreas
,ABK,ALBERTA CITYLINK,Alberta_Citylink,Alberta Citylink
,ABO,AEROEXPRESO,APSA_Colombia,APSA Colombia
,ABU,AEROBUENO,AerovÃ­as_Bueno,Aerovías Bueno
,ACR,AEROCENTER,"Aerocenter,_Escuela_de_FormaciÃ³n_de_Pilotos_Privados_de_AviÃ³n","Aerocenter, Escuela de Formación de Pilotos Privados de Avión"
O4,ABV,ANTRAK,Antrak_Air,Antrak Air
GB,ABX,ABEX,Airborne_Express,Airborne Express
GB,ABX,ABEX,ABX_Air,ABX Air
,ABZ,ATA-BRAZIL,ATA_Brasil,ATA Brasil
,ACC,,Avcard_Services,Avcard Services
,ACD,ACADEMY,Academy_Airlines,Academy Airlines
8V,ACP,ASTRAL CARGO,Astral_Aviation,Astral Aviation
,ACY,ATLAS CARGOLINES,Atlas_Cargo_Airlines,Atlas Cargo Airlines
,ADA,AUSCAL,Airservices_Australia,Airservices Australia
4G*,,,Advance_Leasing_Company,Advance Leasing Company
8T,TID,,Air_Tindi,Air Tindi
,ADB,ANTONOV BUREAU,Antonov_Airlines,Antonov Airlines
,ADD,,Advanced_Air_Co.,Advanced Air Co.
,ADE,ADA AIR,Ada_Air,Ada Air
,ADG,AEREA TRAINING,Aerea_Flying_Training_Organization,Aerea Flying Training Organization
,ADI,AUDELI,Audeli_Air,Audeli Air
,ADJ,ABICAR,Abidjan_Air_Cargo,Abidjan Air Cargo
Z7*,ADK,ADCO,ADC_Airlines,ADC Airlines
,ADL,COTSWOLD,,Aero Dynamics
,ADN,AERODIENST,Aero-Dienst,Aero-Dienst
,ADP,AERODIPLOMATIC,Aerodiplomatic,Aerodiplomatic
,ADY,AERODYNE,Aerodyne_(airline),Aerodyne
,ADQ,AIR DATA,Avion_Taxi,Avion Taxi
JP,ADR,ADRIA,Adria_Airways,Adria Airways
,ADS,SONORAV,Aviones_de_Sonora,Aviones de Sonora
,ADU,AIRDEAL,Airdeal_Oy,Airdeal Oy
,ADV,ADVANCE,Advance_Air_Charters,Advance Air Charters
EM*,AEB,AEROBEN,Aero_Benin,Aero Benin
,AEC,AEROCESAR,AerocÃ©sar,Aerocésar
,AED,,Aerotrans_Airlines,Aerotrans Airlines
,ADX,ANDAX,Anderson_Aviation,Anderson Aviation
A3,AEE,AEGEAN,Aegean_Airlines,Aegean Airlines
,AEG,FUMIGACIONES SAM,Aerofumigaciones_Sam,Aerofumigaciones Sam
2K,GLG,AEROGAL,Aerogal,Aerogal
,AEI,INTERAM,Aeroexpreso_Interamericano,Aeroexpreso Interamericano
,AEJ,KHAKI EXPRESS,Air_Express,Air Express
,AEK,AEROCON,Aerocon,Aerocon
,AEM,AEROMADRID,Aero_Madrid,Aero Madrid
KD,AEN,AIR ENTERPRISE,Air_Enterprise,Air Enterprise
,AEO,AERO OCCIDENTE,Aeroservicios_Ejecutivos_Del_Occidente,Aeroservicios Ejecutivos Del Occidente
,AEP,AEROTEC,Aerotec_Escuela_de_Pilotos,Aerotec Escuela de Pilotos
KO,AER,ACE AIR,Alaska_Central_Express,Alaska Central Express
VX,AES,ACES,ACES_Colombia,ACES Colombia
KH,AAH,ALOHA,Aloha_Air_Cargo,Aloha Air Cargo
,AAK,ALASKA ISLAND,Alaska_Island_Air,Alaska Island Air
AA,AAL,AMERICAN,American_Airlines,American Airlines
AX,—,—,AmericanConnection,AmericanConnection[5]
AN,AAA,ANSETT,Ansett_Australia,Ansett Australia
,AAC,ARMYAIR,Army_Air_Corps_(United_Kingdom),Army Air Corps
5W,AEU,FLYSTAR,Astraeus_(airline),Astraeus
,AEV,AEROVENTAS,Aeroventas,Aeroventas
VV,AEW,AEROSVIT,Aerosvit_Airlines,Aerosvit Airlines
,AEX,AVCO,Airway_Express,Airway Express
,AEZ,AERIAL TRANZ,Aerial_Transit,Aerial Transit
,AFA,BLUE ALFA,Alfa_Air,Alfa Air
WK,AFB,AMERICAN FALCON,American_Falcon,American Falcon
QQ,UTY,UNITY,Alliance_Airlines,Alliance Airlines
,UVS,UNI-LEONE,Air_Universal,Air Universal
,UVT,AUVIA,Auvia_Air,Auvia Air
,AFC,AFRICAN WEST,Africa_West_Airlines,African West Air
,AFE,AIRFAST,Airfast_Indonesia,Airfast Indonesia
FG,AFG,ARIANA,Ariana_Afghan_Airlines,Ariana Afghan Airlines
RV*,AFI,AFRICAWORLD,Africaone,Africaone
Y2,AFJ,JAMBO,Alliance_Air_(Uganda),Alliance Air
,AFK,AFRICA LINKS,Africa_Air_Links,Africa Air Links
SU,AFL,AEROFLOT,Aeroflot_Russian_Airlines,Aeroflot Russian Airlines
,AFO,AERO EMPRESA,Aero_Empresa_Mexicana,Aero Empresa Mexicana
,AFQ,ALBA,Alba_Servizi_Aerotrasporti,Alba Servizi Aerotrasporti
5Z,AFX,,Airfreight_Express,Airfreight Express
,AFY,AFRICA CHARTERED,Africa_Chartered_Services,Africa Chartered Services
,AFZ,AFREIGHT,Africa_Freight_Services,Africa Freight Services
,AGA,,Aeronaves_Del_Centro,Aeronaves Del Centro
,AGC,AGRICO,Arab_Agricultural_Aviation_Company,Arab Agricultural Aviation Company
,AGF,ATLANTIC GULF,Atlantic_Gulf_Airlines,Atlantic Gulf Airlines
5D,SLI,COSTERA,Aerom%C3%A9xico_Connect,Aeroméxico Connect
,AGG,ALGOMA,Algoma_Airways,Algoma Airways
,AGH,ALTAGNA,Altagna,Altagna
,AGO,ANGOLA CHARTER,Angola_Air_Charter,Angola Air Charter
,AGP,AIR TARA,AERFI_Group,AERFI Group
,AGQ,GALASERVICE,Aerogala,Aerogala
1A,AGT,AMADEUS,Amadeus_IT_Group,Amadeus IT Group
,AGU,SARMA,Angara_Airlines,Angara Airlines
,AGW,AERO GAMBIA,Aero_Gambia,Aero Gambia
JJ,AGX,GENEX,Aviogenex,Aviogenex
,BLR,BLUE RIDGE,Atlantic_Coast_Airlines,Atlantic Coast Airlines
,BLZ,AEROLOZ,Aero_Barloz,Aero Barloz
PL,PLI,Aeroperu,Aeroper%C3%BA,Aeroperú
8A,BMM,ATLAS BLUE,Atlas_Blue,Atlas Blue
,BNB,AEROBANOBRAS,Aero_Banobras,Aero Banobras
,AGY,FLIGHT GROUP,Aero_Flight_Service,Aero Flight Service
,AGZ,AGROLET,Agrolet-Mci,Agrolet-Mci
GD,AHA,AIR ALPHA,Air_Alpha_Greenland,Air Alpha Greenland
,AHC,AZALAVIACARGO,Azal_Avia_Cargo,Azal Avia Cargo
,AHE,AIRPORT HELICOPTER,"Airport_Helicopter_Basel,_Muller_&_Co.","Airport Helicopter Basel, Muller & Co."
,CJE,BIRD JET,Aeroservices_Corporate,Aeroservices Corporate
,AHF,ASPEN,Aspen_Helicopters,Aspen Helicopters
,AHG,AEROCHAGO,Aerochago_Airlines,Aerochago Airlines
,AHH,AIRHOLD,Airplanes_Holdings,Airplanes Holdings
,AHP,AEROCHIAPAS,Aerochiapas,Aerochiapas
,AHU,ABC HUNGARY,ABC_Air_Hungary,ABC Air Hungary
HT,AHW,AEROMIST,Aeromist-Kharkiv,Aeromist-Kharkiv
J2,AHY,AZAL,Azerbaijan_Airlines,Azerbaijan Airlines
U3,AIA,AVIES,Avies,Avies
4Y,AIB,AIRBUS INDUSTRIE,Airbus_Industrie,Airbus Industrie
,AIH,AIR INCHEON,Air_Incheon,Air Incheon
RS,ASV,AIR SEOUL,Air_Seoul,Air Seoul
,AIJ,ABC AEROLINEAS,ABC_Aerol%C3%ADneas,ABC Aerolíneas
,AIK,AFRICAN AIRLINES,African_Airlines,African Airlines International Limited
,AIN,FLY CARGO,African_International_Airways,African International Airways
5A,AIP,ALPINE AIR,Alpine_Air_Express,Alpine Air Express
,AIU,ALIA,Alicante_Internacional_Airlines,Alicante Internacional Airlines
,ABP,BAIR,ABS_Jets,ABS Jets
ED*,ABQ,PAKBLUE,Airblue,Airblue
,THM,THAI AIRMARK,Airmark_Aviation,Airmark Aviation
,AIR,AIRLIFT,Airlift_International,Airlift International
,AIT,,Airest,Airest
,AIV,AIRVIAS,Airvias_S/A_Linhas_AÃ©reas,Airvias S/A Linhas Aéreas
W4,BES,BIRD EXPRESS,,Aero Services Executive
,AIW,TARTAN,Atlantic_Island_Airways,Atlantic Island Airways
,AIX,CRUISER,Aircruising_Australia,Aircruising Australia
,AIY,AIRCREW,Aircrew_Check_and_Training_Australia,Aircrew Check and Training Australia
IZ,AIZ,ARKIA,Arkia_Israel_Airlines,Arkia Israel Airlines
,AJA,AFGHAN JET,Afghan_Jet_International_Airlines,Afghan Jet International Airlines
,AJB,AERO JBR,Aero_JBR,Aero JBR
,AJE,JET EXPRESS,Aero_Jet_Express,Aero Jet Express
,AJF,AVIACONSULT,Avia_Consult_Flugbetriebs,Avia Consult Flugbetriebs
,AJI,AMERISTAR,Ameristar_Jet_Charter,Ameristar Jet Charter
,AJK,BAMBI,Allied_Air,Allied Air
,AJO,AEROEXO,Aero_Ejecutivo,Aero Ejecutivo
,AJP,AEROJETS,Aero_Jets_Corporativos,Aero Jets Corporativos
,AJS,AEROEJECUTIVOS,Aeroejecutivos_Colombia,Aeroejecutivos Colombia
M6,AJT,AMERIJET,Amerijet_International,Amerijet International
,AJV,AYJAY CARGO,ANA_%26_JP_Express,ANA & JP Express
,AJW,ALPHAJET,Alpha_Jet_International,Alpha Jet International
,AJY,AYJET,AJet,AJet
,AKB,KARAB,Aktjubavia,Aktjubavia
,AKC,ARCA,Arca_AerovÃ­as_Colombianas_Ltda.,Arca Aerovías Colombianas Ltda.
,AKF,ANIKAY,Anikay_Air_Company,Anikay Air Company
,AKH,AKHAL,Akhal,Akhal
,MNI,AEROMIL,Aeromilenio,Aeromilenio
,AKK,AKLAK,Aklak_Air,Aklak Air
4A,AKL,,Air_Kiribati,Air Kiribati
,AKN,ALKAN AIR,Alkan_Air,Alkan Air
,AKW,ANGKORWAYS,Angkor_Airways,Angkor Airways
,AKZ,ABSOLUTE,AK_Navigator_LLC,AK Navigator LLC
,ALB,ALBATROS,Aero_Albatros_(airline),Aero Albatros
,ALD,ALBION,Albion_Aviation,Albion Aviation
,ALE,AEROALAS,Aeroalas_Colombia,Aeroalas Colombia
,ALF,ACEFORCE,Allied_Command_Europe_(Mobile_Force),Allied Command Europe (Mobile Force)
,BTS,AEROLINEAS ALBATROS,Aerotaxis_Albatros,Aerotaxis Albatros
,FYS,AMERICAN FLYERS,,American Flyers
,DFA,AERO COACH,Aero_Coach_Aviation,Aero Coach Aviation
EV,ASQ,ACEY,Atlantic_Southeast_Airlines,Atlantic Southeast Airlines
,BVR,BAVARIAN,ACM_Air_Charter,ACM Air Charter
,ALG,AIRLOG,Air_Logistics,Air Logistics
,ALL,VALLARTA,Aerovallarta,Aerovallarta
HP,AWE,CACTUS,America_West_Airlines,America West Airlines
,TNO,AEROUNION,Aerotransporte_de_Carga_Union,Aerotransporte de Carga Union
,TND,TAXIS CESSNA,Aero_Taxis_Cessna,Aero Taxis Cessna
,TMP,TEMPE,Arizona_Express_Airlines,Arizona Express Airlines
,ALO,ALLEGHENY,Allegheny_Commuter_Airlines,Allegheny Commuter Airlines
,ALP,ALLPOINTS,Allpoints_Jet,Allpoints Jet
,ALP,ALPINER,Alpliner_AG,Alpliner AG
,ALQ,ALTAIR,Altair_Aviation_(1986),Altair Aviation (1986)
VH,ALV,AEROPOSTAL,Aeropostal_Alas_de_Venezuela,Aeropostal Alas de Venezuela
,ALW,ALNACIONAL,"Alas_Nacionales,_S.A.","Alas Nacionales, S.A."
,ALY,ALYESKA,Alyeska_Air_Service,Alyeska Air Service
,ALZ,,Alta_Flights_(Charters)_Ltd.,Alta Flights (Charters) Ltd.
,AMA,ADIK,ATMA_(airline),ATMA
,AMD,AEROLINEAS MEDELLIN,AerolÃ­neas_MedellÃ­n,Aerolíneas Medellín
,AMF,AMFLIGHT,Ameriflight,Ameriflight
,AMH,MANN,Alan_Mann_Helicopters_Ltd.,Alan Mann Helicopters Ltd.
,AMJ,AVIATION AMOS,Aviation_Amos,Aviation Amos
,AMK,AMER AIR,Amerer_Air,Amerer Air
,AMM,AEROM,Aeroputul_International_Marculesti,Aeroputul International Marculesti
,AMP,ATSA,Aero_Transporte_S.A._(ATSA),Aero Transporte S.A. (ATSA)
,AMQ,LIFELINE,Aeromedicare_Ltd.,Aeromedicare Ltd.
,AMQ,AMEX,Aircraft_Management_and_Consulting,Aircraft Management and Consulting
TZ,AMT,AMTRAN,ATA_Airlines,ATA Airlines
,AMV,,AMC_Airlines,AMC Airlines
AM,AMX,AEROMEXICO,Aerom%C3%A9xico,Aeroméxico
,AMZ,AMIYA AIR,Amiya_Airline,Amiya Airline
,ANC,ANGLO,Anglo_Cargo,Anglo Cargo
,BRP,AEROBRA,AeroBratsk,AeroBratsk
,ANH,ALAJNIHAH,Alajnihah_for_Air_Transport,Alajnihah for Air Transport
,ANM,NORAM,Aerotransportacion_de_Norteamerica,Aerotransportacion de Norteamerica
,ANM,ANTARES,"Antares_Airtransport,_Maintenance_&_Service_GmbH","Antares Airtransport, Maintenance & Service GmbH"
TL,ANO,TOPEND,Airnorth,Airnorth
,ANQ,ANTIOQUIA,Aerol%C3%ADnea_de_Antioquia,Aerolínea de Antioquia
OY,ANS,AEROANDES,Andes_L%C3%ADneas_A%C3%A9reas,Andes Líneas Aéreas
IW,AOM,French Lines,AOM_French_Airlines,AOM French Airlines
,ANW,AVINOR,"Aviaci%C3%B3n_Del_Noroeste,_S.A._de_C.V.","Aviación Del Noroeste, S.A. de C.V."
,SAP,TOBOL,Avia_Jaynar,Avia Jaynar
,EMS,SERVIEMPRESARIAL,Aero_Servicios_Empresariales,Aero Servicios Empresariales
,AOA,ALCON,"Alcon_Servicios_AÃ©reos,_S.A._de_C.V.","Alcon Servicios Aéreos, S.A. de C.V."
J6,AOC,AERO AVCOM,AVCOM,AVCOM
,AOD,AERO CZECH,Aero_Vodochody,Aero Vodochody
,AOF,ATAIR,Atair_Pty_Ltd.,Atair Pty Ltd.
2D,AOG,AVIP,Aero_VIP_(Argentina),Aero VIP
,MUN,AEROMUNDO,Aeromundo_Ejecutivo,Aeromundo Ejecutivo
,MUR,MURI,AerolÃ­nea_Muri,Aerolínea Muri
,AOI,ASTORIA,"Astoria,_Inc.","Astoria, Inc."
,NRO,AEROMASTER,Aero_Rent,Aero Rent JSC
,NRP,AERONORD,Aeronord-Grup,Aeronord-Grup
,AOK,,Aeroatlantico_Colombia,Aeroatlantico Colombia
,AOL,ANGKOR AIR,Angkor_Airlines,Angkor Airlines
,AON,AERO ENTERPRISE,Aero_Entreprise,Aero Entreprise
,AOO,COMPANY AS,"As,_Opened_Joint_Stock_Company","As, Opened Joint Stock Company"
,AOP,AEROPILOTO,Aeropiloto,Aeropiloto
VB,VIV,AEROENLACES,Aeroenlaces_Nacionales,Aeroenlaces Nacionales
,VIZ,AEROVIZ,Aerovis_Airlines,Aerovis Airlines
,VJE,,AvJet_Routing,AvJet Routing
,VGF,VISTA GULF,Aerovista,Aerovista Gulf Express
,VER,ALMAVER,Almaver,Almaver
,AOR,INTER-AFRO,Afro_International_Ent._Limited,Afro International Ent. Limited
,SMX,ALIEXPRESS,Alitalia_Express,Alitalia Express
OE,AOT,ASIA OVERNIGHT,Asia_Overnight_Express,Asia Overnight Express
,AOV,AEROVISION,Aero_Vision,Aero Vision
,AOX,AEROVALLE,Aerotaxi_Del_Valle,Aerotaxi Del Valle
,APC,AIRPAC,"Airpac_Airlines,_Inc.","Airpac Airlines, Inc."
,SVM,SERVIMONTE,Aeroservicios_Monterrey,Aeroservicios Monterrey
,APF,AMAPOLA,Amapola_Flyg_AB,Amapola Flyg AB
,APH,AIRFLIGHT,"Alpha_Aviation,_Inc.","Alpha Aviation, Inc."
,API,ASA PESADA,"ASA_Pesada,_Lda.","ASA Pesada, Lda."
,APJ,AIR PRINT,"Air_Print,_S.A.","Air Print, S.A."
,PET,AEROPETRO,Aerotransporte_Petrolero,Aerotransporte Petrolero
GV,ARF,Aero Fox,Aero_Flight,Aero Flight
,BJT,BAY JET,ACM_Aviation,ACM Aviation
,BKL,BARCOL,Aircompany_Barcol,Aircompany Barcol
,BLA,ALL CHARTER,All_Charter_Limited,All Charter Limited
,APL,APPALACHIAN,"Appalachian_Flying_Service,_Inc.","Appalachian Flying Service, Inc."
,APM,ALASKA PACIFIC,"Airpac,_Inc.","Airpac, Inc."
,APO,AEROPRO,Aeropro,Aeropro
,APP,AEROPERLAS,"AerolÃ­neas_PacÃ­fico_AtlÃ¡ntico,_S.A._(Apair)","Aerolíneas Pacífico Atlántico, S.A. (Apair)"
,APQ,ASPEN BASE,Aspen_Aviation,Aspen Aviation
,APU,AEROPUMA,"Aeropuma,_S.A.","Aeropuma, S.A."
JW,APW,BIG A,Arrow_Air,Arrow Air
,APX,PARCEL EXPRESS,Apex_Air_Cargo,Apex Air Cargo
,APY,APA INTERNACIONAL,APA_Internacional,APA Internacional
,AQA,ATCO,"Aeroatlas,_S.A.","Aeroatlas, S.A."
,AQL,AQUILA,Aquila_Air_Ltd.,Aquila Air Ltd.
,AQO,ALCOA SHUTTLE,Aluminum_Company_Of_America,Aluminum Company Of America
,AQT,AVIOQUINTANA,"Aviones_de_Renta_de_Quintana_Roo,_S.A._de_C.V.","Aviones de Renta de Quintana Roo, S.A. de C.V."
,AQU,QUARIUS,AirQuarius_Aviation,AirQuarius Aviation
,AQZ,QUANZA,Aerodyne_Charter_Company,Aerodyne Charter Company
,ARA,ARIK AIR,Arik_Air,Arik Air
,ARB,AVIAIR,Avia_Air_N.V.,Avia Air N.V.
2B,AWT,ALBAWINGS,Albawings,Albawings
4C,ARE,AIRES,"Aires,_Aerov%C3%ADas_de_Integraci%C3%B3n_Regional,_S.A.","Aires, Aerovías de Integración Regional, S.A."
AR,ARG,ARGENTINA,Aerol%C3%ADneas_Argentinas,Aerolíneas Argentinas
,ARH,ARROWHEAD,Arrowhead_Airways,Arrowhead Airways
,ARI,AEROVICS,Aero_Vics,Aero Vics
,ARJ,,"Aerojet_de_Costa_Rica,_S.A.","Aerojet de Costa Rica, S.A."
,SUN,,Antillana_De_NavegaciÃ³n_AÃ©rea,Antillana De Navegación Aérea
,SUO,SERVICIO SANLUIS,Aeroservicios_De_San_Luis,Aeroservicios De San Luis
,SUP,AEROSUPER,Aerosuper,Aerosuper
,ARK,LINK SERVICE,Aero_Link_Air_Services_S.L.,Aero Link Air Services S.L.
,ARL,AIRLEC,Airlec_-_Air_Aquitaine_Transport,Airlec - Air Aquitaine Transport
,ARM,AMEX,Aeromarket_Express,Aeromarket Express
,ARO,ARROW,Arrow_Aviation_Ltd.,Arrow Aviation Ltd.
,KLD,AIR KLAIPEDA,Air_KlaipÄda,Air Klaipėda
,ARQ,ARMSTRONG,"Armstrong_Air,_Inc.","Armstrong Air, Inc."
,ARS,METSERVICE,Aeromet_Servicios,Aeromet Servicios
,ART,AEROTAL,Aerotal_AerolÃ­neas_Territoriales_de_Colombia_Ltda.,Aerotal Aerolíneas Territoriales de Colombia Ltda.
,ARV,ARAVCO,Aravco_Ltd.,Aravco Ltd.
,ARW,ARIABIRD,Aria_(airline),Aria
,OST,ALANIA,Airline_Alania,Airline Alania
,HUC,LINEAS TEHUACAN,AerolÃ­neas_de_TechuacÃ¡n,Aerolíneas de Techuacán
,HUT,AEROHUITZILIN,Aerotransportes_Huitzilin,Aerotransportes Huitzilin
,HUY,AERO HUMAYA,Aero_Transportes_Del_Humaya,Aero Transportes Del Humaya
,ARY,GOSEY,Argosy_Airways,Argosy Airways
AS,ASA,ALASKA,"Alaska_Airlines,_Inc.","Alaska Airlines, Inc."
PL,ASE,MOROZOV,Airstars,Airstars
,ASF,SCHEFF,"Air_Schefferville,_Inc.","Air Schefferville, Inc."
,ASG,AFRICAN STAR,African_Star_Airways_(PTY)_Ltd.,African Star Airways (PTY) Ltd.
,ASI,AEROSUN,"Aerosun_International,_Inc.","Aerosun International, Inc."
,ASM,AWESOME,Awesome_Flight_Services_(PTY)_Ltd.,Awesome Flight Services (PTY) Ltd.
,ASO,AERO NITRA,Aero_Slovakia,Aero Slovakia
,ASP,AIRSPRINT,Airsprint,Airsprint
,ASR,ALL STAR,"All_Star_Airlines,_Inc.","All Star Airlines, Inc."
,AST,AEROESTE,AerolÃ­neas_Del_Oeste,Aerolíneas Del Oeste
,NRE,AVIONES ARE,Aviones_Are,Aviones Are
,WAP,ARROW PANAMA,Arrow_Panama,Arrow Panama
,ASV,ASTRAVIA,Astravia-Bissau_Air_Transports_Ltd.,Astravia-Bissau Air Transports Ltd.
OB,ASZ,AIR ASTRAKHAN,Astrakhan_Airlines,Astrakhan Airlines
,ATB,STARLITE,Atlantair_Ltd.,Atlantair Ltd.
,ATD,AEROTOURS,Aerotours_Dominicana,Aerotours Dominicana
,ATE,ATLANTIS CANADA,"Atlantis_Transportation_Services,_Ltd.","Atlantis Transportation Services, Ltd."
,ATG,BACHYT,Aerotrans_Airline,Aerotrans Airline
HC,ATI,,Aero-Tropics_Air_Services,Aero-Tropics Air Services
,ATK,AEROTACA,AeroTACA,AeroTACA
FO,ATM,AIRTAS,Airlines_of_Tasmania,Airlines of Tasmania
,CPV,AIRCORPORATE,Air_Corporate,Air Corporate
,ATP,ASTRAL,ASTRAL_Colombia_-_Aerotransportes_Especiales_Ltda.,ASTRAL Colombia - Aerotransportes Especiales Ltda.
,FEO,FERINCO,Aeroferinco,Aeroferinco
,FES,AERO ALFE,Aero_Taxis_Y_Servicios_Alfe,Aero Taxis Y Servicios Alfe
,FFA,AVIALESOOKHRANA,Avialesookhrana,Avialesookhrana
,FFB,FOXTROT FOXTROT,Africair_Service,Africair Service
,ATR,ATLAS-AIR,Atlas_Airlines,Atlas Airlines
,ATT,AERTURAS,Aer_Turas,Aer Turas
,ATU,ATLANT HUNGARY,Atlant_Aerobatics_Ltd.,Atlant Aerobatics Ltd.
,ATV,AVANTI AIR,Avanti_Air,Avanti Air
,ATW,AERO TRADES,Aero_Trades_(Western)_Ltd.,Aero Trades (Western) Ltd.
OS,AUA,AUSTRIAN,Austrian_Airlines,Austrian Airlines
IQ,AUB,AUGSBURG-AIR,Augsburg_Airways,Augsburg Airways
,TUP,TUPOLEVAIR,Aviastar-Tu,Aviastar-Tu
RU,ABW,AIRBRIDGE CARGO,AirBridge_Cargo,AirBridge Cargo
,TUR,,ATUR,ATUR
,TXU,ATESA,ATESA_Aerotaxis_Ecuatorianos,ATESA Aerotaxis Ecuatorianos
,TXX,COWBOY,Austin_Express,Austin Express
,AUD,AUDI AIR,"Audi_Air,_Inc.","Audi Air, Inc."
,AUF,AUGUSTA,Augusta_Air_Luftfahrtunternehmen,Augusta Air Luftfahrtunternehmen
MO,AUH,SULTAN,Abu_Dhabi_Amiri_Flight,Abu Dhabi Amiri Flight
,AUM,ATLAMUR,Air_Atlantic_Uruguay,Air Atlantic Uruguay
,AUN,AVIONES UNIDOS,Aviones_Unidos,Aviones Unidos
,AUP,,Avia_Business_Group,Avia Business Group
,SVE,AEROESPECIAL,Aero_Servicios_Expecializados,Aero Servicios Expecializados
GR,AUR,AYLINE,Aurigny_Air_Services,Aurigny Air Services
NO,AUS,,Aus-Air,Aus-Air
AU,AUT,AUSTRAL,Austral_L%C3%ADneas_A%C3%A9reas,Austral Líneas Aéreas
,AUU,AURORA AIR,"Aurora_Aviation,_Inc.","Aurora Aviation, Inc."
,AUY,AUSA,"AerolÃ­neas_Uruguayas,_S.A.","Aerolíneas Uruguayas, S.A."
AO,AUZ,AUSTRALIAN,Australian_Airlines,Australian Airlines
AV,AVA,AVIANCA,Avianca,Avianca
A0,MCJ,JETMAC,Avianca_Argentina,Avianca Argentina
O6,ONE,OCEAN AIR,Avianca_Brazil,Avianca Brazil
,AVB,BEAUPAIR,Aviation_Beauport,Aviation Beauport
,AVF,CARIBOO,Aviair_Aviation_Ltd.,Aviair Aviation Ltd.
,AVH,KENT HELI,AV8_Helicopters_(UK),AV8 Helicopters
,AVJ,ATOMIC,Avia_Traffic_Company,Avia Traffic Company
,AVK,AVIATE-COPTER,AV8_Helicopters,AV8 Helicopters
,AVM,AVEMEX,"AviaciÃ³n_Ejecutiva_Mexicana,_S.A.","Aviación Ejecutiva Mexicana, S.A."
,AVO,AVIATION WORK,Aviation_at_Work,Aviation at Work
,AVP,AVCORP,Avcorp_Registrations,Avcorp Registrations
,AVP,AVIA PUEBLA,Aviacion_Corporativa_de_Peubla,Aviacion Corporativa de Peubla
,LFP,ALFA-SPACE,Alfa_Aerospace,Alfa Aerospace
,LFR,LANFREIGHT,Atlantic_Airfreight_Aviation,Atlantic Airfreight Aviation
,LFT,LIFTCO,Aerolift_Company,Aerolift Company
,AVQ,AQUILINE,"Aviation_Services,_Inc.","Aviation Services, Inc."
,AVR,ACTIVE AERO,"Active_Aero_Charter,_Inc.","Active Aero Charter, Inc."
,AVS,AVIALSA,Avialsa_T-35,Avialsa T-35
,AVT,ASIAVIA,Asia_Avia_Airlines,Asia Avia Airlines
,AVU,AVIASUD,Avia_Sud_AÃ©rotaxi,Avia Sud Aérotaxi
,AVV,AIRVANTAGE,Airvantage_Incorporated,Airvantage Incorporated
,AVW,AVIATOR,Aviator_Airways,Aviator Airways
,AVX,PASLAUGA,Aviapaslauga,Aviapaslauga
,XXX,,,ASL (Air Service Liege)
,YRG,YUGAIR,Air_Yugoslavia,Air Yugoslavia
K8,ZAK,,Airlink_Zambia,Airlink Zambia
,ZZM,,ANAM,Agence Nationale des Aerodromes et de la Meteorologie
4Y,BGA,BELUGA,Airbus_Transport_International,Airbus Transport International
B9,BGD,AIR BANGLA,Air_Bangladesh,Air Bangladesh
,BGF,BULGARIAN,Aviodetachment-28,Aviodetachment-28
,BGG,AERO BG,Aero_BG,Aero BG
,BHC,BAHIA,Aerotaxis_De_La_Bahia,Aerotaxis De La Bahia
,BIV,AVIASERVICE,Aviaservice,Aviaservice
,SZA,AESA,Aerol%C3%ADneas_de_El_Salvador,Aerolíneas de El Salvador
,AVY,AEROVARADERO,"Aerovaradero,_S.A.","Aerovaradero, S.A."
,AWB,AIRNAT,"Airways_International,_Inc.","Airways International, Inc."
,AWK,AIRWORK,Airwork,Airwork
,AWL,AUSSIEWORLD,Australian_Wetleasing,Australian Wetleasing
,AWO,AWOOD AIR,Awood_Air_Ltd.,Awood Air Ltd.
,AWR,ARCTIC WINGS,Arctic_Wings_And_Rotors_Ltd.,Arctic Wings And Rotors Ltd.
,ISM,STORK,Auo_Airclub_AIST-M,Auo Airclub AIST-M
,AWS,ARAB WINGS,Arab_Wings,Arab Wings
,AWV,AIRWAVE,"Airwave_Transport,_Inc.","Airwave Transport, Inc."
,AWY,AEROWEE,"Aeroway,_S.L.","Aeroway, S.L."
HJ,AXF,FREIGHTEXPRESS,Asian_Express_Airlines,Asian Express Airlines
,AXH,AEROMEXHAGA,Aeromexhaga,Aeromexhaga
,AXI,AIR FREIGHTER,"Aeron_International_Airlines,_Inc.","Aeron International Airlines, Inc."
,AXK,EXPRESS JET,African_Express_Airways,African Express Airways
AK,AXM,RED CAP,AirAsia,AirAsia
D7,XAX,XANADU,AirAsia_X,AirAsia X
DJ,WAJ,WING ASIA,AirAsia_Japan,AirAsia Japan
I5,IAD,ARIYA,AirAsia_India,AirAsia India
,AXN,ALEXANDROS,Alexandair,Alexandair
,AXP,AEROMAX SPAIN,Aeromax,Aeromax
,AXQ,ACTION AIR,Action_Airlines_(Action_Air_Charter),Action Airlines (Action Air Charter)
,BNI,ALBERNI,Alberni_Airways,Alberni Airways
,BNZ,AERO BONANZA,AerolÃ­neas_Bonanza,Aerolíneas Bonanza
,BOC,AEROBONA,Aerobona,Aerobona
,BOI,ABAIR,Aboitiz_Air,Aboitiz Air
,AXR,RENTAXEL,"Axel_Rent,_S.A.","Axel Rent, S.A."
,AXS,ALTUS,Altus_Airlines,Altus Airlines
,AXV,AXAVIA,Aviaxess,Aviaxess
,AXX,IMPEX,Avioimpex_A.D.p.o.,Avioimpex A.D.p.o.
6V,AXY,AXIS,Axis_Airways,Axis Airways
,AYD,AIRLINES ALADIA,Aladia_Airlines,Aladia Airlines
,AYK,,Aykavia_Aircompany,Aykavia Aircompany
,AYM,AIRMAN,"Airman,_S.L.","Airman, S.L."
,NPT,NEPTUNE,Atlantic_Airlines_(United_Kingdom),Atlantic Airlines
,GBN,ATLANTIC GABON,Atlantic_Airlines_(Gabon),Atlantic Airlines
EX,BJK,BLACKJACK,Atlantic_Airlines_(United_States),Atlantic Airlines
,HHA,ATLANTIC HONDURAS,Atlantic_Airlines_de_Honduras,Atlantic Airlines de Honduras
,AYN,ATLANTIC NICARAGUA,"Atlantic_Airlines,_S.A.","Atlantic Airlines, S.A."
,AYS,,Awsaj_Aviation_Services,Awsaj Aviation Services
,AYT,AYEET,Ayeet_Aviation_%26_Tourism,Ayeet Aviation & Tourism
3G,AYZ,ATLANT-SOYUZ,Atlant-Soyuz_Airlines,Atlant-Soyuz Airlines
AZ,AZA,ALITALIA,Alitalia,Alitalia
,AZB,TUMARA,Azamat_(airline),Azamat
ZE,AZE,ARCUS AIR,Arcus-Air_Logistic,Arcus-Air Logistic
A2,AZI,ASTRA,Astra_Airlines,Astra Airlines
,AZK,AZALHELICOPTER,Azalhelikopter,Azalhelikopter
,AZL,SKY AFRICA,Africa_One,Africa One
,AZM,AEROCOZUMEL,Aerocozumel,Aerocozumel
,AZP,ARIZONA PACIFIC,Arizona_Pacific_Airways,Arizona Pacific Airways
,AZS,ZITOTRANS,Aviacon_Zitotrans_Air_Company,Aviacon Zitotrans Air Company
,AZT,AZIMUT,"Azimut,_S.A.","Azimut, S.A."
,AZV,AZOV AVIA,Azov_Avia_Airlines,Azov Avia Airlines
,MHC,AERO JOMACHA,Aero_Jomacha,Aero Jomacha
,AZY,ARIZAIR,"Arizona_Airways,_Inc.","Arizona Airways, Inc."
,AZZ,AZZA TRANSPORT,Azza_Transport,Azza Transport
,NAR,NIGHT AIR,Air_Continental_Inc,Air Continental Inc
,NAU,ANTANIK,Antanik-Air,Antanik-Air
,NER,NEWAIR,Air_Newark,Air Newark
,NFF,,Aircraft_Support_and_Services,Aircraft Support and Services
,OBA,AEROBANANA,Aerobanana,Aerobanana
,OBK,AMAKO AIR,Amako_Airlines,Amako Airlines
R7,OCA,AROSCA,Aserca_Airlines,Aserca Airlines
,NFS,,Afrique_CArgo_Service_Senegal,Afrique CArgo Service Senegal
,NGC,ANGOSERVICE,Angoservice,Angoservice
,NGE,ANGEL AIR,Angel_Airlines_(Thailand),Angel Airlines
,NGF,ANGEL FLIGHT,Angel_Flight_America,Angel Flight America
,OUL,CITY EXPRESS,Air_Atonabee,Air Atonabee
,OVA,AERONOVA,Aero_Nova,Aero Nova
,XPE,EXPERT,Amira_Air,Amira Air
,XSS,INTER EXPRESS,Aero_Express_Intercontinental,Aero Express Intercontinental
,XTJ,,Advance_Aviation_Services,Advance Aviation Services
,TLR,AIR LIBYA,Air_Libya_Tibesti,Air Libya Tibesti
,OVC,,Aerovic,Aerovic
,RVE,AIRVENTURE,Airventure,Airventure
,RVI,AERO SERVICIOS,Aero_Servicios,Aero Servicios
,RVL,AIR VALLEE,Airvallee,Airvallee
,OVE,AEROMOVER,Aeromover,Aeromover
,OVI,VIAS EJECUTIVAS,AerovÃ­as_Ejecutivas,Aerovías Ejecutivas
,PTD,PITY,Aero_Servicio_Pity,Aero Servicio Pity
,PTE,AERO-COP,Aero_Copter,Aero Copter
,PLL,AIRPAL,Air_Pal,Air Pal
,PLM,PULLMANTUR,Air_Pullmantur,Air Pullmantur
RX,AEH,AVEX,Aviaexpress,Aviaexpress
,PSG,SERVIAVIONES,Aviones_Para_Servirle,Aviones Para Servirle
,SLU,AVIO SLUZBA,Avio_Sluzba,Avio Sluzba
,SCU,SCORPIO UNIVERS,Air_Scorpio,Air Scorpio
,SIP,AIR SPIRIT,Air_Spirit,Air Spirit
,BMV,OLIGA,Alatau_Airlines,Alatau Airlines
,GUG,AVIATECA,Aviateca,Aviateca
,PXX,PINE STATE,Aroostook_Aviation,Aroostook Aviation
,PYC,AEROPYCSA,Aeropycsa,Aeropycsa
,PVK,BORIS,Association_of_Private_Pilots_of_Kazakhstan,Association of Private Pilots of Kazakhstan
,BAS,AEROSERV,Aero_Services,Aero Services
,BBE,BABEL AIR,Ababeel_Aviation,Ababeel Aviation
,BBT,AGYDAL,Air_Bashkortostan,Air Bashkortostan
,MCY,MERCY,Ambulance_Air_Africa,Ambulance Air Africa
MQ,EGF,EAGLE FLIGHT,Envoy_Air,American Eagle Airlines
,PUE,PUELCHE,Aeropuelche,Aeropuelche
,PUT,,Aeroput,Aeroput
ZS,AZI,AZZURRA,Azzurra_Air,Azzurra Air
FF,,,Airshop,Airshop
ML,ETC,TRANATTICO,African_Transport_Trading_and_Investment_Company,African Transport Trading and Investment Company
,VUE,FLIGHTVUE,AD_Aviation,AD Aviation
,XCT,AEROCOSTAXI,Aero_Costa_Taxi_AÃ©reo,Aero Costa Taxi Aéreo
,VRO,AEROVITRO,Aerovitro,Aerovitro
,VRI,VILLARICA,Aerotaxi_Villa_Rica,Aerotaxi Villa Rica
,VEG,AEROVEGA,Aerovega,Aerovega
,VVG,AEROVILLA,Aerovilla,Aerovilla
,VLR,VILLAVERDE,AerolÃ­neas_Villaverde,Aerolíneas Villaverde
,WIL,WILLIAMETTE,Aero_Air,Aero Air
,VEJ,VENEJECUTIV,Aero_Ejecutivos,Aero Ejecutivos
,WAB,WABASH,Aero_Industries_Inc,Aero Industries Inc
,VNG,VANGUARDIA,Aero_Servicios_Vanguardia,Aero Servicios Vanguardia
,VAD,VALLES,Aero_Taxi_Los_Valles,Aero Taxi Los Valles
,VMR,AERO VILAMOURA,Aero_Vilamoura,Aero Vilamoura
,VLS,VIREL,Aero_Virel,Aero Virel
,XAA,,Aeronautical_Radio_Inc,Aeronautical Radio Inc
,VUO,AEROVUELOX,Aerovuelox,Aerovuelox
,VTM,AERONAVES TSM,Aeronaves_TSM,Aeronaves TSM
VU,VUN,AIRIVOIRE,Air_Ivoire,Air Ivoire
BP,BOT,BOTSWANA,Air_Botswana,Air Botswana
,XPR,,Air-Rep,Air-Rep
,XLL,TINGA-TINGA,Air_Excel,Air Excel
,VAE,AIR-EVANS,Air_Evans,Air Evans
,WHY,AIR SOREL,Air_Sorel,Air Sorel
,WDR,WIND RIDER,Air_Net_Private_Charter,Air Net Private Charter
,XEC,,Air_Executive_Charter,Air Executive Charter
GS,UPA,FOYLE,Air_Foyle,Air Foyle
,VTY,VICTORY,Air_Midwest_(Nigeria),Air Midwest (Nigeria)
VT,VTA,AIR TAHITI,Air_Tahiti,Air Tahiti
3N,URG,URGA,Air_Urga,Air Urga
,VDR,VARDAR,Air_Vardar,Air Vardar
VL,VIM,CRYSTAL,Air_VIA,Air VIA
,WLR,AIRWALSER,Air_Walser,Air Walser
,URA,ROSAVIA,Aircompany_Rosavia,Aircompany Rosavia
,XLB,,Aircraft_Performance_Group,Aircraft Performance Group
,WLA,AIRLIMITED,Airwaves_Airlink,Airwaves Airlink
,XFX,AIRCORP,Airways_Corporation_of_New_Zealand,Airways Corporation of New Zealand
,WAY,GARONNE,Airways_(airline),Airways
,WGS,AIRWINGS,Airwings_oy,Airwings oy
,XAK,SUNEXPRESS,Airkenya,Airkenya
,WPK,WOLFPACK,Air-Lift_Associates,Air-Lift Associates
,VAB,,Airtrans_Ltd,Airtrans Ltd
,URP,AIR-ARP,ARP_410_Airlines,ARP 410 Airlines
,WPR,WESTPAC RESCUE,Auckland_Regional_Rescue_Helicopter_Trust,Auckland Regional Rescue Helicopter Trust
,URR,AIR AURORA,Aurora_Airlines,Aurora Airlines
,UST,AUSTRO AEREO,Austro_AÃ©reo,Austro Aéreo
,WLT,WINGLET,Aviation_Partners,Aviation Partners
,VLV,VLADLIFT,Avialift_Vladivostok,Avialift Vladivostok
,VME,AVIAMERICA,AviaciÃ³n_Comercial_de_AmÃ©rica,Aviación Comercial de América
,VVA,IALSI,Aviast_Air,Aviast Air
,WLV,WOLVERINE,Aviation_North,Aviation North
FK,WTA,WEST TOGO,Africa_West,Africa West
,VNT,AVIENT,Avient_Air_Zambia,Avient Air Zambia
,VZR,IAZUR,Aviazur,Aviazur
,VID,AVIAPRAD,Aviaprad,Aviaprad
G2,VXG,AVIREX-GABON,Avirex_(airline),Avirex
,VXX,EXPRESSAVIA,Aviaexpress_Aircompany,Aviaexpress Aircompany
,XAM,ALLIANCE,AMR_Services_Corporation,AMR Services Corporation
,XAO,,Airline_Operations_Services,Airline Operations Services
,VAZ,REMONT AIR,Airlines_400,Airlines 400
V8,VAS,ATRAN,ATRAN_Cargo_Airlines,ATRAN Cargo Airlines
,VAM,AMERAVIA,Ameravia,Ameravia
K6,KHV,AIR ANGKOR,Angkor_Air,Angkor Air
,VBC,AIR VICTOR,AVB-2004,AVB-2004 Ltd
,XKX,,ASECNA,ASECNA
,XAT,,AT_and_T_Aviation_Division,AT and T Aviation Division
,VAI,AIR AVALAIR,Avalair,Avalair
,VRT,AVERITT,Averitt_Air_Charter,Averitt Air Charter
VE,AVE,AVENSA,Avensa,Avensa
V5,VLI,AEROVOLAR,Avolar_Aerol%C3%ADneas,Avolar Aerolíneas
,VSA,STARBIRD,Avstar_Aviation,Avstar Aviation
,VSC,AVESCA,AVESCA,AVESCA
,XAV,AVIAPROM,Aviaprom_Enterprises,Aviaprom Enterprises
,VTT,VIATRANSPORT,Avia_Trans_Air_Transport,Avia Trans Air Transport
,VSR,AVIOSTART,Aviostart_AS,Aviostart AS
,VTG,ATACARGO,AviaÃ§Ã£o_Transportes_AÃ©reos_e_Cargas,Aviação Transportes Aéreos e Cargas
,XJA,,Assistance_Aeroportuaire_de_L'Aeroport_de_Paris,Assistance Aeroportuaire de L'Aeroport de Paris
,XFS,,American_Flight_Service_Systems,American Flight Service Systems
,XME,AUS-CARGO,Australian_air_Express,Australian air Express
,XMG,,AMS_Group,AMS Group
,CAJ,CAR LINE,Air_Caraibes_Atlantique,Air Caraibes Atlantique
,CAO,AIRCHINA FREIGHT,Air_China_Cargo,Air China Cargo
,CBE,AEROCARIBE,AerovÃ­as_Caribe,Aerovías Caribe
,CBO,TAXI CABO,Aerotaxi_del_Cabo,Aerotaxi del Cabo
,CBS,AIR COLUMBUS,Air_Columbus,Air Columbus
,CBV,CABOAEREO,Aereo_Cabo,Aereo Cabo
CA,CCA,AIR CHINA,Air_China,Air China
,CDA,CARDAL,Aerocardal,Aerocardal
Q6,CDP,CONDOR-PERU,Aero_Condor_Peru,Aero Condor Peru
,CDU,,Aerotrans,Aerotrans
,CDV,SKOL,Airline_Skol,Airline Skol
,CFF,AEROFAN,Aerofan,Aerofan
,CFM,ACEF,ACEF,ACEF
,CFR,,Africa_One,Africa One
,CFV,CALAFIA,Aero_Calafia,Aero Calafia
,CGB,CARGO BELIZE,Air_Cargo_Belize,Air Cargo Belize
,CGV,CLUBE ALGARVE,Aero_Clube_Do_Algarve,Aero Clube Do Algarve
,CGW,CHANGCHENG,Air_Great_Wall,Air Great Wall
,CHJ,AIR CHAIKA,Aircompany_Chaika,Aircompany Chaika
,CHR,ZAIRE CHARTER,Air_Charter_Services,Air Charter Services
,CHV,CHARTAIR,Air_Charter_Professionals,Air Charter Professionals
,CID,ACID,Asia_Continental_Airlines,Asia Continental Airlines
5F,CIR,AIR ARCTIC,Arctic_Circle_Air_Service,Arctic Circle Air Service
,CKL,CIRCLE CITY,Aviation_Charter_Services,Aviation Charter Services
,CLL,AEROCASTILLO,AerovÃ­as_Castillo,Aerovías Castillo
,CLP,CLUB PORTUGAL,Aero_Club_De_Portugal,Aero Club De Portugal
,CMF,COMPASSION,Air_Care_Alliance,Air Care Alliance
,CND,CONDOMINICANA,Aero_Continente_Dominicana,Aero Continente Dominicana
,CNE,CONNECTOR,Air_Toronto,Air Toronto
,CNH,CHENANGO,Aquila_Air,Aquila Air
,CNU,AIR CONSUL,Air_Consul,Air Consul
,CNX,CANEX,AllCanada_Express,AllCanada Express
,CPF,TECHSERVICE,Airtechservice,Airtechservice
QC,CRD,AIR CORRIDOR,Air_Corridor,Air Corridor
NV,CRF,AIR CENTRAL,Air_Central,Air Central
,CRJ,AIR CRUZAL,Air_Cruzal,Air Cruzal
,CRP,AEROTRANSCORP,Aerotransportes_Corporativos,Aerotransportes Corporativos
,CRQ,CREE,Air_Creebec,Air Creebec
,CTA,CHAR-TRAN,Aero_Charter_and_Transport,Aero Charter and Transport
,CTE,TENGLONG,Air_Tenglong,Air Tenglong
,CTR,CENTAURO,AerolÃ­neas_Centauro,Aerolíneas Centauro
,CUO,CUAHONTE,Aerocuahonte,Aerocuahonte
CV,CVA,CHATHAM,Air_Chathams,Air Chathams
CW,CWM,AIR MARSHALLS,Air_Marshall_Islands,Air Marshall Islands
,CWP,COASTWATCH,Australian_Customs_Service,Australian Customs Service
,CYL,CITYLINER,Air_One_Cityliner,Air One Cityliner
,CYO,COYOTE,Air_Transport,Air Transport
ZA,CYD,CYCLONE,AccessAir,AccessAir
,CYE,AEROCHEYENNE,Aerocheyenne,Aerocheyenne
AH,DAH,AIR ALGERIE,Air_Alg%C3%A9rie,Air Algérie
,DAP,DAP,Aerov%C3%ADas_DAP,Aerovías DAP
,DBA,DOUBLE-A,,Air Alpha
,DBD,AIR NIAGARA,Air_Niagara_Express,Air Niagara Express
,DEF,TIRPA,Aviation_Defense_Service,Aviation Defense Service
,DEG,DEGGER,Air_Service_Groningen,Air Service Groningen
KI,DHI,ADAM SKY,Adam_Air,Adam Air
ER,DHL,D-H-L,Astar_Air_Cargo,Astar Air Cargo
,DHM,ARCHER,Archer_Aviation,Archer Aviation
,DIC,AEROMEDICA,Aeromedica,Aeromedica
,DIN,AERODIN,Aerodin,Aerodin
HO,DJA,ANTINEA,Antinea_Airlines,Antinea Airlines
,DJU,AIR DJIB,Air_Djibouti,Air Djibouti
EN,DLA,DOLOMITI,Air_Dolomiti,Air Dolomiti
,DLS,AEROMODELO,Aero_Modelo,Aero Modelo
,DLU,DEL SUR,Aerol%C3%ADneas_del_Sur,Aerolíneas del Sur
,DMC,DINAMICAMONT,Aerodinamica_de_Monterrey,Aerodinamica de Monterrey
,DMI,AERODINAMICO,Aeroservicios_Dinamicos,Aeroservicios Dinamicos
,DML,,Aerotaxis_Dosmil,Aerotaxis Dosmil
,DNA,AERODESPACHOS,Aerodespachos_de_El_Salvador,Aerodespachos de El Salvador
,DNC,FLYINGOLIVE,Aerodynamics_MÃ¡laga,Aerodynamics Málaga
,DNJ,DYNAJET,Aerodynamics_Incorporated,Aerodynamics Incorporated
NM,DRD,ALADA AIR,Air_Madrid,Air Madrid
,DRM,DARTMOOR,Airways_Flight_Training,Airways Flight Training
,DRO,AERONORESTE,Aeronaves_Del_Noreste,Aeronaves Del Noreste
,DSC,ADDIS CARGO,Addis_Air_Cargo_Services,Addis Air Cargo Services
,DSK,SKYBANNER,Aero_Algarve,Aero Algarve
,DST,DESERT,Aex_Air,Aex Air
,DVI,AERO DAVINCI,Aerodavinci,Aero Davinci International
,DYN,AERO DYNAMIC,Aero_Dynamics,Aero Dynamics
,EAE,AECA,Aeroservicios_Ecuatorianos,Aeroservicios Ecuatorianos
,EAP,AERO-PYRENEES,Aero-Pyrenees,Aero-Pyrenees
,EAT,TRANS EUROPE,Air_Transport,Air Transport
EE,EAY,REVAL,Aero_Airlines,Aero Airlines
,EBC,CALIXJET,Aero_Ejecutivo_De_Baja_California,Aero Ejecutivo De Baja California
4F,ECE,AIRCITY,Air_City,Air City
,ECG,EJECTUIVOS RCG,Aero_Ejecutivos_RCG,Aero Ejecutivos RCG
,ECL,AERO CASTELLANA,AeronÃ¡utica_Castellana,Aeronáutica Castellana
,ECM,AERO COMERCIALES,AerolÃ­neas_Comerciales,Aerolíneas Comerciales
,EDA,ANDES,Aerolineas_Nacionales_Del_Ecuador,Aerolineas Nacionales Del Ecuador
,EET,AESTE,Air_Este,Air Este
,EFC,FLIGHT TAXI,Air_Mana,Air Mana
EI,EIN,SHAMROCK,Aer_Lingus,Aer Lingus
,EJP,EJECCORPORATIVOS,Aeroservicios_Ejecutivos_Corporativos,Aeroservicios Ejecutivos Corporativos
E8,ELG,ALPI EAGLES,Alpi_Eagles,Alpi Eagles
,END,ARRENDADORA,Arrendadora_y_Transportadora_AÃ©rea,Arrendadora y Transportadora Aérea
,ENW,AERONOR,Aeronaves_Del_Noroeste,Aeronaves Del Noroeste
,EOL,EOLE,Airailes,Airailes
,EOM,AERO ERMES,Aero_Ermes,Aero Ermes
,EPL,EMPRESARIALES,Aero_Transportes_Empresariales,Aero Transportes Empresariales
,EPE,AEROEMPRESARIAL,Aero_Empresarial,Aero Empresarial
KY,EQL,EQUATORIAL,Air_S%C3%A3o_Tom%C3%A9_and_Pr%C3%ADncipe,Air São Tomé and Príncipe
,ERG,AVIANERGO,Avianergo,Avianergo
,ERI,ASERGIO,Aero_Servicios_Regiomontanos,Aero Servicios Regiomontanos
,ERK,AEROSEC,Aerosec,Aerosec
,ERM,EOMAAN,Aeromaan,Aeromaan
,ESB,AEREOSABA,Aereosaba,Aereosaba
,ESO,,Avitat,Avitat
,ESU,ALESUR,AerolÃ­neas_Ejecutivas_Del_Sureste,Aerolíneas Ejecutivas Del Sureste
,ESZ,ESPERANZA,AeronÃ¡utica_La_Esperanza,Aeronáutica La Esperanza
,ETE,AEROSIETE,Aero_Siete,Aero Siete
,EUK,SNOWBIRD,Air_Atlanta_Europe,Air Atlanta Europe
,EVE,SUNBEAM,Air_Evex,Air Evex
,EVR,DIANA,Aeronautical_Academy_of_Europe,Aeronautical Academy of Europe
,EWE,,Eurowings_Europe,Eurowings Europe
,EXG,EXCHANGE,Air_Exchange,Air Exchange
,FAC,FAROECOPTER,Atlantic_Helicopters,Atlantic Helicopters
,FAG,FUAER,Argentine_Air_Force,Argentine Air Force
PC,FAJ,FIJIAIR,Air_Fiji,Air Fiji
,FAN,FANBIRD,AF-Air_International,AF-Air International
,FBW,,Aviation_Data_Systems,Aviation Data Systems
,FCI,FAST CHECK,Air_Carriers,Air Carriers
,FCO,AEROFRISCO,Aerofrisco,Aerofrisco
,FCU,,Alfa_4,Alfa 4
,FDA,AIR SANKORE,African_Airlines,African Airlines
,FDS,FLYDOC,African_Medical_and_Research_Foundation,African Medical and Research Foundation
,FGT,FREIAERO,Aero_Freight,Aero Freight
,FIC,AEROSAFIN,Aerosafin,Aerosafin
OF,FIF,AIR FINLAND,Air_Finland,Air Finland
,FII,FLIGHT CHECKER,Aerodata_Flight_Inspection,Aerodata Flight Inspection
,FIX,AIRFIX,Airfix_Aviation,Airfix Aviation
FJ,FJI,PACIFIC,Fiji_Airways,Fiji Airways
,FLD,,Air_Falcon,Air Falcon
RC,FLI,FAROELINE,Atlantic_Airways,Atlantic Airways
QH,FLZ,AIR FLORIDA,Aero_Leasing,Aero Leasing
,FMT,,Air_Fret_De_Mauritanie,Air Fret De Mauritanie
,FNM,,Avio_Nord,Avio Nord
,FNO,RIAZOR,Aeroflota_Del_Noroeste,Aeroflota Del Noroeste
,FNX,AERO FENIX,Aero_Fenix,Aero Fenix
,FPY,AFRICOMPANY,African_Company_Airlines,African Company Airlines
,FRJ,AFRIJET,Afrijet_Airlines,Afrijet Airlines
,FRK,AFRIFAST,Afrika_Aviation_Handlers,Afrika Aviation Handlers
,FRQ,CHARTER AFRIQUE,Afrique_Chart'air,Afrique Chart'air
,FRT,,Aerofreight_Airlines,Aerofreight Airlines
,FST,FAST TRACK,Aeros_Limited,Aeros Limited
,FTC,AFFAIRES TCHAD,Air_Affaires_Tchad,Air Affaires Tchad
,FTY,FLY TYROL,ABC_Bedarfsflug,ABC Bedarfsflug
NY,FXI,FAXI,Air_Iceland,Air Iceland
2P,GAP,ORIENT PACIFIC,Air_Philippines,Air Philippines
,GAU,AEROGAUCHO,Aerogaucho,Aerogaucho
,GBJ,GLOBAL JET,,Aero Business Charter
,GCF,AEROCARTO,Aeronor,Aeronor
,GCK,AEROGEM,Aerogem_Cargo,Aerogem Cargo
,GFO,AEROVIAS GOLFO,AerovÃ­as_del_Golfo,Aerovías del Golfo
,GGL,GIRA GLOBO,Aeron%C3%A1utica_(Angola),Aeronáutica
ZX,GGN,GEORGIAN,Air_Georgian,Air Georgian
,GHL,HANDLING,Aviance,Aviance
,GHN,AIR GHANA,Air_Ghana,Air Ghana
,GIL,AFRICAN TRANSPORT,African_International_Transport,African International Transport
2U,GIP,FUTURE EXPRESS,Air_Guinee_Express,Air Guinee Express
,GIZ,AFRILENS,Africa_Airlines,Africa Airlines
,GLL,TWINS,Air_Gemini,Air Gemini
,GLT,GASLIGHT,Aero_Charter,Aero Charter
,GME,MAYAN EAGLES,Aguilas_Mayas_Internacional,Aguilas Mayas Internacional
,GMM,AEROGUAMUCHIL,Aerotaxis_Guamuchil,Aerotaxis Guamuchil
,GMS,SERVICIOS GAMA,Aeroservicios_Gama,Aeroservicios Gama
0A,GNT,GINTA,Amber_Air,Amber Air
,GOA,ALBERTA,Alberta_Government,Alberta Government
,GRE,GREECE AIRWAYS,Air_Scotland,Air Scotland
DA,GRG,AIR GEORGIA,Air_Georgia,Air Georgia
,GRI,,Air_Cargo_Center,Air Cargo Center
GL,GRL,GREENLAND,Air_Greenland,Air Greenland
LL,GRO,ALLEGRO,Allegro_(airline),Allegro
,GRR,AGROAR,Agroar_-_Trabalhos_A%C3%A9reos,Agroar - Trabalhos Aéreos
,GRX,GRODNO,Aircompany_Grodno,Aircompany Grodno
,GSP,GREEN SPEED,Airlift_Alaska,Airlift Alaska
,GSV,AGRAV,Agrocentr-Avia,Agrocentr-Avia
,GTC,GOLDEN WINGS,Altin_Havayolu_Tasimaciligi_Turizm_Ve_Ticaret,Altin Havayolu Tasimaciligi Turizm Ve Ticaret
5Y,GTI,GIANT,Atlas_Air,Atlas Air
,GTP,GRUPOTAMPICO,Aerotaxi_Grupo_Tampico,Aerotaxi Grupo Tampico
,GUA,AGUASCALIENTES,Aerotaxis_de_Aguascalientes,Aerotaxis de Aguascalientes
GG,GUY,GREEN BIRD,Air_Guyane,Air Guyane
,GVI,IRINA,Air_Victoria_Georgia,Air Victoria Georgia
H9,HAD,HAITI AVIA,Air_d%27Ayiti,Air d'Ayiti
GG,HAH,AIR COMORES,Air_Comores_International,Air Comores International
,HAT,TAXI BIRD,Air_Taxi,Air Taxi
,HEI,AEROHEIN,Aerohein,Aerohein
,HGH,HIGHER,Atlantic_Air_Lift,Atlantic Air Lift
,HID,EJECUTIVA HIDALGO,AviaciÃ³n_Ejecutiva_De_Hildago,Aviación Ejecutiva De Hildago
,HJA,AIRHAITI,Air_Haiti,Air Haiti
HD,HLN,ORANGE,Air_Holland,Air Holland
,HJT,AL-RAIS CARGO,Al_Rais_Cargo,Al Rais Cargo
,HKH,HAWKHUNGARY,Air-Invest,Air-Invest
,HMA,TAHOMA,Air_Tahoma,Air Tahoma
,HMT,HAMILTON,Air_Nova_(UK),Air Nova
,HOM,AERO HOMEX,Aero_Homex,Aero Homex
,HPO,ALMIRON,Almiron_Aviation,Almiron Aviation
,HQO,,Avinor,Avinor
,HYR,HIGHFLYER,Airlink_Airways,Airlink Airways
8C,HZT,HORIZON TOGO,Air_Horizon,Air Horizon
,IAE,,AC_Insat-Aero,AC Insat-Aero
,ICM,INTER-CAMEROUN,Air_Inter_Cameroun,Air Inter Cameroun
,IFI,HELLAS LIFT,Air_Lift,Air Lift
,IKM,EASY SHUTTLE,Aero_Survey,Aero Survey
,ILK,ILEK,Aero_Airline,Aero Airline
,IME,AIRTIME,Airtime_Charters,Airtime Charters
,IMN,TAXI CIMARRON,Aerotaxis_Cimarron,Aerotaxis Cimarron
,INA,AERO-NACIONAL,Aero_Internacional,Aero Internacional
,ING,AEROINGE,Aeroingenieria,Aeroingenieria
,INO,INTENOR,Aeroservicios_Intergrados_de_Norte,Aeroservicios Intergrados de Norte
,IPL,IPULL,Airpull_Aviation,Airpull Aviation
,IRD,ARVAND,Arvand_Airlines,Arvand Airlines
,IRH,ATLAS AVIA,Atlas_Aviation_Group,Atlas Aviation Group
,IRW,ARAM,Aram_Airline,Aram Airline
,IRX,ARIA,Aria_Air,Aria Tour
,ITE,AEROTAXI,Aerotaxi_S.R.O.,Aerotaxi S.R.O.
,ITF,AIR AVITA,Avita-Servicos_AÃ©reos,Avita-Servicos Aéreos
,ITO,AERO CITRO,Aero_Citro,Aero Citro
,IVE,COMPANY EXEC,Air_Executive,Air Executive
,IWS,,Aviainvest,Aviainvest
W9,JAB,AIR BAGAN,Air_Bagan,Air Bagan
,JAD,AEROJAL,Aerojal,Aerojal
,JAR,AIRLINK,Airlink,Airlink
,JEE,AMBJEK AIR,Ambjek_Air_Services,Ambjek Air Services
,UTX,,Avfinity,Avfinity
,JMR,ALEXANDAIR,Alexandair,Alexandair
,JMX,JAMAICA EXPRESS,Air_Jamaica_Express,Air Jamaica Express
,JOA,,Air_Swift_Aviation,Air Swift Aviation
,JOB,JOBENI,Aerojobeni,Aerojobeni
IP,JOL,EDIL,Atyrau_Air_Ways,Atyrau Air Ways
,JPR,JASPER,Aerosmith_Aviation,Aerosmith Aviation
,JTS,AVIONESJETS,Arrendamiento_de_Aviones_Jets,Arrendamiento de Aviones Jets
,JUA,JUAREZ,Aero_Juarez,Aero Juarez
QK,JZA,JAZZ,Air_Canada_Jazz,Air Canada Jazz
,KAA,AASCO,Asia_Aero_Survey_and_Consulting_Engineers,Asia Aero Survey and Consulting Engineers
,KAD,AIR KIROVOGRAD,Air_Kirovograd,Air Kirovograd
,KAM,ICO-AIR,Air_Mach,Air Mach
,KAV,AIRKUFRA,Air_Kufra,Air Kufra
,KEK,ARKHABAY,Arkhabay,Arkhabay
,KFK,KRIFKA AIR,Aero_Charter_Krifka,Aero Charter Krifka
,KFT,AIR KRAFT MIR,Air_Kraft_Mir,Air Kraft Mir
,KGD,CONCORDE AIR,Air_Concorde,Air Concorde
,KHH,,Alexandria_Airlines,Alexandria Airlines
,KIE,TWEETY,Afit,Afit
,KKB,KHAKI BLUE,Air_South_(South_Carolina),Air South
KK,KKK,ATLASJET,Atlasjet,Atlasjet
,KLB,TRANS MALI,Air_Mali_International,Air Mali International
,KLZ,AEROKALUZ,Aerokaluz,Aerokaluz
JS,KOR,AIR KORYO,Air_Koryo,Air Koryo
,KOY,ALEKS,Araiavia,Araiavia
,KRE,AEROSUCRE,AeroSucre,AeroSucre
,KRT,KOKTA,Air_Kokshetau,Air Kokshetau
,KSI,KISSARI,Air_Kissari,Air Kissari
,KTN,AERONAVIGACIYA,Aeronavigaciya,Aeronavigaciya
,KVR,KAVAIR,Alliance_Avia,Alliance Avia
,KYC,DOLPHIN,Av_Atlantic,Av Atlantic
KC,KZR,ASTANALINE,Air_Astana,Air Astana
,LAG,AVILEG,Aviation_Legacy,Aviation Legacy
,,AEROLAGOS,AerovÃ­as_De_Lagos,Aerovías De Lagos
LV,LBC,ALBANIAN,Albanian_Airlines,Albanian Airlines
,LBI,ALBISA,Albisa,Albisa
,LBW,ALBANWAYS,Albatros_Airways,Albatros Airways
,LDG,DURANGO,AerolÃ­neas_AÃ©reas_Ejecutivas_De_Durango,Aerolíneas Aéreas Ejecutivas De Durango
3S,BOX,GERMAN CARGO,Aerologic,Aerologic
,LDN,ALDONAS AIR,Al-Donas_Airlines,Al-Donas Airlines
,LDR,AEROLIDER,Aero_Lider,Aero Lider
,LEM,,Alim_Airlines,Aleem
,LET,MEXEJECUTIV,Aerol%C3%ADneas_Ejecutivas,Aerolíneas Ejecutivas
,LFA,,Air_Alfa,Air Alfa
,LFC,LIFE FLIGHT CANADA,Aero_Control_Air,Aero Control Air
,LGN,AEROLAGUNA,Aerolaguna,Aerolaguna
,LHR,AL AHRAM,Al_Ahram_Aviation,Al Ahram Aviation
D4,LID,ALIDA,Alidaunia,Alidaunia
,LIE,AL-DAWOOD AIR,Al-Dawood_Air,Al-Dawood Air
,LKP,LAKE POWELL,American_Aviation,American Aviation
,LKS,AIRLIN,Airlink_Solutions,Airlink Solutions
,LKY,LUCKY,Air_Solutions,Air Solutions
9I,LLR,ALLIED,Air_India_Regional,Alliance Air
,LMA,AEROLIMA,Aerolima,Aerolima
,LML,ALAMIA AIR,Alamia_Air,Alamia Air
,LMP,AIR FLIGHT,Air_Plus_Argentina,Air Plus Argentina
,LMT,ALMATY,Almaty_Aviation,Almaty Aviation
,LMX,LINEAS MEXICANAS,AerolÃ­neas_Mexicanas_J_S,Aerolíneas Mexicanas J S
,LMY,AGLEB,Air_Almaty,Air Almaty
,LMZ,ALUNK,Air_Almaty_ZK,Air Almaty ZK
XL,LNE,AEROLANE,Aerolane,Aerolane
,LNK,LINK,Airlink,Airlink
,LNT,LINEAINT,Aerol%C3%ADneas_Internacionales,Aerolíneas Internacionales
,LOK,ALOK AIR,Alok_Air,Alok Air
,LOU,AIR SAINTLOUIS,Air_Saint_Louis,Air Saint Louis
,LPC,NETSTAR,Alpine_Aviation,Alpine Aviation
A6,LPV,ALPAV,Air_Alps_Aviation,Air Alps Aviation
,LRO,ALROSA,Alrosa-Avia,Alrosa-Avia
,LRW,AL RIDA,Al_Rida_Airways,Al Rida Airways
,LSK,AURELA,Aurela,Aurela
,LSM,,Aerobusinessservice,Aerobusinessservice
,LSR,ALSAIR,Alsair,Alsair
,LTI,LATINO,Aerotaxis_Latinoamericanos,Aerotaxis Latinoamericanos
,LUC,ALBINATI,Albinati_Aeronautics,Albinati Aeronautics
TD,LUR,,Atlantis_European_Airways,Atlantis European Airways
,LVN,ALIVEN,Aliven,Aliven
,LVR,AVIAVILSA,Aviavilsa,Aviavilsa
L8,LXG,LUXOR GOLF,Air_Luxor_GB,Air Luxor GB
LK,LXR,AIRLUXOR,Air_Luxor,Air Luxor
,LYT,APATAS,Apatas_Air,Apatas Air
,LZP,DOC AIR,Air_Ban,Air Ban
,LZR,LAZUR BEE-GEE,Air_Lazur,Air Lazur
,MAM,AEROMAN,AerÃ³dromo_De_La_Mancha,Aeródromo De La Mancha
MK,MAU,AIRMAURITIUS,Air_Mauritius,Air Mauritius
,MBA,AVAG AIR,Avag_Air,Avag Air
,MBB,AIR MANAS,Air_Manas,Air Manas
,MBC,MABECO,Airjet_Exploracao_Aerea_de_Carga,Airjet Exploracao Aerea de Carga
,MBV,AEREM,Aeriantur-M,Aeriantur-M
,MCB,WESTMID,Air_Mercia,Air Mercia
,MCD,AIR MED,Air_Medical,Air Medical
,MCO,MARCOS,AerolÃ­neas_Marcos,Aerolíneas Marcos
,MDC,NIGHT SHIP,Atlantic_Aero_and_Mid-Atlantic_Freight,Atlantic Aero and Mid-Atlantic Freight
MD,MDG,AIR MADAGASCAR,Air_Madagascar,Air Madagascar
,MDX,MEDAIR,Aerosud_Charter,Aerosud Charter
,MEF,EMPENNAGE,Air_Meridan,Air Meridan
,MFL,MCFLY,Aero_McFly,Aero McFly
,MGE,MAGELLAN,Asia_Pacific_Airlines_(Guam),Asia Pacific Airlines
,MGS,AEROMAGAR,Aeromagar,Aeromagar
,MIE,AEROPREMIER,Aero_Premier_De_Mexico,Aero Premier De Mexico
9U,MLD,AIR  MOLDOVA,Air_Moldova,Air Moldova
,MLF,AMAL,Amal_Airlines,Amal Airlines
L9,MLI,AIR MALI,Air_Mali_(1960%E2%80%931985),Air Mali
,MLN,AIR MADELEINE,Air_Madeleine,Air Madeleine
,MMC,AERMARCHE,Aermarche,Aermarche
,MMD,MERMAID,Air_Alsie,Air Alsie
,MMM,AVIAMERIDIAN,Aviation_Company_Meridian,Aviation Company Meridian
,MMP,AMP-INC,AMP_Incorporated,AMP Incorporated
,MMX,PERUMAX,Air_Max_Africa,Airmax
,MNE,AEROAMANECER,AerolÃ­neas_Amanecer,Aerolíneas Amanecer
,MNG,AERO MONGOLIA,Aero_Mongolia,Aero Mongolia
,MOC,MONARCH CARGO,Air_Monarch_Cargo,Air Monarch Cargo
,MOP,PUBLICITARIA,Aeropublicitaria_De_Angola,Aeropublicitaria De Angola
,MOR,AEROMORELIA,AerolÃ­neas_De_Morelia,Aerolíneas De Morelia
A7,MPD,RED COMET,Air_Plus_Comet,Air Plus Comet
QO,MPX,AEROMEXPRESS,Aeromexpress,Aeromexpress
,MQT,MUSKETEER,Air_ITM,Air ITM
,MRL,AEROMORELOS,Aeromorelos,Aeromorelos
,MRM,MARITIME,Aerocharter,Aerocharter
,MRP,ABAS,Abas_(airline),Abas
MR,MRT,MIKE ROMEO,Air_Mauritanie,Air Mauritanie
,MSC,,Air_Cairo,Air Cairo
,MSK,AIR SPORT,Air_Sport,Air Sport
,MSM,AEROMAS EXPRESS,Aeromas,Aeromas
,MSO,MESO AMERICANAS,Aerol%C3%ADneas_Mesoamericanas,Aerolíneas Mesoamericanas
,MSV,AERAFKAM,Aero-Kamov,Aero-Kamov
,MTB,AEROMETROPOLIS,Aerotaxis_Metropolitanos,Aerotaxis Metropolitanos
,MTE,AEROMET,Aeromet_LÃ­nea_AÃ©rea,Aeromet Línea Aérea
,MTK,AIRMETACK,Air_Metack,Air Metack
,MTY,MONTY,Air_Montegomery,Air Montegomery
,MXO,MAXAERO,Aerotaxi_Mexicano,Aerotaxi Mexicano
,MYS,AERO YAQUI,Aero_Yaqui_Mayo,Aero Yaqui Mayo
,MZK,,AVC_Airlines,AVC Airlines
,MZL,MONTES AZULES,AerovÃ­as_Montes_Azules,Aerovías Montes Azules
3S*,AEN,,Aeroland_Airways,Aeroland Airways
2V*,,,Amtrak,Amtrak
8D*,SUW,,Astair,Astair
F4,NBK,AL-AIR,Albarka_Air,Albarka Air
,NCL,ANCARGO AIR,ACA-Ancargo_Air_Sociedade_de_Transporte_de_Carga_Lda,ACA-Ancargo Air Sociedade de Transporte de Carga Lda
,NEL,AEROLAREDO,Aero_Servicios_de_Nuevo_Laredo,Aero Servicios de Nuevo Laredo
,NGV,ANGOAVIA,Angoavia,Angoavia
,NID,AERONI,Aeroni,Aeroni
,NIE,AERONIETO,Aeroejecutiva_Nieto,Aeroejecutiva Nieto
AJ,NIG,AEROLINE,Aero_Contractors_(Nigeria),Aero Contractors
,NKZ,NOVOKUZNETSK,Aerokuzbass,Aerokuzbass
,NRS,NORTH SLOPE,Atlantic_Richfield_Company,Atlantic Richfield Company
,NSO,SOSA,Aerol%C3%ADneas_Sosa,Aerolíneas Sosa
,NTD,,Aero_Norte,Aero Norte
,NTV,INTER-IVOIRE,Interivoire,Air Inter Ivoire
,NUL,SERVICIOS NUEVOLEON,Aeroservicios_De_Nuevo_Leon,Aeroservicios De Nuevo Leon
,NVI,NEW AVIAL,Avial_NV,Avial NV Aviation Company
,NWG,NORWING,Airwing,Airwing
,NXA,BLUE-DOLPHIN,Air_Next,Air Next
,OAO,DVINA,Arkhangelsk_2_Aviation_Division,Arkhangelsk 2 Aviation Division
,OGI,AEROGISA,Aerogisa,Aerogisa
,OLV,OLVE,AerolÃ­neas_Olve,Aerolíneas Olve
,OMG,OMEGA,Aeromega,Aeromega
,ONR,EDER,Air_One_Nine,Air One Nine
,ONT,ONTARIO,Air_Ontario,Air Ontario
,ORP,CORPSA,Aerocorp,Aerocorp
,ORS,AVIATION SERVICE,Action_Air,Action Air
,OSN,AEROSAN,Aerosan,Aerosan
,OSO,,Aviapartner_Limited_Company,Aviapartner Limited Company
,PAJ,ALIPARMA,Aliparma,Aliparma
,PBT,PARABET,Air_Parabet,Air Parabet
8Y,PBU,AIR-BURUNDI,Air_Burundi,Air Burundi
,PCG,POSTAL CARGO,Aeropostal_Cargo_de_Mexico,Aeropostal Cargo de Mexico
,PCK,AIRPACK EXPRESS,Air_Pack_Express,Air Pack Express
,PCS,AIR PALACE,Air_Palace,Air Palace
OT,PEL,PELICAN,Aeropelican_Air_Services,Aeropelican Air Services
,PEV,PEOPLES,Peoples_Vienna_Line,Peoples Vienna Line
,PFI,PACIFICO CHIHUAHUA,AerolÃ­neas_Chihuahua,Aerolíneas Chihuahua
,PFT,PROFREIGHT,Air_Cargo_Express_International,Air Cargo Express International
,PHR,PHARAOH,Al_Farana_Airline,Al Farana Airline
,PHW,PHOENIX SHARJAH,Ave.com,Ave.com
,PIE,PIRATE,Air_South_West,Air South West
,PIF,AEROCALPA,Aeroservicios_California_Pacifico,Aeroservicios California Pacifico
,PKA,PAKISTAN AIRWAY,AST_Pakistan_Airways,AST Pakistan Airways
,PNL,AEROPERSONAL,Aero_Personal,Aero Personal
,PNU,AERO PLATINUM,Aero_Servicios_Platinum,Aero Servicios Platinum
,POY,APOYO AEREO,Apoyo_AÃ©reo,Apoyo Aéreo
,PRI,AEROPRIV,Aerotransportes_Privados,Aerotransportes Privados
,PRT,PATRIOT,Atlantic_Coast_Jet,Atlantic Coast Jet
AD,PRZ,RADISAIR,Air_Paradise_International,Air Paradise International
,PSE,SIPSE,Aeroservicio_Sipse,Aeroservicio Sipse
,PSL,CORSAN,Aeroservicios_Corporativos_De_San_Luis,Aeroservicios Corporativos De San Luis
,PVA,TRANSPRIVADO,Aerotransportes_Privados,Aerotransportes Privados
,PZA,AEREO PARAZA,AÃ©reo_Taxi_Paraza,Aéreo Taxi Paraza
QD,QCL,ACLA,Air_Class_L%C3%ADneas_A%C3%A9reas,Air Class Líneas Aéreas
,QEA,,Aviation_Consultancy_Office,Aviation Consultancy Office
,QAT,AIR QUASAR,Aero_Taxi,Aero Taxi
,QKC,QUAKER CITY,Aero_Taxi_Aviation,Aero Taxi Aviation
,QLA,QUEBEC LABRADOR,Aviation_Quebec_Labrador,Aviation Quebec Labrador
QS,QSC,ZEBRA,African_Safari_Airways,African Safari Airways
,QUI,QUIMMCO,Aero_Quimmco,Aero Quimmco
,RAD,AIR ALADA,Alada,Alada
,RAI,DIASA,Aerotur_Air,Aerotur Air
,RAP,RAPTOR,Air_Center_Helicopters,Air Center Helicopters
,RBE,RUM BENIN,Aur_Rum_Benin,Aur Rum Benin
,RBJ,AEROBAJIO,Aeroserivios_Del_Bajio,Aeroserivios Del Bajio
4Y,RBU,AIRBUS FRANCE,Airbus_France,Airbus France
,RBV,AIR ROBERVAL,Air_Roberval,Air Roberval
AG,ARU,ARUBA,Aruba_Airlines,Aruba Airlines
,RCC,RACER,Air_Charters_Eelde,Air Charters Eelde
,RCE,AEROCER,Aerocer,Aerocer
,RCF,AEROFLOT-CARGO,Aeroflot-Cargo,Aeroflot-Cargo
MC,RCH,REACH,Air_Mobility_Command,Air Mobility Command
,RCI,AIR CASSAI,Air_Cassai,Air Cassai
,RCO,AEROCOAHUILA,Aero_Renta_De_Coahuila,Aero Renta De Coahuila
,RCP,AEROCORPSA,Aerocorp,Aerocorp
,RCQ,REGIONAL CARGO,AerolÃ­neas_Regionales,Aerolíneas Regionales
,RCU,AIR COURIER,Atlantic_S.L.,Atlantic S.L.
,RCX,SERVICE CENTER,Air_Service_Center,Air Service Center
,RDD,ADLINES,Adygeya_Airlines,Adygeya Airlines
,RDM,AIR ADA,Air_Ada,Air Ada
RE,REA,AER ARANN,Aer_Arann,Aer Arann
,REN,AERORENT,Aero-Rent,Aero-Rent
,RES,RESCUE,Australian_Maritime_Safety_Authority,Australian Maritime Safety Authority
UU,REU,REUNION,Air_Austral,Air Austral
,REY,AEROREY,Aero-Rey,Aero-Rey
,RFC,AERO AFRICA,Aero_Africa,Aero Africa
ZP,AZP,GUARANI,Amaszonas_Paraguay,Amaszonas Paraguay
,RFD,RAFHILER,Aerotransportes_Rafilher,Aerotransportes Rafilher
,RGO,ARGOS,Argo,Argo
,RGT,AIRGOAT,Airbourne_School_of_Flying,Airbourne School of Flying
,RHL,ARCHIPELS,Air_Archipels,Air Archipels
,RIF,INTERMIN AVIA,Aviation_Ministry_of_the_Interior_of_the_Russian_Federation,Aviation Ministry of the Interior of the Russian Federation
,RIS,AERIS,Aeris_GestiÃ³n,Aeris Gestión
6K,RIT,ASIAN SPIRIT,Asian_Spirit,Asian Spirit
,RJS,ASERJET,Aeroservicios_Jet,Aeroservicios Jet
,RKA,AIRAFRIC,Air_Afrique,Air Afrique
A5,RLA,AIRLINAIR,Airlinair,Airlinair
,RLK,,Air_Nelson,Air Nelson
,RLL,AEROLEONE,Air_Leone,Air Leone
QL,RLN,AERO LANKA,Aero_Lanka,Aero Lanka
,RLZ,ALIZE,Air_Alize,Air Alize
,RMD,AIR AMDER,Air_Amder,Air Amder
R3,RME,ARMENIAN,Armenian_Airlines,Armenian Airlines
MV,RML,HELLASMED,Air_Mediterranean,Air Mediterranean
,RMO,ARM-AERO,Arm-Aero,Arm-Aero
,RMX,AEROMAX,Air_Max,Air Max
2O,RNE,AIR SALONE,Air_Salone,Air Salone
,RNM,AEROMNEM,Aeronem_Air_Cargo,Aeronem Air Cargo
,RNR,RUNNER,Air_Cargo_Masters,Air Cargo Masters
U8,RNV,ARMAVIA,Armavia,Armavia
,ROE,ESTE-BOLIVIA,Aeroeste,Aeroeste
,ROH,AEROGEN,Aero_Gen,Aero Gen
,ROI,AVIOR,Avior_Airlines,Avior Airlines
,ROL,AEROEL,Aeroel_Airways,Aeroel Airways
BQ,ROM,BRAVO QUEBEC,Aeromar_L%C3%ADneas_A%C3%A9reas_Dominicanas,Aeromar Lineas Aereas Dominicanas
,ROO,AEROITALIA,Aeroitalia,Aeroitalia
,ROD,AERODAN,Aerodan,Aerodan
P5,RPB,AEROREPUBLICA,AeroRep%C3%BAblica,AeroRepública
,RPC,AEROPACSA,AerolÃ­neas_Del_PacÃ­fico,Aerolíneas Del Pacífico
,RRC,AEROROCA,Aero_Roca,Aero Roca
,RRE,AERO TORREON,Aerotransportes_Internacionales_De_Torreon,Aerotransportes Internacionales De Torreon
,RRM,AIR ROMANIA,Acvila_Air,Acvila Air-Romanian Carrier
,RSC,TARASCAS,AerolÃ­neas_Ejecutivas_Tarascas,Aerolíneas Ejecutivas Tarascas
BF,RSR,CONGOSERV,A%C3%A9ro-Service,Aero-Service
5L,RSU,AEROSUR,Aerosur,AeroSur
,RTE,LUZAVIA,Aeronorte,Aeronorte
,RTH,ARTHELICO,Artis_(airline),Artis
,RTO,ARTOAIR,Arhabaev_Tourism_Airlines,Arhabaev Tourism Airlines
,RTQ,TURQUOISE,Air_Turquoise,Air Turquoise
,RTU,AEROTUCAN,Aerotucan,Aerotucan
,RUD,ANASTASIA,Air_Anastasia,Air Anastasia
,RUM,AIR RUM,Air_Rum,Air Rum
,RUN,CARGO TURK,ACT_Havayollari,ACT Havayollari
,RVP,AEROVIP,Air_VIP,Air VIP
,RVT,AIR-VET,Aircompany_Veteran,Aircompany Veteran
,RWB,,Alliance_Express_Rwanda,Alliance Express Rwanda
,RWC,ARROWEC,Arrow_Ecuador_Arrowec,Arrow Ecuador Arrowec
,RWY,TYNWALD,Aerogroup_98_Limited,Aerogroup 98 Limited
,RXT,AERO-EXTRA,Aeroxtra,Aeroxtra
,RWS,,Air_Whitsunday,Air Whitsunday
,RZL,AERO ZAMBIA,Aero_Zambia,Aero Zambia
,RZN,ZANO,Aero_Zano,Aero Zano
,RZZ,RED ZONE,Anoka_Air_Charter,Anoka Air Charter
,SBH,AEROSAAB,Aerosaab,Aerosaab
,SCD,ASSOCIATED,Associated_Aviation,Associated Aviation
,SCM,SCREAMER,American_Jet_International,American Jet International
EX,SDO,AERO DOMINGO,Air_Santo_Domingo,Air Santo Domingo
,SDP,SUDPACIFICO,Aero_Sudpacifico,Aero Sudpacifico
,SEF,SERVIPACIFICO,Aero_Servicios_Ejecutivas_Del_Pacifico,Aero Servicios Ejecutivas Del Pacifico
JR,SER,AEROCALIFORNIA,Aero_California,Aero California
,SGV,SEGOVIA,Aerosegovia,Aerosegovia
,SHH,AIRSHARE,Airshare_Holdings,Airshare Holdings
,SIY,SIYUSA,Aerosiyusa,Aerosiyusa
,SIZ,AEROSILZA,Aero_Silza,Aero Silza
,SJN,SAN JUAN,Air_San_Juan,Air San Juan
,SKP,SKIPPER,Aero-North_Aviation_Services,Aero-North Aviation Services
,SMI,SAMI,Aero_Sami,Aero Sami
Z3,SMJ,AVAVIA,Avient_Aviation,Avient Aviation
,SOD,ALSOL,AerolÃ­neas_Sol,Aerolíneas Sol
,SOE,AIR SOLEIL,Air_Soleil,Air Soleil
,SOG,AEROSOGA,Aero_Soga,Aero Soga
,SPD,SPEEDLINE,Airspeed_Aviation,Airspeed Aviation
M3,SPJ,AIR SKOPJE,Air_Service_(airline),Air Service
,SPO,EJECTUIV PACIFICO,Aeroservicios_Ejecutivos_Del_Pacifico,Aeroservicios Ejecutivos Del Pacifico
,SPY,THAI SPACE,Asian_Aerospace_Service,Asian Aerospace Service
,SPZ,SPEED SERVICE,Airworld,Airworld
,SQR,ALSAQER AVIATION,Alsaqer_Aviation,Alsaqer Aviation
,SRI,AIRSAFARI,Air_Safaris_and_Services,Air Safaris and Services
,SRV,SERVICORP,Aero_Servicio_Corporativo,Aero Servicio Corporativo
,SSL,SIERRA SULTAN,Air_Sultan,Air Sultan
,SSM,RAPID,Aero_1_Pro-Jet,Aero 1 Pro-Jet
,SSN,SUNSTREAM,Airquarius_Air_Charter,Airquarius Air Charter
,STK,SAT PAK,Aeropac,Aeropac
,STT,PARADISE,Air_St._Thomas,Air St. Thomas
,SUE,AEROSURESTE,AerolÃ­neas_Del_Sureste,Aerolíneas Del Sureste
,SUY,SURVEY,Aerial_Surveys_(1980)_Limited,Aerial Surveys (1980) Limited
GM,SVK,SLOVAKIA,Air_Slovakia,Air Slovakia
,SWH,SHOCKWAVE,Adler_Aviation,Adler Aviation
R3,SYL,AIR YAKUTIA,Aircompany_Yakutia,Aircompany Yakutia
,SYT,SKYTRACK,Aerosud_Aviation,Aerosud Aviation
,TAA,AERO COSTA,Aeroservicios_de_La_Costa,Aeroservicios de La Costa
VW,TAO,TRANS-AEROMAR,Aeromar,Aeromar
,TBL,AEROTREBOL,Aerotrebol,Aerotrebol
,TBO,AERO CABOS,Aero_Taxi_de_Los_Cabos,Aero Taxi de Los Cabos
JY,TCI,KERRMONT,Air_Turks_and_Caicos,Air Turks and Caicos
,TCO,TRANSCOLOMBIA,Aerotranscolombina_de_Carga,Aerotranscolombina de Carga
,TDG,TURBO DOG,Air_Cargo_Express,Air Cargo Express
,TDT,TRIDENT,Atlas_Helicopters,Atlas Helicopters
,TDY,AIR TODAY,Air_Today,Air Today
,TED,AEROAZTECA,Aero_Servicios_Azteca,Aero Servicios Azteca
OR,TFL,ARKEFLY,Arkefly,Arkefly
,TIR,ANTAIR,Antair,Antair
,TLB,TRIPLE-A,Atlantique_Air_Assistance,Atlantique Air Assistance
,TLD,AEREO AUTLAN,Aereo_Taxi_Autlan,Aereo Taxi Autlan
,TLE,AEROUTIL,Aero_Util,Aero Util
,TLU,AEROTOLUCA,Aero_Toluca_Internactional,Aero Toluca Internactional
,TME,TAXICENTRO,Aero_Taxi_del_Centro_de_Mexico,Aero Taxi del Centro de Mexico
,TOC,TROPICMEX,Aerotropical,Aerotropical
,TOH,TOMISKO CARGO,Air_Tomisko,Air Tomisko
CG,TOK,BALUS,Airlines_PNG,Airlines PNG
,TON,AEROTONALA,Aero_Tonala,Aero Tonala
,TPB,AERO TROPICAL,Aero_Tropical,Aero Tropical
TY,TPC,AIRCAL,Air_Cal%C3%A9donie,Air Calédonie
,TPK,TCHAD-HORIZON,Air_Horizon,Air Horizon
,TPO,TAXI-POTOSI,Aero_Taxi_del_Potosi,Aero Taxi del Potosi
,TQS,AEROTURQUESA,Aeroturquesa,Aeroturquesa
,TRH,TRANSTAR,Airmark_Aviation,Airmark Aviation
FL,TRS,CITRUS,AirTran_Airways,AirTran Airways
TS,TSC,AIR TRANSAT,Air_Transat,Air Transat
,TSQ,AIRTRA,Airtransse,airtransse
,TTB,AERO TURISTICAS,AerolÃ­neas_TurÃ­sticas_del_Caribe,Aerolíneas Turísticas del Caribe
,TTE,TETON,Avcenter,Avcenter
,TUN,TUNGARU,Air_Tungaru,Air Tungaru
EC,TWN,TWINARROW,Avialeasing_Aviation_Company,Avialeasing Aviation Company
,TXD,TAXI OESTE,Aerotaxis_del_Noroeste,Aerotaxis del Noroeste
,TXF,ALFE,Aerotaxis_Alfe,Aerotaxis Alfe
,TXI,AEREOTAXIS,Aereotaxis,Aereotaxis
,TZA,AERO TOMZA,Aero_Tomza,Aero Tomza
,TZT,ZAMBEZI,Air_Zambezi,Air Zambezi
,UAG,AFRALINE,Afra_Airlines,Afra Airlines
,UAR,AEROSTAR,Aerostar_Airlines,Aerostar Airlines
,UCK,GALETA,Air_Division_of_the_Eastern_Kazakhstan_Region,Air Division of the Eastern Kazakhstan Region
DW,UCR,CHARTER UKRAINE,Aero-Charter_Ukraine,Aero-Charter Ukraine
U7,UGA,UGANDA,Air_Uganda,Air Uganda
,UED,AIR L-A,Air_LA,Air LA
6U,UKR,AIR UKRAINE,Air_Ukraine,Air Ukraine
,UMB,AIR UMBRIA,Air_Umbria,Air Umbria
,UND,ATUNEROS UNIDOS,Atuneros_Unidos_de_California,Atuneros Unidos de California
,USC,STAR CHECK,AirNet_Express,AirNet Express
,VNR,AVANTAIR,Avantair,Avantair
,FLP,AEROCLUB FLAPS,Aeroclub_Flaps,Aeroclub Flaps
6R,DRU,MIRNY,Alrosa_Air_Company,Alrosa Air Company
AD,AZU,Azul,Azul_Linhas_A%C3%A9reas_Brasileiras,Azul Linhas Aéreas Brasileiras
6A,CHP,AVIACSA,Aviacsa,Aviacsa
JZ,SKX,SKY EXPRESS,Avia_Express,Avia Express
,SRY,,As-Aero,As-Aero
WD*,AAN,AMSTEL,Amsterdam_Airlines,Amsterdam Airlines
,AAT,,Air_Central_Asia,Air Central Asia
,ABI,Anair,Antigua_and_Barbuda_Airways,Antigua and Barbuda Airways
,ACB,AFRICARGO,African_Cargo_Services,African Cargo Services
,ACH,AIR PLUS,Air_Cargo_Plus,Air Cargo Plus
,ACJ,AVICHARTER,Avicon,Avicon
,AAE,AIR EAST,Astec_Air_East,Astec Air East
,AAE,ARIZONA,Arizona_Air,Arizona Air
AQ,AAH,ALOHA,Aloha_Airlines,Aloha Airlines
,AAJ,ALFA SUDAN,Alfa_Airlines,Alfa Airlines
,AAN,ANDALUSAIR,Andalusair,Andalusair
,AAP,ASTRO AIR,Astro_Air_International,Astro Air International
,AAP,ARABASCO,Arabasco_Air_Services,Arabasco Air Services
3J,AAQ,LIAISON,Air_Alliance,Air Alliance
,AAR,PATRIOT,Americ_Air,Americ Air
,AAS,Aviaservice,Airtransservice,Airtransservice
,AAS,AIR SERVICES,Austrian_Air_Services,Austrian Air Services
OB,AAT,AUSTRIAN CHARTER,Austrian_Airtransport,Austrian Airtransport
,AAV,,Aly_Aviation,Aly Aviation
,AAW,AUSTIN,Austin_Airways,Austin Airways
SM,AAW,,Allied_Airways,Aberdeen Airways
,AAW,ALMETA AIR,Almeta_Air,Almeta Air
KJ,AAZ,,Asian_Air,Asian Air
BX,ABL,AIR BUSAN,Air_Busan,Air Busan
AK,ABR,AIRBRIDGE,Air_Bridge_Carriers,Air Bridge Carriers
,ABT,AMBITION,Ambeo,Ambeo
,ABE,ARBERIA AIRLINES,Arberia_Airlines,Arberia Airlines
,ACE,,Air_Charter_Express,Air Charter Express
YE,ACQ,,Aryan_Cargo_Express,Aryan Cargo Express
,ACS,AIRCRAFT SALES,Aircraft_Sales_and_Services,Aircraft Sales and Services
,ABS,AIR CENTRAL,Air_Central_(US),Air Central
,ADF,ADE AVIACION,"Ade,_AviaciÃ³n_Deportiva","Ade, Aviación Deportiva"
,ADT,ARRENDA-TRANS,Arrendaminetos_y_Transportes_Turisticos,Arrendaminetos y Transportes Turisticos
,ADZ,,Avio_Delta,Avio Delta
,ADW,,ADC_Airways,ADC Airways
,AED,AIE EXPERIENCE,Aie_Experience_Flight,Aie Experience Flight
,AED,Aernspa,Aernord,Aernord
,AEG,AIR EAST,Air_East,Air East
,AEI,POLISH BIRD,Air_Italy_Polska,Air Italy Polska
,AEK,ACORISA,Aero_Costa_Rica,Aero Costa Rica
,AES,AEROPARAGUAY,Aerosur_Paraguay,Aerosur Paraguay
,AET,AEROTIME,Aerotime,Aerotime
VJ,AFF,AFRIWAYS,Africa_Airways,Africa Airways
,AFM,,Air_Four,Air Four
QH,FLA,PALM,Air_Florida,Air Florida
3O,MAC,ARABIA MAROC,Air_Arabia_Maroc,Air Arabia Maroc
,ADO,AIR DO,Air_Do,Air Do
,MRY,AIR MARINE,Air_Marine,Air Marine
,PNK,AIRPINK,Air_Pink,Air Pink
,AFN,SIMBA,African_International_Airlines,African International Airlines
,BDV,ABERDAV,Aberdair_Aviation,Aberdair Aviation
,TTN,TITANIUM,Absolute_Flight_Services,Absolute Flight Services
,ATZ,ACE TAXI,Ace_Air,Ace Air
,ARO,ACERO,Acero_Taxi,Acero Taxi
,CRV,ACROPOLIS,Acropolis_Aviation,Acropolis Aviation
,ADC,AD ASTRA,AD_Astra_Executive_Charter,AD Astra Executive Charter
,DDS,ADDIS LINE,Addis_Airlines,Addis Airlines
,TEC,TECHJET,ADI_Shuttle_Group,ADI Shuttle Group
,AXX,SKY SHUTTLE,Advance_Air_Luftfahrtgesellschaft,Advance Air Luftfahrtgesellschaft
,AAX,ADVANCE AVIATION,Advance_Aviation,Advance Aviation
,ADV,ADVANCED,Advanced_Flight_Training,Advanced Flight Training
,DVN,,Adventia,Adventia
,EBS,,AEG_Aviation_Services,AEG Aviation Services
,ALS,AERALP,Aeralp,Aeralp
,AEF,,Aerea,Aerea
,DRD,AEREO DORADO,Aereo_Dorado,Aereo Dorado
,FUT,AEREO FUTURO,Aereo_Futuro,Aereo Futuro
,MMG,RUTA MAYA,Aereo_Ruta_Maya,Aereo Ruta Maya
,AGI,ANGELES AMERICA,Aereo_Transportes_Los_Angeles_de_America,Aereo Transportes Los Angeles de America
,WWG,AERO-W,Aereo_WWG,Aereo WWG
,AKR,AERO CLINKER,Aero_Clinker,Aero Clinker
,RRB,,Aero_Club_de_Castellon,Aero Club de Castellon
,ARP,IVORYCORP,Aero_Corporate,Aero Corporate
,EPU,ELITACAPULCO,Aero_Elite_Acapulco,Aero Elite Acapulco
,XPN,,Aero_Express_(airline),Aero Express
,TEJ,,AeroJet,AeroJet
,GHM,,Aero_Service_Bolivia,Aero Service Bolivia
,GLM,GLOBAL MALI,Aero_Services_Mali,Aero Services Mali
,GUE,AERO GUERRERO,Aero_Servicio_Guerrero,Aero Servicio Guerrero
,ABA,AEROBETA,Aero-Beta,Aero-Beta
,AJH,ALJARAFE,Aeroaljarafe,Aeroaljarafe
,AOB,CARIBE CORO,Aerocaribe_Coro,Aerocaribe Coro
,ERC,,Aerocharter,Aerocharter
,BSO,,Aeroclub_Barcelona-Sabadell,Aeroclub Barcelona-Sabadell
,NVA,,Aeroclub_de_Albacete,Aeroclub de Albacete
,LUE,,Aeroclub_de_Alicante,Aeroclub de Alicante
,MLL,MALLORCA,Aeroclub_de_Mallorca,Aeroclub de Mallorca
,AVX,,Aeroclub_de_Vitoria,Aeroclub de Vitoria
,CTD,AEROCORPORATIVOS,Aerocorporativos,Aerocorporativos
,RRO,,Aerocredo,Aerocredo
,FAQ,,Aerofaq,Aerofaq
,PLS,AEROPLUS,Aeroflot-Plus,Aeroflot-Plus
,LIN,AEROLIMOUSINE,Aerolimousine,Aerolimousine
,PCP,PRINCIPAL,AerolÃ­nea_Principal_Chile,Aerolínea Principal Chile
,ALT,AERLINEAS CENTRALES,AerolÃ­neas_Centrales,Aerolíneas Centrales
,DMJ,DAMOJH,Aerol%C3%ADneas_Damojh,Aerolíneas Damojh
,AHL,HIDALGO,AerolÃ­neas_Hidalgo,Aerolíneas Hidalgo
,LDL,,Aerologic_(Russia),Aerologic
,APR,AEROPERLAS,AerolÃ­neas_Primordiales,Aerolíneas Primordiales
,VSC,,Aeronautic_de_Los_Pirineos,Aeronautic de Los Pirineos
HD,ADO,AIR DO,Air_Do,AIRDO
,AGM,ANGEL MED,Aviation_West_Charters,Aviation West Charters
X9,NVD,NORDVIND,Avion_Express,Avion Express
,BAH,BAHRAIN,The_Amiri_Flight,The Amiri Flight
UJ,LMU,ALMASRIA,Al_Masria_Universal_Airlines,Al Masria Universal Airlines
,AVD,ALAMO,"Ãlamo_AviaciÃ³n,_S.L.","Álamo Aviación, S.L."
,PNX,SPINNER,AIS_Airlines,AIS Airlines
,AHO,AIR HAMBURG,Air_Hamburg,Air Hamburg
,BBF,SPEEDCHARTER,B-Air_Charter,B-Air Charter
JD,CBJ,CAPITAL JET,Beijing_Capital_Airlines,Beijing Capital Airlines
,OUF,,Beijing_Eofa_International_Jet,Beijing Eofa International Jet
,OTA,OUTLAW,Business_Aviators,Business Aviators
,BWD,BLUEWEST,Bluewest_Helicopters-Greenland,Bluewest Helicopters-Greenland
,TBL,TELCO,Bell_Aliant_Regional_Communications,Bell Aliant Regional Communications
,CWR,CITY WORLD,Beijing_City_International_Jet,Beijing City International Jet
,BJV,BEIJING VISTA,Beijing_Vistajet_Aviation,Beijing Vistajet Aviation
,BHK,BLUEHAKIN,,Blu Halkin
,BXJ,BRIXTEL JET,Brixtel_Group,Brixtel Group
,BYG,BYGONE,Bygone_Aviation,Bygone Aviation
,BOB,BACKBONE,Backbone_A/S,Backbone A/S
ID,BTK,BATIK,Batik_Air,Batik Air
,BBJ,BLUE KOREA,Blue_Air_Lines,Blue Air Lines
,BCJ,BLUE BOY,Blue_Jet_Charters,Blue Jet Charters
,BNA,BUN AIR,Bun_Air_Corporation,Bun Air Corporation
,PVO,PROVOST,Bearing_Supplies_Limited,Bearing Supplies Limited
,BAA,BALKAN AGRO,Balkan_Agro_Aviation,Balkan Agro Aviation
,BAB,AWAL,Bahrain_Air_BSC_(Closed),Bahrain Air BSC (Closed)
,BAC,,BAC_Leasing_Limited,BAC Leasing Limited
,BAE,FELIX,BAE_Systems,BAE Systems
,BAF,BELGIAN AIRFORCE,Belgian_Air_Force,Belgian Air Force
8Q*,BAJ,RODEO,Baker_Aviation,Baker Aviation
,BAK,BLACKHAWK,Blackhawk_Airways,Blackhawk Airways
L9,BAL,BELLEAIR EUROPE,Belle_Air_Europe,Belle Air Europe
,BAL,BRITANNIA,Britannia_Airways,Britannia Airways
,BAM,BUSINESS AIR,Business_Air_Services,Business Air Services
,BAN,PENGUIN,British_Antarctic_Survey,British Antarctic Survey
,BAR,BRADLEY,Bradly_Air_(Charter)_Services,Bradly Air (Charter) Services
,BAU,AIR BISSAU,Bissau_Airlines,Bissau Airlines
,BAV,BAY AIR,Bay_Aviation_Ltd,Bay Aviation Ltd
BA,BAW,SPEEDBIRD,British_Airways,British Airways
,BAX,,Best_Aero_Handling_Ltd,Best Aero Handling Ltd
,BAY,,Bravo_Airways,Bravo Airways
,BBA,BANAIR,Bannert_Air,Bannert Air
BG,BBC,BANGLADESH,Biman_Bangladesh_Airlines,Biman Bangladesh Airlines
BO,BBD,BLUE CARGO,Bluebird_Cargo,Bluebird Nordic
,BBS,BEIBARS,Beibars_CJSC,Beibars CJSC
,BBV,BRAVO EUROPE,Bravo_Airlines,Bravo Airlines
,BBZ,COBRA,Bluebird_Aviation,Bluebird Aviation
B4,BCF,BACH,BACH_Flugbetriebsges,BACH Flugbetriebsges
,BCI,BLUE ISLAND,Blue_Islands,Blue Islands
,BCL,,British_Caribbean_Airways,British Caribbean Airways
,BCR,BACKER,British_Charter,British Charter
,BCT,BOBCAT,BCT_Aviation,BCT Aviation
,BCV,BUSINESS AVIATION,Business_Aviation_Center,Business Aviation Center
WX,BCY,CITY JET,CityJet,CityJet
BZ,BDA,BLUE DART,Blue_Dart_Aviation,Blue Dart Aviation
JA,BON,Air Bosna,B%26H_Airlines,B&H Airlines
,BDF,BISSAU DISCOVERY,Bissau_Discovery_Flying_Club,Bissau Discovery Flying Club
,AYB,BELGIAN ARMY,Belgian_Army,Belgian Army
,BDR,BADR AIR,Badr_Airlines,Badr Airlines
,BEA,BEST AIR,Best_Aviation_Ltd,Best Aviation Ltd
,BED,BELOGORYE,Belgorod_Aviation_Enterprise,Belgorod Aviation Enterprise
,AJC,BAR HARBOR,Bar_Harbor_Airlines,Bar Harbor Airlines
,BEF,BALEAR EXPRESS,Balear_Express,Balear Express
,BEH,BLUECOPTER,Bel_Air_Helicopters,Bel Air Helicopters
,BEK,BERKUT,Berkut_Air,Berkut Air
,BET,BETA CARGO,BETA_-_Brazilian_Express_Transportes_AÃ©reos,BETA - Brazilian Express Transportes Aéreos
,BFC,BASLER,Basler_Flight_Service,Basler Flight Service
,BFG,BEARFLIGHT,Bear_Flight,Bear Flight
J4,BFL,BUFFALO,Buffalo_Airways,Buffalo Airways
,BFO,BOMBARDIER,Bombardier_Aerospace,Bombardier
,BFR,BURKLINES,Burkina_Airlines,Burkina Airlines
,BFS,BUSINESS FLIGHT,Business_Flight_Sweden,Business Flight Sweden
,BFW,SUMMAN,Bahrain_Defence_Force,Bahrain Defence Force
8H,BGH,BALKAN HOLIDAYS,BH_Air,BH Air
,BGI,BRITISH GULF,British_Gulf_International,British Gulf International
,BGK,GULF INTER,British_Gulf_International-Fez,British Gulf International-Fez
A8,BGL,BENIN GOLF,Benin_Golf_Air,Benin Golf Air
,BGM,BUGAVIA,Bugulma_Air_Enterprise,Bugulma Air Enterprise
,BGR,BUDGET AIR,Budget_Air_Bangladesh,Budget Air Bangladesh
,BGT,BERGEN AIR,Bergen_Air_Transport,Bergen Air Transport
,BHA,BUDDHA AIR,Buddha_Air,Buddha Air
,BHI,SHARIF,Balkh_Airlines,Balkh Airlines
,BHL,BRISTOW,Bristow_Helicopters,Bristow Helicopters
,BHN,BRISTOW HELICOPTERS,Bristow_Helicopters_Nigeria,Bristow Helicopters Nigeria
,BHO,BHOJA,Bhoja_Airlines,Bhoja Airlines
4T,BHP,BELAIR,Belair_Airlines,Belair Airlines
,BHR,BIGHORN AIR,Bighorn_Airways,Bighorn Airways
UP,BHS,BAHAMAS,Bahamasair,Bahamasair
,BHT,BRIGHTAIR,Bright_Air,Bright Air
E6,,,Bringer_Air_Cargo_Taxi_AÃ©reo,Bringer Air Cargo Taxi Aéreo
LZ,,,Balkan_Bulgarian_Airlines,Balkan Bulgarian Airlines
TH,,,BA_Connect,BA Connect
,BHY,BOSPHORUS,Bosphorus_European_Airways,Bosphorus European Airways
,BID,BINAIR,Binair,Binair
,BIG,BIG ISLE,Big_Island_Air,Big Island Air
BS,BIH,BRINTEL,British_International_Helicopters,British International Helicopters
,BIL,BILAIR,Billund_Air_Center,Billund Air Center
,BIN,BISON-AIR,Boise_Interagency_Fire_Center,Boise Interagency Fire Center
,BIO,BIOFLIGHT,,Bioflight A/S
,BIR,BIRD AIR,Bird_Leasing,Bird Leasing
,BIZ,BIZZ,Bizjet_Ltd,Bizjet Ltd
,BJA,BAJA AIR,Baja_Air,Baja Air
,BJC,BALTIC JET,Baltic_Jet_Aircompany,Baltic Jet Aircompany
,BJS,SOLUTION,Business_Jet_Solutions,Business Jet Solutions
B4,BKA,BANKAIR,Bankair,Bankair
,BKF,BAKERFLIGHT,BF-Lento_OY,BF-Lento OY
,BKK,BLINKAIR,Blink_(airline),Blink
,BKJ,BARKEN JET,Barken_International,Barken International
PG,BKP,BANGKOK AIR,Bangkok_Airways,Bangkok Airways
,BKV,BUKOVYNA,Bukovyna,Bukovyna
,BLB,BLUEBIRD SUDAN,Blue_Bird_Aviation,Blue Bird Aviation
,BLC,BELLESAVIA,Bellesavia,Bellesavia
,BLE,BLUE BERRY,Blue_Line_(airline),Blue Line
KF,BLF,BLUEFIN,Blue1,Blue1
,BLG,BELGAVIA,Belgavia,Belgavia
,BLH,BLUE HORIZON,Blue_Horizon_Travel_Club,Blue Horizon Travel Club
,BLJ,BLUEWAY,Blue_Jet,Blue Jet
,BLL,BALTIC AIRLINES,Baltic_Airlines,Baltic Airlines
,BLN,BIAR,Bali_International_Air_Service,Bali International Air Service
JV,BLS,BEARSKIN,Bearskin_Lake_Air_Service,Bearskin Lake Air Service
,BLT,BALTAIR,Baltic_Aviation,Baltic Aviation
B3,BLV,BELLVIEW AIRLINES,Bellview_Airlines_(Nigeria),Bellview Airlines
BD,BMA,MIDLAND,BMI_(airline),BMI
BM,BMR,MIDLAND,BMI_Regional,BMI Regional
,BMD,BRITISH MEDICAL,British_Medical_Charter,British Medical Charter
,BME,BRIGGS,Briggs_Marine_Environmental_Services,Briggs Marine Environmental Services
,BMH,MASAYU,Bristow_Masayu_Helicopters,Bristow Masayu Helicopters
WW,BMI,BABY,Bmibaby,Bmibaby
CH,BMJ,BEMIDJI,Bemidji_Airlines,Bemidji Airlines
5Z,BML,BISMILLAH,Bismillah_Airlines,Bismillah Airlines
,BMN,BOWMAN,Bowman_Aviation,Bowman Aviation
,BMW,BMW-FLIGHT,BMW,BMW
,BMX,BANXICO,Banco_de_Mexico,Banco de Mexico
,BND,BOND,Bond_Offshore_Helicopters,Bond Offshore Helicopters
,BNE,BENINA AIR,Benina_Air,Benina Air
,BNG,VECTIS,BN_Group_Limited,BN Group Limited
,BNJ,JET BELGIUM,Air_Service_LiÃ¨ge_(ASL),Air Service Liège (ASL)
,BNL,NILE TRADING,Blue_Nile_Ethiopia_Trading,Blue Nile Ethiopia Trading
,BNR,BONAIR,Bonair_Aviation,Bonair Aviation
,BNS,BANCSTAR,Bancstar_(airline),Bancstar - Valley National Corporation
,BNT,BENTIU AIR,Bentiu_Air_Transport,Bentiu Air Transport
,BNV,BENANE,Benane_Aviation_Corporation,Benane Aviation Corporation
,BNW,BRITISH NORTH,British_North_West_Airlines,British North West Airlines
,BOA,KUMANOVO,Boniair,Boniair
,BOD,UGABOND,Bond_Air_Services,Bond Air Services
,BOE,BOEING,Boeing,Boeing
,BOF,BORDAIR,Bordaire,Bordaire
,BOO,BOOKAJET,Bookajet,Bookajet Limited
BO,BOU,BOURAQ,Bouraq_Indonesia_Airlines,Bouraq Indonesia Airlines
BV,BPA,BLUE PANOROMA,Blue_Panorama_Airlines,Blue Panorama Airlines
,BPK,VENERA,Berkhut_ZK,Berkhut ZK
,BPO,PIROL,Bundespolizei-Fliegertruppe,Bundespolizei-Fliegertruppe
,BPS,BASE,Budapest_Aircraft_Services,Budapest Aircraft Services/Manx2
,BPT,BONUS,Bonus_Aviation,Bonus Aviation
,BPX,,British_Petroleum_Exploration,British Petroleum Exploration
7R,BRB,BRA-TRANSPAEREOS,BRA-Transportes_A%C3%A9reos,BRA-Transportes Aéreos
,BRD,BROCK AIR,Brock_Air_Services,Brock Air Services
,BRE,AVIABREEZE,Breeze_Ltd,Breeze Ltd
8E,BRG,BERING AIR,Bering_Air,Bering Air
,BRK,BRIANSK-AVIA,Briansk_State_Air_Enterprise,Briansk State Air Enterprise
,BRN,BRANSON,Branson_Airlines,Branson Airlines
,BRO,COASTRIDER,BASE_Regional_Airlines,BASE Regional Airlines
,BRS,BRAZILIAN AIR FORCE,Brazilian_Air_Force,Brazilian Air Force
,BRT,BRITISH,British_Regional_Airlines,British Regional Airlines
B2,BRU,BELARUS AVIA,Belavia,Belavia Belarusian Airlines
,BRV,BRAVO,Bravo_Air_Congo,Bravo Air Congo
,BRW,BRIGHT SERVICES,Bright_Aviation_Services,Bright Aviation Services
BN,BNF,Braniff,Braniff_International_Airways,Braniff International Airways
,BRX,BUFF EXPRESS,Buffalo_Express_Airlines,Buffalo Express Airlines
,BRY,BURAIR,Burundayavia,Burundayavia
,BSC,BIG SHOT,Bistair_-_Fez,Bistair - Fez
,BSD,AIRLINES STAR,Blue_Star_Airlines,Blue Star Airlines
,BSI,BRASAIR,Brasair_Transportes_AÃ©reos,Brasair Transportes Aéreos
,BSJ,BLUE SWAN,Blue_Swan_Aviation,Blue Swan Aviation
,BSM,,Blue_Sky_Aviation,Blue Sky Aviation
,BSS,BISSAU AIRSYSTEM,Bissau_Aero_Transporte,Bissau Aero Transporte
,BST,TUNCA,Best_Air,Best Air
,BSW,SKY BLUE,Blue_Sky_Airways,Blue Sky Airways
GQ,BSY,BIG SKY,Big_Sky_Airlines,Big Sky Airlines
V9,BTC,BASHKIRIAN,BAL_Bashkirian_Airlines,BAL Bashkirian Airlines
,BTH,BALTIJAS HELICOPTERS,Baltijas_Helicopters,Baltijas Helicopters
,BTK,BALTYKA,Baltyka,Baltyka
BQ,BTL,BALTIA,Baltia_Air_Lines,Baltia Air Lines
,BTR,BOTIR-AVIA,Botir-Avia,Botir-Avia
,GAA,BIZEX,Business_Express,Business Express
,BTT,BEETEE-SLAVUTA,BT-Slavuta,BT-Slavuta
Y6,BTV,BATAVIA,Batavia_Air,Batavia Air
L9,BTZ,BRISTOW,Bristow_U.S._LLC,Bristow U.S. LLC
1T,BUC,BULGARIAN CHARTER,Bulgarian_Air_Charter,Bulgarian Air Charter
,BUL,BLUE AIRLINES,Blue_Airlines,Blue Airlines
BU,BUN,BURAL,Buryat_Airlines_Aircompany,Buryat Airlines Aircompany
,BUZ,BUZZ,Buzz_(airline),Buzz Stansted
,BVA,BUFFALO AIR,Buffalo_Airways,Buffalo Airways
,BVC,BULGARIAN WINGS,Bulgarian_Aeronautical_Centre,Bulgarian Aeronautical Centre
,BVN,SHOW-ME,Baron_Aviation_Services,Baron Aviation Services
J8,BVT,BERJAYA,Berjaya_Air,Berjaya Air
,BVU,,Bellview_Airlines_(Sierra_Leone),"Bellview Airlines, Sierra Leone"
QW,BWG,BLUE WINGS,Blue_Wings,Blue Wings
,BWI,BLUE TAIL,Blue_Wing_Airlines,Blue Wing Airlines
,BWL,BRITWORLD,British_World_Airlines,British World Airlines
,BXA,BEXAIR,Bahrain_Executive_Air_Services,Bahrain Executive Air Services
,BXH,PALLISER,Bar_XH_Air,Bar XH Air
SN,BXI,XENIA,Brussels_International_Airlines,Brussels International Airlines
8W?,,,BAX_Global,BAX Global
,BYA,BERRY,Berry_Aviation,Berry Aviation
,BYC,Bayon Air,Cambodia_Bayon_Airlines,Cambodia Bayon Airlines
,BYF,BAY FLIGHT,,San Carlos Flight Center
,BYL,BYLINA,Bylina,Bylina Joint-Stock Company
,BYR,,Berytos_Airlines,Berytos Airlines
,BYE,BAYU,Bayu_Indonesia_Air,Bayu Indonesia Air
,BZA,BERLIN BEAR,Bizair_Fluggesellschaft,Bizair Fluggesellschaft
DB,BZH,BRITAIR,Brit_Air,Brit Air
,BZZ,BUZZARD,Butane_Buzzard_Aviation_Corporation,Butane Buzzard Aviation Corporation
,AUJ,AUSTROJET,Business_Flight_Salzburg,Business Flight Salzburg
,CKM,COSMOS,BKS_Air,BKS Air (Rivaflecha)
,CLF,CLIFTON,Bristol_Flying_Centre,Bristol Flying Centre
,CLN,SEELINE,Barnes_Olsen_Aeroleasing,Barnes Olsen Aeroleasing
,CPJ,CORPJET,Baltimore_Air_Transport,Baltimore Air Transport
E9,CXS,CLIPPER CONNECTION,Boston-Maine_Airways,Boston-Maine Airways
SN,BEL,BEE-LINE,Brussels_Airlines,Brussels Airlines
,EAH,EASTERN,Baltimore_Airways,Baltimore Airways
,EBA,BOND AVIATION,Bond_Aviation,Bond Aviation
,EXB,BRAZILIAN ARMY,Brazilian_Army_Aviation,Brazilian Army Aviation
,EXP,EXPRESS AIR,Business_Express_Delivery,Business Express Delivery
,FOS,,Bel_Limited,Bel Limited
,HAW,THAI HAWK,Bangkok_Aviation_Center,Bangkok Aviation Center
,HAX,SCOOP,Benair,Benair
NT,IBB,BINTER,Binter_Canarias,Binter Canarias
,IRJ,BONYAD AIR,Bonyad_Airlines,Bonyad Airlines
,IVR,RERUN,Burundaiavia,Burundaiavia
0B,BMS,BLUE MESSENGER,Blue_Air,Blue Air
KJ,LAJ,BEE MED,British_Mediterranean_Airways,British Mediterranean Airways
,LBY,ALBAN-BELLE,Belle_Air,Belle Air
,LED,SWEEPER,Blom_Geomatics,Blom Geomatics
,LTL,LITTORAL,Benin_Littoral_Airways,Benin Littoral Airways
,LXJ,FLEXJET,Bombardier_Business_Jet_Solutions,Bombardier Business Jet Solutions
FB,LZB,FLYING BULGARIA,Bulgaria_Air,Bulgaria Air
,MBR,BRAZILIAN NAVY,Brazilian_Navy_Aviation,Brazilian Navy Aviation
8N,NKF,NORDFLIGHT,Barents_AirLink,Barents AirLink
,NYB,BELGIAN NAVY,Belgian_Navy,Belgian Navy
,OGJ,BAKO AIR,Bakoji_Airlines_Services,Bakoji Airlines Services
,PEB,PALEMA,Benders_Air,Benders Air
,PNT,PORTNET,Balmoral_Central_Contracts,Balmoral Central Contracts
,POI,BOJBAN,BGB_Air,BGB Air
,PPS,PIPESTONE,Butte_Aviation,Butte Aviation
,RHD,RED HEAD,Bond_Air_Services,Bond Air Services
,RLR,RATTLER,Business_Airfreight,Business Airfreight
,RPX,RAPEX,BAC_Express_Airlines,BAC Express Airlines
,RRS,BLACKBOX,Boscombe_Down_DERA,Boscombe Down DERA (Formation)
,SCJ,SCANJET,Business_Jet_Sweden,Business Jet Sweden
,SHT,SHUTTLE,British_Airways,British Airways Shuttle
,SKH,SKYNEWS,British_Sky_Broadcasting,British Sky Broadcasting
,TXB,TEXTRON,Bell_Helicopter_Textron,Bell Helicopter Textron
,UKA,UKAY,Buzzaway_Limited,Buzzaway Limited
,VLX,AVOLAR,Biz_Jet_Charter,Biz Jet Charter
,VOL,BLUE SPEED,Blue_Chip_Jet,Blue Chip Jet
,WFD,AVRO,BAE_Systems,BAE Systems
,WTN,TARNISH,BAE_Systems,BAE Systems
,XBO,,Baseops_International,Baseops International
,XDA,,Bureau_Veritas,Bureau Veritas
,XMS,SANTA,British_Airways,British Airways Santa
,ZBA,BOSKY,Boskovic_Air_Charters,Boskovic Air Charters Limited
,BLM,BLUE ARMENIA,Blue_Sky_Airlines,Blue Sky Airlines
,JMP,JUMP RUN,Businesswings,Businesswings
CJ,CFE,FLYER,BA_CityFlyer,BA CityFlyer
OB,BOV,BOLIVIANA,Boliviana_de_Aviaci%C3%B3n,Boliviana de Aviación
YB,BRJ,BORA JET,Borajet,Borajet
BW,BWA,CARIBBEAN AIRLINES,Caribbean_Airlines,Caribbean Airlines
,QAI,CHICKPEA,Conquest_Air,Conquest Air
,CNW,,Continental_Airways_(Modova),Continental Airways
,CRF,CROIX ROUGE,Croix_Rouge_Francais,Croix Rouge Francais
,CBH,CLUB HOUSE,Corporate_Eagle_Management_Services,Corporate Eagle Management Services
,CCT,CONNECT,Connect_Air_(Canada),Connect Air
,CME,COMMERCE BANK,Commerce_Bancshares,Commerce Bank
,CAA,INSPECTOR,Civil_Aviation_Authority_of_the_Czech_Republic,Civil Aviation Authority of the Czech Republic
,BKR,BOX KAR,Civil_Air_Patrol,Civil Air Patrol South Carolina Wing
,BFN,,Compagnie_Nationale_Naganagani,Compagnie Nationale Naganagani
,AWX,ALLWEATHER,Civil_Aviation_Authority_(United_Kingdom),Civil Aviation Authority Directorate of Airspace Policy
,BBN,BRABAZON,Civil_Aviation_Authority_(United_Kingdom),Civil Aviation Authority Airworthiness Division
,ATQ,COLIBRI,CHC_Helicopters_Nigeria,CHC Helicopters Nigeria
,APL,AEREO PRINCIPAL,Corporativo_Aereo_Principal,Corporativo Aereo Principal
,AIO,AIR CHIEF,United_States_Air_Force,"Chief of Staff, United States Air Force"
,AID,CENTURY AIRBIRD,Christian_Konig_-_Century_Airbirds,Christian Konig - Century Airbirds
,,CAMBRIDGE,Cambridge_Recurrent_Training,Cambridge Recurrent Training
,,CALFIRE,California_Department_of_Forestry_and_Fire_Protection,California Department of Forestry and Fire Protection
,,SAMARITAN,Careflight_Queensland,Careflight Queensland
7N,PWD,Pawa Dominicana,Pawa_Dominicana,Pawa Dominicana
,,CASTLEFILM,Castle_Air_Charter,Castle Air Charter
,,TIGER,Challenges_Aviation,Challenges Aviation
,,HERITAGE,Challenges_Limited,Challenges Limited
,,ASTON,Cheqair,Cheqair
,,CLASSIC WINGS,Clacton_Aero_Club,Clacton Aero Club
,,CORAL SUN,Coral_Sun_Airways,Coral Sun Airways
,,COWAN,Cowan_Group,Cowan Group
,SMW,SMART WINGS,Carpatair_Flight_Service,Carpatair Flight Service
,CCL,ANGKOR WAT,Cambodia_Airlines,Cambodia Airlines
,CCB,DOLPHIN,Caricom_Airways_(Barbados),Caricom Airways
,CYH,YUHAO,China_Southern_Airlines_Henan,China Southern Airlines Henan
,CFB,FOREBASE,Chongqing_Forebase_General_Aviation,Chongqing Forebase General Aviation
,XCA,COLT,Colt_Transportes_Aereos,Colt Transportes Aereos
,GCY,HELIBIRD,CHC_Global_Operations_International,CHC Global Operations International
,TIP,TRANSPAC,C_and_M_Aviation,C and M Aviation
,SRJ,SYRJET,C_Air_Jet_Airlines,C Air Jet Airlines
,ORO,CAPRI,C_N_Air,C N Air
,RWG,RED WING,C&M_Airways,C&M Airways
,RMU,AIR-MAUR,"C.S.P.,_Societe","C.S.P., Societe"
,CJZ,CALIBER JET,Caliber_Jet,Caliber Jet
,CRC,CAMAIRCO,Cameroon_Airlines_Corporation,Cameroon Airlines Corporation
,AUN,COMMON SKY,Common_Sky,Common Sky
,CBI,CABI,Cabi_(airline),Cabi
,CPI,AIRCAI,Compagnia_Aerea_Italiana,Compagnia Aerea Italiana
5C,ICL,CAL,CAL_Cargo_Air_Lines,CAL Cargo Air Lines
,CMR,CAMEO,CAM_Air_Management,CAM Air Management
,CTZ,CATA,CATA_L%C3%ADnea_A%C3%A9rea,CATA Línea Aérea
,CCF,TOMCAT,CCF_Manager_Airline,CCF Manager Airline
,CED,CEDTA,CEDTA,CEDTA (Compañía Ecuatoriana De Transportes Aéreos)
,HBI,HELIBIRD,CHC_Denmark,CHC Denmark
,HEM,HEMS,CHC_Helicopter,CHC Helicopter
AW,SCH,SCHREINER,CHC_Airways,CHC Airways
,HKS,HELIBUS,CHC_Helikopter_Service,CHC Helikopter Service
,VCI,CI-TOURS,CI-Tours,CI-Tours
,CKC,,CKC_Services,CKC Services
,CMZ,CEE-EM STAIRS,CM_Stair,CM Stair
,CNT,KNET,Centre_national_d%27%C3%A9tudes_des_t%C3%A9l%C3%A9communications,Centre national d'études des télécommunications - C.N.E.T.
,OAP,COAPA,COAPA_AIR,COAPA AIR
,PDR,SPEEDSTER,COMAV,COMAV
,CRH,HELI-MEX,CRI_Helicopters_Mexico,CRI Helicopters Mexico
,IRO,IRON AIR,CSA_Air,CSA Air
,CSE,OXFORD,CSE_Aviation,CSE Aviation
,CTQ,CITYLINK,CTK_-_CiTylinK,CTK Network Aviation
,CBR,CABAIR,Cabair_College_of_Air_Training,Cabair College of Air Training
,CVE,KABEX,Cabo_Verde_Express,Cabo Verde Express
,CWD,AMBASSADOR,Caernarfon_Airworld,Caernarfon Airworld
,CXE,CAICOS,Caicos_Express_Airways,Caicos Express Airways
,CCE,,Cairo_Air_Transport_Company,Cairo Air Transport Company
,CGC,CAL-GULF,Cal_Gulf_Aviation,Cal Gulf Aviation
,REZ,CAL AIR,Cal-West_Aviation,Cal-West Aviation
,CSL,CALIFORNIA SHUTTLE,California_Air_Shuttle,California Air Shuttle
3C,CMV,CALIMA,Calima_AviaciÃ³n,Calima Aviación
MO,CAV,CALM AIR,Calm_Air,Calm Air
R9,CAM,AIR CAMAI,Camai_Air,Camai Air
K6,KHV,ANGKOR AIR,Cambodia_Angkor_Air,Cambodia Angkor Air
UY,UYC,CAM-AIR,Cameroon_Airlines,Cameroon Airlines
,HSO,HELIASTURIAS,Campania_Helicopteros_De_Transporte,Campania Helicopteros De Transporte
C6,CJA,CANJET,CanJet,CanJet
,PIL,PINNACLE,Canada_Jet_Charters,Canada Jet Charters
CP,CDN,CANADIAN,Canadian_Airlines,Canadian Airlines
,CTG,CANADIAN COAST GUARD,Canadian_Coast_Guard,Canadian Coast Guard
,HIA,HAIDA,Canadian_Eagle_Aviation,Canadian Eagle Aviation
,CFC,CANFORCE,Canadian_Forces,Canadian Forces
,BZD,BLIZZARD,Canadian_Global_Air_Ambulance,Canadian Global Air Ambulance
,CDN,CANADIAN,CHC_Helicopter,Canadian Helicopters
,TKR,TANKER,Canadian_Interagency_Forest_Fire_Centre,Canadian Interagency Forest Fire Centre
,XNC,,Canadian_National_Telecommunications,Canadian National Telecommunications
5T,MPE,EMPRESS,Canadian_North,Canadian North
CP,CPC,EMPRESS,Canadian_Pacific_Airlines,Canadian Pacific Airlines
,CDR,CANADIAN REGIONAL,Canadian_Regional_Airlines,Canadian Regional Airlines
,CWH,WARPLANE HERITAGE,Canadian_Warplane_Heritage_Museum,Canadian Warplane Heritage Museum
W2,CWA,CANADIAN WESTERN,Canadian_Western_Airlines,Canadian Western Airlines
,CWW,CANAIR,Canair,Canair
,CUI,CAN-AIR,Cancun_Air,Cancun Air
9K,KAP,CAIR,Cape_Air,Cape Air
,CTO,,Cape_Air_Transport,Cape Air Transport
,SEM,SEMO,Cape_Central_Airways,Cape Central Airways
,CMY,CAPE SMYTHE AIR,Cape_Smythe_Air,Cape Smythe Air
,CPX,CAPAIR,Capital_Air_Service,Capital Air Service
,CPD,CAPITAL DELTA,Capital_Airlines,Capital Airlines
,NCP,CAPITAL SHUTTLE,Capital_Airlines_(Nigeria),Capital Airlines Limited
PT,CCI,CAPPY,Capital_Cargo_International_Airlines,Capital Cargo International Airlines
,CCQ,CAP CITY,Capital_City_Air_Carriers,Capital City Air Carriers
,EGL,PRESTIGE,Capital_Trading_Aviation,Capital Trading Aviation
,CEX,CAPITOL EXPRESS,Capitol_Air_Express,Capitol Air Express
,CWZ,CAPWINGS,Capitol_Wings_Airline,Capitol Wings Airline
,VAN,CAMEL,Caravan_Air,Caravan Air
,CWN,CAMBRIAN,Cardiff_Wales_Flying_Club,Cardiff Wales Flying Club
,FVA,AIR VIRGINIA,Cardinal/Air_Virginia,Cardinal/Air Virginia
,GOL,CARGOLAAR,Cardolaar,Cardolaar
,CDI,CARDS,Cards_Air_Services,Cards Air Services
,CFH,CARE FLIGHT,Care_Flight,Care Flight
,CDM,CARGA AEREA,Carga_AÃ©rea_Dominicana,Carga Aérea Dominicana
,EST,CARGAINTER,Carga_Express_Internacional,Carga Express Internacional
GG,GGC,LONG-HAUL,Cargo_360,Cargo 360
,MCX,MAURICARGO,Cargo_Express,Cargo Express
,CRV,CARGOIV,Cargo_Ivoire,Cargo Ivoire
,CLM,CARGO LINK,Cargo_Link_(Caribbean),Cargo Link (Caribbean)
,CLA,FIREBIRD,Cargo_Logic_Air,Cargo Logic Air
,CTW,THIRD CARGO,Cargo_Three,Cargo Three
2G,CRG,WHITE PELICAN,Cargoitalia,Cargoitalia
W8,CJT,CARGOJET,Cargojet_Airways,Cargojet Airways
CV,CLX,CARGOLUX,Cargolux,Cargolux
C8,ICV,CARGO MED,Cargolux_Italia,Cargolux Italia
,CGM,HOTEL CHARLIE,Cargoman,Cargoman
,DEL,RED TAIL,Carib_Aviation,Carib Aviation
,BCB,WAVEBIRD,Carib_Express,Carib Express
,PWD,CARIBAIR,CARIBAIR,CARIBAIR
,DCC,CARICARGO,Caribbean_Air_Cargo,Caribbean Air Cargo
,CLT,CARIBBEAN,Caribbean_Air_Transport,Caribbean Air Transport
,CLT,,Club_Aerocelta_de_Vuelo_Con_Motor,Club Aerocelta de Vuelo Con Motor
BW,BWA,CARIBBEAN,Caribbean_Airlines,Caribbean Airlines
,IQQ,CARIBJET,Caribbean_Airways,Caribbean Airways
,CSX,CHOICE AIR,Choice_Airways,Choice Airways
,TLC,CARIB-X,Caribbean_Express,Caribbean Express
8B,GFI,CARIB STAR,Caribbean_Star_Airlines,Caribbean Star Airlines
,CRB,CARIBBEAN COMMUTER,Caricom_Airways,Caricom Airways
,CRT,CARIBINTAIR,Caribintair,Caribintair
,CVG,CARILL,Carill_Aviation,Carill Aviation
V3,KRP,CARPATAIR,Carpatair,Carpatair
,CRR,CARRANZA,Carranza_(airline),Carranza
,ULS,ULSTER,Carroll_Air_Service,Carroll Air Service
,CMT,CASEMENT,Casement_Aviation,Casement Aviation
,CSO,CASAIR,Casino_Airline,Casino Airline
,CSP,CASPER AIR,Casper_Air_Service,Casper Air Service
RV,CPN,CASPIAN,Caspian_Airlines,Caspian Airlines
,CSJ,CASTLE,Castle_Aviation,Castle Aviation
,CAZ,EUROCAT,Cat_Aviation,Cat Aviation
,CBT,CATALINA AIR,Catalina_Flying_Boats,Catalina Flying Boats
,TEX,CATEX,Catex,Catex
KA,HDA,DRAGON,Cathay_Dragon,Cathay Dragon
CX,CPA,CATHAY,Cathay_Pacific,Cathay Pacific
,CJR,CAVERTON AIR,Caverton_Helicopters,Caverton Helicopters
KX,CAY,CAYMAN,Cayman_Airways,Cayman Airways
5J,CEB,CEBU,Cebu_Pacific,Cebu Pacific
,CIL,CECIL,Cecil_Aviation,Cecil Aviation
,CEG,CEGA,Cega_Aviation,Cega Aviation
,CEC,CELTAIR,Celtic_Airways,Celtic Airways
,CWE,CELTIC,Celtic_West,Celtic West
,CEV,CENTEV,Centre_d%27Essais_en_Vol,Centre d'Essais en Vol
,CNL,WYO-AIR,Centennial_Airlines,Centennial Airlines
,CNS,CHRONOS,Cobalt_Air_LLC,Cobalt Air LLC
,CVO,CENTERVOL,Center_Vol,Center Vol
,CTS,CENTER-SOUTH,Center-South,Center-South
,CET,CENTRAFRICAIN,Centrafrican_Airlines,Centrafrican Airlines
,CAX,CENTRAL EXPRESS,Central_Air_Express,Central Air Express
,CTL,CENTRAL COMMUTER,Central_Airlines,Central Airlines
,CNY,CENTRAL LEONE,Central_Airways,Central Airways
,ACN,AEROCENTRO,Central_American_Airlines,Central American Airlines
,YOG,YOGAN AIR,Central_Aviation,Central Aviation
,DRN,DISCOS REYNOSA,Central_De_Discos_De_Reynosa,Central De Discos De Reynosa
,CMA,EUROCENTRAL,Central_European_Airlines,Central European Airlines
,CHA,CHARTER CENTRAL,Central_Flying_Service,Central Flying Service
,CEM,CENTRAL MONGOLIA,Central_Mongolia_Airways,Central Mongolia Airways
9M,GLR,GLACIER,Central_Mountain_Air,Central Mountain Air
,CQC,,Central_Queensland_Aviation_College,Central Queensland Aviation College
,CSI,SKYPORT,Central_Skyport,Central Skyport
,CLW,CENTRALWINGS,Centralwings,Centralwings
,DTV,DUTCH VALLEY,Centre_Airlines,Centre Airlines
,CGS,GEO CENTRE,Centre_of_Applied_Geodynamica,Centre of Applied Geodynamica
J7,CVC,AVIACENTRE,Centre-Avia,Centre-Avia
,CCV,HELICORPORATIVO,Centro_De_Helicopteros_Corporativos,Centro De Helicopteros Corporativos
,ACF,FORCAN,Centro_de_FormaciÃ³n_AeronÃ¡utica_de_Canarias,Centro de Formación Aeronáutica de Canarias
WE,CWC,CHALLENGE CARGO,Centurion_Air_Cargo,Centurion Air Cargo
,URY,CENTURY AVIA,Century_Aviation,Century Aviation
,HAI,,Century_Aviation_International,Century Aviation International
,XAD,,Certified_Air_Dispatch,Certified Air Dispatch
,CER,CETRACA,Cetraca_Aviation_Service,Cetraca Aviation Service
,IRU,CHABAHAR,Chabahar_Airlines,Chabahar Airlines
,CLG,CHALLAIR,Chalair_Aviation,Chalair Aviation
OP,CHK,CHALKS,Chalk%27s_International_Airlines,Chalk's International Airlines
,CLS,AIRISTO,Challenge_Air_Transport,Challenge Air Transport
,CHS,CHALLENGE AVIATION,Challenge_Aviation,Challenge Aviation
,OFF,CHALLENGE AIR,Challenge_International_Airlines,Challenge International Airlines
,CHG,SKY CHALLENGER,Challenge_Aero,Challenge Aero
,CPH,CHAMPAGNE,Champagne_Airlines,Champagne Airlines
MG,CCP,CHAMPION AIR,Champion_Air,Champion Air
,NCH,CHANCHANGI,Chanchangi_Airlines,Chanchangi Airlines
2Z,CGN,CHANGAN,Chang_An_Airlines,Chang An Airlines
,CHN,CHANNEL,Channel_Island_Aviation,Channel Island Aviation
,WML,MARLIN,Chantilly_Air,Chantilly Air
,CPL,CHAPARRAL,Chaparral_Airlines,Chaparral Airlines
S8,CSU,CHARI SERVICE,Chari_Aviation_Services,Chari Aviation Services
,CAH,CHARLAN,Charlan_Air_Charter,Charlan Air Charter
,HMD,HAMMOND,Charlie_Hammonds_Flying_Service,Charlie Hammonds Flying Service
,CGD,,Charlotte_Air_National_Guard,Charlotte Air National Guard
,CHW,CHARTER WIEN,Charter_Air,Charter Air
,HRT,CHARTRIGHT,Chartright_Air,Chartright Air
RP*,CHQ,CHAUTAUQUA,Chautauqua_Airlines,Chautauqua Airlines
,CBB,CHEBAIR,Cheboksary_Airenterprise,Cheboksary Airenterprise JSC
,CHM,,Chemech_Aviation,Chemech Aviation
,CHZ,CHERL,Cherline,Cherline
,CMK,CHERAVIA,Chernomor-Avia,Chernomor-Avia
,CBM,BLUE MAX,Cherokee_Express,Cherokee Express
,CCY,CHERRY,Cherry_Air,Cherry Air
,CAB,CHESAPEAKE AIR,Chesapeake_Air_Service,Chesapeake Air Service
,CVR,CHEVRON,Chevron_U.S.A,Chevron U.S.A
,CYA,CHEYENNE AIR,Cheyenne_Airways,Cheyenne Airways
,CGO,WILD ONION,Chicago_Air,Chicago Air
C8,WDY,WINDY CITY,Chicago_Express_Airlines,Chicago Express Airlines
,RAT,RIVERRAT,Chief_Rat_Flight_Services,Chief Rat Flight Services
,CCH,CHILCHOTA,Chilchota_Taxi_AÃ©reo,Chilchota Taxi Aéreo
,DES,CHILCOTIN,Chilcotin_Caribou_Aviation,Chilcotin Caribou Aviation
,CAD,CHILLIWACKAIR,Chilliwack_Aviation,Chilliwack Aviation
,ETN,CHIMNIR,Chim-Nir_Aviation,Chim-Nir Aviation
CI,CAL,DYNASTY,China_Airlines,China Airlines
CK,CKK,CARGO KING,China_Cargo_Airlines,China Cargo Airlines
MU,CES,CHINA EASTERN,China_Eastern_Airlines,China Eastern Airlines
G5,HXA,CHINA EXPRESS,China_Express_Airlines,China Express Airlines
,CFA,FEILONG,China_Flying_Dragon_Aviation,China Flying Dragon Aviation
,CTH,TONGHANG,China_General_Aviation,China General Aviation Corporation
,CAG,CHINA NATIONAL,China_National_Aviation_Corporation,China National Aviation Corporation
CJ,CBF,CHINA NORTHERN,China_Northern_Airlines,China Northern Airlines
WH,CNW,CHINA NORTHWEST,China_Northwest_Airlines,China Northwest Airlines
,CHC,CHINA HELICOPTER,China_Ocean_Helicopter_Corporation,China Ocean Helicopter Corporation
8Y,CYZ,CHINA POST,China_Postal_Airlines,China Postal Airlines
CZ,CSN,CHINA SOUTHERN,China_Southern_Airlines,China Southern Airlines
,CXN,CHINA SOUTHWEST,China_Southwest_Airlines,China Southwest Airlines
KN,CUA,LIANHANG,China_United_Airlines,China United Airlines
XO,CXH,XINHUA,China_Xinhua_Airlines,China Xinhua Airlines
3Q,CYH,YUNNAN,China_Yunnan_Airlines,China Yunnan Airlines
,CGU,CHINGUETTI,Chinguetti_Airlines,Chinguetti Airlines
,CEP,CHIPOLA,Chipola_Aviation,Chipola Aviation
,CPW,CHIPPEWA-AIR,Chippewa_Air_Commuter,Chippewa Air Commuter
X7,CHF,CHITA,Chitaavia,Chitaavia
,CAS,CHRISTMAN,Christman_Air_System,Christman Air System
,OEC,CHRISTOPHORUS,Christophorus_Flugrettungsverein,Christophorus Flugrettungsverein
,CHO,CHROME AIR,Chrome_Air_Service,Chrome Air Services
,CHU,CHURCHAIR,Church_Aircraft,Church Aircraft
A2,CIU,CIELOS,Cielos_Airlines,Cielos Airlines
QI,CIM,CIMBER,Cimber_Sterling,Cimber Sterling
C7,CIN,CINNAMON,Cinnamon_Air,Cinnamon Air
,RRU,HELICIRRUS,Cirrus_(Chile),Cirrus
,NTS,NITE STAR,Cirrus_Air,Cirrus Air
C9,RUS,CIRRUS AIR,Cirrus_Airlines,Cirrus Airlines
,JTI,,Cirrus_Middle_East,Cirrus Middle East
,FIV,FIVE STAR,CitationAir,CitationAir
,XCX,,Citibank,Citibank
,HZX,ZHONGXIN,Citic_General_Aviation,Citic General Aviation
CF,SDR,SWEDESTAR,City_Airline,City Airline
G3,CIX,CONNEXION,City_Connexion_Airlines,City Connexion Airlines
,XBG,,City_of_Bangor_(airline),City of Bangor
WX,BCY,CITY-IRELAND,CityJet,CityJet
,CAQ,AIR CHESTER,Cityair_(Chester)_Limited,Cityair (Chester) Limited
,CII,CITYFLY,Cityfly,Cityfly
CJ,CFE,FLYER,CityFlyer_Express,CityFlyer Express
,CNB,CITYHUN,Cityline_Hungary,Cityline Hungary
,HSR,HOOSIER,Citylink_Airlines,Citylink Airlines
,CIW,CIVFLIGHT,Civair,Civair Airways
,CAP,CAP,Civil_Air_Patrol,Civil Air Patrol
CT,CAT,Mandarin,Civil_Air_Transport,Civil Air Transport
,CIA,CALIMERA,Civil_Aviation_Authority_(Slovak_Republic),Civil Aviation Authority
,CIV,CIVAIR,Civil_Aviation_Authority_of_New_Zealand,Civil Aviation Authority of New Zealand
,CBA,CALIBRA,Civil_Aviation_Inspectorate_of_the_Czech_Republic,Civil Aviation Inspectorate of the Czech Republic
,FMC,CLAESSENS,Claessens_International_Limited,Claessens International Limited
,CLK,CLARKAIR,Clark_Aviation,Clark Aviation
,CSF,CALEDONIAN,Clasair,Clasair
,CLY,CLAY-LACY,Clay_Lacy_Aviation,Clay Lacy Aviation
,CGK,CLICK AIR,Click_Airways,Click Airways
,CLZ,CLOUDLINE,Cloud_9_Air_Charters,Cloud 9 Air Charters
,CLD,CLOWES,Clowes_Estates_Limited,Clowes Estates Limited
,SDJ,SPACEJET,Club_328,Club 328
6P,ISG,CLUBAIR,Club_Air,Club Air
BX,CST,COAST CENTER,Coast_Air,Coast Air
DQ,,U.S. Virgin Islands,Coastal_Air,Coastal Air
,TCL,TRANS COASTAL,Coastal_Air_Transport,Coastal Air Transport
,CNG,SID-AIR,Coastal_Airways,Coastal Airways
,CSV,COASTAL TRAVEL,Coastal_Aviation,Coastal Travels
,CHL,COHLMIA,Cohlmia_Aviation,Cohlmia Aviation
,OLR,COLAEREOS,ColaÃ©reos,Colaéreos
,CLE,COLEMILL,Colemill_Enterprises,Colemill Enterprises
9L,CJC,COLGAN,Colgan_Air,Colgan Air
,CAE,HUMMINGBIRD,Colibri_Aviation,Colibri Aviation
YD,CAT,,Cologne_Air_Transport_GmbH,Cologne Air Transport GmbH
,CCX,,Colt_International,Colt International
,WCO,COLUMBIA HELI,Columbia_Helicopters,Columbia Helicopters
,KLR,KAY-LER,Columbus_Air_Transport,Columbus Air Transport
,GHP,GRASSHOPPER EX,Colvin_Aviation,Colvin Aviation
OH,COM,COMAIR,Comair,Comair
MN,CAW,COMMERCIAL,Comair_(South_Africa),Comair
,GCM,GLOBECOM,Comair_Flight_Services,Comair Flight Services
,CDE,COMEX,Comed_Group,Comed Group
,CVV,COMERAVIA,Comeravia,Comeravia
,CRS,COMERCIAL AEREA,Comercial_AÃ©rea,Comercial Aérea
,CMG,SUNSPY,Comet_Airlines,Comet Airlines
,FYN,FLYNN,Comfort_Air,Comfort Air
,CMJ,COMFORT JET,Comfort_Jet_Services,Comfort Jet Services
,CLA,COMLUX,Comlux_Aviation,Comlux Aviation
,CMH,COMMODORE,Commair_Aviation,Commair Aviation
,CTM,COTAM,Commandement_Du_Transport_Aerien_Militaire_Francais,Commandement Du Transport Aerien Militaire Francais
,CML,COMMANDAIR,Commander_Air_Charter,Commander Air Charter
,CRM,COMMANDERMEX,Commander_Mexicana,Commander Mexicana
,CMS,ACCESS,Commercial_Aviation_(Canadian_airline),Commercial Aviation
,GAR,,Commodore_Aviation,Commodore Aviation
,CJS,COMMONWEALTH,Commonwealth_Jet_Service,Commonwealth Jet Service
C5,UCA,COMMUTAIR,CommutAir,CommutAir
KR,CWK,CONTICOM,Comores_Airlines,Comores Airlines
,CGR,COMPRIP,Compagnia_Generale_Ripreseaeree,Compagnia Generale Ripreseaeree
,CMM,CAMALI,Air_Mali_(2005),Compagnie Aérienne du Mali
,CPM,,Compagnie_Mauritanienne_Des_Transports,Compagnie Mauritanienne Des Transports
,GIC,CEBEGE,Compagnie_de_Bauxites_de_Guinee,Compagnie de Bauxites de Guinee
,AIF,,CompaÃ±Ã­a_AÃ©rea_de_Valencia,Compañía Aérea de Valencia
,ATF,AEROTECNICAS,CompaÃ±Ã­a_AerotÃ©cnicas_FotogrÃ¡ficas,Compañía Aerotécnicas Fotográficas
,LCT,STELLAIR,CompaÃ±Ã­a_De_Actividades_Y_Servicios_De_AviaciÃ³n,Compañía De Actividades Y Servicios De Aviación
,EJV,EJECUTIVA,Compania_Ejecutiva,Compania Ejecutiva
,HSE,HELISURESTE,Compania_Helicopteros_Del_Sureste,Compania Helicopteros Del Sureste
,MDR,AEROPLANOS,Compania_Mexicana_De_Aeroplanos,Compania Mexicana De Aeroplanos
GJ,MXC,MEXICARGO,Compania_Mexicargo,Compania Mexicargo
,HSS,TAS HELICOPTEROS,CompaÃ±Ã­a_Transportes_AÃ©reos_Del_Sur,Compañía Transportes Aéreos Del Sur
,TAV,TAVISA,CompaÃ±Ã­a_de_Servicios_AÃ©reos_Tavisa,Compañía de Servicios Aéreos Tavisa
,CYF,COMPANY FLIGHT,Company_Flight,Company Flight
CP,CPZ,COMPASS ROSE,Compass_Airlines_(North_America),Compass Airlines
,CPS,COMPASS,Compass_International_Airways,Compass International Airways
,XCO,,Compuflight_Operations_Service,Compuflight Operations Service
,XCS,,Compuserve_Incorporated,Compuserve Incorporated
,CRC,CONAIR-CANADA,Conair_Group,Conair Aviation
,COD,CONCORDAVIA,Concordavia,Concordavia
,CNR,CONAERO,Condor_Aero_Services,Condor Aero Services
,CIB,CONDOR BERLIN,Condor_Flugdienst,Condor
DE,CFG,CONDOR,Condor_Flugdienst,Condor Flugdienst
,COF,CONFORT,Confort_Air,Confort Air
,CAK,,Congo_Air,Congo Air
,CGA,CONGRESSIONAL,Congressional_Air,Congressional Air
,ROY,,Conifair_Aviation,Conifair Aviation
,BSN,BASTION,Connectair_Charters,Connectair Charters
,CAC,CONQUEST AIR,Conquest_Airlines,Conquest Airlines
,CXO,CONROE AIR,Conroe_Aviation_Services,Conroe Aviation Services
,VCH,CONSORCIO HELITEC,Consorcio_Helitec,Consorcio Helitec
,UZA,CONSTANTA,Constanta_Airline,Constanta Airline
,KIS,CONTACTAIR,Contact_Air,Contactair
,XCL,,Contel_ASC,Contel ASC
CO,COA,CONTINENTAL,Continental_Airlines,Continental Airlines
PC,PVV,CONTAIR,Continental_Airways,Continental Airways
CO,,JETLINK,Continental_Express,Continental Express
CS,CMI,AIR MIKE,Continental_Micronesia,Continental Micronesia
,CON,CONOCO,ConocoPhillips,Continental Oil
V0,VCV,CONVIASA,Conviasa,Conviasa
,CKA,COOK-AIR,Cook_Inlet_Aviation,Cook Inlet Aviation
,SVY,SURVEYOR,Cooper_Aerial_Surveys,Cooper Aerial Surveys
CM,CMP,COPA,Copa_Airlines,Copa Airlines
,CAT,AIRCAT,Copenhagen_Air_Taxi,Copenhagen Air Taxi
,COP,COPPER STATE,Copper_State_Air_Service,Copper State Air Service
,AAQ,COPTERLINE,Copterline,Copterline
CQ,CCW,CENTRAL CHARTER,Central_Charter,Central Charter
XC,CAI,CORENDON,Corendon_Airlines,Corendon Airlines
CD,CND,DUTCH CORENDON,Corendon_Dutch_Airlines,Corendon Dutch Airlines
,CRA,CORAL,Coronado_AerolÃ­neas,Coronado Aerolíneas
,CPB,PENTA,Corpac_Canada,Corpac Canada
,CNC,CENCOR,CorporaciÃ³n_AÃ©reo_Cencor,Corporación Aéreo Cencor
,CPG,CORPORANG,Corporacion_Aeroangeles,Corporacion Aeroangeles
,CGY,,Corporacion_Paraguaya_De_Aeronautica,Corporacion Paraguaya De Aeronautica
,CPT,AIR SPUR,Corporate_Air_(Billings),Corporate Air
,CPR,CORPAIR,Corporate_Air_(Hartford),Corporate Air
,CPO,MOKAN,Corporate_Aircraft_Company,Corporate Aircraft Company
,COO,CORPORATE,Corporate_Airlink,Corporate Airlink
,CKE,CHECKMATE,Corporate_Aviation_Services,Corporate Aviation Services
,VHT,VEGAS HEAT,Corporate_Flight_International,Corporate Flight International
,VTE,VOLUNTEER,Corporate_Flight_Management,Corporate Flight Management
,CJI,SEA JET,Corporate_Jets,Corporate Jets
SS,CRL,CORSAIR,Corsairfly,Corsairfly
XK,CCM,CORSICA,Corse_MÃ©diterranÃ©e,Corse Méditerranée
F5,COZ,COSMIC AIR,Cosmic_Air,Cosmic Air
,COT,,Costa_Airlines,Costa Airlines
,CHI,COUGAR,Cougar_Helicopters,Cougar Helicopters
,MGB,MOCKINGBIRD,Coulson_Flying_Service,Coulson Flying Service
,NSW,,Country_Connection_Airlines,Country Connection Airlines
,CIK,COUNTRY AIR,Country_International_Airlines,Country International Airlines
,CSD,DELIVERY,Courier_Services,Courier Services
,CUT,COURT AIR,Court_Helicopters,Court Helicopters
,OU,COURTLINE,Court_Line,Court Line
,CVL,COVAL,Coval_Air,Coval Air
,COW,COWI,COWI_A/S,COWI
7C,COY,COYNE AIR,Coyne_Aviation,Coyne Aviation
,CFD,AERONAUT,Cranfield_University,Cranfield University
,CRE,CREE AIR,Cree_Airways,Cree Airways
,ELM,CRELAM,Crelam,Crelam
,CAN,CREST,Crest_Aviation,Crest Aviation
,KRM,TRANS UNIVERSAL,Crimea_Universal_Avia,Crimea Universal Avia
OU,CTN,CROATIA,Croatia_Airlines,Croatia Airlines
,HRZ,CROATIAN AIRFORCE,Croatian_Air_Force_and_Defense,Croatian Air Force
,CRX,CROSSAIR,Cross_Aviation,Cross Aviation
QE,ECC,Cigogne,Crossair_Europe,Crossair Europe
,CWX,CROW EXPRESS,Crow_Executive_Air,Crow Executive Air
,CKR,CROWN AIR,Crown_Air_Systems,Crown Air Systems
,CRO,CROWN AIRWAYS,Crown_Airways,Crown Airways
,CRW,REGAL,Crownair,Crownair
,VCR,VOE CRUISER,Cruiser_Linhas_A%C3%A9reas,Cruiser Linhas Aéreas
,CTY,CENTURY,Cryderman_Air_Service,Cryderman Air Service
,CYT,CRYSTAL-AIR,Crystal_Shamrock,Crystal Shamrock
,IRO,IRON AIR,CSA_Air,CSA Air
CU,CUB,CUBANA,Cubana_de_Aviaci%C3%B3n,Cubana de Aviación
,CTF,CUTTER FLIGHT,Cutter_Aviation,Cutter Aviation
,CBL,CUMBERLAND,Cumberland_Airways,Cumberland Airways (Nicholson Air Service)
,CTT,CATT,Custom_Air_Transport,Custom Air Transport
,RGN,CYGNUS AIR,Gestair_Cargo,Cygnus Air
,CYC,CYPRAIR,Cyprair_Tours,Cyprair Tours
,CYS,SKYBIRD,Cypress_Airlines,Cypress Airlines
CY,CYP,CYPRUS,Cyprus_Airways,Cyprus Airways
YK,KYV,AIRKIBRIS,Cyprus_Turkish_Airlines,Cyprus Turkish Airlines
,CEF,CZECH AIR FORCE,Czech_Air_Force,Czech Air Force
,AHD,AIRHANDLING,Czech_Air_Handling,Czech Air Handling
OK,CSA,CSA,Czech_Airlines,Czech Airlines
,CIE,CZECH REPUBLIC,Czech_Government_Flying_Service,Czech Government Flying Service
8L,CGP,,Cargo_Plus_Aviation,Cargo Plus Aviation
GJ,CDC,HUALONG,CDI_Cargo_Airlines,CDI Cargo Airlines
,HNL,MAPLELEAF,CHC_Helicopters_Netherlands,CHC Helicopters Netherlands
5Z,KEM,CEMAIR,,CemAir
,JLH,CESA,Centro_de_Servicio_Aeronautico,Centro de Servicio Aeronautico
,FCB,NEW AGE,Cobalt_Air,Cobalt
,CVK,CARGO LINE,CAVOK_Air,CAVOK Airlines
XG,CLI,CLICKJET,Clickair,Clickair
,ABA,Aerobeta,,Aero-Beta
,DJT,DREAMJET,Dreamjet,Dreamjet
,DPJ,JET CARD,Delta_Private_Jets,Delta Private Jets
,DJR,DESERT FLIGHT,Desert_Jet,Desert Jet
,DLC,SOARCOPTER,Dehong_South_Asian_General_Aviation,Dehong South Asian General Aviation
,DNS,,Dniproaviaservis_Company,Dniproaviaservis Company
,DRF,,Dream_Flyers_Training_Center,Dream Flyers Training Center
,DMF,DEMLY,DMCFLY,DMCFLY
,NAU,DANAUS,Danaus_Lineas_Aereas,Danaus Lineas Aereas
,GDF,,DF_Aviation,DF Aviation
,DDA,DUSTY,D_&_D_Aviation,D & D Aviation
,DNK,DIRECT JET,D&K_Aviation,D&K Aviation
V5,VPA,VIP TAXI,DanubeWings,DanubeWings
,DHE,HELIDAP,DAP_Helicopteros,DAP Helicopteros
,VLF,VOLANTE,DFS_UK_Limited,DFS UK Limited
WD,DSR,DAIRAIR,DAS_Air_Cargo,DAS Air Cargo
,RKC,DAS CONGO,DAS_Airlines,DAS Airlines
DX,DTR,DANISH,DAT_Danish_Air_Transport,DAT Danish Air Transport
,ENT,DATENT,DAT_Enterprise_Limited,DAT Enterprise Limited
,BDN,GAUNTLET,DERA_Boscombe_Down,DERA Boscombe Down
,DSN,DESNA,DESNA,DESNA
,DET,SAMAL,DETA_Air,DETA Air
,DGO,DGO JET,DGO_Jet,DGO Jet
,DAE,YELLOW,DHL_Aero_Expreso,DHL Aero Expreso
D0,DHK,WORLD EXPRESS,DHL_Air_Limited,DHL Air Limited
,DHV,WORLDSTAR,DHL_Aviation,DHL Aviation
ES,DHX,DILMUN,DHL_International,DHL International
L3,JOS,,DHL_de_Guatemala,DHL de Guatemala
,RSK,REDSKIN,DSWA,DSWA
D3,DAO,DALO AIRLINES,Daallo_Airlines,Daallo Airlines
N2,DAG,DAGAL,Dagestan_Airlines,Dagestan Airlines
,DHA,,Dahla_Airlines,Dahla Airlines
,DCS,TWIN STAR,DC_Aviation,DC Aviation
,DCX,DAIMLER,Daimler-Chrysler,Daimler-Chrysler
,DKA,,Daka_(airline),Daka
,DLR,DALA AIR,Dala_Air_Services,Dala Air Services
H8,KHB,DALAVIA,Dalavia,Dalavia
,DXP,DALLAS EXPRESS,Dallas_Express_Airlines,Dallas Express Airlines
,DAS,AIRDAM,Damascene_Airways,Damascene Airways
,DSA,DANBURY AIRWAYS,Danbury_Airways,Danbury Airways
,DOP,DANCOPTER,Dancopter,Dancopter
,DAF,DANISH AIRFORCE,Danish_Air_Force,Danish Air Force
DD,DDL,,Danish_Air_Lines,Danish Air Lines
,DAR,DANISH ARMY,Danish_Army,Danish Army
,DNY,DANISH NAVY,Danish_Navy,Danish Navy
,DNU,DANU,Danu_Oro_Transportas,Danu Oro Transportas
,DRT,DARTA,Darta,Darta
0D,DWT,DARWIN,Darwin_Airline,Darwin Airline
,DSQ,DASAB AIR,Dasab_Airlines,Dasab Airlines
,DSH,DASH CHARTER,Dash_Air_Charter,Dash Air Charter
,GOB,PILGRIM,Dash_Aviation,Dash Aviation
,DGX,DASNA,Dasnair,Dasnair
,DAB,,Dassault_Aviation,Dassault Aviation
,CVF,CLOVERLEAF,Dassault_Falcon_Jet_Corporation,Dassault Falcon Jet Corporation
,DSO,DASSAULT,Dassault_Falcon_Service,Dassault Falcon Service
,DTN,DATA AIR,Data_International,Data International
,XDT,,Date_Transformation_Corp,Date Transformation Corp
D5,DAU,DAUAIR,Dauair,Dauair
,DCO,,David_Crawshaw_Consultants_Limited,David Crawshaw Consultants Limited
,DWN,DAWN AIR,Dawn_Air,Dawn Air
,DJS,DAYJET,DayJet,DayJet
,DAY,DAYA,Daya_Aviation,Daya Aviation
,DHC,DEHAVILLAND,De_Havilland,De Havilland
,IAY,IASON,Deadalos_Flugtbetriebs,Deadalos Flugtbetriebs
,DAA,DECUR,Decatur_Aviation,Decatur Aviation
,DKN,DECCAN,Deccan_Charters,Deccan Charters
,JDC,JOHN DEERE,Deere_and_Company,Deere and Company
,DWR,DELAWARE,Delaware_Skyways,Delaware Skyways
,DEA,JET SERVICE,Delta_Aerotaxi,Delta Aerotaxi
,SNO,SNOWBALL,Delta_Air_Charter,Delta Air Charter
,ELJ,ELITE JET,Delta_Private_Jets,Delta Private Jets
DL,DAL,DELTA,Delta_Air_Lines,Delta Air Lines
,KMB,KEMBLEJET,Delta_Engineering_Aviation,Delta Engineering Aviation
,DLI,DELTA EXPRESS,Delta_Express_International,Delta Express International
,DSU,DELTA STATE,Delta_State_University,Delta State University
J7,DNM,DENIM,Denim_Air,Denim Air
,FEC,FALCON EXPRESS,Denver_Express,Denver Express
,DJT,DENVER JET,Denver_Jet,Denver Jet
,FGC,FORESTALS,Generalitat_de_Catalunya,Departament d'Agricultura de la Generalitat de Catalunya
,DRY,DERAYA,Deraya_Air_Taxi,Deraya Air Taxi
,DRX,,Des_R_Cargo_Express,Des R Cargo Express
,MIZ,MILAZ,Desarrollo_Milaz,Desarrollo Milaz
,DTY,DESTINY,Destiny_Air_Services,Destiny Air Services
2A,,,Deutsche_Bahn,Deutsche Bahn
1I,AMB,CIVIL AIR AMBULANCE,Deutsche_Rettungsflugwacht,Deutsche Rettungsflugwacht
,LFO,LUFO,German_Aerospace_Center,Deutsches Zentrum fur Luft-und Raumfahrt EV
,DIS,DI AIR,Di_Air,Di Air
,SPK,SPARKLE,Diamond_Aviation,Diamond Aviation
,DRB,DIDIER,Didier_Rousset_Buy,Didier Rousset Buy
,DGT,DIGITAL,Digital_Equipment_Corporation,Digital Equipment Corporation
,DIP,DIPFREIGHT,Diplomatic_Freight_Services,Diplomatic Freight Services
,ENA,ENA,DirecciÃ³n_General_de_AviaciÃ³n_Civil_y_Telecomunicasciones,Dirección General de Aviación Civil y Telecomunicasciones
,DIA,BLUE SKY,Direct_Air,Direct Air
,XAP,MID-TOWN,Midway_Connection,Direct Air trading as Midway Connection
,DCT,,Direct_Flight,Direct Flight
,SXP,EXPRESS SKY,Direct_Fly,Direct Fly
AW,DIR,DIRGANTARA,Dirgantara_Air_Service,Dirgantara Air Service
,DCV,DISCOVER,Discover_Air,Discover Air
DH,DVA,DISCOVERY AIRWAYS,Discovery_Airways,Discovery Airways
,XDS,,Dispatch_Services,Dispatch Services
,DIX,DIX FLIGHT,Dix_Aviation,Dix Aviation
,DEE,TACAIR,Dixie_Airways,Dixie Airways
Z6,UDN,DNIEPRO,Dniproavia,Dniproavia
,FDN,FLYING DOLPHIN,Dolphin_Air,Dolphin Air
,IXX,ISLAND EXPRESS,Dolphin_Express_Airlines,Dolphin Express Airlines
,DPL,DOME,Dome_Petroleum,Dome Petroleum
YU,ADM,DOMINAIR,Dominair,Dominair
,MYO,MAYORAL,Dominguez_Toledo,Dominguez Toledo (Grupo Mayoral)
DO,DOA,DOMINICANA,Dominicana_de_Aviaci%C3%B3n,Dominicana de Aviación
E3,DMO,DOMODEDOVO,Domodedovo_Airlines,Domodedovo Airlines
,DVB,DONSEBAI,Don_Avia,Don Avia
,DON,DONAIR,Donair_Flying_Club,Donair Flying Club
D9,DNV,DONAVIA,Donavia,Donavia
5D,UDC,DONBASS AERO,DonbassAero,DonbassAero
,DAD,DORADO AIR,Dorado_Air,Dorado Air
,DOR,DORNIER,Dornier_Flugzeugwerke,Dornier
,DAV,DANA AIR,Dornier_Aviation_Nigeria,Dornier Aviation Nigeria
,DOM,DOS MUNDOS,Dos_Mundos_(airline),Dos Mundos
,DCA,DREAM CATCHER,Dreamcatcher_Airways,Dreamcatcher Airways
KB,DRK,ROYAL BHUTAN,Druk_Air,Druk Air
,DRE,MICHIGAN,Drummond_Island_Air,Drummond Island Air
,DUB,DUBAI,Dubai_Airwing,Dubai Airwing
,DBK,SEAGULL,Dubrovnik_Air,Dubrovnik Air
,DUK,LION KING,Ducair,Ducair
,DBJ,DUCHESS,Duchess_of_Britany_(Jersey)_Limited,Duchess of Britany (Jersey) Limited
,LPD,LEOPARD,Duke_of_York,UK Royal/HRH Duke of York
,DUN,DUNAIR,Dun'Air,Dun'Air
,PHD,PANHANDLE,Duncan_Aviation,Duncan Aviation
,VVF,WORLDFOCUS,Dunyaya_Bakis_Hava_Tasimaciligi,Dunyaya Bakis Hava Tasimaciligi
,DUO,FLY DUO,Duo_Airways,Duo Airways
,DJE,DURANGO JET,Durango_Jet,Durango Jet
,DNL,DUTCH ANTILLES,Dutch_Antilles_Express,Dutch Antilles Express
,DCE,DUTCH CARIBBEAN,Dutch_Caribbean_Express,Dutch Caribbean Express
,DBR,DUTCHBIRD,Dutchbird,Dutchbird
,DBR,DOBROLET,Dobrolet_(low-cost_airline),Dobrolet LLC
,DFS,DWYAIR,Dwyer_Aircraft_Services,Dwyer Aircraft Services
,XDY,,Dynair_Services,Dynair Services
,DNR,DYNAMAIR,Dynamair_Aviation,Dynamair Aviation
,DYE,DYNAMIC,Dynamic_Airlines,Dynamic Air
,DYA,DYNAMIC AIR,Dynamic_Airways,Dynamic Airways
DI,BAG,SPEEDWAY,DBA_(airline),Dba
,EES,,Eagle_Express,Eagle Express
,EBF,,Echo_Airlines,Echo Airlines
MQ,ENY,ENVOY,Envoy_Air,Envoy Air
,EVK,EVERETT,Everett_Aviation,Everett Aviation
,ENK,SUNBIRD,Executive_Airlink,Executive Airlink
EL,ELB,ELLINAIR HELLAS,Ellinair,Ellinair
,EVT,,Everett_Limited,Everett Limited
,ELN,ELERON,Eleron_Aviation_Company,Eleron Aviation Company
,ECC,ECLAIR,Eclair_Aviation,Eclair Aviation
,ELU,EGYPTIAN LEISURE,Egyptian_Leisure_Airlines,Egyptian Leisure Airlines
,LTD,LIGHT SPEED,Executive_Express_Aviation,Executive Express Aviation/JA Air Charter
,EPR,EMPEROR,Express_Airways,Express Airways
,XSL,EXCELAIRE,Excel-Aire_Service,Excel-Aire Service
,XSR,AIRSHARE,Executive_Flight_Services,Executive Flight Services
,EZJ,GUYANA JET,Ezjet_GT,Ezjet GT
9A,EZX,EAGLEXPRESS,Eagle_Express_Air_Charter,Eagle Express Air Charter
,MNU,MAINER,Elite_Airways,Elite Airways
,,EUROPE JET,Europe_Jet,Europe Jet
E1,,,Everbread,Everbread
,EHD,PLATINUM AIR,E_H_Darby_Aviation,E H Darby Aviation
1C,,,Electronic_Data_Systems,Electronic Data Systems
1Y,,,Electronic_Data_Systems,Electronic Data Systems
,EXW,ECHOLINE,EAS_Airlines,Executive Airlines Services
,EFS,EFAOS,EFAOS,EFAOS- Agencia De Viagens e Turismo
,EFD,EVER FLIGHT,Eisele_Flugdienst,Eisele Flugdienst
,FSD,FLUGSERVICE,EFS-Flugservice,EFS-Flugservice
,EIS,COOL,EIS_Aircraft,EIS Aircraft
,IAG,EPAG,EPAG,EPAG
,ESI,ELISERVIZI,ESI_Eliservizi_Italiani,ESI Eliservizi Italiani
,EUY,EUROAIRWAYS,EU_Airways,EU Airways
VE,EUJ,UNION JET,EUjet,EUjet
BR,EVA,EVA,EVA_Air,EVA Air
,ICR,ICARUS FLIGHTS,Eagle_Aero,Eagle Aero
,FEI,ARCTIC EAGLE,Eagle_Air_(Iceland),Eagle Air
,,,Eagle_Air_(Indonesia),Eagle Air
,EGR,EAGLE SIERRA,Eagle_Air_(Sierra_Leone),Eagle Air
EY,EFL,FLYING EAGLE,Eagle_Air_(Tanzania),Eagle Air
H7,EGU,AFRICAN EAGLE,Eagle_Air_(Uganda),Eagle Air
NZ,EAG,EAGLE,Eagle_Airways,Eagle Airways
,EGX,THAI EAGLE,Eagle_Air_Company,Eagle Air Company
,GYP,GYPSY,Eagle_Aviation_(UK),Eagle Aviation
,EGN,FRENCH EAGLE,Eagle_Aviation_France,Eagle Aviation France
,SEG,SEN-EAGLE,Eagle_International,Eagle International
,EGJ,EAGLE JET,Eagle_Jet_Charter,Eagle Jet Charter
,EMD,EAGLEMED,Eaglemed,Eaglemed (Ballard Aviation)
,ERX,EARTH AIR,Earth_Airlines,Earth Airlines Services
QU,UTN,UT Ukraine,UTair-Ukraine,UTair-Ukraine
S9,HSA,DUMA,East_African_Safari_Air,East African Safari Air
,EXZ,TWIGA,East_African_Safari_Air_Express,East African Safari Air Express
,EMU,,East_Asia_Airlines,East Asia Airlines
,ECT,EASTWAY,East_Coast_Airways,East Coast Airways
,ECJ,EASTCOAST JET,East_Coast_Jets,East Coast Jets
,EHA,AIRE HAMPTON,East_Hampton_Aire,East Hampton Aire
,EKC,BLUE GOOSE,East_Kansas_City_Aviation,East Kansas City Aviation
,CTK,COSTOCK,East_Midlands_Helicopters,East Midlands Helicopters
,DXH,EAST STAR,East_Star_Airlines,East Star Airlines
,EWA,EASTWEST,East-West_Airlines_(Australia),East-West Airlines
ZE,ESR,EASTAR,Eastar_Jet,Eastar Jet
,EAZ,EASAIR,Eastern_Air,Eastern Air
,EAX,EASTEX,Eastern_Air_Executive,Eastern Air Executive
EA,EAL,EASTERN,Eastern_Air_Lines,Eastern Air Lines
,EAL,EASTERN,Eastern_Air_Lines_(2015),Eastern Air Lines
T3,EZE,EASTFLIGHT,Eastern_Airways,Eastern Airways
QF,EAQ,EASTERN,Eastern_Australia_Airlines,Eastern Australia Airlines
,ECI,EASTERN CAROLINA,Eastern_Carolina_Aviation,Eastern Carolina Aviation
,GNS,GENESIS,Eastern_Executive_Air_Charter,Eastern Executive Air Charter
,LIS,LARISA,Eastern_Express,Eastern Express
,EME,EMAIR,Eastern_Metro_Express,Eastern Metro Express
,EPB,EAST PAC,Eastern_Pacific_Aviation,Eastern Pacific Aviation
,ESJ,EASTERN SKYJETS,Eastern_SkyJets,Eastern SkyJets
DK,ELA,,Eastland_Air,Eastland Air
W9,SGR,STINGER,Eastwind_Airlines,Eastwind Airlines
,FYE,FLYME,Easy_Link_Aviation,Easy Link Aviation Services
,CMN,CIMMARON AIRE,Eckles_Aircraft,Eckles Aircraft
,EJT,ECLIPSE JET,Eclipse_Aviation,Eclipse Aviation
,ECQ,SKYBRIDGE,Eco_Air,Eco Air
,DEI,,Ecoair_International,Ecoair
,ECX,AIR ECOMEX,Ecomex_Air_Cargo,Ecomex Air Cargo
,ECD,ECOTOUR,Ecotour_(airline),Ecotour
,XCC,XCALAK,Ecoturistica_de_Xcalak,Ecoturistica de Xcalak
,ECV,EQUATOGUINEA,Ecuato_Guineana,Ecuatoguineana De Aviación (EGA)
,EQC,ECUA-CARGO,Ecuatorial_Cargo,Ecuatorial Cargo
,ECU,ECUAVIA,Ecuavia,Ecuavia
WK,EDW,EDELWEISS,Edelweiss_Air,Edelweiss Air
,SLO,SLOW,Edgartown_Air,Edgartown Air
,EDC,SALTIRE,Air_Charter_Scotland,Air Charter Scotland
,EDJ,EDWARDS,Edwards_Jet_Center_of_Montana,Edwards Jet Center of Montana
,EIJ,EFATA,Efata_Papua_Airlines,Efata Papua Airlines
,EUW,EUROWEST,EFS_European_Flight_Service,EFS European Flight Service
MS,MSR,EGYPTAIR,Egyptair,Egyptair
,MSX,EGYPTAIR CARGO,Egyptair_Cargo,Egyptair Cargo
,EGY,,Egyptian_Air_Force,Egyptian Air Force
,EJX,,Egyptian_Aviation,Egyptian Aviation
,EMA,,Egyptian_Aviation_Company,Egyptian Aviation Company
,EIX,AIR EXPORTS,Ei_Air_Exports,Ei Air Exports
,EIR,EIRJET,Eirjet,Eirjet
LY,ELY,ELAL,El_Al_Israel_Airlines,El Al Israel Airlines
,CMX,EL CAMINANTE,El_Caminante_Taxi_AÃ©reo,El Caminante Taxi Aéreo
,GLQ,QUILADA,El_Quilada_International_Aviation,El Quilada International Aviation
,ELS,EL SAL,El_Sal_Air,El Sal Air
,ESC,SOLAMERICA,El_Sol_De_AmÃ©rica,El Sol De América
UZ,BRQ,BURAQAIR,El-Buraq_Air_Transport,El-Buraq Air Transport
,ELX,ELAN,Elan_Express,Elan Express
,LBR,MOTION,Elbe_Air_Transport,Elbe Air Transport
,NLK,ELAVIA,Elbrus-Avia,Elbrus-Avia Air Enterprise
,DND,DINDER,Eldinder_Aviation,Eldinder Aviation
,PDV,ELICAR,Elicar,Elicar
,EDO,ELIDOLOMITI,Elidolomiti,Elidolomiti
,ELB,ELILOBARDIA,Elieuro,Elieuro
,EFG,ELIFRIULIA,Elifriulia,Elifriulia
,ELH,LARIO,Elilario_Italia,Elilario Italia
,EOA,LOMBARDA,Elilombarda,Elilombarda
,MEE,ELIMEDITERRANEA,Elimediterranea,Elimediterranea
,VUL,ELIOS,Elios_(airline),Elios
,IEP,ELIPIU,Elipiu',Elipiu'
,RSA,ESRA,Elisra_Airlines,Elisra Airlines
,EAI,ELAIR,Elite_Air,Elite Air
,EJD,ELITE DUBAI,Elite_Jets,Elite Jets
,FGS,ELITELLINA,Elitellina,Elitellina
,ELT,ELLIOT,Elliott_Aviation,Elliott Aviation
,MGG,ELMAGAL,Elmagal_Aviation_Services,Elmagal Aviation Services
,ELR,,Elrom_Aviation_and_Investments,Elrom Aviation and Investments
,EAM,EMBASSY AIR,Embassy_Airlines,Embassy Airlines
,EFT,EMBASSY FREIGHT,Embassy_Freight_Company,Embassy Freight Company
,EMB,EMBRAER,Embraer,Empresa Brasileira De Aeronáutica
,XSL,SATSLAB,Embry-Riddle_Aeronautical_University,Embry-Riddle Aeronautical University
,JEM,GEMSTONE,Emerald_Airways,Emerald Airways
,EWW,EMERY,Emery_Worldwide_Airlines,Emery Worldwide Airlines
,EMT,EMETEBE,Emetebe,Emetebe
EK,UAE,EMIRATES,Emirates_Airlines,Emirates Airlines
,SBC,SABIAN AIR,Emoyeni_Air_Charter,Emoyeni Air Charter
,EMP,EMPIRE,Empire_Air_Service,Empire Air Service
EM,CFS,EMPIRE AIR,Empire_Airlines,Empire Airlines
,MPR,,Empire_Aviation_Services,Empire Aviation Services
,ETP,TESTER,Empire_Test_Pilots%27_School,Empire Test Pilots' School
,AUO,UNIFORM OSCAR,"Empresa_(Aero_Uruguay),_S.A.","Empresa (Aero Uruguay), S.A."
,PRG,ASPAR,Empresa_Aero-Servicios_Parrague,Empresa Aero-Servicios Parrague
,CRN,AEROCARIBBEAN,Empresa_Aerocaribbean,Empresa Aerocaribbean
,VNA,EBBA,Empresa_AviaciÃ³n_Interamericana,Empresa Aviación Interamericana
EU,EEA,ECUATORIANA,Empresa_Ecuatoriana_De_Aviaci%C3%B3n,Empresa Ecuatoriana De Aviación
,CNI,SERAER,Empresa_Nacional_De_Servicios_AÃ©reos,Empresa Nacional De Servicios Aéreos
,VNE,VENEZOLANA,Empresa_Venezolana,Empresa Venezolana
,GTV,GAVIOTA,Empresa_de_AviaciÃ³n_Aerogaviota,Empresa de Aviación Aerogaviota
,XLT,INFRAERO,Empressa_Brasileira_de_Infra-Estrutura_Aeroportuaria-Infraero,Empressa Brasileira de Infra-Estrutura Aeroportuaria-Infraero
,ENC,ENDECOTS,Endecots,Endecots
,ENI,ENIMEX,Enimex,Enimex
G8,ENK,ENKOR,Enkor,Enkor JSC
,EGV,GLEISNER,Enrique_Gleisner_Vivanco,Enrique Gleisner Vivanco
,ESE,ENSENADA ESPECIAL,Ensenada_Vuelos_Especiales,Ensenada Vuelos Especiales
OF,ENT,ENTER,Enter_Air,Enter Air
,ENS,ENTERGY SHUTTLE,Entergy_Services,Entergy Services
,EWS,WORLD ENTERPRISE,Enterprise_World_Airways,Enterprise World Airways
E0,ESS,NEW DAWN,Eos_Airlines,Eos Airlines
,EKA,EQUAFLIGHT,Equaflight_Service,Equaflight Service
,EQZ,ZAMBIA CARGO,Equatair_Air_Services_(Zambia),Equatair Air Services (Zambia)
,EQT,,Equatorial_Airlines,Equatorial Airlines
,ERH,ERAH,Era_Helicopters,Era Helicopters
,IRY,ERAM AIR,Eram_Air,Eram Air
,ERF,ERFOTO,Erfoto,Erfoto
,ERE,AIR ERIE,Erie_Airways,Erie Airways
B8,ERT,ERITREAN,Eritrean_Airlines,Eritrean Airlines
,EAD,AERO-ESCOLA,Escola_De_Aviacao_Aerocondor,Escola De Aviacao Aerocondor
,CTV,ARE AVIACION,Escuela_De_Pilotos_Are_AviaciÃ³n,Escuela De Pilotos Are Aviación
,EPC,ESPACE,Espace_Aviation_Services,Espace Aviation Services
,ERC,ESSO,Esso_Resources_Canada,Esso Resources Canada
E7,ESF,,Estafeta_Carga_A%C3%A9rea,Estafeta Carga Aérea
,EEF,ESTONIAN AIR FORCE,Estonian_Air_Force,Estonian Air Force
OV,ELL,ESTONIAN,Estonian_Air,Estonian Air
,ETA,ESTRELLAS,Estrellas_Del_Aire,Estrellas Del Aire
ET,ETH,ETHIOPIAN,Ethiopian_Airlines,Ethiopian Airlines
,MJM,ELCO ETI,,Eti 2000
EY,ETD,ETIHAD,Etihad_Airways,Etihad Airways
,ETM,ETRAM,Etram_Air_Wing,Etram Air Wing
,EVN,EURAVIATION,Euraviation,Euraviation
,ECN,EURO CONTINENTAL,Euro_Continental_AIE,Euro Continental AIE
RZ,,,Euro_Exec_Express,Euro Exec Express
,ESN,EURO SUN,Euro_Sun,Euro Sun
,EAK,EAKAZ,Euro-Asia_Air,Euro-Asia Air
,KZE,KAZEUR,Euro-Asia_Air_International,Euro-Asia Air International
MM,MMZ,EUROATLANTIC,EuroAtlantic_Airways,EuroAtlantic Airways
,GOJ,GOJET,EuroJet_Aviation,EuroJet Aviation
,EUP,EUROSTAR,Euroair,Euroair
,EUU,EUROAMERICAN,Euroamerican_Air,Euroamerican Air
,ECY,ECHELON,Euroceltic_Airways,Euroceltic Airways
,EUC,,Eurocontrol,Eurocontrol
,ECF,EUROCOPTER,Eurocopter,Eurocopter
UI,ECA,EUROCYPRIA,Eurocypria_Airlines,Eurocypria Airlines
GJ,EEZ,E-FLY,Eurofly,Eurofly
,EEU,EUROFLY,Eurofly_Service,Eurofly Service
,EUG,EUROGUINEA,Euroguineana_de_AviaciÃ³n,Euroguineana de Aviación
,ERJ,JET ITALIA,,Eurojet Italia
,JLN,JET LINE,Eurojet_Limited,Eurojet Limited
,RDP,JET-ARROW,Eurojet_Romania,Eurojet Romania
,EJS,EEJAY SERVICE,Eurojet_Servis,Eurojet Servis
K2,ELO,EUROLOT,Eurolot,Eurolot
3W,EMX,EUROMANX,Euromanx_Airways,Euromanx Airways
,GED,LANGUEDOC,Europe_Air_Lines,Europe Air Lines
5O,FPO,FRENCH POST,Europe_Airpost,Europe Airpost
,EUT,FIESTA,European_2000_Airlines,European 2000 Airlines
,EAG,,European_Aeronautical_Group_UK,European Aeronautical Group UK
EA,EAL,STAR WING,European_Air_Express,European Air Express
QY,BCS,EUROTRANS,European_Air_Transport,European Air Transport
E7,EAF,EUROCHARTER,European_Aviation_Air_Charter,European Aviation Air Charter
,EBJ,,European_Business_Jets,European Business Jets
,ECB,COASTAL CLIPPER,European_Coastal_Airlines,European Coastal Airlines
,ETV,EURO EXEC,European_Executive,European Executive
,EXC,ECHO EXPRESS,European_Executive_Express,European Executive Express
,EBG,EUROSENSE,Eurosense,Eurosense
,ESX,CATFISH,Euroskylink,Euroskylink
EW,EWG,EUROWINGS,Eurowings,Eurowings
,EVE,EVELOP,Evelop_Airlines,Evelop Airlines
EZ,EIA,EVERGREEN,Evergreen_International_Airlines,Evergreen International Airlines
,VTS,EVERTS,Everts_Air_Alaska,Everts Air Alaska/Everts Air Cargo
,EMN,AGENCY,Examiner_Training_Agency,Examiner Training Agency
JN,XLA,EXPO,Excel_Airways,Excel Airways
,XEL,HELI EXCEL,Excel_Charter,Excel Charter
,GZA,EXCELLENT AIR,Excellent_Air,Excellent Air
MB,EXA,CANADIAN EXECAIRE,Execair_Aviation,Execair Aviation
,VCN,AVCON,Execujet_Charter,Execujet Charter
,EJO,MIDJET,Execujet_Middle_East,Execujet Middle East
,VMP,VAMPIRE,Execujet_Scandinavia,Execujet Scandinavia
,LFL,LIFE FLIGHT,Executive_Air,Executive Air
,EAC,EXECAIR,Executive_Air_Charter,Executive Air Charter
,XAF,,Executive_Air_Fleet,Executive Air Fleet
,ECS,ECHO,Executive_Aircraft_Charter_and_Charter_Services,Executive Aircraft Charter and Charter Services
,XAH,,Executive_Aircraft_Services,Executive Aircraft Services
OW,EXK,EXECUTIVE EAGLE,Executive_Airlines,Executive Airlines
,EXU,SACAIR,,Executive Airlines
,JTR,JESTER,Executive_Aviation_Services,Executive Aviation Services
,EXE,EXEC,Executive_Flight,Executive Flight
,TRI,TRILLIUM,Ontario_Government,Executive Flight Operations Ontario Government
,EXJ,,Executive_Jet_Charter,Executive Jet Charter
,EJM,JET SPEED,Executive_Jet_Management,Executive Jet Management
,TEA,TRAVELMAX,Executive_Turbine_Aviation,Executive Turbine Aviation
,EXF,EXIMFLIGHT,Eximflight,Eximflight
,EXN,EXIN,Exin,Exin
,EXR,EXPERTOS ENCARGA,Expertos_En_Carga,Expertos En Carga
,FXA,EFFEX,Express_Air,Express Air
,EIC,EXCARGO,Express_International_Cargo,Express International Cargo
,XPL,EXPRESSLINE,Express_Line_Aircompany,Express Line Aircompany
,XNA,EXPRESSNET,Express_Net_Airlines,Express Net Airlines
EO,LHN,LONGHORN,Express_One_International,Express One International
,XTO,EXPRESS TOURS,Express_Tours,Express Tours
EV,ASQ,ACEY,ExpressJet_Airlines,ExpressJet
,JTM,SKYMAN,Exxavia_Limited,Exxavia Limited
U2,EZY,EASY,EasyJet,easyJet
DS,EZS,TOPSWISS,EasyJet_Switzerland,easyJet Switzerland
,AAE,ARIZONA,Western_Express_Air,Western Express Air
,EAV,MAYFLOWER,Eagle_Airlines_Luftverkehrsges,Eagle Airlines Luftverkehrsges
,ABU,,Eagle_Aviation_Services,Eagle Aviation Services
,XRO,CRAMER,ExxAero,ExxAero
EC,EJU,ALPINE,EasyJet_Europe,easyJet Europe
,EVL,EVOLEM,Evolem_Aviation,Evolem Aviation
,KWX,KAY DUB,Florida_Aerocharter,Florida Aerocharter
,FAS,FLORIDA CARGO,Florida_Air_Cargo,Florida Air Cargo
,VNX,VANCE,Fly_Advance,Fly Advance
,FUM,FUNLINE,Fuxion_Line_Mexico,Fuxion Line Mexico
,FRB,RAKWAY,Fly_Rak,Fly Rak
,FMI,FIRST MYANMAR,FMI_Air,FMI Air
,FRX,FORT AERO,Fort_Aero,Fort Aero
,PBR,POLAR BEAR,Fast_Air,Fast Air
,SRE,STREAMJET,Fly_Jetstream_Aviation,Fly Jetstream Aviation
,FTZ,GREY BIRD,Fastjet,Fastjet
,FAP,FAIR SCHOOL,F_Air,F Air
,FLS,RAPIDOS,Flying_Fast,Flying Fast
,EYE,SOCKEYE,F.S._Air_Service,F.S. Air Service
,IFA,RED ANGEL,FAI_Air_Service,FAI Air Service
,FLC,FLIGHT CHECK,FINFO_Flight_Inspection_Aircraft,FINFO Flight Inspection Aircraft
,FKI,KIEL AIR,FLM_Aviation_Mohrdieck,FLM Aviation Mohrdieck
,DCM,DOT COM,Fltplan.com,FLTPLAN (anonymized service)
,FLW,QUICKFLOW,FLowair_Aviation,FLowair Aviation
,FMG,HUSKY,FMG_Verkehrsfliegerschule_Flughafen_Paderborn-Lippstadt,FMG Verkehrsfliegerschule Flughafen Paderborn-Lippstadt
,FRA,RUSHTON,FR_Aviation,FR Aviation
,FSB,SEABIRD,FSB_Flugservice_&_Development,FSB Flugservice & Development
,LEJ,LEIPZIG FAIR,FSH_Luftfahrtunternehmen,FSH Luftfahrtunternehmen
,FBA,FAB AIR,Fab_Air,Fab Air
,FCS,MEXFACTS,Facts_Air,Facts Air
,FAV,FAIRAVIA,Fair_Aviation,Fair Aviation
,FWD,FAIR WIND,Fair_Wind_Air_Charter,Fair Wind Air Charter
,FLS,FAIRLINE,Fairlines,Fairlines
,FFC,FAIROAKS,Fairoaks_Flight_Centre,Fairoaks Flight Centre
,FWY,FAIRWAYS,Fairways_Corporation,Fairways Corporation
,FCN,FALCON,Falcon_Air,Falcon Air
,FAR,FALCAIR,Falcon_Air_(Slovenia),Falcon Air
,FAO,PANTHER,Falcon_Air_Express,Falcon Air Express
,FAU,FALCON AIRLINE,Falcon_Airline,Falcon Airline
IH,,,Falcon_Aviation,Falcon Aviation
,FVS,FALCON AVIATION,Falcon_Aviation_Services,Falcon Aviation Services
,FJC,FALCONJET,Falcon_Jet_Centre,Falcon Jet Centre
,FAW,FALWELL,Falwell_Aviation,Falwell Aviation
FE,FEA,Far Eastern,Far_Eastern_Air_Transport,Far Eastern Air Transport
FD,,FLYDOC,Royal_Flying_Doctor_Service,Royal Flying Doctor Service
,FDL,FARMINGDALE STATE,Farmingdale_State_University,Farmingdale State University
,FAH,BLUE STRIP,Farnair_Hungary,Farnair Hungary
,FRN,FARNED,Farnair_Netherlands,Farnair Netherlands
,FAT,FARNER,ASL_Airlines_Switzerland,ASL Airlines Switzerland
,RAF,FARNAS,Farnas_Aviation_Services,Farnas Aviation Services
,HBL,HELIBLUE,Faroecopter,Faroecopter
F6,RCK,ROCKROSE,FaroeJet,FaroeJet
,FRW,FARWEST,Farwest_Airlines,Farwest Airlines
F3,FSW,FASO,Faso_Airways,Faso Airways
,FHL,FINDON,Fast_Helicopters,Fast Helicopters
,FAY,FAYBAN AIR,Fayban_Air_Services,Fayban Air Services
,SKM,SKYTEM,Fayetteville_Flying_Service_and_Scheduled_Skyways_System,Fayetteville Flying Service and Scheduled Skyways System
,FDR,FEDAIR,Federal_Air,Federal Air
,FLL,FEDERAL AIRLINES,Federal_Airlines,Federal Airlines
,DCN,DIPLOMATIC CLEARANCE,Federal_Armed_Forces,Federal Armed Forces
,FRM,FEDARM,Federal_Armored_Service,Federal Armored Service
,NHK,NIGHTHAWK,Federal_Aviation_Administration,Federal Aviation Administration
FX,FDX,FEDEX,FedEx_Express,FedEx Express
,FNK,AURIKA,Feniks_Airline,Feniks Airline
,FER,FERIA,Feria_AviaciÃ³n,Feria Aviación
N8,HGK,SALAAMA,Fika_Salaama_Airlines,Fika Salaama Airlines
4S,FNC,FINALAIR CONGO,Finalair_Congo,Finalair Congo
,FAK,FACTS,Financial_Airxpress,Financial Airxpress
,FBF,FINE AIR,Fine_Airlines,Fine Airlines
,FTR,FINISTAIR,Finist%27air,Finist'air
AY,FIN,FINNAIR,Finnair,Finnair
FC,WBA,WESTBIRD,Finncomm_Airlines,Finncomm Airlines
,FNF,FINNFORCE,Finnish_Air_Force,Finnish Air Force
FY,FFM,FIREFLY,Firefly_(airline),Firefly
7F,FAB,FIRST AIR,First_Air,First Air
,JRF,,First_Air_Transport,First Air Transport
,FCC,FIRST CAMBODIA,First_Cambodia_Airlines,First Cambodia Airlines
DP,FCA,JETSET,First_Choice_Airways,First Choice Airways
,MBL,FIRST CITY,First_City_Air,First City Air
,GGA,JAWJA,First_Flying_Squadron,First Flying Squadron
,FIR,FIRSTLINE AIR,First_Line_Air,First Line Air
,FTS,FIRST SABRE,First_Sabre,First Sabre
8F,FFR,FISCHER,Fischer_Air,Fischer Air
,FFP,FLYING FISH,Fischer_Air_Polska,Fischer Air Polska
8D,EXV,EXPOAVIA,FitsAir,FitsAir
,FSX,FLAG,Flagship_Express_Services,Flagship Express Services
,FLE,FLAIR,Flair_Airlines,Flair Airlines
,WAF,FLAMENCO,Flamenco_Airways,Flamenco Airways
,FMR,FLAMINGO AIR,Flamingo_Air,Flamingo Air
,FLN,ILIAS,Flamingo_Air-Line,Flamingo Air-Line
,FSH,FLASH,Flash_Airlines,Flash Airlines
,BWY,BROADWAY,Fleet_Requirements_Air_Direction_Unit,Fleet Requirements Air Direction Unit
,FLR,FLEETAIR,Fleetair,Fleetair
,FXY,FLEXY,Flexair,Flexair
,FXT,,Flexflight,Flexflight
,TUD,TUNDRA,Flight_Alaska,Flight Alaska
,FCK,NAV CHECKER,FCS_Flight_Calibration_Services,FCS Flight Calibration Services
,VOR,FLIGHT CAL,Flight_Calibration_Services_Ltd.,Flight Calibration Services Ltd.
,FCV,NAVAIR,Flight_Centre_Victoria,Flight Centre Victoria
,FCP,FLIGHTCORP,Nelson_Aviation_College_Ltd,Nelson Aviation College Ltd
,FDP,,Flight_Dispatch_Services,Flight Dispatch Services
,FLX,FLIGHT EXPRESS,"Flight_Express,_Inc.","Flight Express, Inc."
,CFI,CHINA JET,Flight_Inspection_Center_of_the_General_Administration_of_Civil_Aviation_in_China,Flight Inspection Center of the General Administration of Civil Aviation in China
,LTS,SPECAIR,Flight_Inspections_and_Systems,Flight Inspections and Systems
,IVJ,INVADER JACK,Flight_International_(airline),Flight International
,MIT,MATCO,Flight_Line,Flight Line
,FOI,,Flight_Ops_International,Flight Ops International
,OPT,OPTIONS,Flight_Options,Flight Options
,CLB,CALIBRATOR,Flight_Precision_Limited,Flight Precision Limited
,FSL,FLIGHTSAFETY,Flight_Safety_Limited,Flight Safety Limited
,FSU,,Flight_Support_Sweden,Flight Support Sweden
,CCK,CABLE CHECK,Flight_Trac,Flight Trac
,AYR,CYGNET,Flight_Training_Europe,Flight Training Europe
,FWQ,UNITY,Flight_West_Airlines,Flight West Airlines
,KLO,KLONDIKE,Flight-Ops_International,Flight-Ops International
,CSK,CASCADE,Flightcraft,Flightcraft
,FEX,FLIGHTEXEC,,Flightexec
B5,FLT,FLIGHTLINE,Flightline_(UK),Flightline
,FTL,FLIGHT-AVIA,Flightline,Flightline
,FPS,FLIGHTPASS,Flightpass_Limited,Flightpass Limited
,FSR,FLIGHTSTAR,Flightstar_Corporation,Flightstar Corporation
,KDZ,KUDZU,Flightworks,Flightworks
,FAZ,FLINT AIR,Flint_Aviation_Services,Flint Aviation Services
,OJY,OHJAY,Florida_Air,Florida Air
PA,FCL,FLORIDA COASTAL,Florida_Coastal_Airlines,Florida Coastal Airlines
,FFS,FORESTRY,Florida_Department_of_Agriculture,Florida Department of Agriculture
,FJS,FLORIDAJET,Florida_Jet_Service,Florida Jet Service
RF,FWL,FLO WEST,Florida_West_International_Airways,Florida West International Airways
,FFG,WITCHCRAFT,Flugdienst_Fehlhaber,Flugdienst Fehlhaber
,FLU,YELLOW FLYER,Flugschule_Basel,Flugschule Basel
,EZB,EICHENBURGER,Flugschule_Eichenberger,Flugschule Eichenberger
,FWZ,,Flugwerkzeuge_Aviation_Software,Flugwerkzeuge Aviation Software
8W,EDR,BIRDVIEW,Fly_All_Ways,Fly All Ways
F2,FLM,FLY WORLD,Fly_Air,Fly Air
,FCT,DEALER,Fly_CI_Limited,Fly CI Limited
,FEE,FLY EURO,Fly_Europa_Limited,Fly Europa Limited
,FXL,FLY EXCELLENT,Fly_Excellent,Fly Excellent
,NVJ,NOUVINTER,Fly_International_Airways,Fly International Airways
OJ,FJM,GREENHEART,Fly_Jamaica_Airways,Fly Jamaica Airways
,FIL,FLYLINE,Fly_Line,Fly Line
SH,FLY,FLYBIRD,Fly_Me_Sweden,Fly Me Sweden
,IAD,FLYWEX,Fly_Wex,Fly Wex
D7,XFA,FAX AIR,FlyAsianXpress,FlyAsianXpress
TE,LIL,LITHUANIA AIR,FlyLal,FlyLal
LF,NDC,NORDIC,FlyNordic,FlyNordic
F7,BBO,BABOO,Flybaboo,Flybaboo
BE,BEE,JERSEY,Flybe,Flybe
,FCE,FLYCOLUMBIA,Flycolumbia,Flycolumbia
,FLO,FLYCOM,Flycom,Flycom
,GVG,BLUECRAFT,Flygaktiebolaget_Gota_Vingar,Flygaktiebolaget Gota Vingar
Y2,GSM,GLOBESPAN,Flyglobespan,Flyglobespan
,FPA,,Flygprestanda_AB,Flygprestanda
,ETS,EXTRANS,Flygtransporter_I_Nykoping,Flygtransporter I Nykoping
,INU,INSTRUCTOR,Flyguppdraget_Backamo,Flyguppdraget Backamo
W3,FYH,FLY HIGH,Flyhy_Cargo_Airlines,Flyhy Cargo Airlines
,FCR,FLYING CARPET,Flying_Carpet_Company,Flying Carpet Company
,FYG,FLYING GROUP,Flying_Service,Flying Service
,FGP,FLYING CENTER,Flying-Research_Aerogeophysical_Center,Flying-Research Aerogeophysical Center
,FLK,FLYLINK,Flylink_Express,Flylink Express
FP,FRE,PELICAN,FlyPelican,FlyPelican
,FTM,FLYTEAM,Flyteam_Aviation,Flyteam Aviation
,FKS,FOCUS,Focus_Air,Focus Air
,FOP,,Fokker,Fokker
,NOF,FONNA,Fonnafly,Fonnafly
,FOB,FORDAIR,Ford_Motor_Company,Ford Motor Company
VY,FOS,,Formosa_Airlines,Formosa Airlines
,FOR,FORMULA,Formula_One_Management,Formula One Management
,FHS,HELISCOT,Forth_and_Clyde_Helicopter_Services,Forth and Clyde Helicopter Services
,FXC,AIR FUTURE,Fortunair_Canada,Fortunair Canada
BN,,,Forward_Air_International_Airlines,Forward Air International Airlines
,FSA,FOSTER-AIR,Foster_Aviation,Foster Aviation
,JFY,YEOMAN,Foster_Yeoman,Foster Yeoman
,FTE,FOTOGRAFIA,Fotografia_F3,Fotografia F3
5F,FIA,Fly One,Fly_One,Fly One
HK,FSC,FOUR STAR,Four_Star_Aviation,Four Star Aviation / Four Star Cargo
,WDS,WINDS,Four_Winds_Aviation,Four Winds Aviation
,FXR,WILDFOX,Foxair_(airline),Foxair
,FDO,FRENCH CUSTOM,France_Douanes,France Douanes
FH,FHY,FREEBIRD AIR,Freebird_Airlines,Freebird Airlines
SJ,FOM,FREE AIR,Freedom_Air,Freedom Air
FP,FRE,FREEDOM,Freedom_Air_(Guam),Freedom Air
,FFF,INTER FREEDOM,Freedom_Air_Services,Freedom Air Services
,FRL,FREEDOM AIR,Freedom_Airlines,Freedom Airlines
,FAS,FREEDOM AIRWAYS,Freedom_Airways,Freedom Airways
,FWC,FREEWAY,Freeway_Air,Freeway Air
,FRG,FREIGHT RUNNERS,Freight_Runners_Express,Freight Runners Express
,FAF,FRENCH AIR FORCE,French_Air_Force,Force Aerienne Francaise
,FMY,FRENCH ARMY,French_Army,Aviation Legere De L'Armee De Terre
,FNY,FRENCH NAVY,French_Navy,France Marine Nationale
,FRR,FRESH AIR,Fresh_Air_airline,Fresh Air
,BZY,BREEZY,Fresh_Air_Aviation,Fresh Air Aviation
,FAE,WILDGOOSE,Freshaer,Freshaer
,FAL,FRIENDSHIP,Friendship_Air_Alaska,Friendship Air Alaska
,FLF,FRIEND AIR,Friendship_Airlines,Friendship Airlines
,FGY,,Froggy_Corporate_Aviation,Froggy Corporate Aviation
F9,FFT,FRONTIER FLIGHT,Frontier_Airlines,Frontier Airlines
,ITR,OUT BACK,Frontier_Commuter,Frontier Commuter
2F,FTA,FRONTIER-AIR,Frontier_Flying_Service,Frontier Flying Service
,FNG,FINNGUARD,Frontier_Guard,Frontier Guard
,FUJ,FUJAIRAH,Fujairah_Aviation_Centre,Fujairah Aviation Centre
,CFJ,FUJIAN,Fujian_Airlines,Fujian Airlines
,GAX,GRAND AIRE,Full_Express,Full Express
,FAM,FAASA,FumigaciÃ³n_AÃ©rea_Andaluza,Fumigación Aérea Andaluza
,FFY,FUN FLYING,Fun_Flying_Thai_Air_Service,Fun Flying Thai Air Service
,ROG,REGO,FundaciÃÂ³_Rego,FundaciÃ³ Rego
,FUN,FUNTSHI,Funtshi_Aviation_Service,Funtshi Aviation Service
,FGL,Applewood,Futura_Gael,Futura Gael
FH,FUA,FUTURA,Futura_International_Airways,Futura International Airways
FZ,FDB,SKYDUBAI,Flydubai,Flydubai
,FWK,,Flightworks,Flightworks
,ACT,AMERICAN CHECK,Flight_Line,Flight Line
9Y,FGE,GEORGIA WING,Fly_Georgia,Fly Georgia
,FRF,FAIRFLEET,Fleet_Air_International,Fleet Air International
VK,FVK,BALDER,FlyViking,FlyViking
GL,GRL,GREENLAND,Air_Greenland,Air Greenland
,DBC,DIAMOND BACK,Gemini_Air_Group,Gemini Air Group
,GCH,GAMA SWISS,Gama_Aviation_Switzerland,Gama Aviation Switzerland
GX,GBC,SPRAY,Guangxi_Beibu_Gulf_Airlines,Guangxi Beibu Gulf Airlines
,GOP,GOSPA AIR,Gospa_Air,Gospa Air
,HGT,HIGHTECH,GMJ_Air_Shuttle,GMJ Air Shuttle
,GRE,,Greenlandcopter,Greenlandcopter
,GMQ,CORGI,Germania_Express,Germania Express
,KNM,KINGDOM,GB_Helicopters,GB Helicopters
,GCW,GLOBALCREW,Global_Air_Crew,Global Air Crew
,GBH,,Global_Avia_Handling,Global Avia Handling
Y5,GMR,GOLDEN MYANMAR,Golden_Myanmar_Airlines,Golden Myanmar Airlines
,GML,GEEANDEL,G_&_L_Aviation,G & L Aviation
,EXH,BATMAN,G5_Executive,G5 Executive
,MTA,GAK AVIATION,GAK/Mitchell_Aero,GAK/Mitchell Aero
,GGS,GATSA,GATSA,GATSA
,GBX,ISLAND TIGER,GB_Airlink,GB Airlink
GT,GBL,GEEBEE AIRWAYS,GB_Airways,GB Airways
,GCS,GALION,GCS_Air_Service,GCS Air Service
,FFU,FERRANTI,GEC_Marconi_Avionics,GEC Marconi Avionics
,GCC,GECAS,GECAS,GECAS
,GEN,GENSA-BRASIL,GENSA,GENSA
,GET,AIR FLOW,,Get High
,GET,GETRA,GETRA,GETRA
,GFW,,GFW_Aviation,GFW Aviation
,GGT,THUNDERBALL,,Trans Island Airways
,GHI,,GH_Stansted_Limited,GH Stansted Limited
Z5,GMG,GMG,GMG_Airlines,GMG Airlines
,GPE,REGIONAL EXPRESS,GP_Express_Airlines,GP Express Airlines
,GPR,GPM AEROSERVICIO,GPM_Aeroservicio,GPM Aeroservicio
,GIB,GRAVIA,GR-Avia,GR-Avia
,BMK,MURAT,GST_Aero,GST Aero Aircompany
,GTX,BIG-DEE,GTA_Air,GTA Air
,GAH,GAMHELICO,Ga-Ma_Helicoptere,Ga-Ma Helicoptere
,GBE,GABEX,Gabon_Express,Gabon Express
,GRT,,Gabon-Air-Transport,Gabon-Air-Transport
,GIG,GACELA AIR,Gacela_Air_Cargo,Gacela Air Cargo
,GFC,GAIL FORCE,Gail_Force_Express,Gail Force Express
,GNJ,HERCULES JET,Gain_Jet_Aviation,Gain Jet Aviation
,SWF,GALAIR,Galair_International,Galair International
,GLS,GALS,Galaircervis,Galaircervis
7O,GAL,GALAXY,Galaxy_Air,Galaxy Air
,GXY,GALAX,Galaxy_Airlines_(Japan),Galaxy Airlines
,GAS,GALENA AIR SERVICE,Galena_Air_Service,Galena Air Service
1G,,,Galileo_CRS,Galileo International
,GMA,GAMA,Gama_Aviation,Gama Aviation
GC,GNR,GAMBIA INTERNATIONAL,Gambia_International_Airlines,Gambia International Airlines
,NML,NEWMILL,Gambia_New_Millennium_Air,Gambia New Millennium Air
,GMJ,GAMISA,Gamisa_AviaciÃ³n,Gamisa Aviación
G7,GNF,Gandalf,Gandalf_Airlines,Gandalf Airlines
,GAN,GANAIR,Gander_Aviation,Gander Aviation
,GSA,GARDEN STATE,Garden_State_Airlines,Garden State Airlines
,AHM,AIR HURON,Garrison_Aviation,Garrison Aviation
GA,GIA,INDONESIA,Garuda_Indonesia,Garuda Indonesia
,GHS,GATARI,Gatari_Hutama_Air_Services,Gatari Hutama Air Services
,EGO,GAUTENG,Gauteng_Air_Cargo,Gauteng Air Cargo
,GVN,GAVINA,Gavina,Gavina
4G,GZP,GAZPROMAVIA,Gazpromavia,Gazpromavia
,GEE,GEESAIR,Geesair,Geesair
GR,GCO,GEMINI,Gemini_Air_Cargo,Gemini Air Cargo
,GAB,GENDALL,Gendall_Air,Gendall Air
,GDB,BELGIAN GENERMERIE,Gendarmerie_Belge,Gendarmerie Belge
,FGN,FRANCE GENDARME,National_Gendarmerie,National Gendarmerie
,SWK,SKYWALKER,General_Aerospace,General Aerospace
,GWS,GENAIR,General_Airways,General Airways
,GNZ,GONZO,General_Aviation,General Aviation
,GTH,GOTHAM,General_Aviation_Flying_Services,General Aviation Flying Services
,XGA,,General_Aviation_Terminal,General Aviation Terminal
,GMC,GENERAL MOTORS,General_Motors,General Motors
,GNX,,Genex,Genex
,GSL,SURVEY-CANADA,Geographic_Air_Surveys,Geographic Air Surveys
A9,TGZ,TAMAZI,Georgian_Airways,Georgian Airways
,FGA,GEORGIA FED,Georgian_Aviation_Federation,Georgian Aviation Federation
,GGF,GEORGIAN AFRICA,Georgian_Cargo_Airlines_Africa,Georgian Cargo Airlines Africa
QB,GFG,NATIONAL,Georgian_National_Airlines,Georgian National Airlines
,,,Great_Barrier_Airlines,Great Barrier Airlines
,GAF,GERMAN AIR FORCE,German_Air_Force,German Air Force
,GAM,GERMAN ARMY,German_Army,German Army
,GNY,GERMAN NAVY,German_Navy,German Navy
,GHY,GERMAN SKY,German_Sky_Airlines,German Sky Airlines
ST,GMI,GERMANIA,Germania_(airline),Germania
4U,GWI,GERMAN WINGS,Germanwings,Germanwings
,GDN,,Gerry's_Dnata,Gerry's Dnata
,GFD,KITE,Gesellschaft_Fur_Flugzieldarstellung,Gesellschaft Fur Flugzieldarstellung
GP,RIV,RIVERA,APG_Airlines,APG Airlines
,GES,GESTAIR,Gestair,Gestair
,GTR,STAR GESTAR,Gestar,Gestar
,GJT,BANJET,GestiÃ³n_AÃ©rea_Ejecutiva,Gestión Aérea Ejecutiva
,GHT,,Ghadames_Air_Transport,Ghadames Air Transport
GH,GLP,GLOBUS,Globus_Airlines,Globus Airlines
GH,GHA,GHANA,Ghana_Airways,Ghana Airways
G0,GHB,GHANA AIRLINES,Ghana_International_Airlines,Ghana International Airlines
,NTC,NIGHT CHASE,Gibson_Aviation,Gibson Aviation
,RPS,RESPONSE,Global_Air_Charter,Global Air Charter
,GAG,GEEBIRD,Greybird_Pilot_Academy,Greybird Pilot Academy
,GBS,GLOBAL SERVE,Global_Air_Services_Nigeria,Global Air Services Nigeria
,GLC,,Global_Aircargo,Global Aircargo
,,,Global_Airways_(Turks_and_Caicos),Global Airways (Turks and Caicos)
,BSP,,Global_Airways_(BSP),Global Airways (BSP)
,GLB,GLO-AIR,Global_Airways_(GLB),Global Airways (GLB)
,GBB,GLOBE,Global_Aviation_Operations,Global Aviation Operations
,GAK,AVIAGROUP,Global_Aviation_and_Services_Group,Global Aviation and Services Group
,GGZ,GLOBAL GEORGIAN,Global_Georgian_Airways,Global Georgian Airways
,GLJ,GLOBAL AUSTRIA,Global_Jet_Austria,Global Jet Austria
,NSM,THUNDERCLOUD,Global_Jet_Corporation,Global Jet Corporation
,SVW,SILVER ARROWS,Global_Jet_Luxembourg,Global Jet Luxembourg
,GSK,GLOBAL SKY,Global_Sky_Aircharter,Global Sky Aircharter
,GSS,JET LIFT,Global_Supply_Systems,Global Supply Systems
,XGS,,Global_System,Global System
,XGW,,Global_Weather_Dynamics,Global Weather Dynamics
,GLW,,Global_Wings,Global Wings
,GJA,,Globe_Jet,Globe Jet
G8,GOW,GOAIR,GoAir,GoAir
GK,,,Go_One_Airways,Go One Airways
G7,GJS,LINDBERGH,GoJet_Airlines,GoJet Airlines
,GGE,,Gobierno_De_Guinea_Ecuatorial,Gobierno De Guinea Ecuatorial
,GOF,GOF-AIR,Gof-Air,Gof-Air
,GOI,SWISS HAWK,Gofir,Gofir
G3,GLO,GOL TRANSPORTE,Gol_Transportes_A%C3%A9reos,Gol Transportes Aéreos
,GBT,GOLD BELT,Gold_Belt_Air_Transport,Gold Belt Air Transport
,GDA,AIR PARTNER,GoldAir,GoldAir
,GDK,GOLDECK FLUG,Goldeck-Flug,Goldeck-Flug
DC,GAO,GOLDEN,Golden_Air,Golden Air
,GDD,GOLDEN AIRLINES,Golden_Airlines,Golden Airlines
,GPA,GOLDEN PAC,Golden_Pacific_Airlines,Golden Pacific Airlines
,GRS,GOLDEN RULE,Golden_Rule_Airlines,Golden Rule Airlines
,GLD,GOLDEN STAR,Golden_Star_Air_Cargo,Golden Star Air Cargo
,GOS,,Goldfields_Air_Services,Goldfields Air Services
,GAQ,GOLFAIR,Golfe_Air_Quebec,Golfe Air Quebec
,GLE,GOLIAF AIR,Goliaf_Air,Goliaf Air
,GOM,GOMEL,Gomel_Airlines,Gomel Airlines
5Z,GON,GONINI,Gonini_Air_Services,Gonini Air Services
,RDR,RED STAR,Goodridge_(UK)_Limited,Goodridge (UK) Limited
G1,,GORKHA AIRLINES,Gorkha_Airlines,Gorkha Airlines
,GOR,GORLITSA,Gorlitsa_Airlines,Gorlitsa Airlines
,HKG,HONGKONG GOVERNMENT,Government_Flying_Service,Government Flying Service
,GRZ,COM FLIGHT,Zambia,Government of Zambia Communications Flight
,HLD,GRANITE,Grampian_Flight_Centre,Grampian Flight Centre
,GAV,GRANAVI,Granada_AviaciÃ³n,Granada Aviación
,GAE,GRAND EXPRESS,Grand_Aire_Express,Grand Aire Express
,GND,GRAND VEGAS,Grand_Airways,Grand Airways
,CVU,CANYON VIEW,Grand_Canyon_Airlines,Grand Canyon Airlines
GV,GUN,HOOT,Grant_Aviation,Grant Aviation
,LMK,LANDMARK,Grantex_Aviation,Grantex Aviation
,GRA,GREAT AMERICAN,Great_American_Airways,Great American Airways
,GRA,FLEX,,Guardian Air Assest Management
ZK,GLA,LAKES AIR,Great_Lakes_Airlines,Great Lakes Airlines
,GLU,LAKES CARGO,Great_Lakes_Airways_(Uganda),Great Lakes Airways (Uganda)
,GRP,GREAT PLAINS,Great_Plains_Airlines,Great Plains Airlines
IJ,GWL,GREAT WALL,Great_Wall_Airlines,Great Wall Airlines
,GWA,G-W AIR,Great_Western_Air,Great Western Air
,HAF,HELLENIC AIR FORCE,Greek_Air_Force,Hellenic Air Force
,HNA,HELLENIC NAVY,Greek_Navy,Greek Navy
,GFF,GRIFFIN AIR,Griffin_Aviation,Griffin Aviation
,GXA,GRIXONA,Grixona,Grixona
,GZD,GRIZODUBOVA AIR,Grizodubova_Air_Company,Grizodubova Air Company
,HTG,GROSSMANN,Grossmann_Air_Service,Grossmann Air Service
,GSJ,GROSSJET,Grossmann_Jet_Service,Grossmann Jet Service
,GHV,GROUND HANDLING,Ground_Handling_Service_de_Mexico,Ground Handling Service de Mexico
,GPM,GRUPOMED,Grup_Air-Med,Grup Air-Med
,EJC,GRUPOEJECUTIVA,Grupo_De_AviaciÃ³n_Ejecutiva,Grupo De Aviación Ejecutiva
,TAT,TACA-COSTARICA,Grupo_TACA,Grupo TACA
,VMM,VUELOS MED,Grupo_Vuelos_Mediterraneo,Grupo Vuelos Mediterraneo
,GMT,GRUPOMONTERREY,Grupo_AÃ©reo_Monterrey,Grupo Aéreo Monterrey
,GSY,GUARD AIR,Guard_Systems,Guard Systems
G6,BSR,BISSAU AIRLINES,Guine_Bissaur_Airlines,Guine Bissaur Airlines
,GIJ,GUINEA AIRWAYS,Guinea_Airways,Guinea Airways
,GNC,GUINEA CARGO,Guinea_Cargo,Guinea Cargo
J9,GIF,GUINEE AIRLINES,Guinee_Airlines,Guinee Airlines
,GEA,GEASA,Guinee_Ecuatorial_Airlines,Guinea Ecuatorial Airlines
,GIQ,GUIPAR,Guinee_Paramount_Airlines,Guinee Paramount Airlines
,CGH,GUIZHOU,Guizhou_Airlines,Guizhou Airlines
,GUS,GUJA,Guja,Guja
G8,GUJ,GUJARATAIR,Gujarat_Airways,Gujarat Airways
,TSU,TRANSAUTO,Contract_Air_Cargo,Gulf & Caribbean Cargo / Contract Air Cargo
,GUF,GULF AFRICAN,Gulf_African_Airlines,Gulf African Airlines
GF,GFA,GULF AIR,Gulf_Air,Gulf Air
,GAT,GULF TRANS,Gulf_Air_Inc,Gulf Air Inc
,GCN,GULF CENTRAL,Gulf_Central_Airlines,Gulf Central Airlines
,SFY,SKY FLITE,Gulf_Flite_Center,Gulf Flite Center
,GPC,AIR GULFPEARL,Gulf_Pearl_Air_Lines,Gulf Pearl Air Lines
,GLF,GULFSTREAM TEST,Gulfstream_Aerospace,Gulfstream Aerospace
,GFS,GULFSTAR,Gulfstream_Airlines,Gulfstream Airlines
,GFT,GULF FLIGHT,Gulfstream_International_Airlines,Gulfstream International Airlines
,GUL,GULL-AIR,Gull_Air,Gull Air
,GUM,GUM AIR,Gum_Air,Gum Air
,GDH,RISING SUN,Guneydogu_Havacilik_Isletmesi,Guneydogu Havacilik Isletmesi
GY,,,Guyana_Airways_2000,Guyana Airways 2000
,GWN,GWYN,Gwyn_Aviation,Gwyn Aviation
,GAC,DREAM TEAM,GlobeAir,GlobeAir
3S,GUY,GREEN BIRD,Air_Guyane_Express,Air Guyane Express
,HTB,HELIX-CRAFT,Helix-Craft_Aviation,Helix-Craft Aviation
,HCK,HELI-CHARTER,Heli-Charter,Heli-Charter
MR,MML,TRANS MONGOLIA,Hunnu_Air,Hunnu Air
,HAY,HAMBURG AIRWAYS,Hamburg_Airways,Hamburg Airways
,HSN,H.S.AVIATION,"H.S.AVIATION_CO.,_LTD.","H.S.AVIATION CO., LTD."
,HRN,HERONAIR,Heron_Luftfahrt,Heron Luftfahrt
,HYP,HYPERION,Hyperion_Aviation,Hyperion Aviation
,HFM,MOONRAKER,Hi_Fly_Malta,Hi Fly Malta
A5,HOP,AIR HOP,Hop!,Hop!
,HLA,HEAVYLIFT,HC_Airlines,HC Airlines
,HWD,FLITEWISE,HPM_Investments,HPM Investments
,KTR,COPTER TRANS,HT_Helikoptertransport,HT Helikoptertransport
,AHT,HELIAPRA,HTA_Helicopteros,HTA Helicopteros
,,,Hacienda_Airlines,Hacienda Airlines
,FMS,HADI,Hadison_Aviation,Hadison Aviation
H6,HAG,HAGELAND,Hageland_Aviation_Services,Hageland Aviation Services
,POW,AIRNET,Hagondale_Limited,Hagondale Limited
HR,HHN,ROOSTER,Hahn_Air,Hahn Air
H1,,,,Hahn Air Systems
HU,CHH,HAINAN,Hainan_Airlines,Hainan Airlines
1R,,,Hainan_Phoenix_Information_Systems,Hainan Phoenix Information Systems
,HLS,,Haiti_Air_Freight,Haiti Air Freight
2T,HAM,,Haiti_Ambassador_Airlines,Haiti Ambassador Airlines
,HTI,HAITI INTERNATIONAL,Haiti_International_Air,Haiti International Air
,HRB,HAITI AIRLINE,Haiti_International_Airline,Haiti International Airline
,HNR,HANAIR,Haiti_National_Airlines,Haiti National Airlines (HANA)
,HTN,,Haiti_North_Airline,Haiti North Airline
,HTC,HAITI TRANSAIR,Haiti_Trans_Air,Haiti Trans Air
,HBC,HALISA,Haitian_Aviation_Line,Haitian Aviation Line
,HAJ,HAJVAIRY,Hajvairy_Airlines,Hajvairy Airlines
,HKL,HAK AIRLINE,Hak_Air,Hak Air
,HLH,HALA AIR,Hala_Air,Hala Air
,HCV,CREOLE,Halcyonair,Halcyonair
4R,HHI,HAMBURG JET,Hamburg_International,Hamburg International
,HJL,BIZJET,Hamlin_Jet,Hamlin Jet
,HMM,HAMRA,Hamra_Air,Hamra Air
,WVA,WABASH VALLEY,Hand_D_Aviation,Hand D Aviation
,HGR,HANG,Hangar_8,Hangar 8
,HGD,HANGARD,Hangard_Aviation,Hangard Aviation
,HAN,HANSUNG AIR,Hansung_Airlines,Hansung Airlines
X3,HLX,YELLOW CAB,Hapag-Lloyd_Express,Hapag-Lloyd Express (TUIfly)
HF,HLF,HAPAG LLOYD,Hapagfly,Hapagfly
HB,HAR,HARBOR,Harbor_Airlines,Harbor Airlines
HQ,HMY,HARMONY,Harmony_Airways,Harmony Airways
,NBR,NORBROOK,Haughey_Air,Haughey Air
,PYN,POYSTON,Haverfordwest_Air_Charter_Services,Haverfordwest Air Charter Services
,HAV,HAVILAH,Havilah_Air_Services,Havilah Air Services
HA,HAL,HAWAIIAN,Hawaiian_Airlines,Hawaiian Airlines
HP,,,Hawaiian_Pacific_Airlines,Hawaiian Pacific Airlines
,HKR,AIR HAW,Hawk_Air,Hawk Air
,HMX,HAWK MEXICO,Hawk_De_Mexico,Hawk De Mexico
BH,,,Hawkair,Hawkair
,HKI,HAWKEYE,Hawkaire,Hawkaire
,HZL,HAZELTON,Hazelton_Airlines,Hazelton Airlines
HN,HVY,HEAVY CARGO,Heavylift_Cargo_Airlines,Heavylift Cargo Airlines
,HVL,HEAVYLIFT INTERNATIONAL,Heavylift_International,Heavylift International
,HDC,HELICATALUNA,Helcopteros_De_Cataluna,Helcopteros De Cataluna
,HCB,HELEN,Helenair_(Barbados),Helenair (Barbados)
,HCL,HELENCORP,Helenair_Corporation,Helenair Corporation
,HHP,HELENIA,Helenia_Helicopter_Service,Helenia Helicopter Service
,HLR,HELI BULGARIA,Heli_Air_Services,Heli Air Services
,ALJ,ALPIN HELI,Heli_Ambulance_Team,Heli Ambulance Team
,HEB,HELIBERNINA,Heli_Bernina,Heli Bernina
,HFR,HELIFRANCE,Heli_France,Heli France
,HYH,HELIHUNGARY,Heli_Hungary,Heli Hungary
,HLM,HELIMIDWEST,Heli_Medwest_De_Mexico,Heli Medwest De Mexico
,HLI,HELI SAINT-TROPEZ,Heli_Securite,Heli Securite
,HTP,HELI TRIP,Heli_Trip,Heli Trip
,HLU,HELI UNION,Heli_Union_Heli_Prestations,Heli Union Heli Prestations
,MCM,HELI AIR,Heli-Air-Monaco,Heli-Air-Monaco
,HHE,HELI HOLLAND,Heli-Holland,Heli-Holland
,HRA,ERICA,Heli-Iberica,Heli-Iberica
,HIF,HIFSA,Heli-Iberica_Fotogrametria,Heli-Iberica Fotogrametria
,HIG,INTER GUYANNE,Heli-Inter_Guyane,Heli-Inter Guyane
,HLK,HELI-LINK,Heli-Link,Heli-Link
,HMC,HELIAMERICA,Heliamerica_De_Mexico,Heliamerica De Mexico
,HEA,HELIAVIA,Heliavia-Transporte_AÃ©reo,Heliavia-Transporte Aéreo
,CDY,CADDY,Heliaviation_Limited,Heliaviation Limited
,HIB,HELIBRAVO,Helibravo_Aviacao,Helibravo Aviacao
,HLC,HELICAP,Helicap,Helicap
,COV,HELICENTRE,Helicentre_Coventry,Helicentre Coventry
,HEL,HELICOL,Helicol,Helicol
,HCP,HELI CZECH,Helicopter,Helicopter
,JKY,JOCKEY,Helicopter_&_Aviation_Services,Helicopter & Aviation Services
,MVK,MAVRIK,Helicopter_Training_&_Hire,Helicopter Training & Hire
,HAP,HELIPERSONAL,Helicopteros_Aero_Personal,Helicopteros Aero Personal
,HAA,AGROFORESTAL,Helicopteros_Agroforestal,Helicopteros Agroforestal
,HNT,HELICOP INTER,Helicopteros_Internacionales,Helicopteros Internacionales
,HEN,HELINAC,HelicÃ³pteros_Y_VehÃ­culos_Nacionales_AÃ©reos,Helicópteros Y Vehículos Nacionales Aéreos
,HHH,HELICSA,Helicsa,Helicsa
JB,JBA,HELIJET,Helijet,Helijet
,HDR,HELIDRIFT,Helikopterdrift,Helikopterdrift
,SCO,SWEDCOPTER,Helikopterservice_Euro_Air,Helikopterservice Euro Air
,OCE,HELIOCEAN,Heliocean,Heliocean
ZU,HCY,HELIOS,Helios_Airways,Helios Airways
,HLP,HELIPISTAS,Helipistas,Helipistas
,HPL,HELIPORTUGAL,Heliportugal,Heliportugal
,HEC,HELICAMPECHE,Heliservicio_Campeche,Heliservicio Campeche
,HSU,HELIS,Helisul,Helisul
,HSI,HELISWISS,Heliswiss,Heliswiss
,HLT,HELITAFE,Helitafe,Helitafe
,HIT,HELITALIA,Helitalia,Helitalia
,OFA,OFAVI,Helitaxi_Ofavi,Helitaxi Ofavi
,HLT,HELITOURS,Helitours,Helitours
9I,HTA,SCANBIRD,Helitrans,Helitrans
,HTS,HELITRANS,Helitrans_Air_Service,Helitrans Air Service
,HLW,HELIWORKS,Heliworks,Heliworks
HJ,HEJ,HELLAS JET,Hellas_Jet,Hellas Jet
HW,FHE,FLYHELLO,Hello_(airline),Hello
,HLG,HELOG,Helog,Helog
2L,OAW,HELVETIC,Helvetic_Airways,Helvetic Airways
,HMS,HEMUS AIR,Hemus_Air,Hemus Air
,HAC,,Henebury_Aviation,Henebury Aviation
,SSH,SNOWSHOE,Heritage_Flight,Heritage Flight (Valley Air Services)
,MRX,SPEEDMARK,Herman's_Markair_Express,Herman's Markair Express
,HED,FLAPJACK,Heritage_Aviation_Developments,Heritage Aviation Developments
EO,ALX,ALLCONGO,Hewa_Bora_Airways,Hewa Bora Airways
UD,HER,HEX AIRLINE,Hex%27Air,Hex'Air
,HHS,HIJET,Hi-Jet_Helicopter_Services,Hi-Jet Helicopter Services
5K,HFY,SKY FLYER,Hi_Fly_(airline),Hi Fly
,HLB,HIGH-LINE,High-Line_Airways,High-Line Airways
,HWY,HIWAY,Highland_Airways,Highland Airways
,HSH,HASA,HispÃ¡nica_de_AviaciÃ³n,Hispánica de Aviación
,HIS,HISPANIOLA,Hispaniola_Airways,Hispaniola Airways
,HGA,HOGAN AIR,Hogan_Air,Hogan Air
,NTH,NORTH AIR,Hokkaido_Air_System,Hokkaido Air System
,ABH,,Hokuriki-Koukuu_Company,Hokuriki-Koukuu Company
H5,HOA,HOLA,Hola_Airlines,Hola Airlines
,HIN,HOLDING GROUP,Holding_International_Group,Holding International Group
,HOL,HOLIDAY,Holiday_Airlines,Holiday Airlines
HC,HCC,CZECH HOLIDAYS,Holidays_Czech_Airlines,Holidays Czech Airlines
,HTR,HOLSTEN,Holstenair_Lubeck,Holstenair Lubeck
,HMV,HOMAC,Homac_Aviation,Homac Aviation
,HAS,HONDURAS AIR,Honduras_Airlines,Honduras Airlines
HX,CRK,BAUHINIA,Hong_Kong_Airlines,Hong Kong Airlines
RH,HKC,MASCOT,Hong_Kong_Air_Cargo,Hong Kong Air Cargo
UO,HKE,HONGKONG SHUTTLE,Hong_Kong_Express_Airways,Hong Kong Express Airways
A6,HTU,HONGLAND,Hongtu_Airlines,Hongtu Airlines
,HEX,HONIARA CARGO,Honiara_Cargo_Express,Honiara Cargo Express
,HPJ,HOPA-JET,Hop-A-Jet,Hop-A-Jet
HH,,HOPE AIR,Hope_Air,Hope Air
QX,QXE,HORIZON AIR,Horizon_Air,Horizon Air
,KOK,KOKO,Horizon_Air_Service,Horizon Air Service
,HSM,ALOFUKAIR,Horizon_Air_for_Transport_and_Training,Horizon Air for Transport and Training
,HOR,HORIZON,Horizon_Air-Taxi,Horizon Air-Taxi
BN,HZA,HORIZON,Horizon_Airlines_(Australia),Horizon Airlines
,HPS,HORIZON PLUS,Horizon_Plus,Horizon Plus
,HUD,HUD,Horizons_Unlimited,Horizons Unlimited
,HOZ,HORIZONTES AEREOS,Horizontes_AÃ©reos,Horizontes Aéreos
,HDI,DINAMICOS,Hoteles_Dinamicos,Hoteles Dinamicos
,HHO,HOUSTON HELI,Houston_Helicopters,Houston Helicopters
,GGV,GREGG AIR,Houston_Jet_Services,Houston Jet Services
,OZU,HOZAVIA,Hozu-Avia,Hozu-Avia
,HUB,HUB,Hub_Airlines,Hub Airlines
,HUS,HUESSLER,Huessler_Air_Service,Huessler Air Service
,GMH,HUGHES EXPRESS,Hughes_Aircraft_Company,Hughes Aircraft Company
,WHR,WHIRLEYBIRD,Hummingbird_Helicopter_Service,Hummingbird Helicopter Service
,HUV,SILVER EAGLE,Hunair_Hungarian_Airlines,Hunair Hungarian Airlines
,HUF,HUNGARIAN AIRFORCE,Hungarian_Air_Force,Hungarian Air Force
,HYA,HYACK,Hyack_Air,Hyack Air
,HYC,HYDRO CARGO,Hydro_Air_Flight_Operations,Hydro Air Flight Operations
,HYD,HYDRO,Hydro-Qu%C3%A9bec,Hydro-Québec
H4,,,HÃ©li_SÃ©curitÃ©_Helicopter_Airlines,Héli Sécurité Helicopter Airlines
,HKB,CLASSIC,Hawker_Beechcraft,Hawker Beechcraft
,ETI,JETHAWK,H-Bird_Aviation_Services_AB,H-Bird Aviation Services AB
,IJW,JET WEST,,Interjet Inc.
,IIC,,Inter_Island_Air_Charter,Inter Island Air Charter
,EXP,ISLAND EXPRESS,Island_Air_Express,Island Air Express
,IWL,,Island_Wings,Island Wings
IK,KAR,IKAR,Ikar_(airline),Ikar
,ICN,,Iconair,Iconair
,IAC,INTERCHARTER,Interaviation_Charter,Interaviation Charter
,IFM,ICOPTER,Ifly,Ifly
,ITC,,International_Air_Carrier_Association,International Air Carrier Association
II,CSQ,CHASQUI,IBC_Airways,IBC Airways
0C,IBL,CATOVAIR,IBL_Aviation,IBL Aviation
,BBL,BLUE,IBM,IBM Euroflight Operations
,YYY,,,ICAO
C3,IPR,ICAR,ICAR_(aircraft_manufacturer),Independent Carrier (ICAR)
,CIC,AIR TRADER,ICC_Canada,ICC Canada
,IDG,INDIGO,IDG_Technology_Air,IDG Technology Air
,IFL,EIFEL,IFL_Group,IFL Group
,RDE,FLIGHT RED,II_Lione_Alato_Arl,II Lione Alato Arl
,IJM,JET MANAGEMENT,IJM_International_Jet_Management,IJM International Jet Management
,IKK,IKIAIR,IKI_International_Airlines,IKI International Airlines
,IKN,IKON,IKON_FTO,IKON FTO
,BLU,BLUENOSE,IMP_Aviation_Services,IMP Aviation Services
,XGG,,IMP_Group_Aviation_Services,IMP Group Aviation Services
1F,,,INFINI_Travel_Information,INFINI Travel Information
,IPA,IPEC,IPEC_Aviation,IPEC Aviation
,IPM,SHIPEX,IPM_Europe,IPM Europe
,LVB,SILVERBIRD,IRS_Airlines,IRS Airlines
,ISD,ISDAVIA,ISD_Avia,ISD Avia
1U,,,ITA_Software,ITA Software
,FDF,,IVV_Femida,IVV Femida
IB,IBE,IBERIA,Iberia_Airlines,Iberia Airlines
I2,IBS,IBEREXPRESS,Iberia_Express,Iberia Express
,IBR,IBERTOUR,Ibertour_Servicios_AÃ©reos,Ibertour Servicios Aéreos
,IBT,IBERTRANS,Ibertrans_A%C3%A9rea,Ibertrans Aérea
TY,IWD,IBERWORLD,Iberworld,Iberworld
FW,IBX,IBEX,Ibex_Airlines,Ibex Airlines
,IBC,IBICENCA,Ibicenca_Air,Ibicenca Air
,AKI,,Ibk-Petra,Ibk-Petra
,RAC,TUZLA AIR,Icar_Air,Icar Air
,FRC,FRANCHE COMPTE,Icare_Franche_Compte,Icare Franche Compte
D8,IBK,NORTRANS,Norwegian_Air_International_(Norway),Norwegian Air International (Norway)
,ICD,ICARO,Icaro_(Ecuador),Icaro Air
,ICA,ICARFLY,Icaro_(Italy),Icaro
,IUS,ICARUS,Icarus_(airline),Icarus
,ICJ,ICEJET,Icejet,Icejet
FI,ICE,ICEAIR,Icelandair,Icelandair
,ICG,ICELAND COAST,Icelandic_Coast_Guard,Icelandic Coast Guard
,IKR,IKAROS,Ikaros_DK,Ikaros DK
,CIO,CIOCCO,Il_Ciocco_International_Travel_Service,Il Ciocco International Travel Service
,ILV,ILAVIA,Il-Avia,Il-Avia
,IDL,ILDEFONSO,Ildefonso_Redriguez,Ildefonso Redriguez
V8,IAR,ILIAMNA AIR,Iliamna_Air_Taxi,Iliamna Air Taxi
,ILP,,Ilpo_Aruba_Cargo,Ilpo Aruba Cargo
,ILL,ILYICHAVIA,Ilyich-Avia,Ilyich-Avia
,IMR,IMAER,Imaer,Imaer
,ITX,IMPROTEX,Imair_Airlines,Imair Airlines
,PNX,PHOENIX,Imperial_Airways,Imperial Airways
,IMT,IMTREC,Imtrec_Aviation,Imtrec Aviation
DH,IDE,INDEPENDENCE AIR,Independence_Air,Independence Air
,IDP,INDEPENDENT,Independent_Air_Freighters,Independent Air Freighters
6E,IGO,IFLY,IndiGo_Airlines,IndiGo Airlines
,IIL,INDIA INTER,India_International_Airways,India International Airways
,IFC,INDIAN AIRFORCE,Indian_Air_Force,Indian Air Force
IC,IAC,INDAIR,Indian_Airlines,Indian Airlines
,IDR,INDICATOR,Indicator_Company,Indicator Company
I9,IBU,INDIGO BLUE,Indigo_Airlines_(American_airline),Indigo Airlines
,IDA,INTRA,Indonesia_Air_Transport,Indonesia Air Transport
QZ,AWQ,WAGON AIR,Indonesia_AirAsia,Indonesia AirAsia
IO,IAA,INDO LINES,Indonesian_Airlines,Indonesian Airlines
,IPN,NUSANTARA,Industri_Pesawat_Terbang_Nusantara,Industri Pesawat Terbang Nusantara
,ITN,TITANLUX,Industrias_Titan,Industrias Titan
,FFI,INFINIT,Infinit_Air,Infinit Air
C4,IMX,ZIMEX,Zimex_Aviation,Zimex Aviation
,INS,,Inflite_The_Jet_Centre,Inflite The Jet Centre
,IVA,INNOTECH,Innotech_Aviation,Innotech Aviation
,INC,INSELAIR,Insel_Air_International,Insel Air International
,ICC,CARTO,Institut_CartogrÃ fic_de_Catalunya,Institut Cartogràfic de Catalunya
,INT,INTAIRCO,Intair,Intair
,INL,INTAL AVIA,Intal_Avia,Intal Avia
,FFL,,Intavia_Limited,Intavia Limited
,XRA,INTENSIVE,Intensive_Air,Intensive Air
,ITW,INTER WINGS,Inter_Air,Inter Air
,INX,INTER-EURO,Inter_Airlines,Inter Express
H4,IIN,,Inter_Island_Airways,Inter Island Airways
,CAR,QUEBEC ROMEO,Inter_RCA,Inter RCA
,NTT,INTER-TROPIC,Inter_Tropic_Airlines,Inter Tropic Airlines
,TCU,TROPAIR,Inter_Tropical_Aviation,Inter Tropical Aviation
,ITA,CAFEX,Inter-Air,Inter-Air
,ICN,INTER-CANADIAN,Inter-Canadian,Inter-Canadian
,UGL,UGLY VAN,Inter-Island_Air,Inter-Island Air
,IMA,INTER-MOUNTAIN,Inter-Mountain_Airways,Inter-Mountain Airways
,ITS,INTER-STATE,Inter-State_Aviation,Inter-State Aviation
D6,ILN,INLINE,Interair_South_Africa,Interair South Africa
,NTE,INTERMEX,Interaire,Interaire
ZA,SUW,ASTAIR,Interavia_Airlines,Interavia Airlines
,IVT,INTERAVIA,Interaviatrans,Interaviatrans
RS,ICT,CONTAVIA,Intercontinental_de_Aviaci%C3%B3n,Intercontinental de Aviación
,ICP,CHOPER,Intercopters,Intercopters
,IFT,INTERFLIGHT,Interflight,Interflight
,IJT,,Interflight_(Learjet),Interflight (Learjet)
,RFL,INFLY,Interfly,Interfly
,IFF,INTERFREIGHT,Interfreight_Forwarding,Interfreight Forwarding
,IGN,DIVINE AIR,Interguide_Air,Interguide Air
,ISN,TRI-BIRD,Interisland_Airlines,Interisland Airlines
,IWY,ISLANDWAYS,Interisland_Airways_Limited,Interisland Airways Limited
,MTF,INTERJET,Interjet,Interjet
,IHE,INTERCOPTER,Interjet_Helicopters,Interjet Helicopters
ID,ITK,INTERLINK,Interlink_Airlines,Interlink Airlines
,IAK,AIR CARGO EGYPT,International_Air_Cargo_Corporation,International Air Cargo Corporation
,EXX,EXPRESS INTERNATIONAL,International_Air_Corporation,International Air Corporation
,IAS,STARFLEET,International_Air_Service,International Air Service
,IAX,INTERAIR SERVICES,International_Air_Services,International Air Services
6I,IBZ,INTERBIZ,International_Business_Air,International Business Air
,IBY,CENTRAL STAGE,International_Business_Aircraft,International Business Aircraft
,ICS,INTERSERVI,International_Charter_Services,International Charter Services
,ICX,INTEX,International_Charter_Xpress,International Charter Xpress
,RED,RED CROSS,International_Committee_of_the_Red_Cross,International Committee of the Red Cross
,IIG,ALDAWLYH AIR,"International_Company_for_Transport,_Trade_and_Public_Works","International Company for Transport, Trade and Public Works"
,IFX,IFTA,International_Flight_Training_Academy,International Flight Training Academy
,IJA,I-JET,International_Jet_Aviation_Services,International Jet Aviation Services
,HSP,HOSPITAL,,International Jet Charter
,RSQ,SKYMEDIC,International_SOS_WIndhoek,International SOS WIndhoek
,ISF,,International_Stabilisation_Assistance_Force,International Stabilisation Assistance Force
,THN,ATHENA,International_Security_Assistance_Force,International Security Assistance Force
,ITH,INTRANS NIGERIA,International_Trans-Air,International Trans-Air
,IPT,INTERPORT,Interport_Corporation,Interport Corporation
,IKY,GENERAL SKY,Intersky_Bulgary,Intersky Bulgary
3L,ISK,INTERSKY,Intersky,Intersky
I4,FWA,FREEWAYAIR,Interstate_Airlines,Interstate Airlines
,ITU,INTERLOS,Intervuelos,Intervuelos
,INV,INVER,Inversija,Inversija
,IND,IONA,Iona_National_Airways,Iona National Airways
,IOA,IOWA AIR,Iowa_Airways,Iowa Airways
IR,IRA,IRANAIR,Iran_Air,Iran Air
EP,IRC,ASEMAN,Iran_Aseman_Airlines,Iran Aseman Airlines
,IRB,,Iranair_Tours,Iranair Tours
,IRG,NAFT,Iranian_Naft_Airlines,Iranian Naft Airlines
IA,IAW,IRAQI,Iraqi_Airways,Iraqi Airways
,BIS,IRBIS,Irbis_Air,Irbis Air
,IRL,IRISH,Irish_Air_Corps,Irish Air Corps
,RDK,IRISH TRANS,Irish_Air_Transport,Irish Air Transport
,XMR,AUTHORITY,Irish_Aviation_Authority,Irish Aviation Authority
IH,MZA,IRTYSH AIRLINES,Irtysh_Air,Irtysh Air
,KCE,KACEY,Irving_Oil,Irving Oil
,ISI,ISLANDMEX,Island_Air_(Mexico),Island Air
,ILF,ISLAND FLIGHT,Island_Air_Charters,Island Air Charters
,XYZ,RAINBIRD,Island_Air_Express,Island Air Express
,ISA,,Island_Airlines,Island Airlines
,SOY,SORIANO,Island_Aviation,Island Aviation
,DQA,,Island_Aviation_Services,Island Aviation Services
,IOM,ISLE AVIA,Island_Aviation_and_Travel,Island Aviation and Travel
2S,SDY,SANDY ISLE,Island_Express_(airline),Island Express
,MTP,METROCOPTER,Island_Helicopters,Island Helicopters
,ILC,,ILAS_Air,ILAS Air
,IAJ,JARLAND,Islandair_Jersey,Islandair Jersey
CN,,,Islands_Nationair,Islands Nationair
,ICB,ICEBIRD,Islandsflug,Icebird Airline
IF,ISW,PINTADERA,Islas_Airways,Islas Airways
,IGS,ISLA GRANDE,Isle_Grande_Flying_School,Isle Grande Flying School
WC,ISV,,Islena_De_Inversiones,Islena De Inversiones
,IOS,SCILLONIA,Isles_of_Scilly_Skybus,Isles of Scilly Skybus
,IAI,ISRAEL AIRCRAFT,Israel_Aircraft_Industries,Israel Aircraft Industries
,IAF,,Israeli_Air_Force,Israeli Air Force
6H,ISR,ISRAIR,Israir,Israir
,IST,ISTANBUL,Istanbul_Airlines,Istanbul Airlines
FS,ACL,SPADA,Itali_Airlines,Itali Airlines
,IFS,RIVIERA,Italy_First_(airline),Italy First
GI,IKA,ITEK-AIR,Itek_Air,Itek Air
,IVS,IVOIRE AERO,Ivoire_Aero_Services,Ivoire Aero Services
,IVW,IVOIRAIRWAYS,Ivoire_Airways,Ivoire Airways
,IJE,IVOIRE JET,Ivoire_Jet_Express,Ivoire Jet Express
,OIC,,Iwamoto_Crane_Co_Ltd,Iwamoto Crane Co Ltd
,IXR,X-BIRD,Ixair,Ixair
H9,IZM,IZMIR,Izair,Izair
,IZA,IZHAVIA,Izhavia,Izhavia
HC,,,Iceland_Express,Iceland Express
,IMG,IMPERIAL AIRLINES,Imperial_Cargo_Airlines,Imperial Cargo Airlines
,JTF,JETFIN,Jet_Time_(Finland),Jet Time
,JAS,,Jet_Aviation_Flight_Services_Inc,Jet Aviation Flight Services Inc
,JTN,JET TEST,Jet_Test_International,Jet Test International
,JGJ,GLOBAL JINGGONG,Jinggong_Jet,Jinggong Jet
,JNY,UNIJET-ROCKBAND,Journey_Aviation,Journey Aviation
,JKR,JOKER,Justice_Air_Charter,Justice Air Charter
,JWD,,Jayawijaya_Dirgantara,Jayawijaya Dirgantara
,JSH,STREAM AIR,Jet-stream,Jet-stream
,JCB,JAYSEEBEE,J_C_Bamford_(Excavators),J C Bamford (Excavators)
,RFX,REFLEX,J_P_Hunt,J P Hunt Air Carriers
XM,,J AIR,J-Air,J-Air
JC,JEX,JANEX,JAL_Express,JAL Express
JO,JAZ,JALWAYS,JALways,JALways
,JDA,JAY DEE,JDAviation,JDAviation
,JDP,RED PELICAN,JDP_Lux,JDP Lux
,JHM,,JHM_Cargo_Expreso,JHM Cargo Expreso
,TQM,TACOMA,JM_Family_Aviation,JM Family Aviation
MT,JMC,JAYEMMSEE,JMC_Airlines,JMC Airlines
1M,,,JSC_Transport_Automated_Information_Systems,JSC Transport Automated Information Systems
,JSJ,JS CHARTER,JS_Air,JS Air
,JES,JAY-ESS AVIATION,JS_Aviation,JS Aviation
,JCK,JACKSON,Jackson_Air_Services,Jackson Air Services
JI,JAE,JADE CARGO,Jade_Cargo_International,Jade Cargo International
,JAW,JAW,Jamahiriya_Airways,Jamahiriya Airways
,JMB,JAMBOAFRICA,Jambo_Africa_Airlines,Jambo Africa Airlines
,WWW,JANET,Janet_(airline),Janet
,FJX,,Jet_Sky_Cargo_and_Air_Charter,Jet Sky Cargo and Air Charter
,JAK,YANZAR,Jana-Arka,Jana-Arka
,JAX,JANAIR,Janair,Janair
3X,JAC,COMMUTER,Japan_Air_Commuter,Japan Air Commuter
,JSV,,Japan_Aircraft_Service,Japan Aircraft Service
JL,JAL,JAPANAIR,Japan_Airlines,Japan Airlines
JL,JAL,J-BIRD,Japan_Airlines_Domestic,Japan Airlines Domestic
EG,JAA,ASIA,Japan_Asia_Airways,Japan Asia Airways
NU,JTA,JAI OCEAN,Japan_Transocean_Air,Japan Transocean Air
JA,JAT,ROCKSMART,Jat_Airways,JetSMART
VJ,JTY,JATAYU,Jatayu_Airlines,Jatayu Airlines
J9,JZR,JAZEERA,Jazeera_Airways,Jazeera Airways
7C,JJA,JEJU AIR,Jeju_Air,Jeju Air
,JNY,JENAIR,Jenney_Beechcraft,Jenney Beechcraft
,XLD,,Jeppesen_Data_Plan,Jeppesen Data Plan
,JPN,JETPLAN,Jeppesen_UK,Jeppesen UK
O2,JEA,JETA,OLT_Express,OLT Express
,JSI,SISTEMA,Jet_Air_Group,Jet Air Group
9W,JAI,JET AIRWAYS,Jet_Airways,Jet Airways
QJ,,,Jet_Airways_(USA),Jet Airways
,JTX,JET ASPEN,Jet_Aspen_Air_Lines,Jet Aspen Air Lines
PP,PJS,JETAVIATION,Jet_Aviation,Jet Aviation
,BZF,BIZFLEET,Jet_Aviation_Business_Jets,Jet Aviation Business Jets
,JCF,JET CENTER,Jet_Center_Flight_Training,Jet Center Flight Training
,JCT,JET CHARTER,Jet_Charter,Jet Charter
,JCX,JET CONNECT,Jet_Connection,Jet Connection
,DWW,DON JUAN,Jet_Courier_Service,Jet Courier Service
,JED,JET EAST,Jet_East_International,Jet East International
,JEI,JET EXECUTIVE,Jet_Executive_International_Charter,Jet Executive International Charter
,RZA,RAZOR,Jet_Fighter_Flights,Jet Fighter Flights
,CFT,CASPER FREIGHT,Jet_Freighters,Jet Freighters
,JGD,JET GEE-AND-DEE,Jet_G&D_Aviation,Jet G&D Aviation
,MJL,MOLDJET,Jet_Line_International,Jet Line International
,JEK,JET OPS,Jet_Link,Jet Link
,HTL,HEARTLAND,Jet_Link_Aviation,Jet Link Aviation
,JNR,JET NORTE,Jet_Norte,Jet Norte
,JRN,JET RENT,Jet_Rent,Jet Rent
,JDI,JEDI,Jet_Story,Jet Story
3K,JSA,JETSTAR ASIA,Jetstar_Asia_Airways,Jetstar Asia Airways
,JSM,JET STREAM,Jet_Stream,Jet Stream
,JTC,JETRANS,Jet_Trans_Aviation,Jet Trans Aviation
,JTT,MOSCOW JET,Jet-2000,Jet-2000
,OPS,OPS-JET,Jet-Ops,Jet-Ops
LS,EXS,CHANNEX,Jet2.com,Jet2.com
,JFU,ARGAN,Jet4You,Jet4You
,OSW,BEVO,JetAfrica_Swaziland,JetAfrica Swaziland
B6,JBU,JETBLUE,JetBlue_Airways,JetBlue Airways
,JMG,JET MAGIC,JetMagic,JetMagic
,JMK,,Jetmagic,Jetmagic
JF,JAA,JET ASIA,Jet_Asia_Airways,Jet Asia Airways
,JAF,BEAUTY,Jetairfly,Jetairfly
,JTL,FIREFLY,Jetall_Holdings,Jetall Holdings
,JAG,JETALLIANCE,Jetalliance,Jetalliance
0J,JCS,JETCLUB,Jetclub,Jetclub
,QNZ,QANTAS JETCONNECT,Jetconnect,Jetconnect
,UEJ,JETCORP,Jetcorp,Jetcorp
,JCC,JETCRAFT,Jetcraft_Aviation,Jetcraft Aviation
,JXA,,Jetex_Aviation,Jetex Aviation
,JEF,JETFLITE,Jetflite,Jetflite
,JFL,LINEFLYER,Jetfly_Airlines,Jetfly Airlines
,JIC,JIC-JET,Jetgo_International,Jetgo International
,JLX,KEN JET,Jetlink_Express,Jetlink Express
,JLH,,Jetlink_Holland,Jetlink Holland
,JNV,JETNOVA,Jetnova_de_AviaciÃ³n_Ejecutiva,Jetnova de Aviación Ejecutiva
,JPO,JETPRO,Jetpro,Jetpro
,MDJ,JETRAN AIR,Jetran_Air,Jetran Air
,JRI,JETRIDER,Jetrider_International,Jetrider International
,JEJ,MEXJETS,Jets_Ejecutivos,Jets Ejecutivos
,JEP,JET PERSONALES,Jets_Personales,Jets Personales
,JSE,SERVIJETS,Jets_Y_Servicios_Ejecutivos,Jets Y Servicios Ejecutivos
SG,JGO,JETSGO,JetsGo,JetsGo
JQ,JST,JETSTAR,Jetstar_Airways,Jetstar Airways
GK,JJP,ORANGE LINER,JetStar_Japan,JetStar Japan
JM,JKT,KAITAK,Jetstar_Hong_Kong_Airways,Jetstar Hong Kong Airways
,JXT,VANNIN,Jetstream_Executive_Travel,Jetstream Executive Travel
,RSP,REDSTRIPE,JetSuite,JetSuite
,JPQ,JETT PAQUETERIA,Jett_Paqueteria,Jett Paqueteria
JX,JEC,TAIPAN,Jett8_Airlines_Cargo,Jett8 Airlines Cargo
JO,JTG,JETTIME,Jettime,Jettime
,JTN,JETTRAIN,Jettrain_Corporation,Jettrain Corporation
,JWY,JETWAYS,Jetways_of_Iowa,Jetways of Iowa
GX,JXX,JETBIRD,JetX_Airlines,JetX Airlines
,JIB,JIBAIRLINE,Jibair,Jibair
,JSW,JIGSAW,,Jigsaw Project
,HKN,HANKINS,Jim_Hankins_Air_Service,Jim Hankins Air Service
,RAS,SHANHIL,Jim_Ratliff_Air_Service,Jim Ratliff Air Service
,JDG,LADYBLUE,Joanas_Avialinijos,Joanas Avialinijos
,JBR,JOBAIR,Job_Air,Job Air
,JHN,AIR JOHNSON,Johnson_Air,Johnson Air
,JON,JOHNSONSAIR,Johnsons_Air,Johnsons Air
,JMJ,JOHNSTON,Johnston_Airways,Johnston Airways
,JMM,JOICOMAR,Joint_Military_Commission,Joint Military Commission
,JMT,JOMARTAXI,Jomartaxi_Aereo,Jomartaxi Aereo
,ODI,ODINN,"Jonsson,_H_Air_Taxi","Jonsson, H Air Taxi"
R5,JAV,JORDAN AVIATION,Jordan_Aviation,Jordan Aviation
J4,JCI,,Jordan_International_Air_Cargo,Jordan International Air Cargo
,JVK,ISLANDIC,Jorvik_(airline),Jorvik
,ENZ,ENZO,Jota_Aviation,Jota Aviation
,JNJ,JOURNEY JET,Journey_Jet,Journey Jet
,JUR,JUNKERS,Ju-Air,Ju-Air
,JFS,JAEMCO,Juanda_Flying_School,Juanda Flying School
,JUC,JUBA CARGO,Juba_Cargo_Services_&_Aviation_Company,Juba Cargo Services & Aviation Company
6J,JUB,JUBBA,Jubba_Airways,Jubba Airways
,DKE,DUKE,Jubilee_Airways,Jubilee Airways
HO,DKH,JUNEYAO AIRLINES,Juneyao_Airlines,Juneyao Airlines
,MEY,MELODY,Justair_Scandinavia,Justair Scandinavia
,DOJ,JUSTICE,Justice_Prisoner_and_Alien_Transportation_System,Justice Prisoner and Alien Transportation System
,JNL,JETNETHERLANDS,JetNetherlands,JetNetherlands
,JFA,MOSQUITO,Jetfly_Aviation,Jetfly Aviation
3Q,KCH,CAM AIR,,KC International Airlines
,KCR,KOLOB,Kolob_Canyons_Air_Services,Kolob Canyons Air Services
KW,KHK,SUNRAY,Kharkiv_Airlines,Kharkiv Airlines
,KGZ,BERMET,Kyrgyz_Airlines,Kyrgyz Airlines
,KDC,KAY DEE,K_D_Air_Corporation,K D Air Corporation
,KSA,SKY CAMEL,K_S_Avia,K S Avia
,KMI,KAY-MILE AIR,K-Mile_Air,K-Mile Air
KD,KLS,KALSTAR,Kalstar_Aviation,Kalstar Aviation
KD,KNI,KALININGRAD AIR,KD_Avia,KD Avia
WA,KLC,CITY,KLM_Cityhopper,KLM Cityhopper
,KLH,KLM HELI,KLM_Helicopter,KLM Helicopter
KL,KLM,KLM,KLM,KLM
N2,QNK,KABO,Kabo_Air,Kabo Air
,KMC,KAHAMA,Kahama_Mining_Corporation,Kahama Mining Corporation
,KAI,KAISER,Kaiser_Air,Kaiser Air
K4,CKS,CONNIE,Kalitta_Air,Kalitta Air
K9,KFS,KALITTA,Kalitta_Charters,Kalitta Charters
CB*,KFS,KALITTA,Kalitta_Charters_II,Kalitta Charters II
,KES,KALLAT EL SKER,Kallat_El_Saker_Air_Company,Kallat El Saker Air Company
RQ,KMF,KAMGAR,Kam_Air,Kam Air
E2,KMP,KAMPUCHEA,Kampuchea_Airlines,Kampuchea Airlines
,KIZ,,Arkia_Israel_Airlines,Kanaf-Arkia Airlines
,KHE,KANFEY HAEMEK,Kanfey_Ha'emek_Aviation,Kanfey Ha'emek Aviation
,KSU,K-STATE,Kansas_State_University,Kansas State University
V2,AKT,AVIAKARAT,Karat_(airline),Karat
,KRB,KARIBU AIR,Karibu_Airways_Company,Karibu Airways Company
,KLG,KARLOG,Karlog_Air_Charter,Karlog Air Charter
,KAJ,KARTHAGO,Karthago_Airlines,Karthago Airlines
,KAE,KARTIKA,Kartika_Airlines,Kartika Airlines
,KTV,KATAVIA,Kata_Transportation,Kata Transportation
,KTK,KATEKAVIA,Katekavia,Katekavia
,KAT,KATO-AIR,Kato_Airline,Kato Airline
KV,MVD,AIR MINVODY,Kavminvodyavia,Kavminvodyavia
,XKA,,Kavouras_Inc,Kavouras Inc
,KRN,ANTOL,Kaz_Agros_Avia,Kaz Agros Avia
,KAW,KAZWEST,Kaz_Air_West,Kaz Air West
,KAO,KAZAVAIA,Kazan_Aviation_Production_Association,Kazan Aviation Production Association
,KPH,KAMA,Kazan_Helicopters,Kazan Helicopters
,KKA,KAKAIR,Kazavia,Kazavia
,KZS,SPAKAZ,Kazaviaspas,Kazaviaspas
,JFK,KEENAIR,Keenair_Charter,Keenair Charter -
,KLX,KELIX,Kelix_Air,Kelix Air
,FKL,KELNER,Kelner_Airways,Kelner Airways
,KFA,FLIGHTCRAFT,Kelowna_Flightcraft_Air_Charter,Kelowna Flightcraft Air Charter
,KDA,KENDELL,Kendell_Airlines,Kendell Airlines
M5,KEN,KENMORE,Kenmore_Air,Kenmore Air
,KBA,BOREK AIR,Kenn_Borek_Air,Kenn Borek Air
,KAH,DEKAIR,Kent_Aviation,Kent Aviation
KQ,KQA,KENYA,Kenya_Airways,Kenya Airways
,KVS,KEVIS,Kevis,Kevis
,KEY,KEY AIR,Key_Airlines,Key Airlines
,LYM,KEY LIME,Key_Lime_Air,Key Lime Air
,FTP,FOOTPRINT,Keystone_Aerial_Surveys,Keystone Aerial Surveys
BZ,KEE,KEYSTONE,Keystone_Air_Service,Keystone Air Service
K6,KZW,KHALIFA AIR,Khalifa_Airways,Khalifa Airways
,WKH,WEST-KHARKOV,Kharkov_Aircraft_Manufacturing_Company,Kharkov Aircraft Manufacturing Company
,KHR,KHAZAR,Khazar,Khazar
,KHP,PHOTROS AIR,Khoezestan_Photros_Air_Lines,Khoezestan Photros Air Lines
,KRV,KHORIV-AVIA,Khoriv-Avia,Khoriv-Avia
,KHO,AIRCOMPANY KHORS,Khors_Aircompany,Khors Aircompany
,KHY,KHYBER,Khyber_Afghan_Airlines,Khyber Afghan Airlines
,UAK,AVIATION PLANT,Kiev_Aviation_Plant,Kiev Aviation Plant
,KNG,KING,King_Aviation,King Aviation
,BEZ,SEA BREEZE,Kingfisher_Air_Services,Kingfisher Air Services
IT,KFR,KINGFISHER,Kingfisher_Airlines,Kingfisher Airlines
,KNX,KNIGHT FLIGHT,Knighthawk_Air_Express,Knighthawk Air Express
,KAS,KINGSTON AIR,Kingston_Air_Services,Kingston Air Services
,KIP,KINNARPS,Kinnarps,Kinnarps
,KNS,KINSHASA AIRWAYS,Kinshasa_Airways,Kinshasa Airways
,KTA,VYATKA-AVIA,Kirov_Air_Enterprise,Kirov Air Enterprise
Y9,IRK,KISHAIR,Kish_Air,Kish Air
,KHA,AIR KITTYHAWK,Kitty_Hawk_Aircargo,Kitty Hawk Aircargo
,KHC,CARGO HAWK,Kitty_Hawk_Airways,Kitty Hawk Airways
KP,KIA,KIWI AIR,Kiwi_International_Air_Lines,Kiwi International Air Lines
,KRA,REGIONAL,Kiwi_Regional_Airlines,Kiwi Regional Airlines
,KNA,KNIGHTAIR,Knight_Air,Knight Air
,KHX,RIZZ,Knighthawk_Express,Knighthawk Express
,KGT,KNIGHT-LINER,Knights_Airlines,Knights Airlines
,KOA,KOANDA,Koanda_Avacion,Koanda Avacion
,OYE,KODA AIR,Koda_International,Koda International
7K,KGL,KOGALYM,Kogalymavia_Air_Company,Kogalymavia Air Company
,KOM,COMJET,Kom_Activity,Kom Activity
,KMA,KOMI AVIA,Komiaviatrans_State_Air_Enterprise,Komiaviatrans State Air Enterprise
8J,KMV,KOMIINTER,Komiinteravia,Komiinteravia
,KNM,KNAAPO,Komsomolsk-on-Amur_Air_Enterprise,Komsomolsk-on-Amur Air Enterprise
,KOB,AUTOFLEX,Koob-Corp_-_96_KFT,Koob-Corp - 96 KFT
KE,KAL,KOREANAIR,Korean_Air,Korean Air
,KMG,KOSMAS CARGO,Kosmas_Air,Kosmas Air
,KSM,KOSMOS,Kosmos_Airlines,Kosmos
,KOS,KOSOVA,Kosova_Airlines,Kosova Airlines
,WOK,WOKAIR,Kovar_Air,Kovar Air
7B,KJC,KRASNOJARSKY AIR,Krasnojarsky_Airlines,Krasnojarsky Airlines
,KFC,KREMENCHUK,Kremenchuk_Flight_College,Kremenchuk Flight College
,KRG,AVIAMONTAG,Krimaviamontag,Krimaviamontag
,KRO,KROONK,Kroonk_Air_Agency,Kroonk Air Agency
K9,KRI,Krylo,Krylo_Airlines,Krylo Airlines
,KYM,CRIMEA AIR,Krym_(airline),Krym
,OPC,OPTIC,Krystel_Air_Charter,Krystel Air Charter
GW,KIL,AIR KUBAN,Kuban_Airlines,Kuban Airlines
VD,KPA,KUNPENG,Kunpeng_Airlines,Kunpeng Airlines
,KZA,,Kurzemes_Avio,Kurzemes Avio
,KBV,SWECOAST,Kustbevakningen,Kustbevakningen
KU,KAC,KUWAITI,Kuwait_Airways,Kuwait Airways
GO,KZU,KUZU CARGO,Kuzu_Airlines_Cargo,Kuzu Airlines Cargo
,QVR,PEGASO,Kvadro_Aero,Kvadro Aero
,KWN,KWENA,Kwena_Air,Kwena Air
N5,KGZ,BERMET,Kyrgyz_Airlines,Kyrgyz Airlines
,KTC,DINARA,Kyrgyz_Trans_Avia,Kyrgyz Trans Avia
QH,LYN,ALTYN AVIA,Kyrgyzstan_(airline),Kyrgyzstan
R8,KGA,KYRGYZ,Kyrgyzstan_Airlines,Kyrgyzstan Airlines
,DAM,FLIGHT RESCUE,Kyrgyzstan_Department_of_Aviation,Kyrgyzstan Department of Aviation
,KGB,KEMIN,Kyrgz_General_Aviation,Kyrgz General Aviation
FK,KEW,BLIZZARD,Keewatin_Air,Keewatin Air
,AOE,LIVINGSTONE AIR,Livingstone_Executive,Livingstone Executive
,LZF,SKYLEASE,Lease_Fly,Lease Fly
,LBR,,Lepl_Project_Limited,Lepl Project Limited (Air Costa)
,LHB,FAMILY,Liebherr_Geschaeftreiseflugzeug,Liebherr Geschaeftreiseflugzeug
,LLD,,Lloyd_Aviation,Lloyd Aviation
,LGA,LOGAIR,Logistic_Air,Logistic Air
,LGC,LEGACY AIR,Legacy_Air,Legacy Air
,JKA,JACKET,LeTourneau_University,LeTourneau University
,LTY,SKYDECK,Liberty_Air,Liberty Air
,LWL,CUB DRIVER,Lowlevel,Lowlevel
,LAX,,Laminar_Air,Laminar Air
,LWA,LIBYAN WINGS,Libyan_Wings,Libyan Wings
YQ,LCT,TAR,TAR_Aerolineas,TAR Aerolineas
,LYB,HIGHLANDS,Lynden_Air_Cargo_(PNG),Lynden Air Cargo
,LAH,STAR SHIP,L_A_Helicopter,L A Helicopter
,LJY,ELJAY,L_J_Aviation,L J Aviation
,LRB,LADY RACINE,L_R_Airlines,L R Airlines
,PHO,PHOTOFLIGHT,L&L_Flygbildteknik,L&L Flygbildteknik
,LEX,LEX,L%27Express_Airlines,L'Express Airlines
,FNT,FLIGHT INTERNATIONAL,L-3_Communications_Flight_International_Aviation,L-3 Communications Flight International Aviation
JF,LAB,LAB,L.A.B._Flying_Service,L.A.B. Flying Service
LR,LRC,LACSA,LACSA,LACSA
,LDE,LADE,LADE_-_L%C3%ADneas_A%C3%A9reas_Del_Estado,LADE - Líneas Aéreas Del Estado
KG,BNX,AIR BARINAS,LAI_-_L%C3%ADnea_A%C3%A9rea_IAACA,LAI - Línea Aérea IAACA
LA,LAN,LAN CHILE,LATAM_Chile,LATAM Chile
4M,DSM,LAN AR,LATAM_Argentina,LATAM Argentina
UC,LCO,LAN CARGO,LATAM_Cargo_Chile,LATAM Cargo Chile
,ARE,LAN COLOMBIA,LATAM_Colombia,LATAM Colombia
PZ,LAP,PARAGUAYA,LATAM_Paraguay,LATAM Paraguay
,LNC,LANCANA,LAN_Dominicana,LAN Dominicana
LU,LXP,LANEX,LATAM_Express,LATAM Express
LP,LPE,LANPERU,LATAM_Peru,LATAM Peru
NI,,,LANICA,LANICA
,LSA,INTERNACIONAL,L%C3%ADneas_A%C3%A9reas_Nacionales_S._A._(Peru),LANSA
,APT,LAP,"LAP_Colombia_-_LÃ­neas_AÃ©reas_Petroleras,_S.A.","LAP Colombia - Líneas Aéreas Petroleras, S.A."
,OTN,LASTP,LASTP,LASTP
,LCB,BUSRE,LC_Busre,LC Busre
L5,,,L%C3%ADnea_A%C3%A9rea_Cuencana,Línea Aérea Cuencana
LO,LOT,POLLOT,LOT_Polish_Airlines,LOT Polish Airlines
,JKA,JACKET,LeTourneau_University,LeTourneau University
XO,LTE,FUN JET,LTE_International_Airways,LTE International Airways
L3,LTO,BILLA TRANSPORT,LTU_Austria,LTU Austria
LT,LTU,LTU,LTU_International,LTU International
,JFC,JET-FLEET,LTV_Jet_Fleet_Corporation,LTV Jet Fleet Corporation
,LUK,LUKOIL,LUKoil-Avia,LUKoil-Avia
,ASK,AIR SASK,La_Ronge_Aviation_Services,La Ronge Aviation Services
,LVT,TAXIVALENCIANA,La_Valenciana_Taxi_AÃ©reo,La Valenciana Taxi Aéreo
,SKQ,SKYLAB,Labcorp,Labcorp
,LAL,LAB AIR,Labrador_Airways,Labrador Airways
N6,JEV,,Lagun_Air,Lagun Air
,HCA,HAVASU,Lake_Havasu_Air_Service,Lake Havasu Air Service
,LKL,LAKELAND,Lakeland_Aviation,Lakeland Aviation
,LKR,LAKER,Laker_Airways,Laker Airways
,LBH,LAKER BAHAMAS,Laker_Airways_(Bahamas),Laker Airways (Bahamas)
,LMR,LAMAIR,Lamra,Lamra
,TCR,TICOS,Lanaes_Aereas_Trans_Costa_Rica,Lanaes Aereas Trans Costa Rica
,ISL,ISLANDIA,Landsflug,Landsflug
,PAP,PROFLIGHT,Langtry_Flying_Group,Langtry Flying Group
IK,LKN,Lankair,Lankair,Lankair
QL,RLN,AERO LANKA,Lankan_Cargo,Lankan Cargo
,LZA,AEROLANZA,Lanza_Air,Lanza Air
,LZT,BARAKA,Lanzarote_Aerocargo,Lanzarote Aerocargo
,LLL,LAVIE,Lao_Skyway,Lao Skyway
QV,LAO,LAO,Lao_Airlines,Lao Airlines
,LKA,NAKLAO,Lao_Central_Airlines,Lao Capricorn Air
L7,LPN,LAOAG AIR,Laoag_International_Airlines,Laoag International Airlines
,LRD,LAREDO AIR,Laredo_Air,Laredo Air
,LTC,LATCHARTER,LatCharter,LatCharter
,LAF,LATVIAN AIRFORCE,Latvian_Air_Force,Latvian Air Force
NG,LDA,LAUDA AIR,Lauda_Air,Lauda Air
OE[28],LDM,LAUDA MOTION,LaudaMotion,LaudaMotion
,LDI,LAUDA ITALY,Lauda_Air_Italy,Lauda Air Italy
,LEP,LAUGHLIN EXPRESS,Laughlin_Express,Laughlin Express
,LSU,LAUS AIR,Laus_(airline),Laus
,LAR,LAWRENCE,Lawrence_Aviation,Lawrence Aviation
,LAY,LAYANG,Layang-Layang_Aerospace,Layang-Layang Aerospace
,LPL,LEASE-A-PLANE,Lease-a-Plane_International,Lease-a-Plane International
LQ,LAQ,LAT,Lebanese_Air_Transport,Lebanese Air Transport
,LAT,LEBANESE AIR,Lebanese_Air_Transport,Lebanese Air Transport (charter)
,LAD,LADCO-AIR,Lebanon_Airport_Development_Corporation,Lebanon Airport Development Corporation
,LEB,LEBAP,Lebap,Lebap
,LCA,LECONTE,Leconte_Airlines,Leconte Airlines
LI,LIA,LIAT,Leeward_Islands_Air_Transport,Leeward Islands Air Transport
,LGD,LEGENDARY,Legend_Airlines,Legend Airlines
,LWD,LEISURE WORLD,Leisure_Air,Leisure Air
,LEN,LENTINI,Lentini_Aviation,Lentini Aviation
,LOR,LEO CHARTER,Leo-Air,Leo-Air
,LEL,LEONAVIA,Leonsa_De_AviaciÃ³n,Leonsa De Aviación
LV,LVL,LEVEL,Level_(airline),Level
,LYW,LIBYAN AIRWAYS,Libyan_Airlines,Libyan Airlines
LN,LAA,LIBAIR,Libyan_Arab_Airlines,Libyan Arab Airlines
,LCR,LIBAC,Libyan_Arab_Air_Cargo,Libyan Arab Air Cargo
,LIQ,,Lid_Air,Lid Air
,LCG,CONGOLAISE,Lignes_Aeriennes_Congolaises,Lignes Aeriennes Congolaises
,LKD,LATCHAD,Lignes_Aeriennes_Du_Tchad,Lignes Aeriennes Du Tchad
,LME,LIMAIR EXPRESS,Lignes_Mauritaniennes_Air_Express,Lignes Mauritaniennes Air Express
,GCB,LINACONGO,Linacongo,Lignes Nationales Aeriennes - Linacongo
,GDQ,,Lincoln_Air_National_Guard,Lincoln Air National Guard
,LRT,,Lincoln_Airlines,Lincoln Airlines
,LSY,LINDSAY AIR,Lindsay_Aviation,Lindsay Aviation
,NOT,COSTA NORTE,LÃ­nea_AÃ©rea_Costa_Norte,Línea Aérea Costa Norte
,LMC,LINEAS DECARGA,LÃ­nea_AÃ©rea_Mexicana_de_Carga,Línea Aérea Mexicana de Carga
L7,LNP,SAPSA,LÃ­nea_AÃ©rea_SAPSA,Línea Aérea SAPSA
,NEG,AGUAS NEGRAS,LÃ­nea_AÃ©rea_de_Fumig_Aguas_Negras,Línea Aérea de Fumig Aguas Negras
QL,LER,LASER,L%C3%ADnea_A%C3%A9rea_de_Servicio_Ejecutivo_Regional,Línea Aérea de Servicio Ejecutivo Regional
,LSE,,LÃ­nea_De_Aeroservicios,Línea De Aeroservicios
,TUY,AEREOTUY,L%C3%ADnea_Tur%C3%ADstica_Aereotuy,Línea Turística Aereotuy
,ALR,AEROLAIRE,LÃ­neas_AÃ©reas_Alaire_S.L.,Líneas Aéreas Alaire S.L.
ZE,LCD,LINEAS AZTECA,L%C3%ADneas_A%C3%A9reas_Azteca,Líneas Aéreas Azteca
,LCN,CANEDO,LÃ­neas_AÃ©reas_Canedo,Líneas Aéreas Canedo LAC
,LCM,LINEAS COMERCIALES,LÃ­neas_AÃ©reas_Comerciales,Líneas Aéreas Comerciales
,EDD,LINEAS DURANGO,LÃ­neas_AÃ©reas_Ejectuivas_De_Durango,Líneas Aéreas Ejectuivas De Durango
,EDR,ELDORADRO,LÃ­neas_AÃ©reas_Eldorado,Líneas Aéreas Eldorado
,FED,FEDERALES,L%C3%ADneas_A%C3%A9reas_Federales,Líneas Aéreas Federales
,LMN,LINEAS MONARCA,LÃ­neas_AÃ©reas_Monarca,Líneas Aéreas Monarca
,LIJ,LINEAS JOSE,LÃ­neas_AÃ©reas_San_Jose,Líneas Aéreas San Jose
,UMA,HUMAYA,LÃ­neas_AÃ©reas_del_Humaya,Líneas Aéreas del Humaya
,LEC,LECA,Linex,Linex
,SMS,SANTOMENSES,Linhas_AÃ©reas_Santomenses,Linhas Aéreas Santomenses
TM,LAM,MOZAMBIQUE,Linhas_A%C3%A9reas_de_Mo%C3%A7ambique,Linhas Aéreas de Moçambique
,LAW,,Link_Airways_of_Australia,Link Airways of Australia
,WGT,WORLDGATE,Lion_Air_Services,Lion Air Services
JT,LNI,LION INTER,Lion_Air,Lion Mentari Airlines
,LEU,LIONSAIR,Lions-Air,Lions-Air
,LYF,LITHUANIAN AIRFORCE,Lithuanian_Air_Force,Lithuanian Air Force
,LRA,LITTLE RED,Little_Red_Air_Service,Little Red Air Service
LM,LVG,LIVINGSTON,Livingston_Energy_Flight,Livingston Energy Flight
LB,LLB,LLOYDAEREO,Lloyd_A%C3%A9reo_Boliviano,Lloyd Aéreo Boliviano
,LNA,ELNAIR,Lnair_Air_Services,Lnair Air Services
,XLG,,Lockheed_Air_Terminal,Lockheed Air Terminal
,LAC,LOCKHEED,Lockheed_Corporation,Lockeed Aircraft Corporation
,XDD,,Lockheed_DUATS,Lockheed DUATS
,CBD,CATBIRD,Lockheed_Martin_Aeronautics,Lockheed Martin Aeronautics
,LNG,LIGHTNING,Lockheed_Martin_Aeronautics_Company,Lockheed Martin Aeronautics Company
LC,LOG,LOGAN,Loganair,Loganair
,CLV,AEROTRAINING,Lom_Praha_Flying_School,Lom Praha Flying School
,LMS,LOMAS,Lomas_Helicopters,Lomas Helicopters
,LCY,LONDON CITY,London_City_Airport_Jet_Centre,London City Airport Jet Centre
,LNX,LONEX,London_Executive_Aviation,London Executive Aviation
,LOV,LOVEAIR,London_Flight_Centre_(Stansted),London Flight Centre (Stansted)
,LHC,MUSTANG,London_Helicopter_Centres,London Helicopter Centres
,LSS,LONE STAR,Lone_Star_Airlines,Lone Star Airlines
,ORA,LONG ISLAND,Long_Island_Airlines,Long Island Airlines
,LGT,LONGTAIL,Longtail_Aviation,Longtail Aviation
,LRR,LORRAINE,Lorraine_Aviation,Lorraine Aviation
,LSC,CEDROS,Los_Cedros_AviaciÃ³n,Los Cedros Aviación
,TAS,LOTUS FLOWER,Lotus_Air,Lotus Air
,LTW,TWENTAIR,Luchtvaartmaatschappij_Twente,Luchtvaartmaatschappij Twente
,LKE,LUCKY AIR,Lucky_Air,Lucky Air
,LUT,LUGO,Luft_Carago,Luft Carago
,LVD,AIR SANTE,Luftfahrt-Vermietungs-Dienst,Luftfahrt-Vermietungs-Dienst
HE,LGW,WALTER,Luftfahrtgesellschaft_Walter,Luftfahrtgesellschaft Walter
LH,DLH,LUFTHANSA,Lufthansa,Lufthansa
LH,GEC,LUFTHANSA CARGO,Lufthansa_Cargo,Lufthansa Cargo
CL,CLH,HANSALINE,Lufthansa_CityLine,Lufthansa CityLine
L1,,,Lufthansa_Systems,Lufthansa Systems
,LHT,LUFTHANSA TECHNIK,Lufthansa,Lufthansa Technik
DV,LTF,Garfield,,Lufttaxi Fluggesellschaft
L5,LTR,LUFT TRANSPORT,Lufttransport,Lufttransport
,LHS,ENTERPRISE LUHANSK,Luhansk,Luhansk
,UNY,UNIVERSITY,Lund_University_School_of_Aviation,Lund University School of Aviation
LG,LGL,LUXAIR,Luxair,Luxair
,LXA,RED LION,,Luxaviation
,LUV,LUX RESCUE,Luxembourg_Air_Rescue,Luxembourg Air Rescue
,LFE,LUX EXPRESS,Luxflight_Executive,Luxflight Executive
,LXO,,Luxor_Air,Luxor Air
,LUZ,LISBON JET,Luzair,Luzair
5V,UKW,UKRAINE WEST,Lviv_Airlines,Lviv Airlines
,LYD,LYDDAIR,Lydd_Air,Lydd Air
,LCH,LYNCH AIR,Lynch_Flying_Service,Lynch Flying Service
L2,LYC,LYNDEN,Lynden_Air_Cargo,Lynden Air Cargo
,LXF,LYNX FLIGHT,Lynx_Air_International,Lynx Air International
,LYX,LYNX AIR,Lynx_Aviation_(Pakistan),Lynx Aviation
L4,SSX,SHASTA,Lynx_Aviation_(United_States),Lynx Aviation
Z8,AZN,,L%C3%ADnea_A%C3%A9rea_Amaszonas,Línea Aérea Amaszonas
MJ,LPR,LAPA,L%C3%ADneas_A%C3%A9reas_Privadas_Argentinas,Líneas Aéreas Privadas Argentinas
,LAU,SURAMERICANO,L%C3%ADneas_A%C3%A9reas_Suramericanas,Líneas Aéreas Suramericanas
LT,SNG,SNOW EAGLE,LongJiang_Airlines,LongJiang Airlines
LJ,JNA,JIN AIR,Jin_Air,Jin Air
Q2,DQA,SKYSURFER,Maldivian_(airline),Maldivian (airline)
,MMH,NIGHT RIDER,McMahon_Helicopter,McMahon Helicopter
,HOG,HOGAN AIR,Mahogany_Air_Charters,Mahogany Air Charters
,MTS,,Med_Jets,Med Jets
,MSF,MEINSHENG,Minsheng_International_Jet,Minsheng International Jet
,MXS,MILLON EXPRESS,Millon_Express,Millon Express
,MHF,AIR MARITIME,Maritime_Helicopters,Maritime Helicopters
,MWM,MODERNAIR,Modern_Transporte_Aereo_De_Carga,Modern Transporte Aereo De Carga
,MSJ,MAGNUM AIR,Magnum_Air,Magnum Air
,MWI,MALAWIAN,Malawian_Airlines_2014,Malawian Airlines 2014
,MYP,MANN ROYAL,Mann_Yadanarpon_Airlines,Mann Yadanarpon Airlines
,RDK,RED DUKE,Memorial_Hermann_Hospital_System,Memorial Hermann Hospital System
,MLV,MULTI VALLE,Multiservicios_Aereos_Del_Valle,Multiservicios Aereos Del Valle
,MMJ,MACAUJET,Macau_Jet_International,Macau Jet International
,MXF,MAXFLIGHT,Maximum_Flight_Advantages,Maximum Flight Advantages
OD,MXD,MALINDO EXPRESS,Malindo_Airways,Malindo Airways
,MJC,AIR MANDA,Mandarin_Air,Mandarin Air
,PLG,PILGRIM,MBK-S,MBK-S
,DZR,DOZER,Midwest_Aviation_(Omaha),Midwest Aviation
,MSN,,Milenium_Air_Servicios_Aereos_Integrados,Milenium Air Servicios Aereos Integrados
,MFB,MOUNTAINHELI,Mountain_Flyers_80,Mountain Flyers 80
,HTL,HOTLINE,My_Fair_Jet,My Fair Jet
,JNH,JONAH,M_&_N_Aviation,M & N Aviation
,MCF,MAC FOTO,MAC_Fotografica,MAC Fotografica
,MRG,MANAG'AIR,MANAG'AIR,MANAG'AIR
AQ,MPJ,MAPJET,MAP-Management_and_Planung,MAP-Management and Planung
,TFG,TRAFALGAR,MAS_Airways,MAS Airways
M7,MAA,MAS CARGA,MasAir,MasAir
MH,MWG,MASWINGS,MASwings,MASwings
IN,MAK,MAKAVIO,MAT_Macedonian_Airlines,MAT Macedonian Airlines
,MCC,DISCOVERY,,MCC Aviation
,MGA,MAG AVACION,MG_AviaciÃ³n,MG Aviación
,JLA,SALLINE,MIA_Airlines,MIA Airlines
OM,MGL,MONGOL AIR,MIAT_Mongolian_Airlines,MIAT Mongolian Airlines
,MNC,MUNCIE,MIT_Airlines,MIT Airlines
,MKA,KRUGER-AIR,MK_Airline,MK Airline
MB,MNB,BLACK SEA,MNG_Airlines,MNG Airlines
,EBF,SKYRUNNER,MSR_Flug-Charter,MSR Flug-Charter
,MCV,MTC AVIACION,MTC_AviaciÃ³n,MTC Aviación
,MAQ,MAC AVIATION,Mac_Aviation,Mac Aviation
,MCN,MAC DAN,Mac_Dan_Aviation_Corporation,Mac Dan Aviation Corporation
,MTD,,MacKnight_Airlines,MacKnight Airlines
CC,MCK,,Macair_Airlines,Macair Airlines
,MCS,MACAIR,Macedonian_Airlines_(OA),Macedonian Airlines
,MDH,MADINA AIR,Madina_Air,Madina Air
DM,,,Maersk,Maersk
,MJB,MAGIC BLUE,Magic_Blue_Airlines,Magic Blue Airlines
,MGR,MAGNA AIR,Magna_Air,Magna Air
,MLH,MAHALO,Mahalo_Air,Mahalo Air
W5,IRM,MAHAN AIR,Mahan_Air,Mahan Air
M2,MZS,MAHFOOZ,Mahfooz_Aviation,Mahfooz Aviation
,MAT,MAINE-AV,Maine_Aviation,Maine Aviation
,MAJ,MAGIC AIR,Majestic_Airlines,Majestic Airlines
,AKM,MAKAIR,Mak_Air,Mak Air
,MLG,,Malagasy_Airlines,Malagasy Airlines
,MLX,MALAWI EXPRESS,Malawi_Express,Malawi Express
,MKK,AEROKEY,Malaya_Aviatsia_Dona,Malaya Aviatsia Dona
MH,MAS,MALAYSIAN,Malaysia_Airlines,Malaysia Airlines
,MAE,MALI AIREXPRESS,Mali_Air,Mali Air
,VXP,AVION EXPRESS,Mali_Air_Express,Mali Air Express
,MTZ,MALI AIRWAYS,Mali_Airways,Mali Airways
,MLC,MALILA,Malila_Airlift,Malila Airlift
,MLS,MALL-AIRWAYS,Mall_Airways,Mall Airways
,LOD,LOGIC,Malmoe_Air_Taxi,Malmoe Air Taxi
TF,SCW,SCANWING,Malm%C3%B6_Aviation,Malmö Aviation
R5,MAC,MALTA CHARTER,Malta_Air_Charter,Malta Air Charter
,MWS,MALTA WINGS,Malta_Wings,Malta Wings
MA,MAH,MALEV,Mal%C3%A9v_Hungarian_Airlines,Malév Hungarian Airlines
,MLB,MANAF,Manaf_International_Airways,Manaf International Airways
RI,MDL,MANDALA,Mandala_Airlines,Mandala Airlines
AE,MDA,MANDARIN,Mandarin_Airlines,Mandarin Airlines
JE,MNO,TULCA,Mango_(airline),Mango
,MHN,MANHATTAN,Manhattan_Air,Manhattan Air
,MTO,MANITOULIN,Manitoulin_Air_Services,Manitoulin Air Services
,MNR,TEEMOL,Mann_Air,Mann Air
,MAN,MANNION,Mannion_Air_Charter,Mannion Air Charter
,MTS,MANTRUST,Mantrust_Asahi_Airways,Mantrust Asahi Airways
,MNX,MANX,Manx_Airlines,Manx Airlines
,MAD,MAPLE AIR,Maple_Air_Services,Maple Air Services
,MAR,MARCH,March_Helicopters,March Helicopters
,MCP,MARCOPOLO,Marcopolo_Airways,Marcopolo Airways
,MGI,MARGHI,Marghi_Air,Marghi Air
,MRK,MARKAIR,Markair,Markair
,MKO,GOSHAWK,Markoss_Aviation,Markoss Aviation
6V,MRW,AVIAMARS,Mars_RK,Mars RK
,MCE,MARSHALL,Marshall_Aerospace,Marshall Aerospace
M7,MSL,MARSLANDAIR,Marsland_Aviation,Marsland Aviation
,XMA,,Martin_Aviation_Services,Martin Aviation Services
,MBE,MARTIN,Martin-Baker,Martin-Baker
MP,MPH,MARTINAIR,Martinair,Martinair
,MRA,MARTEX,Martinaire,Martinaire
,MFA,SEAHORSE,Martyn_Fiddler_Associates,Martyn Fiddler Associates
,MVN,MARVIN,Marvin_Limited,Marvin Limited
,TRP,TROOPER,Maryland_State_Police,Maryland State Police
,MTH,RESEARCH,Massachusetts_Institute_of_Technology,Massachusetts Institute of Technology
,MSY,MASSEY,Massey_University_School_of_Aviation,Massey University School of Aviation
,MSW,MASTER AIRWAYS,Master_Airways,Master Airways
,MPL,,Master_Planner,Master Planner
,LMJ,MASTERJET,Masterjet,Masterjet
Q4,,,Mastertop_Linhas_AÃ©reas,Mastertop Linhas Aéreas
,MIA,MAURIA,Mauria,Mauria
,MNV,NAVALE,Mauritanienne_Aerienne_Et_Navale,Mauritanienne Aerienne Et Navale
,MRF,MAUR-FRET,Mauritanienne_Air_Fret,Mauritanienne Air Fret
,MWY,MAURITANIENNE,Mauritanienne_Airways,Mauritanienne Airways
,MDE,MAURI-TRANS,Mauritanienne_De_Transport_Aerien,Mauritanienne De Transport Aerien
,MVR,MAV-AIR,Maverick_Airways,Maverick Airways
H5,MVL,Mavial,Mavial_Magadan_Airlines,Mavial Magadan Airlines
,MAI,MAX AVIA,Max_Avia,Max Avia
,MSF,MAXESA,Max_Sea_Food,Max Sea Food
,MAX,MAX AVIATION,Max-Aviation,Max-Aviation
8M,MXL,MAXAIR,Maxair_(aviation),Maxair
,MXU,CARGO MAX,Maximus_Air_Cargo,Maximus Air Cargo
MY,MXJ,MAX-JET,Maxjet_Airways,Maxjet Airways
,MXS,MAXSUS-AVIA,Maxsus-Avia,Maxsus-Avia
,MXP,BEECHNUT,May_Air_Xpress,May Air Xpress
MW,MYD,MYLAND,Maya_Island_Air,Maya Island Air
7M,MYI,MAYAIR,Mayair,Mayair
,MBS,MBACHI AIR,Mbach_Air,Mbach Air
,MCH,MACLINE,McAlpine_Helicopters,McAlpine Helicopters
,MKL,MCCALL,McCall_Aviation,McCall Aviation
,DAC,DACO,McDonnell_Douglas,McDonnell Douglas
,MDS,MID-SOUTH,McNeely_Charter_Services,McNeely Charter Services
,MEK,MED-TRANS,Med-Trans_of_Florida,Med-Trans of Florida
,MDM,MEDAVIA,Medavia,Medavia
,MRZ,MARS,Medical_Air_Rescue_Services,Medical Air Rescue Services
,MCL,MEDIC,Medical_Aviation_Services,Medical Aviation Services
,MDF,MED-FREIGHT,Mediterranean_Air_Freight,Mediterranean Air Freight
,MDY,,Mediterranean_Airways,Mediterranean Airways
,MEJ,MEDJET,Medjet_International,Medjet International
,MGK,MEGLA,Mega_Aircompany,Mega
,MEL,MEGA AIR,Mega_Linhas_AÃ©reas,Mega Linhas Aéreas
M8,MKN,MEKONG AIRLINES,Mekong_Airlines,Mekong Airlines
IM,MNJ,MENAJET,Menajet,Menajet
,,,Mercer_Airlines,Mercer Airlines
,MXX,MERCHANT,Merchant_Express_Aviation,Merchant Express Aviation
,MEC,MERCAIR,Mercury_Aircourier_Service,Mercury Aircourier Service
,POV,AIR POLTAVA,Meridian_(airline),Meridian
,MRD,MERIDIAN,Meridian_Air_Cargo,Meridian Air Cargo
,MHL,HASSIMAIR,Meridian_Airlines,Meridian Airlines
,DSL,DIESEL,Meridian_Aviation,Meridian Aviation
,MEM,MERIDIAN CHERRY,Meridian_Limited,Meridian Limited
IG,ISS,MERIDIANA,Meridiana,Meridiana
,MEI,AVALON,Merlin_Airways,Merlin Airways
MZ,MNA,MERPATI,Merpati_Nusantara_Airlines,Merpati Nusantara Airlines
YV,ASH,AIR SHUTTLE,Mesa_Airlines,Mesa Airlines
XJ,MES,MESABA,Mesaba_Airlines,Mesaba Airlines
,MSQ,META,Meta_Linhas_A%C3%A9reas,Meta Linhas Aéreas
,MET,METMAN,Meteorological_Research_Flight,Meteorological Research Flight
,MER,METHOW,Methow_Aviation,Methow Aviation
,MVI,,Metro_Business_Aviation,Metro Business Aviation
,MEX,EAGLE EXPRESS,Metro_Airlines,Metro Express
,MTR,METRO,Metroflight,Metroflight
,MTJ,METROJET,Metrojet_Ltd.,Metrojet
,PIX,METROPIX,Metropix_UK,Metropix UK
,MPS,METRO REGIONAL,Metropolis_(airline),Metropolis
,MXB,MEX BLUE,Mex_Blue,Mex Blue
,MJT,MEJETS,Mex-Jet,Mex-Jet
MX,MXA,MEXICANA,Mexicana_de_Aviaci%C3%B3n,Mexicana de Aviación
,MXT,TRANSMEX,MÃ©xico_Transportes_AÃ©reos,México Transportes Aéreos
,HUR,HURRICANE CHARTER,Miami_Air_Charter,Miami Air Charter
GL,BSK,BISCAYNE,Miami_Air_International,Miami Air International
,OWL,NIGHT OWL,Miami_Valley_Aviation,Miami Valley Aviation
,MPT,MIAPET,Miapet-Avia,Miapet-Avia
,BIB,,Michelin_Air_Services,Michelin Air Services
,WIZ,WIZARD,Micromatter_Technology_Solutions,Micromatter Technology Solutions
,NYL,NILE,Mid_Airlines_(Sudan),Mid Airlines
,MPA,MID PAC,Mid-Pacific_Airlines,Mid-Pacific Airlines
,MJR,MAJOR,Midamerica_Jet,Midamerica Jet
ME,MEA,CEDAR JET,Middle_East_Airlines,Middle East Airlines
,MID,,Midland_Airport_Services,Midland Airport Services
,MFR,MIDLINE FREIGHT,Midline_Air_Freight,Midline Air Freight
,MIS,MIDSTATE,Midstate_Airlines,Midstate Airlines
JI,MDW,MIDWAY,Midway_Airlines_(1993%E2%80%932003),Midway Airlines (1993–2003)
ML,MDW,MIDWAY,Midway_Airlines_(1976%E2%80%931991),Midway Airlines (1976–1991)
,FLA,PALM,Midway_Express,Midway Express
,FAX,FAIRFAX,Midwest_Air_Freighters,Midwest Air Freighters
YX,MEP,MIDEX,Midwest_Airlines,Midwest Airlines
MY,MWA,,Midwest_Airlines_(Egypt),Midwest Airlines (Egypt)
,NIT,NIGHTTRAIN,Midwest_Aviation,Midwest Aviation
,MWT,MIDWEST,Midwest_Aviation_Division,Midwest Aviation Division
,HTE,HELICOPTERSMEXICO,Midwest_Helicopters_De_Mexico,Midwest Helicopters De Mexico
MJ,MLR,MIHIN LANKA,Mihin_Lanka,Mihin Lanka
,MAB,MILLARDAIR,Millardair,Millardair
,RJM,MILLEN,The_Millen_Corporation,Millen Corporation
,MLK,NIGERJET,Millennium_Air,Millennium Air
,DLK,DEKKANLANKA,Millennium_Airlines,Millennium Airlines
,MFS,MILLER TIME,Miller_Flying_Services,Miller Flying Services
,OXO,MILL AIR,Million_Air,Million Air
,MIM,MIMINO,Mimino,Mimino
,NAB,,Mina_Airline_Company,Mina Airline Company
,OMR,ORMINE,Minair,Minair
,EBE,MINEBEA,Minebea_Technologies,Minebea Technologies
,MAZ,MINES,Mines_Air_Services_Zambia,Mines Air Services Zambia
,MNL,MINILINER,Miniliner,Miniliner
,MNS,MINISTIC,Ministic_Air,Ministic Air
,WDG,WATCHDOG,"Ministry_of_Agriculture,_Fisheries_and_Food_(United_Kingdom)","Ministry of Agriculture, Fisheries and Food"
,LIR,LISLINE,Minsk_Aircraft_Overhaul_Plant,Minsk Aircraft Overhaul Plant
,MIC,MINT AIRWAYS,Mint_Airways,Mint Airways
,MIR,MIRAMICHI,Miramichi_Air_Service,Miramichi Air Service
,MIF,MIRAS,Miras,Miras
,MOS,,Misr_Overseas_Airways,Misr Overseas Airways
,MAF,MISSI,Mission_Aviation_Fellowship,Mission Aviation Fellowship
,MSN,MISIONAIR,Missionair,Missionair
,MRN,MARIANNE,Missions_Gouvernemtales_Francaises,Missions Gouvernemtales Francaises
,BDG,BULLDOG,Mississippi_State_University,Mississippi State University
,MVA,VALAIR,Mississippi_Valley_Airways,Mississippi Valley Airways
M4,MSA,AIRMERCI,Mistral_Air,Mistral Air
,MBO,MOBIL,Mobil_Oil,Mobil Oil
,MXE,MOZAMBIQUE EXPRESS,Mocambique_Expresso,Mocambique Expresso
,MFZ,MOFAZ AIR,Mofaz_Air,Mofaz Air
,MOW,MOHAWK AIR,Mohawk_Airlines,Mohawk Airlines
MW,MUL,MUKULELE,Mokulele_Airlines,Mokulele Airlines
,MLE,MOLDAERO,Moldaeroservice,Moldaeroservice
2M,MDV,MOLDAVIAN,Moldavian_Airlines,Moldavian Airlines
,MVG,MOLDOVA-STATE,Moldova,Moldova
,RRV,SKYROVER,Mombasa_Air_Safari,Mombasa Air Safari
ZB,MON,MONARCH,Monarch_Airlines,Monarch Airlines
,MNH,MONARCH AIR,Monarch_Airlines_(USA),Monarch Airlines
8I,,,Myway_Airlines,Myway Airlines
,MFC,EAST WIND,Moncton_Flying_Club,Moncton Flying Club
,MDB,MONDEAIR CARGO,Monde_Air_Charters,Monde Air Charters
,MTI,MONTERREY AIR,Monerrey_Air_Taxi,Monerrey Air Taxi
,MKY,MONKY,Monky_Aerotaxis,Monky Aerotaxis
YM,MGX,MONTENEGRO,Montenegro_Airlines,Montenegro Airlines
5M,MNT,MONTSERRAT,Montserrat_Airways,Montserrat Airways
,MNY,MOONEY FLIGHT,Mooney_Aircraft_Corporation,Mooney Aircraft Corporation
,MAL,MORNINGSTAR,Morningstar_Air_Express,Morningstar Air Express
,MSS,WASATCH,Morris_Air,Morris Air Service
,MRO,MORRISON,Morrison_Flying_Service,Morrison Flying Service
3R,GAI,GROMOV AIRLINE,Moskovia_Airlines,Moskovia Airlines
,MPI,MOSPHIL,Mosphil_Aero,Mosphil Aero
M9,MSI,MOTOR SICH,Motor_Sich,Motor Sich
NM,NZM,MOUNTCOOK,Mount_Cook_Airlines,Mount Cook Airlines
,MTN,MOUNTAIN,Mountain_Air_Cargo,Mountain Air Cargo
N4,MTC,MOUNTAIN LEONE,Mountain_Air_Company,Mountain Air Company
,PKP,PIKES PEAK,Mountain_Air_Express,Mountain Air Express
,BRR,MOUNTAIN AIR,Mountain_Air_Service,Mountain Air Service
,MBI,MOUNTAIN BIRD,Mountain_Bird,Mountain Bird
,MHA,MOUNTAIN HIGH,Mountain_High_Aviation,Mountain High Aviation
,MPC,MOUNTAIN PACIFIC,Mountain_Pacific_Air,Mountain Pacific Air
,MTV,MOUNTAIN VALLEY,Mountain_Valley_Air_Service,Mountain Valley Air Service
,MDN,,Mudan_Airlines,Mudan Airlines
,CMJ,MUDANJIANG,Mudanjiang_General_Aviation,Mudanjiang General Aviation
,MTX,MULTITAXI,Multi_Taxi,Multi Taxi
,WBR,WEBER,Multi-Aero,Multi-Aero
,MFT,YORKAIR,Multiflight,Multiflight
,MNZ,MURMAN AIR,Murmansk_Aircompany,Murmansk Aircompany
,MUA,MURRAY AIR,Murray_Air,Murray Air
,MMR,MUSRATA AIR,Musrata_Air_Transport,Musrata Air Transport
,MAW,MUSTIQUE,Mustique_Airways,Mustique Airways
,MYW,FRANKY,My_Way_Airlines,My Way Airlines
VZ,MYT,KESTREL,MyTravel_Airways,MyTravel Airways
UB,UBA,UNIONAIR,Myanma_Airways,Myanma Airways
8M,MMA,MYANMAR,Myanmar_Airways_International,Myanmar Airways International
,MAV,MINOAN,Minoan_Air,Minoan Air
,MYA,MYFLUG,Myflug,Myflug
,VKG,VIKING,MyTravel_Airways,MyTravel Airways
,AAD,Ambassador,Mann_Air_Ltd,Mann Air Ltd
M2,MHV,SNOWCAP,MHS_Aviation_(Germany),MHS Aviation GmbH
,NAK,ENAC SCHOOL,%C3%89cole_Nationale_de_l%27Aviation_Civile,École Nationale de l'Aviation Civile
,NSR,AIR STAR,Nine_Star_Airways,Nine Star Airways
,NEW,MEADOW FLIGHT,Northeastern_Aviation,Northeastern Aviation
,SIQ,SCIENCE QUEST,National_Center_for_Atmospheric_Research,National Center for Atmospheric Research
,NEJ,NET BUSINESS,Netjets_Business_Aviation,Netjets Business Aviation
,NHV,,NHV_Aviation,NHV Aviation
,NHC,NORTHERN,Northern_Helicopter,Northern Helicopter
,DMD,DIAMONDJET,Namdeb_Diamond_Corporation,Namdeb Diamond Corporation
6N,NIN,NIGER AIRLINES,Niger_Airlines,Niger Airlines
N5,FEY,FLYEASY,Fly_Easy,Fly Easy
,NUB,VALLETTA,Nomad_Aviation,Nomad Aviation
,NJA,SHIN NIHON,New_Japan_Aviation,New Japan Aviation
,ROW,ROTORWING,Nor_Aviation,Nor Aviation
,NLG,NELCARGO,NEL_Cargo,NEL Cargo
,NHG,HELGA,NHT_Linhas_A%C3%A9reas,NHT Linhas Aéreas
,WAR,WARBIRDS,NZ_Warbirds_Association,NZ Warbirds Association
,ANL,AIR NACOIA,Nacoia_Lda,Nacoia Lda
,NHZ,NADA AIR,Nada_Air_Service,Nada Air Service
,BFN,,Naganagani,Compangnie Nationale Naganagani
,NAH,NAHANNI,Nahanni_Air_Services_Ltd,Nahanni Air Services Ltd
,NKL,NAKHEEL,Nakheel_Aviation,Nakheel Aviation
,MRE,MED RESCUE,Namibia_Commercial_Aviation,Namibia Commercial Aviation
,NDF,NAMIBIAN AIR FORCE,Namibian_Defence_Force,Namibian Defence Force
,CNJ,NINGHANG,Nanjing_Airlines,Nanjing Airlines
DV,ACK,ACK AIR,Nantucket_Airlines,Nantucket Airlines
,NYA,NANYAH,Nanyah_Aviation,Nanyah Aviation
,NAP,NAPIER,Napier_Air_Service_Inc,Napier Air Service Inc
,NCM,AIR BANE,Nas_Air_(Angola),Nas Air
P9,,,Nas_Air_(Mali),Nas Air
UE,NAS,NASAIRWAYS,Nasair,Nasair
,,,Nas_Air_(Saudi_Arabia),Nas Air
,NJC,NASHVILLE JET,Nashville_Jet_Charters,Nashville Jet Charters
,NCO,NATALCO,Natalco_Air_Lines,Natalco Air Lines
,NTK,NATCA,National_Air_Traffic_Controllers_Association,National Air Traffic Controllers Association
,NSR,NASAIR,National_Air_Charter,National Air Charter
,RFI,SHERLOCK,National_Air_Traffic_Services,National Air Traffic Services
,NAN,NATION AIR,National_Airlines_(1983),National Airlines
N4,NCN,,National_Airlines_(N4),National Airlines
N7,ROK,RED ROCK,National_Airlines_(N7),National Airlines
NA,NAL,NATIONAL,National_Airlines_(NA),National Airlines
N8,NCR,NATIONAL CARGO,National_Airlines_(N8),National Air Cargo dba National Airlines
IN,NIH,NAM,NAM_Air,NAM Air
,KUS,KUSWAG,National_Airlines_(YJ),National Airlines
9O,,,National_Airways_Cameroon,National Airways Cameroon
,LFI,AEROMED,National_Airways_Corporation_(Helicopters),National Airways Corporation
,GTY,,National_Aviation_Company,National Aviation Company
,TNC,NATCOM,National_Aviation_Consultants,National Aviation Consultants
,NXT,NATIONAL FREIGHT,National_Express_(airline),National Express
,GRD,GRID,National_Grid_plc,National Grid plc
,JTE,JETEX,National_Jet_Systems,National Jet Express
,AND,AIR INDIANA,National_Jet_Service,National Jet Service
NC,NJS,NATIONAL JET,National_Jet_Systems,National Jet Systems
,NOL,NAT AIRLINE,National_Overseas_Airlines_Company,National Overseas Airlines Company
,NLS,PANDER,Nationale_Luchtvaartschool,Nationale Luchtvaartschool
,NAE,NATIONS EXPRESS,Nations_Air_Express_Inc,Nations Air Express Inc
CE,NTW,NATIONWIDE,Nationwide_Airlines_(South_Africa),Nationwide Airlines
,NWZ,ZAMNAT,Nationwide_Airlines_(Zambia),Nationwide Airlines (Zambia)
,EVM,SCIENCE,Natural_Environment_Research_Council,Natural Environment Research Council
,NRR,NATUREAIR,Natureair,Natureair
,NRK,NATURELINK,Naturelink_Charter,Naturelink Charter
,NVC,NAV CAN,Nav_Canada,Nav Canada
,NAV,NAV DISPATCH,Nav_Flight_Planning,Nav Flight Planning
,NVP,,Navegacao_AÃ©rea_De_Portugal,Navegacao Aérea De Portugal
,NAY,NAYSA,NavegaciÃ³n_Servicios_AÃ©reos_Canarios_S.A.,Navegación Servicios Aéreos Canarios S.A.
,IRI,NAVID,Navid_Air,Navid
,NVM,NAVIERA,Naviera_Mexicana,Naviera Mexicana
,NVL,NAVLINES,Navigator_Airlines,Navigator Airlines
,XNV,,Navinc_Airlines_Services,Navinc Airlines Services
1N,,,Navitaire,Navitaire
,XNS,,Navtech_System_Support,Navtech System Support
,NZA,,Nayzak_Air_Transport,Nayzak Air Transport
,NEB,NEBRASKA,Nebraska,State of Nebraska
,NEC,NECON AIR,Necon_Air,Necon Air
,NCG,NETHERLANDS COASTGUARD,Nederlandse_Kustwacht,Nederlandse Kustwacht
,NFT,NEFTEAVIA,Nefteyugansk_Aviation_Division,Nefteyugansk Aviation Division
,NLA,NEILTOWN AIR,Neiltown_Air,Neiltown Air
,NLC,NELAIR,Nelair_Charters,Nelair Charters
,CGE,COLLEGE,Nelson_Aviation_College,Nelson Aviation College
RA,RNA,ROYAL NEPAL,Nepal_Airlines,Nepal Airlines
NO,NOS,MOONFLOWER,Neos_(airline),Neos
,TOX,SKY KINGDOM,Neosiam_Airways,Neosiam Airways
,NSL,NERICAIR,Neric,Neric
1I,EJA,EXECJET,NetJets,NetJets
,NET,NETWORK,Network_Aviation_Services,Network Aviation Services
,NEZ,ENGAIR,New_England_Air_Express,New England Air Express
EJ,NEA,NEW ENGLAND,New_England_Airlines,New England Airlines
,NHT,NEWHEIGHTS,New_Heights_291,New Heights 291
,NWD,NEW WORLD,New_World_Jet_Corporation,New World Jet Corporation
,NYH,NEW YORK,New_York_Helicopter,New York Helicopter
,GRY,GRAY RIDER,New_York_State_Police,New York State Police
,KRC,KIWI RESCUE,Royal_New_Zealand_Air_Force,Royal New Zealand Air Force
,HVA,HAVEN-AIR,Newair,Newair
,NLT,NALAIR,Newfoundland_Labrador_Air_Transport,Newfoundland Labrador Air Transport
2N,NTJ,NEXTJET,NextJet,NextJet
,NXF,NEXTFLIGHT,Nextflight_Aviation,Nextflight Aviation
,NXS,NEXUS AVIATION,Nexus_Aviation,Nexus Aviation
,NIS,NICA,Nicarag%C3%BCense_de_Aviaci%C3%B3n,Nicaragüense de Aviación
,NCN,NICON AIRWAYS,Nicon_Airways,Nicon Airways
,NGA,NIGERIA,Nigeria_Airways,Nigeria Airways
,NGR,NIGERIAN AIRFORCE,Nigerian_Air_Force,Nigerian Air Force
,NGX,AIR GLOBAL,Nigerian_Global,Nigerian Global
,EXT,EXECUTIVE,Night_Express,Night Express
HG,NLY,FLYNIKI,Niki_(airline),Niki
,NKV,AIR NIKOLAEV,Nikolaev-Air,Nikolaev-Air
,NSA,NILE SAFARIS,Nile_Safaris_Aviation,Nile Safaris Aviation
,NVA,,Nile_Valley_Aviation_Company,Nile Valley Aviation Company
,NLW,NILE WINGS,Nile_Wings_Aviation_Services,Nile Wings Aviation Services
,NBS,NIMBUS,Nimbus_Aviation,Nimbus Aviation
KZ,NCA,NIPPON CARGO,Nippon_Cargo_Airlines,Nippon Cargo Airlines
,NVK,VARTOSKAVIA,Nizhnevartovskavia,Nizhnevartovskavia
,SRG,,No._202_Squadron_RAF,Search and Rescue 202
,SRD,,No._22_Squadron_RAF,Search and Rescue 22
,NOH,NORTHOLT,No._32_Squadron_RAF,No. 32 (The Royal) Squadron
,AKG,GRIFTER,No._84_Squadron_RAF,No. 84 Squadron RAF
,NBL,NOBIL AIR,Nobil_Air,Nobil Air
DD,NOK,NOK AIR,Nok_Air,Nok Air
XW,NCT,BIG BIRD,NokScoot,NokScoot
,NRL,NOLINOR,Nolinor_Aviation,Nolinor Aviation
,NMD,NOMAD AIR,Nomad_Aviation,Nomad Aviation
,NOC,NORCOPTER,Norcopter,Norcopter
,NEF,NORDEX,Nord-Flyg,Nord-Flyg
5N,AUL,ARCHANGELSK AIR,Nordavia,Nordavia
JH,NES,NORDESTE,Nordeste_Linhas_A%C3%A9reas_Regionais,Nordeste Linhas Aéreas Regionais
6N,NRD,NORTH RIDER,Nordic_Regional,Nordic Regional
,NDS,,Nordstree_(Australia),Nordstree (Australia)
N4,NWS,NORDLAND,Nordwind_Airlines,Nordwind Airlines
,NRT,NORESTAIR,Norestair,Norestair
N5,,,Norfolk_Air,Norfolk Air
,NCF,COUNTY,Norfolk_County_Flight_College,Norfolk County Flight College
,FNA,NORLAND,Norlandair,Norlandair
,NOA,NORONTAIR,Norontair,Norontair
,HMF,LIFEGUARD SWEDEN,Norrlandsflyg,Norrlandsflyg
,NRX,NORSE AIR,Norse_Air_Charter,Norse Air Charter
,NIR,NORSEMAN,Norsk_Flytjeneste,Norsk Flytjeneste
,NOR,NORSKE,Norsk_Helikopter,Norsk Helikopter
,DOC,HELIDOC,Norsk_Luftambulanse,Norsk Luftambulanse
,RTV,TIC-TAC,Nortavia,Nortavia
,NAI,NORTH-ADRIA,North_Adria_Aviation,North Adria Aviation
NA,NAO,NORTH AMERICAN,North_American_Airlines,North American Airlines
,HMR,HAMMER,North_American_Charters,North American Charters
,NAJ,JET GROUP,North_American_Jet_Charter_Group,North American Jet Charter Group
,NAT,MASS AIR,North_Atlantic_Air_Inc,North Atlantic Air Inc
,NFC,NORTH ATLANTIC,North_Atlantic_Cargo,North Atlantic Cargo
,NBN,TEESAIR,North_British_Airlines,North British Airlines
,NCB,NORTH CARIBOU,North_Caribou_Flying_Service_Ltd,North Caribou Flying Service Ltd
,NCC,NORTH COAST,North_Coast_Air_Services_Ltd,North Coast Air Services Ltd
N9,,,North_Coast_Aviation,North Coast Aviation
M3,NFA,NORTH FLYING,North_Flying,North Flying
,NRC,NORTH SEA,North_Sea_Airways,North Sea Airways
,SBX,SKY BOX,North_Star_Air_Cargo,North Star Air Cargo
,NRV,NORVAN,North_Vancouver_Airlines,North Vancouver Airlines
,NWW,HALANT,North_West_Airlines_(Australia),North West Airlines
,PTO,PHOTO,North_West_Geomatics,North West Geomatics
,NEN,NORTHEAST SWAN,North-East_Airlines,North-East Airlines
,VBG,VYBORG AIR,North-West_Air_Transport_Company_-_Vyborg,North-West Air Transport Company - Vyborg
HW,NWL,NORTHWRIGHT,North-Wright_Airways,North-Wright Airways
,NLL,NORTHAFRICAN AIR,Northafrican_Air_Transport,Northafrican Air Transport
,NFL,GREAT LAKES,Northaire_Freight_Lines,Northaire Freight Lines
,NSF,NORTON,Northamptonshire_School_of_Flying,Northamptonshire School of Flying
,NCE,TOP HAT,Northcoast_Executive_Airlines,Northcoast Executive Airlines
,NEE,NORTHEAST,Northeast_Airlines,Northeast Airlines
,NPX,NORTHEAST EXPRESS,Northeast_Aviation,Northeast Aviation
NC,NAC,YUKON,Northern_Air_Cargo,Northern Air Cargo
,BYC,BEIYA,Northern_Airlines_Sanya,Northern Airlines Sanya
,NDA,NORTHERN DAKOTA,Northern_Airways,Northern Airways
,CMU,LANNA AIR,Northern_Aviation_Service,Northern Aviation Service
U7,,,Northern_Dene_Airways,Northern Dene Airways
,NEX,NEATAX,Northern_Executive_Aviation,Northern Executive Aviation
,NIC,ILLINOIS COMMUTER,Northern_Illinois_Commuter,Northern Illinois Commuter
,NTX,NORTAX,Northern_Jet_Management,Northern Jet Management
,NTA,THUNDERBIRD,Northern_Thunderbird_Air,Northern Thunderbird Air
,KOE,KOKEE,Northland_Aviation,Northland Aviation
,NSS,NORTHSTAR,Northstar_Aviation,Northstar Aviation
,NHL,NORTHUMBRIA,Northumbria_Helicopters,Northumbria Helicopters
,NAL,NORTHWAY,Northway_Aviation_Ltd,Northway Aviation Ltd
,NWE,,Northwest_Aero_Associates,Northwest Aero Associates
NW,NWA,NORTHWEST,Northwest_Airlines,Northwest Airlines
FY,NWR,,Northwest_Regional_Airlines,Northwest Regional Airlines
,NWT,TERRITORIAL,Northwest_Territorial_Airways,Northwest Territorial Airways
J3,PLR,POLARIS,Northwestern_Air,Northwestern Air
,NWN,NORTHWINDS,Northwinds_Northern,Northwinds Northern
,NAM,MANITOBA,Nortland_Air_Manitoba,Nortland Air Manitoba
DY,NAX,NOR SHUTTLE,Norwegian_Air_Shuttle,Norwegian Air Shuttle
DI,NRS,REDNOSE,Norwegian_Air_UK,Norwegian Air UK
DN,NAA,NORUEGA,Norwegian_Air_Argentina,Norwegian Air Argentina
D8,IBK,NORTRANS,Norwegian_Air_International,Norwegian Air International
DU,NLH,NORSTAR,Norwegian_Long_Haul,Norwegian Long Haul
DH,NAN,NORSHIP,Norwegian_Air_Norway,Norwegian Air Norway
,TFN,SPIRIT,Norwegian_aviation_college,Norwegian Aviation College
,XNT,,Notams_International,Notams International
BJ,LBT,NOUVELAIR,Nouvelair,Nouvel Air Tunisie
O9,NOV,NOVANILE,Nova_Airline,Nova Airline
,PTR,PATROL,Nova_Scotia_Department_of_Lands_and_Forests,Nova Scotia Department of Lands and Forests
1I,NVR,NAVIGATOR,Novair,Novair
VQ,NVQ,NOVO AIR,Novo_Air,Novo Air
,NVG,SADKO AVIA,Novogorod_Air_Enterprise,Novogorod Air Enterprise
,NSP,NARPAIR,Novosibirsk_Aircraft_Repairing_Plant,Novosibirsk Aircraft Repairing Plant
,NBE,NAKAIR,Novosibirsk_Aviaenterprise,Novosibirsk Aviaenterprise
,NPO,NOVSIB,Novosibirsk_Aviation_Production_Association,Novosibirsk Aviation Production Association
,NOY,NOY AVIATION,Noy_Aviation,Noy Aviation
N6,ACQ,AERO CONTINENTE,Nuevo_Continente,Nuevo Continente
,NHR,NUEVO HORIZONTE,Nuevo_Horizonte_Internacional,Nuevo Horizonte Internacional
,NUN,NUNASI,Nunasi-Central_Airlines,Nunasi-Central Airlines
,NIN,NURVINDO,Nurman_Avia_Indopura,Nurman Avia Indopura
,NYS,NYASA,Nyasa_Express,Nyasa Express
1I,NJE,FRACTION,NetJets_Europe,NetJets Europe
NK,NKS,SPIRIT WINGS,Spirit_Airlines,Spirit Airlines
,ORN,ORANGE JET,Orange_Air,Orange Air
,ONS,AIR DREAMS,One_Airlines,One Airlines
,FET,FREIGHT LINE,Old_Dominion_Freight_Lines,Old Dominion Freight Lines
,OCN,O-BIRD,O_Air,O Air
UQ,OCM,OCONNOR,O%27Connor_Airlines,O'Connor Airlines
,DRL,DRILLER,Omni_Air_Transport,Omni Air Transport
,OWE,OWENAIR,Owenair,Owenair
,JPA,J-PAT,OSACOM,OSACOM
O8,OHK,OASIS,Oasis_Hong_Kong_Airlines,Oasis Hong Kong Airlines
,BCN,BLUE OCEAN,Ocean_Air,Ocean Air
VC,VCX,OCEANCARGO,Ocean_Airlines,Ocean Airlines
,OCS,OCEANSKY,Ocean_Sky_(UK),Ocean Sky (UK)
,TUK,TUCKERNUCK,Ocean_Wings_Commuter_Service,Ocean Wings Commuter Service
O6,ONE,OCEANAIR,Avianca_Brazil,Avianca Brazil
O2,,HARPOON,Linear_Air,Linear Air
,ODS,ODESSA AIR,Odessa_Airlines,Odessa Airlines
,ODY,ODYSSEY,Odyssey_International,Odyssey International
,FOC,FOCA,Office_Federal_De'Aviation_Civile,Office Federal De'Aviation Civile
,GBO,,Ogooue_Air_Cargo,Ogooue Air Cargo
,OKJ,OKADA AIR,Okada_Airlines,Okada Airlines
,OKP,OKAPI,Okapi_Airways,Okapi Airways
,OKA,OKAYJET,Okay_Airways,Okay Airways
,OKL,OKLAHOMA,Oklahoma_Department_of_Public_Safety,Oklahoma Department of Public Safety
,OLX,OLIMEX,Olimex_Aerotaxi,Olimex Aerotaxi
,KVK,PONTA,Olimp_Air,Olimp Air
OA,OAL,OLYMPIC,Olympic_Air,Olympic Air
,OLY,OLAVIA,Olympic_Aviation,Olympic Aviation
WY,OMA,OMAN AIR,Oman_Air,Oman Air
,ORF,OMAN,Oman_Royal_Flight,Oman Royal Flight
OV,OMS,MAZOON,SalamAir,SalamAir
,OAV,OMNI,Omni_-_Aviacao_e_Tecnologia,Omni - Aviacao e Tecnologia
OE,LDM,LAUDA MOTION,Laudamotion,Laudamotion
OY,OAE,OMNI-EXPRESS,Omni_Air_International,Omni Air International
,OMF,OMNIFLYS,Omniflys,Omniflys
,ORL,ON AIR,On_Air_Limited,On Air Limited
,OTG,THAI EXPRESS,One_Two_Go_Airlines,One Two Go Airlines
,OTM,ZEDTIME,Onetime_Airlines_Zambia,Onetime Airlines Zambia
,MED,MEDICAL,Ontario_Ministry_of_Health,Ontario Ministry of Health
8Q,OHY,ONUR AIR,Onur_Air,Onur Air
,OPA,,Opal_Air,Opal Air
,OSA,,Open_Sky_Aviation,Open Sky Aviation
,BOS,MISTRAL,OpenSkies,OpenSkies
,ORR,TURISTICA AURORA,Operadora_Turistica_Aurora,Operadora Turistica Aurora
,OLE,OPERADORA,Operadora_de_Lineas_Ejecutivas,Operadora de Lineas Ejecutivas
,OTP,OPERADORA AEREO,Operadora_de_Transportes_AÃ©reos,Operadora de Transportes Aéreos
,OPV,OPERADORA DE VUELOS,Operadora_de_Veulos_Ejectutivos,Operadora de Veulos Ejectutivos
,LLO,APOLLO,Operation_Enduring_Freedom,Operation Enduring Freedom
,OAX,,Operational_Aviation_Services,Operational Aviation Services
,ORD,ORANGE SERVICES,Orange_Air_Services,Orange Air Services
,ORJ,ORANGE SIERRA,Orange_Air_Sierra_Leone,Orange Air Sierra Leone
,ORE,ORANGE AVIATION,Orange_Aviation,Orange Aviation
,ORX,OREX,Orbit_Express_Airlines,Orbit Express Airlines
,ORK,ORCA TAXI,Orca_Air,Orca Air
,BUE,BLUELIGHT,Orebro_Aviation,Orebro Aviation
,ORM,ORPRISE,Orel_State_Air_Enterprise,Orel State Air Enterprise
R2,ORB,ORENBURG,Orenburg_Airlines,Orenburg Airlines
,OTA,ORGANIZACION,Organizacion_De_Transportes_AÃ©reos,Organizacion De Transportes Aéreos
,OML,MAMBRA,Organizacoes_Mambra,Organizacoes Mambra
,OVV,ORIENTSYR,Orient_Air,Orient Air
,OTR,ORIENTROC,Orient_Airlines,Orient Airlines
,ORN,ORIENT LINER,Orient_Airways,Orient Airways
OX,OEA,ORIENT THAI,Orient_Thai_Airlines,Orient Thai Airlines
,NGK,ORIENTAL BRIDGE,Oriental_Air_Bridge,Oriental Air Bridge
,OAC,ORIENTAL AIR,Oriental_Airlines,Oriental Airlines
QO,OGN,ORIGIN,Origin_Pacific_Airways,Origin Pacific Airways
,OED,ORION CHARTER,Orion_Air_Charter,Orion Air Charter
,OIX,ORIONIX,Orion-x,Orion-x
,KOV,ORLAN,Orlan-2000,Orlan-2000
,RNG,ORANGE,Orange_Aircraft_Leasing,Orange Aircraft Leasing
,OAD,ORSCOM,Orscom_Tourist_Installations_Company,Orscom Tourist Installations Company
,OSH,OSH AVIA,Osh_Avia,Osh Avia
,OCO,AIR COLLEGE,Ostend_Air_College,Ostend Air College
OL,OLT,OLTRA,OLT_Express_Germany,OLT Express Germany
,FNL,FINN FLIGHT,Oulun_Tilauslento,Oulun Tilauslento
ON,RON,OUR AIRLINE,Our_Airline,Our Airline
,OOT,OOTBAS,Out_Of_The_Blue_Air_Safaris,Out Of The Blue Air Safaris
OJ,OLA,OVERLAND,Overland_Airways,Overland Airways
,OAR,BOSS AIR,ONE_AIR,ONE AIR
,OXE,OXOE,Oxaero,Oxaero
,WDK,WOODSTOCK,Oxford_Air_Services,Oxford Air Services
,OAA,,Oxley_Aviation,Oxley Aviation
OZ,OZR,OZARK,Ozark_Air_Lines,Ozark Air Lines
O7,OZJ,AUSJET,Ozjet_Airlines,Ozjet Airlines
OA,OAL,OLYMPIC,Olympic_Airlines,Olympic Airlines
OB,AAN,OASIS,Oasis_International_Airlines,Oasis International Airlines Now assigned to Boliviana de Aviacion (BoA)
,HRS,HORSEMAN,,Pursuit Aviation
,PNA,SEBUS,Palau_National_Airlines,Palau National Airlines
,RSL,PANAMA RENTAL,Panama_Aircraft_Rental_and_Sales,Panama Aircraft Rental and Sales
,NCT,PETE AIR,Pete_Air,Pete Air
,PRT,PRIME ITALIA,Prime_Service_Italia,Prime Service Italia
,PXT,PACK COAST,Pacific_Coast_Jet,Pacific Coast Jet
,BPH,BLACK PHOENIX,Phoenix_Helicopter_Academy,Phoenix Helicopter Academy
,PFY,PELFLIGHT,Pel-Air_Aviation,Pel-Air Aviation
,PXR,PIXAIR,Pixair_Survey,Pixair Survey
,PNC,PRINCE,Prince_Aviation,Prince Aviation
,PMI,AEROEPRIM,Primero_Transportes_Aereos,Primero Transportes Aereos
,KTL,KNOTTSBERRY,P_&_P_Floss_Pick_Manufacturers,P & P Floss Pick Manufacturers
,PCR,PACAIR,PAC_Air,PAC Air
PV,PNR,SKYJET,PAN_Air,PAN Air
9Q,PBA,PEEBEE AIR,PB_Air,PB Air
,PDQ,DISPATCH,PDQ_Air_Charter,PDQ Air Charter
,XAS,,PHH_Aviation_System,PHH Aviation System
,PDG,OSPREY,PLM_Dollar_Group,PLM Dollar Group
PU,PUA,PLUNA,PLUNA,PLUNA
U4,PMT,MULTITRADE,PMTair,PMTair
,PRP,PRONTO,PRT_Aviation,PRT Aviation
OH,JIA,BLUE STREAK,PSA_Airlines,PSA Airlines
,KST,KING STAR,PTL_Luftfahrtunternehmen,PTL Luftfahrtunternehmen
,WIS,WISCAIR,Paccair,Paccair
Y5,PCE,PACE,Pace_Airlines,Pace Airlines
,PAB,AIR BOATS,Pacific_Air_Boats,Pacific Air Boats
,PRC,PACIFIC CHARTER,Pacific_Air_Charter,Pacific Air Charter
,PAQ,SOLPAC,Pacific_Air_Express,Pacific Air Express
,PXP,PAK EXPRESS,Pacific_Air_Transport,Pacific Air Transport
BL,PIC,PACIFIC AIRLINES,Pacific_Airlines,Pacific Airlines
,PAK,PACIFIC ALASKA,Pacific_Alaska_Airlines,Pacific Alaska Airlines
,PCV,PACAV,Pacific_Aviation_(Australia),Pacific Aviation (Australia)
,PCX,,Pacific_Aviation_(United_States),Pacific Aviation (United States)
DJ,PBN,BLUEBIRD,Pacific_Blue_(airline),Pacific Blue
,PQA,SAGE BRUSH,Pacific_Coast_Airlines,Pacific Coast Airlines
8P,PCO,PASCO,Pacific_Coastal_Airlines,Pacific Coastal Airlines
Q8,PEC,PAC-EAST CARGO,Pacific_East_Asia_Cargo_Airlines,Pacific East Asia Cargo Airlines
,PFA,PACIFIC SING,Pacific_Flight_Services,Pacific Flight Services
,PIN,ROAD RUNNERS,Pacific_International_Airlines,Pacific International Airlines
,PSA,PACIFIC ISLE,Pacific_Island_Aviation,Pacific Island Aviation
,PCJ,PACIFIC JET,Pacific_Jet,Pacific Jet
,PPM,PACIFIC PEARL,Pacific_Pearl_Airways,Pacific Pearl Airways
,PAR,PACRIM,Pacific_Rim_Airways,Pacific Rim Airways
LW,NMI,TSUNAMI,Pacific_Wings,Pacific Wings
GX,,,Pacificair,Pacificair
,PFR,PACIFIC WEST,Pacificair_Airlines,Pacificair Airlines
,RCY,RACE CITY,Package_Express,Package Express
,PAE,PAISAJES,Paisajes_EspaÃ±oles,Paisajes Españoles
,PKW,PLATINUM WEST,Pak_West_Airlines,Pak West Airlines
PK,PIA,PAKISTAN,Pakistan_International_Airlines,Pakistan International Airlines
,PKR,PAKKER AVIO,Pakker_Avio,Pakker Avio
,LPA,LINEASPAL,Pal_AerolÃ­neas,Pal Aerolíneas
,PPC,PALAU ASIAPAC,Palau_Asia_Pacific_Airlines,Palau Asia Pacific Airlines
GP,PTP,TRANS PACIFIC,Palau_Trans_Pacific_Airlines,Palau Trans Pacific Airlines
PF,PNW,PALESTINIAN,Palestinian_Airlines,Palestinian Airlines
,JSP,PALMER,Palmer_Aviation,Palmer Aviation
NR,PIR,PAMIR,Pamir_Airways,Pamir Airways
,PFN,PANAFRICAN,,Pan African Air Services
,ODM,,Pan_African_Airways,Pan African Airways
,PAX,PANNEX,Pan_Air,Pan Air
,XPA,,Pan_Am_Weather_Systems,Pan Am Weather Systems
7N,PWD,,PAWA_Dominicana,PAWA Dominicana
PN,,,Pan_American_Airways_(1996-1998),Pan American Airways
PA,PAA,,Pan_American_Airways_(1998-2004),Pan American Airways
PA,PAA,CLIPPER,Pan_American_World_Airways,Pan American World Airways
,PEA,PEACH,Pan_Europeenne_Air_Service,Pan Europeenne Air Service
,PHT,PANANK,Pan_Havacilik_Ve_Ticaret,Pan Havacilik Ve Ticaret
,PMA,PAN MALAYSIA,Pan_Malaysian_Air_Transport,Pan Malaysian Air Transport
,PNC,PANAIRSA,Pan-Air,Pan-Air
,PNF,PANWAYS,Panafrican_Airways,Panafrican Airways
,PGI,PANAGRA,Panagra_Airways,Panagra Airways
,PEI,PANAMEDIA,Panamedia,Panamedia
,PVI,,Panavia,Panavia
,PNH,KUBAN LIK,Panh,Panh
,PHU,PANNON,Pannon_Air_Service,Pannon Air Service
,PNM,PANORAMA,Panorama,Panorama
,PAH,LANI,Panorama_Air_Tour,Panorama Air Tour
,AFD,AIRFED,Panorama_Flight_Service,Panorama Flight Service
P8,PTN,PANTANAL,Pantanal_Linhas_A%C3%A9reas,Pantanal Linhas Aéreas
,HMP,PAPAIR TERMINAL,Papair_Terminal,Papair Terminal
,PAI,SEA RAY,Paradise_Airways,Paradise Airways
,PDI,PARADISE ISLAND,Paradise_Island_Airways,Paradise Island Airways
,PGX,PARAGON EXPRESS,Paragon_Air_Express,Paragon Air Express
,PGF,,Paragon_Global_Flight_Support,Paragon Global Flight Support
,PRR,PARAMOUNT,Paramount_Airlines,Paramount Airlines
I7,PMW,PARAWAY,Paramount_Airways,Paramount Airways
,APE,AIR PARCEL,Parcel_Express,Parcel Express
,IRE,PARIZAIR,Pariz_Air,Pariz Air
,PRA,PARSAVIA,Pars_Aviation_Service,Pars Aviation Service
,PST,TURISMO REGIONAL,Air_Panama,Parsa
,FAP,,Parsons_Airways_Northern,Parsons Airways Northern
,PSC,PASCAN,Pascan_Aviation,Pascan Aviation
P3,PTB,PASSAREDO,Passaredo_Transportes_A%C3%A9reos,Passaredo Transportes Aéreos
,PTC,PATRIA,Patria_Cargas_AÃ©reas,Patria Cargas Aéreas
,BYT,BYTE,Patriot_Aviation_Limited,Patriot Aviation Limited
,ETL,ENTEL,Patterson_Aviation_Company,Patterson Aviation Company
,PHE,PAWAN HANS,Pawan_Hans,Pawan Hans
,IRP,PAYAMAIR,Payam_Air,Payam Air
,KGC,GOLDCREST,Peach_Air,Peach Air
,PRL,PEARL LINE,Pearl_Air,Pearl Air
,PBY,PEARL SERVICES,Pearl_Air_Services,Pearl Air Services
HP,HPA,PEARL AIRWAYS,Pearl_Airways,Pearl Airways
,PVU,PEAU,Peau_Vava%CA%BBu,Peau Vavaʻu
,PXA,PECOTOX,Pecotox_Air,Pecotox Air
PC,PGT,SUNTURK,Pegasus_Airlines,Pegasus Airlines
PE,PEV,PEOPLES,People%27s,People's
1I,,Sunturk,Pegasus_Hava_Tasimaciligi,Pegasus Hava Tasimaciligi
,HAK,HELIFALCON,Pegasus_Helicopters,Pegasus Helicopters
,PDF,PELICAN AIRWAYS,Pelican_Air_Services,Pelican Air Services
,PEX,PELICAN EXPRESS,Pelican_Express,Pelican Express
,PAS,PELITA,Pelita_Air_Service,Pelita Air Service
,PEM,PEM-AIR,Pem-Air,Pem-Air
,PDY,PENDLEY,Pen-Avia,Pen-Avia
KS,PEN,PENINSULA,Peninsula_Airways,Peninsula Airways
,PNE,PENINTER,Peninter_AÃ©rea,Peninter Aérea
,PCA,PENA DEL AIRE,Penya_De_L'Aire,Penya De L'Aire
,CVT,CVETA,Peran,Peran
,PCC,PERFORADORA CENTRAL,Perforadora_Central,Perforadora Central
,PAG,PERIMETER,Perimeter_Aviation,Perimeter Aviation
P9,PGP,PERM AIR,Perm_Airlines,Perm Airlines
,PPQ,PERSONSPAQ,Personas_Y_Pasquetes_Por_Air,Personas Y Pasquetes Por Air
P9,PVN,,Peruvian_Airlines,Peruvian Airlines
,FPR,,Peruvian_Air_Force,Peruvian Air Force
,INP,,Peruvian_Navy,Peruvian Navy
,PEO,PETRO AIR,Petro_Air,Petro Air
,PMX,PEMEX,Petroleos_Mexicanos,Petroleos Mexicanos
,PHM,PETROLEUM,Petroleum_Helicopters,Petroleum Helicopters
,PHC,HELICOPTERS,Petroleum_Helicopters_de_Colombia,Petroleum Helicopters de Colombia
,PTK,PETROKAM,Petropavlovsk-Kamchatsk_Air_Enterprise,Petropavlovsk-Kamchatsk Air Enterprise
,PTY,PETTY,Petty_Transport,Petty Transport
,PHV,NEW BIRD,Phenix_Aviation,Phenix Aviation
,PMY,PHETCHABUN AIR,Phetchabun_Airline,Phetchabun Airline
Z2,EZD,REDHOT,Philippines_AirAsia,Philippines AirAsia
PR,PAL,PHILIPPINE,Philippine_Airlines,Philippine Airlines
,PHI,PHILAIR,Philips_Aviation_Services,Philips Aviation Services
,BCH,BEACHBALL,Phillips_Air,Phillips Air
,PDD,PADA,Phillips_Alaska,Phillips Alaska
,PHL,PHILLIPS,Phillips_Michigan_City_Flying_Service,Phillips Michigan City Flying Service
,PHB,PHOEBUS,Phoebus_Apollo_Aviation,Phoebus Apollo Aviation
,KZM,CARZAM,Phoebus_Apolloa_Zambia,Phoebus Apolloa Zambia
,PHA,GRAY BIRD,Phoenix_Air_(USA),Phoenix Air Group
,PHN,PHOENIX BRASIL,Phoenix_Air_Lines,Phoenix Air Lines
,PAM,PHOENIX,Phoenix_Air_(PAM),Phoenix Air
,PPG,PAPAGO,Phoenix_Air_Transport,Phoenix Air Transport
,WDY,WINDYCITY,Phoenix_Airline_Services,Phoenix Airline Services
HP,,,Phoenix_Airways,Phoenix Airways
,PHY,PHOENIX ARMENIA,Phoenix_Avia,Phoenix Avia
,PHG,PHOENIX GROUP,Phoenix_Aviation,Phoenix Aviation
,XPX,,Phoenix_Flight_Operations,Phoenix Flight Operations
9R,VAP,PHUKET AIR,Phuket_Air,Phuket Air
PI,PDT,PIEDMONT,Piedmont_Airlines_(1948-1989),Piedmont Airlines (1948-1989)
,PDT,PIEDMONT,Piedmont_Airlines,Piedmont Airlines
,PCH,PILATUS WINGS,Pilatus_Flugzeugwerke,Pilatus Flugzeugwerke
,PLU,PILATUS MEXICO,Pilatus_PC-12_Center_De_Mexico,Pilatus PC-12 Center De Mexico
,MKS,MIKISEW,Pimichikamac_Air,Pimichikamac Air
,PNP,PINEAPPLE AIR,Pineapple_Air,Pineapple Air
,PIM,PINFRAMAT,Pinframat,Pinframat
,PCL,PINNACLE GROUP,Pinnacle_Air_Group,Pinnacle Air Group
9E,FLG,FLAGSHIP,Pinnacle_Airlines,Pinnacle Airlines
,PIO,PIONEER,Pioneer_Airlines,Pioneer Airlines
,PER,,Pioneers_Limited,Pioneers Limited
,PRN,PRINAIR EXPRESS,Pirinair_Express,Pirinair Express
,PLN,PLANAR,Planar_(airline),Planar
,PMS,PLANEMASTER,Planemaster_Services,Planemaster Services
,PLZ,PLANET,Planet_Airways,Planet Airways
,PYZ,PLAYERS AIR,Players_Air,Players Air
,LIB,LIBELLE,Polizeihubschrauberstaffel_Hamburg,Polizeihubschrauberstaffel Hamburg
,PSF,LIZARD,Plymouth_School_of_Flying,Plymouth School of Flying
,POC,POCONO,Pocono_Air_Lines,Pocono Air Lines
,PDA,PODILIA,Podilia-Avia,Podilia-Avia
,PAZ,POINTAIR NIGER,Point-Afrique,Point Afrique Niger
,RMI,POINT AIRLINE,Point_Airlines,Point Airlines
,PAW,POINTAIR BURKINA,Pointair_Burkina,Pointair Burkina
,PTS,POINTSCALL,Points_of_Call_Airlines,Points of Call Airlines
PO,PAC,POLAR,Polar_Air_Cargo,Polar Air Cargo
,PMO,POLAR MEXICO,Polar_Airlines_de_Mexico,Polar Airlines de Mexico
,PSR,POLESTAR,Polestar_Aviation,Polestar Aviation
,POT,POLET,Polet_Airlines,Polet Airlines
,POF,AIRPOL,Police_Aux_Fronti%C3%A8res,Police Aux Frontières
,PLC,SPECIAL,Police_Aviation_Services,Police Aviation Services
,PLF,POLISH AIRFORCE,Polish_Air_Force,Polish Air Force
,PNY,POLISH NAVY,Polish_Navy,Polish Navy
,NRW,HUMMEL,Polizeifliegerstaffel_Nordrhein-Westfalen,Polizeifliegerstaffel Nordrhein-Westfalen
,PPH,POLICE PHOENIX,Polizeihubschrauberstaffel_Niedersachsen,Polizeihubschrauberstaffel Niedersachsen
,PIK,POLICE IKARUS,Polizeihubschrauberstaffel_Sachsen-Anhalt,Polizeihubschrauberstaffel Sachsen-Anhalt
,SRP,SPERBER,Polizeihubschauberstaffel_Rheinland-Pfalz,Polizeihubschauberstaffel Rheinland-Pfalz
,PBW,BUSSARD,Polizeihubschrauberstaffel_Baden-WÃ¼rttemberg,Polizeihubschrauberstaffel Baden-Württemberg
,EDL,POLICE EDELWEISS,Polizeihubschrauberstaffel_Bayern,Polizeihubschrauberstaffel Bayern
,PBB,ADEBAR,Polizeihubschrauberstaffel_Brandenburg,Polizeihubschrauberstaffel Brandenburg
,PHH,IBIS,Polizeihubschrauberstaffel_Hessen,Polizeihubschrauberstaffel Hessen
,PMV,POLICE MERLIN,Polizeihubschrauberstaffel_Mecklenburg-Vorpommern,Polizeihubschrauberstaffel Mecklenburg-Vorpommern
,PHS,PASSAT,Polizeihubschrauberstaffel_Sachsen,Polizeihubschrauberstaffel Sachsen
,HBT,HABICHT,Polizeihubschrauberstaffel_ThÃ¼ringen,Polizeihubschrauberstaffel Thüringen
,CUK,CHUKKA,Polo_Aviation,Polo Aviation
,PLA,POLYAIR,Polynesian_Air-Ways,Polynesian Air-Ways
PH,PAO,POLYNESIAN,Polynesian_Airlines,Polynesian Airlines
DJ,PLB,POLYBLUE,Polynesian_Blue,Polynesian Blue
1U,,,Polyot_Sirena,Polyot Sirena
,PND,POND AIR,Pond_Air_Express,Pond Air Express
,PSI,PONT,Pont_International_Airline_Services,Pont International Airline Services
,PLX,POOLEX,Pool_Aviation,Pool Aviation
,PTQ,TOWNSEND,Port_Townsend_Airways,Port Townsend Airways
,POR,PORTEADORA,Porteadora_De_Cosola,Porteadora De Cosola
PD,POE,PORTER,Porter_Airlines,Porter Airlines
NI,PGA,PORTUGALIA,Portugalia,Portugalia
,AFP,PORTUGUESE AIR FORCE,Portuguese_Air_Force,Portuguese Air Force
,POA,PORTUGUESE ARMY,Portuguese_Army,Portuguese Army
,PON,PORTUGUESE NAVY,Portuguese_Navy,Portuguese Navy
BK,PDC,DISTRICT,Potomac_Air,Potomac Air
,PSN,POTOSINA,Potosina_Del_Aire,Potosina Del Aire
,PWL,POWELL AIR,Powell_Air,Powell Air
,PFS,PRAIRIE,Prairie_Flying_Service,Prairie Flying Service
,PWC,PRATT,Pratt_and_Whitney_Canada,Pratt and Whitney Canada
PW,PRF,PRECISION AIR,Precision_Air,Precision Air
,PRE,PRECISION,Precision_Airlines,Precision Airlines
,BAT,BALLISTIC,Premiair,Premiair
,PGL,PREMIERE,Premiair_Aviation_Services,Premiair Aviation Services
,PME,ADUR,Premiair_Fliyng_Club,Premiair Fliyng Club
,EMI,BLUE SHUTTLE,Premium_Air_Shuttle,Premium Air Shuttle
,PMU,PREMIUM,Premium_Aviation,Premium Aviation
,BFA,,Presidence_Du_Faso,Presidence Du Faso
,ONM,,Presidencia_de_La_Republica_de_Guinea_Ecuatorial,Presidencia de La Republica de Guinea Ecuatorial
TO,PSD,,President_Airlines,President Airlines
,PRD,PRESIDENTIAL,Presidential_Aviation,Presidential Aviation
,PWA,PRIESTER,Priester_Aviation,Priester Aviation
,PMM,PRIMAVIA,Primair,Primair
FE,WCP,WHITECAP,Primaris_Airlines,Primaris Airlines
,PMC,PRIMAC,Primas_Courier,Primas Courier
,CRY,CARRIERS,Primavia_Limited,Primavia Limited
,PRM,PRIME AIR,Prime_Airlines,Prime Airlines
,PKZ,PRAVI,Prime_Aviation,Prime Aviation
,CME,COMET,Prince_Edward_Air,Prince Edward Air
,PJP,PRINCELY JETS,Princely_Jets,Princely Jets
8Q,,,Princess_Air,Princess Air
,PCN,PRINCETON,Princeton_Aviation_Corporation,Princeton Aviation Corporation
,PRY,PRIORITY AIR,Priority_Air_Charter,Priority Air Charter
,PAT,PAT,Priority_Air_Transport,Priority Air Transport
,BCK,BANKCHECK,Priority_Aviation_Company,Priority Aviation Company
,PTI,PRIVATAIR,Privatair,Privatair
,PJE,PEE JAY,Private_Jet_Expeditions,Private Jet Expeditions
,PJA,PRIVATE FLIGHT,Private_Jet_Management,Private Jet Management
8W,PWF,PRIVATE WINGS,Private_Wings_Flugcharter,Private Wings Flugcharter
P6,PVG,PRIVILEGE,Privilege_Style_L%C3%ADneas_A%C3%A9reas,Privilege Style Líneas Aéreas
,PRH,PROHAWK,Pro_Air,Pro Air
,PSZ,POP-AIR,Pro_Air_Service,Pro Air Service
,GIY,PROBIZ,Probiz_Guinee,Probiz Guinee
,PAD,AIR PROFESSIONAL,Professional_Express_Courier_Service,Professional Express Courier Service
,PVL,VOLARE,Professione_VOlare,Professione VOlare
P0,PFZ,PROFLIGHT-ZAMBIA,Proflight_Zambia,Proflight Zambia
,PTT,TOTOLAPA,Promotora_Industria_Totolapa,Promotora Industria Totolapa
,PRO,PROPAIR,Propair,Propair
,PPA,AIR PROP,Propheter_Aviation,Propheter Aviation
,PTH,PROTEUS,Proteus_Helicopteres,Proteus Helicopteres
,PTL,PLANTATION,Providence_Airline,Providence Airline
,AWD,,Providence_Aviation_Services,Providence Aviation Services
,SPR,SPEEDAIR,Provincial_Airlines,Provincial Airlines
,PRV,PROVINCIAL,Provincial_Express,Provincial Express
,PSW,PSKOVAVIA,Pskovavia,Pskovavia
,UDA,UDARA,Psudiklat_Perhubungan_Udara/PLP,Psudiklat Perhubungan Udara/PLP
,PTA,PTARMIGAN,Ptarmigan_Airways,Ptarmigan Airways
,PSP,PUBLISERVICIOS,Publiservicios_AÃ©reos,Publiservicios Aéreos
,PUV,PUBLIVOO,Publivoo,Publivoo
,PNG,,Puerto_Rico_National_Guard,Puerto Rico National Guard
,TXV,TAXIVALLARTA,Puerto_Vallarta_Taxi_AÃ©reo,Puerto Vallarta Taxi Aéreo
,PGH,,Pulkovo_Aircraft_Services,Pulkovo Aircraft Services
,PLY,PUMA BRASIL,Puma_Linhas_AÃ©reas,Puma Linhas Aéreas
,PTV,PUNTAVIA,Puntavia_Air_Services,Puntavia Air Services
,MGO,MANGO,Punto_Fa,Punto Fa
,PYR,PYAIR,Pyramid_Air_Lines,Pyramid Air Lines
FV,PLK,PULKOVO,Pulkovo_Aviation_Enterprise,Pulkovo Aviation Enterprise
,PRI,PRIMERA,Primera_Air,Primera Air Scandinavia
,PRW,JETBIRD,Primera_Air_Nordic,Primera Air Nordic
,FQA,QUIK LIFT,Quikjet_Cargo_Airlines,Quikjet Cargo Airlines
,QQE,,Qatar_Executive,Qatar Executive
,QNT,QANAT SHARQ,Qanot_Sharq,Qanot Sharq
QF,QFA,QANTAS,Qantas,Qantas
QF,QLK,QLINK,Qantaslink,Qantaslink
QF,QJE,QJET,Qantaslink,Qantaslink
,QAC,QATAR CARGO,Qatar_Air_Cargo,Qatar Air Cargo
QR,QTR,QATARI,Qatar_Airways,Qatar Airways
,QAF,AMIRI,Qatar_Amiri_Flight,Qatar Amiri Flight
,IRQ,QESHM AIR,Qeshm_Air,Qeshm Air
,QTX,AIR QUANTEX,Quantex_Environmental,Quantex Environmental
,QUE,QUEBEC,Quebec_Government_Air_Service,Quebec Government Air Service
,QNA,QUEEN AIR,Queen_Air,Queen Air
,LBQ,LABQUEST,Quest_Diagnostics,Quest Diagnostics
,QAJ,DAGOBERT,Quick_Air_Jet_Charter,Quick Air Jet Charter
,QAH,QUICK,Quick_Airways_Holland,Quick Airways Holland
,QAS,QUISQUEYA,Quisqueya_Airlines,Quisqueya Airlines
,QAQ,QURINEA AIR,Qurinea_Air_Service,Qurinea Air Service
,QCC,QWEST AIR,Qwest_Commuter_Corporation,Qwest Commuter Corporation
,QWA,,Qwestair,Qwestair
,QWL,Q-CHARTER,Qwila_Air,Qwila Air
,RVF,RAVEN FLIGHT,Ravn_Alaska,Ravn Alaska
,RGV,,RG_Aviation,RG Aviation
,WHH,,Richy_Skylark,Richy Skylark
,KRS,,Rosen_Aviation,Rosen Aviation
,VCB,,Real_Aero_Club_de_Vizcaya,Real Aero Club de Vizcaya
,RLH,SENDI,Ruili_Airlines,Ruili Airlines
,RIX,RECTRIX,Rectrix_Aviation,Rectrix Aviation
,RTO,RACCOON,Rectimo_Air_Transports,Rectimo Air Transports
,RUT,YADID,Reut_Airways,Reut Airways
,WES,WEST INDIAN,Rainbow_International_Airlines,Rainbow International Airlines
,VRD,REDWOOD,Virgin_America,Virgin America
,RGE,REGENT,Regent_Airways,Regent Airways
,RJT,RA JET,RA_Jet_Aeroservicios,RA Jet Aeroservicios
R6,,,RACSA_(airline),RACSA
,BKH,,RAF_Barkston_Heath,RAF Barkston Heath
,CFN,CHURCH FENTON,RAF_Church_Fenton,RAF Church Fenton
,COH,COLT,RAF_Coltishall,RAF Coltishall
,CBY,TYPHOON,RAF_Coningsby,RAF Coningsby
,COT,COTTESMORE,RAF_Cottesmore,RAF Cottesmore
,CWL,CRANWELL,RAF_Cranwell,RAF Cranwell
,KIN,KINLOSS,RAF_Kinloss,RAF Kinloss
,LEE,JAVELIN,RAF_Leeming,RAF Leeming
,LCS,LEUCHARS,RAF_Leuchars,RAF Leuchars
,LOP,LINTON ON OUSE,RAF_Linton-on-Ouse,RAF Linton-on-Ouse
,LOS,LOSSIE,RAF_Lossiemouth,RAF Lossiemouth
,MRH,MARHAM,RAF_Marham,RAF Marham
,NWO,,RAF_Northwood,RAF Northwood
,SMZ,SCAMPTON,RAF_Scampton,RAF Scampton
,STN,SAINT ATHAN,RAF_St_Athan,RAF St Athan
,SMG,,RAF_St_Mawgan,RAF St Mawgan Search and Rescue
,TOF,TOPCLIFFE,RAF_Topcliffe,RAF Topcliffe Flying Training Unit
,VYT,ANGLESEY,RAF_Valley,RAF Valley Flying Training Unit
,VLL,,RAF_Valley,RAF Valley SAR Training Unit
,WAD,VULCAN,RAF_Waddington,RAF Waddington
,WIT,STRIKER,RAF_Wittering,RAF Wittering
,MTL,MITAVIA,RAF-Avia,RAF-Avia
,RKM,RAKAIR,RAK_Airways,RAK Airways
,RWL,RHEINTRAINER,RWL_Luftfahrtgesellschaft,RWL Luftfahrtgesellschaft
,RBB,RABBIT,Rabbit-Air,Rabbit-Air
,ACE,Fastcargo,Race_Cargo_Airlines,Race Cargo Airlines
,GBR,GREENBRIER AIR,Rader_Aviation,Rader Aviation
1D,,,Radixx,Radixx
,RAJ,RAJI,Raji_Airlines,Raji Airlines
,RFA,RALEIGH SERVICE,Raleigh_Flying_Service,Raleigh Flying Service
,REX,RAM EXPRESS,Ram_Air_Freight,Ram Air Freight
,RMT,RAM FLIGHT,Ram_Aircraft_Corporation,Ram Aircraft Corporation
,PPK,PELICAN,Ramp_66,Ramp 66
,RGM,RANGEMILE,Rangemile_Limited,Rangemile Limited
,MWR,RASLAN,Raslan_Air_Service,Raslan Air Service
,RAQ,RATH AVIATION,Rath_Aviation,Rath Aviation
,CSM,LORRY,Ratkhan_Air,Ratkhan Air
,RVR,RAVEN,Raven_Air,Raven Air
,RVN,RAVEN U-S,Raven_Air,Raven Air
,REI,RAY AVIATION,Ray_Aviation,Ray Aviation
,RYT,,Raya_Jet,Raya Jet
,RTN,RAYTHEON,Raytheon_Aircraft_Company,Raytheon Aircraft Company
,RCJ,NEWPIN,Raytheon_Corporate_Jets,Raytheon Corporate Jets
,KSS,KANSAS,Raytheon_Travel_Air,Raytheon Travel Air
,RCB,BALEARES,Real_Aero_Club_De_Baleares,Real Aero Club De Baleares
,CDT,AEROREUS,Real_Aero_Club_de_Reus-Costa_Dorado,Real Aero Club de Reus-Costa Dorado
,RCD,AEROCLUB,Real_Aeroclub_De_Ternerife,Real Aeroclub De Ternerife
,RLV,REAL,Real_Aviation,Real Aviation
,REB,REBUS,Rebus,Rebus
,PSH,PASSION,Red_Aviation,Red Aviation
,RBN,RED BARON,Red_Baron_Aviation,Red Baron Aviation
,DEV,RED DEVILS,Red_Devils_Parachute_Display_Team,Red Devils Parachute Display Team
,RDV,RED AVIATION,Red_Sea_Aviation,Red Sea Aviation
,RSV,RED SKY,Red_Sky_Ventures,Red Sky Ventures
,STR,STARLINE,Red_Star,Red Star
8L,RHC,REDAIR,Redhill_Aviation,Redhill Aviation
,RAV,REED AVIATION,Reed_Aviation,Reed Aviation
,REF,REEF AIR,Reef_Air,Reef Air
V4,REK,REEM AIR,Reem_Air,Reem Air
,RVV,REEVE,Reeve_Aleutian_Airways,Reeve Aleutian Airways
,RBH,CALYPSO,Regal_Bahamas_International_Airways,Regal Bahamas International Airways
,RGY,REGENCY,Regency_Airlines,Regency Airlines
,RAH,REGENT,Regent_Air,Regent Air
,RAG,GERMAN LINK,Regio_Air,Regio Air
,RGR,REGIONAIR,Region_Air,Region Air
,TSH,TRANSCANADA,Regional_1,Regional 1
,RIL,,Regional_Air,Regional Air
,REW,REGIONAL WINGS,Regional_Air_Express,Regional Air Express
,REG,REGIONAL SERVICES,Regional_Air_Services_(Tanzania),Regional Air Services
FN,RGL,MAROC REGIONAL,Regional_Air_Lines,Regional Air Lines
ZL,RXA,REX,Regional_Express_Airlines,Regional Express
,JJM,GEODATA,Regional_Geodata_Air,Regional Geodata Air
P7,REP,REGIOPAR,Regional_Paraguaya,Regional Paraguaya
3C,CEA,CORP-X,RegionsAir,RegionsAir
,REL,RELIANCE AIR,Reliance_Aviation,Reliance Aviation
,RLT,RELIANT,Reliant_Airlines,Reliant Airlines
,RTS,RELIEF,Relief_Transport_Services,Relief Transport Services
,RAN,RENAN,Renan_Airways,Renan
QQ,ROA,RENO AIR,Reno_Air,Reno Air
,RGS,RENOWN,Renown_Aviation,Renown Aviation
RW,RPA,BRICKYARD,Republic_Airlines,Republic Airlines
RH,RPH,PUBLIC EXPRESS,Republic_Express_Airlines,Republic Express Airlines
,RBC,REPUBLICAIR,Republicair,Republicair
,RST,RESORT AIR,Resort_Air,Resort Air
,RDS,RHOADES EXPRESS,Rhoades_Aviation,Rhoades Aviation
,RIU,RIAU AIR,Riau_Airlines,Riau Airlines
,RIA,RICHAIR,Rich_International_Airways,Rich International Airways
,RVC,RIVER CITY,Richards_Aviation,Richards Aviation
,RIC,RICHARDSON,Richardson's_Airway,Richardson's Airway
,RCA,RICHLAND,Richland_Aviation,Richland Aviation
,HPR,HELIPRO,Rick_Lucas_Helicopters,Rick Lucas Helicopters
C7,RLE,RICO,Rico_Linhas_A%C3%A9reas,Rico Linhas Aéreas
,RID,AKRID,Ridder_Avia,Ridder Avia
,RAK,SPORT CLUB,Riga_Airclub,Riga Airclub
,RAZ,RIJNMOND,Rijnmond_Air_Services,Rijnmond Air Services
,POL,,Rikspolisstyrelsen,Rikspolisstyrelsen
,RIM,RIMROCK,Rimrock_Airlines,Rimrock Airlines
,SKA,RIO EXPRESS,Rio_Air_Express,Rio Air Express
,REO,RIO,Rio_Airways,Rio Airways
E2,GRN,GRANDE,Rio_Grande_Air,Rio Grande Air
RL,RIO,RIO,Rio_Linhas_A%C3%A9reas,Rio Linhas Aéreas
SL,RSL,RIO SUL,Rio_Sul_Servi%C3%A7os_A%C3%A9reos_Regionais,Rio Sul Serviços Aéreos Regionais
,RVM,RIVER,River_Ministries_Air_Charter,River Ministries Air Charter
,RGP,GARDEN CITY,River_State_Government_of_Nigeria,River State Government of Nigeria
,UNR,RIVNE UNIVERSAL,Rivne_Universal_Avia,Rivne Universal Avia
,RDL,ROADAIR,Roadair_Lines,Roadair Lines
,RBT,ROBIN,Robinton_Aero,Robinton Aero
V2,RBY,RUBY,Vision_Airlines,Vision Airlines
,ROX,ROBLEX,Roblex_Aviation,Roblex Aviation
,RKW,ROCKWELL,Rockwell_Collins_Avionics,Rockwell Collins Avionics
,ROC,,Rocky_Mountain_Airlines,Rocky Mountain Airlines
,RMA,ROCKY MOUNTAIN,Rocky_Mountain_Airways,Rocky Mountain Airways
,LIF,LIFECARE,Rocky_Mountain_Holdings,Rocky Mountain Holdings
,RDZ,RODZE AIR,Rodze_Air,Rodze Air
,FAD,AIR FRONTIER,Rog-Air,Rog-Air
,RRZ,ROLLRIGHT,Rollright_Aviation,Rollright Aviation
,RRL,MERLIN,Rolls-Royce_Limited,Rolls-Royce Limited
,BTU,ROLLS,Rolls-Royce_plc,Rolls-Royce plc
,ROF,ROMAF,Romanian_Air_Force,Romanian Air Force
,RMV,AEROMAVIA,Romavia,Romavia
,RNS,RONSO,Ronso,Ronso
,ROR,RORAIMA,Roraima_Airways,Roraima Airways
,RNB,ROSBALT,Rosneft-Baltika,Rosneft-Baltika
,NRG,ENERGY,Ross_Aviation,Ross Aviation
,RFS,,Rossair_(Australia),Rossair
,RSS,ROSS CHARTER,Rossair_(South_Africa),Rossair
,ROS,CATCHER,Rossair_Europe,Rossair Europe
FV,SDM,RUSSIA,Rossiya_(airline),Rossiya
,RAL,ROSWELL,Roswell_Airlines,Roswell Airlines
GZ,RAR,AIR RAROTONGA,Air_Rarotonga,Air Rarotonga
,RTR,ROTATUR,Rotatur,Rotatur
,RKT,ROCKET,Rotormotion,Rotormotion
,JCR,ROTTERDAM JETCENTER,Rotterdam_Jet_Center,Rotterdam Jet Center
,ROV,ROVERAIR,Rover_Airways_International,Rover Airways International
,VOS,ROVOS,Rovos_Air,Rovos Air
,RCG,ROYAL CARGO,Royal_Air_Cargo,Royal Air Cargo
RR,RFR,RAFAIR,Royal_Air_Force,Royal Air Force
RS,MJN,MAJAN,Royal_Air_Force_of_Oman,Royal Air Force of Oman
,ACW,AIR CADET,Royal_Air_Force,Royal Air Force
,RRR,ASCOT,Royal_Air_Force,Royal Air Force
,RRF,KITTY,Royal_Air_Force,Royal Air Force
,SHF,VORTEX,Royal_Air_Force,Royal Air Force
,RAX,AIR ROYAL,Royal_Air_Freight,Royal Air Freight
AT,RAM,ROYALAIR MAROC,Royal_Air_Maroc,Royal Air Maroc
R0,RPK,ROYAL PAKISTAN,Royal_Airlines,Royal Airlines
,RLM,ROYAL AMERICAN,Royal_American_Airways,Royal American Airways
V5,RYL,ROYAL ARUBAN,Royal_Aruban_Airlines,Royal Aruban Airlines
,ASY,AUSSIE,Royal_Australian_Air_Force,Royal Australian Air Force
,RXP,ROY EXPRESS,Royal_Aviation_Express,Royal Aviation Express
,RYB,ROYAL BAHRAIN,Royal_Bahrain_Airlines,Royal Bahrain Airlines
BI,RBA,BRUNEI,Royal_Brunei_Airlines,Royal Brunei Airlines
,KDR,DARLINES,Royal_Daisy_Airlines,Royal Daisy Airlines
,RGA,ROYAL GHANA,Royal_Ghanaian_Airlines,Royal Ghanaian Airlines
,ROJ,ROYALJET,Royal_Jet,Royal Jet
RJ,RJA,JORDANIAN,Royal_Jordanian,Royal Jordanian
,RJZ,JORDAN AIR FORCE,Royal_Jordanian_Air_Force,Royal Jordanian Air Force
RK,RCT,GREENSKY,Skyview_Airways,Skyview Airways
RK,RKH,KHMER AIR,Royal_Khmer_Airlines,Royal Khmer Airlines
,RMF,ANGKASA,Royal_Malaysian_Air_Force,Royal Malaysian Air Force
,NVY,NAVY,Royal_Navy,Royal Navy
,NRN,NETHERLANDS NAVY,Royal_Netherland_Navy,Royal Netherland Navy
,NAF,NETHERLANDS AIR FORCE,Royal_Netherlands_Air_Force,Royal Netherlands Air Force
,KIW,KIWI,Royal_New_Zealand_Air_Force,Royal New Zealand Air Force
,NOW,NORWEGIAN,Royal_Norwegian_Air_Force,Royal Norwegian Air Force
,ROP,,Royal_Oman_Police,Royal Oman Police
RL,PPW,PHNOM-PENH AIR,Royal_Phnom_Penh_Airways,Royal Phnom Penh Airways
,RRA,ROYAL RWANDA,Royal_Rwanda_Airlines,Royal Rwanda Airlines
,RSF,ARSAF,Royal_Saudi_Air_Force,Royal Saudi Air Force
,RYS,ROYAL SKY,Royal_Sky,Royal Sky
,RSN,SWAZI NATIONAL,Royal_Swazi_National_Airways,Royal Swazi National Airways
WR,HRH,TONGA ROYAL,Royal_Tongan_Airlines,Royal Tongan Airlines
,RWE,ROYAL WEST,Royal_West_Airlines,Royal West Airlines
,RSB,RUBYSTAR,Rubystar,Rubystar
,RMG,RUMUGU AIR,Rumugu_Air_&_Space_Nigeria,Rumugu Air & Space Nigeria
,RUR,,Rusaero,Rusaero
,KLE,,Rusaero,Rusaero
,CGI,CGI-RUSAIR,Rusair,Rusair JSAC
,RUH,,Rusich-T,Rusich-T
,RLU,RUSLINE AIR,Rusline,Rusline
,MIG,MIG AVIA,Russian_Aircraft_Corporation-MiG,Russian Aircraft Corporation-MiG
,RFF,RUSSIAN AIRFORCE,Russian_Federation_Air_Force,Russian Federation Air Force
P7,ESL,RADUGA,Russian_Sky_Airlines,Russian Sky Airlines
,RUZ,ROSTUERTOL,Rusuertol,Rusuertol
,RUC,RUTACA,Rutas_A%C3%A9reas,Rutas Aéreas
,RND,RUTLAND,Rutland_Aviation,Rutland Aviation
,RUA,,Rwanda_Airlines,Rwanda Airlines
,RWA,,Rwanda_Airways,Rwanda Airways
WB,RWD,RWANDAIR,Rwandair_Express,Rwandair Express
7S,RCT,ARCTIC TRANSPORT,Ryan_Air_Service,Ryan Air Service
,RYA,RYAN AIR,Ryan_Air_Services,Ryan Air Services
RD,RYN,RYAN INTERNATIONAL,Ryan_International_Airlines,Ryan International Airlines
FR,RYR,RYANAIR,Ryanair,Ryanair
,RYZ,RYAZAN AIR,Ryazan_State_Air_Enterprise,Ryazan State Air Enterprise
,RAA,RYNES AVIATION,Rynes_Aviation,Rynes Aviation
YS,RAE,REGIONAL EUROPE,R%C3%A9gional_Compagnie_A%C3%A9rienne_Europ%C3%A9enne,Régional Compagnie Aérienne Européenne
,REV,ENDURANCE,RVL_Aviation,RVL Group
RT,BUG,,UVT_Aero,UVT Aero
,OMN,SERVIOMNIA,Servicios_Aereos_Ominia,Servicios Aereos Ominia
,SEN,SERVISIERRA,Servicios_de_Aviacion_Sierra,Servicios de Aviacion Sierra
,SGC,SAINT GEORGE,SGC_Aviation,SGC Aviation
,SCJ,SIAMJET,Siamjet_Aviation,Siamjet Aviation
,SIX,DRIVE ORANGE,Sixt_Rent_A_Car,Sixt Rent A Car
,SOG,,Solenta_Aviation_Ghana,Solenta Aviation Ghana
,QSR,SPARKLE ROLL,SR_Jet,SR Jet
SS,CRL,CORSAIR,Corsairfly,Corsairfly
,KBN,KABIN,Spiracha_Aviation,Spiracha Aviation
,CBN,CARBONDALE,Southern_Illinois_University,"Southern Illinois University as ""Aviation Flight"""
,IBG,ICE BRIDGE,Springfield_Air,Springfield Air
,BZQ,STING,Seneca_College,Seneca College
,BVV,SPARC,Sparc_Avia,Sparc Avia
,SJM,SINO SKY,Sino_Jet_Management,Sino Jet Management
,SCH,OCEAN BIRD,Seychelles_Airlines,Seychelles Airlines
,BYF,BAY FLIGHT,San_Carlos_Flight_Center,San Carlos Flight Center
,SXT,SERTAXI,Servicios_de_Taxi_Aereos,Servicios de Taxi Aereos
TR,TGW,SCOOTER,Scoot,Scoot
IJ,SJO,JEY SPRING,Spring_Airlines_Japan,Spring Airlines Japan
,SBD,SIBIA,Sibia_(Airline),SIBIA Aircompany Ltd
6Y,ART,SMART LYNX,Smartlynx_Airlines,Smartlynx Airlines
,MYX,TALLINN CAT,Smartlynx_Airlines_Estonia,Smartlynx Airlines Estonia
,DES,DESTINA,Servicios_Aereos_Especializados_Destina,Servicios Aereos Especializados Destina
,LSV,,,Slovenian Ministry of Defence
,FUF,SERVIFUN,Servicios_Aereos_Fun_Fly,Servicios Aereos Fun Fly
E3,VGO,VIRGO,Sabaidee_Airways,Sabaidee Airways
,SAQ,,Safe_Air_Company,Safe Air Company
,SMU,SPRINGER,Sanborn_Map_Company,Sanborn Map Company
,RBR,SIAM AIRNET,Siam_Airnet,Siam Airnet
,SVB,SIAVIA,Siavia,Siavia
,MHQ,HELICARE,Skargardshavets_Helikoptertjanst,Skargardshavets Helikoptertjanst
,BIS,JUMA AIR,Sky_Bishek,Sky Bishek
,SYH,,Sky_Handling,Sky Handling
GG,KYE,SKY CUBE,Sky_Lease_Cargo,Sky Lease Cargo
,KPM,SKY PRIMAIR,Sky_Prim_Air,Sky Prim Air
,BSJ,SKYBUS JET,Skybus_Jet,Skybus Jet
GW,SGR,SKYGREECE,SkyGreece_Airlines,SkyGreece Airlines
,USW,AKSAR,Special_Aviation_Works,Special Aviation Works
N9,SHA,SHREEAIR,Shree_Airlines,Shree Airlines
7E,AWU,SYLT-AIR,Sylt_Air_GmbH,Sylt Air GmbH
,BDS,SOUTH ASIAN,South_Asian_Airlines,South Asian Airlines
OL,SZB,SAMOA,Samoa_Air,Samoa Air
,BEC,,State_Air_Company_Berkut,State Air Company Berkut
S4,RZO,AIR AZORES,SATA_International,SATA International
SA,SAA,SPRINGBOK,South_African_Airways,South African Airways
,KYD,SKYAD,Sky_Messaging,Sky Messaging
,SAB,SKY WORKER,Sky_Way_Air,Sky Way Air
KV,SKV,MAPLE,Sky_Regional_Airlines,Sky Regional Airlines
,SAC,SASCO,SASCO_Airlines,SASCO Airlines
,SAG,MEDICAL AIR,SOS_Flygambulans,SOS Flygambulans
W7,SAH,SAYAKHAT,Sayakhat_Airlines,Sayakhat Airlines
NL,SAI,SHAHEEN AIR,Shaheen_Air_International,Shaheen Air International
MM,SAM,SAM,SAM_Colombia,SAM Colombia
,SAN,AEREOS,SAN_Ecuador,Servicios Aéreos Nacionales (SAN)
,SAO,SAVSER,Sahel_Aviation_Service,Sahel Aviation Service
,ANX,SECRETARIA DEMARINA,Secretaria_de_Marina,Secretaria de Marina
,SAQ,SPRINGBANK,Springbank_Aviation,Springbank Aviation
SK,SAS,SCANDINAVIAN,Scandinavian_Airlines,Scandinavian Airlines
,SAV,,Samal_Air,Samal Air
,SAW,SHAMWING,Sham_Wing_Airlines,Sham Wing Airlines
,SAX,SABAH AIR,Sabah_Air,Sabah Air
,SAY,SUCKLING,ScotAirways,ScotAirways
,SAZ,SWISS AMBULANCE,Swiss_Air-Ambulance,Swiss Air-Ambulance
,SBA,SOL,Sol_Linhas_A%C3%A9reas,SOL Linhas Aéreas
PI,SGU,SOLPARAGUAYO,Sol_del_Paraguay,Sol del Paraguay
,SBA,STA-MALI,STA-MALI,STA-MALI
,SBB,SABER EXPRESS,Steinman_Aviation,Steinman Aviation
UG,SEN,S-BAR,SevenAir,SevenAir
,SBF,SEVENAIR,Seven_Bar_Flying_Service,Seven Bar Flying Service
,SBG,,Sabre_Incorporated,Sabre Incorporated
S7,SBI,SIBERIAN AIRLINES,S7_Airlines,S7 Airlines
,SBL,SOBGHANA,Sobel_Airlines_of_Ghana,Sobel Airlines of Ghana
Q7,SBM,SKY BAHAMAS,SkyBahamas,SkyBahamas
,SBO,STABAIR,Stabo_Air,Stabo Air
,SBQ,SKIBBLE,SmithKline_Beecham_Clinical_Labs,SmithKline Beecham Clinical Labs
,SBR,FREIGHTER,Saber_Aviation,Saber Aviation
BB,SBS,SEABORNE,Seaborne_Airlines,Seaborne Airlines
PV,SBU,BLACK FIN,St_Barth_Commuter,St Barth Commuter
,URJ,STARAV,Star_Air_Aviation,Star Air
,SBZ,SCIBE AIRLIFT,Scibe_Airlift,Scibe Airlift
,AME,AIRMIL,Spanish_Air_Force,Spanish Air Force
,SCA,SOUTH CENTRAL,South_Central_Air,South Central Air
,SCC,SEA-COASTER,Seacoast_Airlines,Seacoast Airlines
K5,SQH,SASQUATCH,SeaPort_Airlines,SeaPort Airlines
,SCE,SCENIC,Scenic_Airlines,Scenic Airlines
,SCF,SOCOFER,Socofer,Socofer
,SCI,SAN CRISTOBAL,Servicios_AÃ©reos_San_CristÃ³bal,Servicios Aéreos San Cristóbal
,SCK,SKYCAM,Sky_Cam,Sky Cam
W3,SCL,SWIFTAIR,Switfair_Cargo,Switfair Cargo
,SCB,SAIGON,Saigon_Capital_Aircraft_Management,Saigon Capital Aircraft Management
,SCN,SOUTH AMERICAN,South_American_Airlines,South American Airlines
,AHI,AEROCHISA,Servicios_AÃ©reos_de_Chihuahua_Aerochisa,Servicios Aéreos de Chihuahua Aerochisa
FP,AND,SERVI ANDES,Servicios_A%C3%A9reos_de_los_Andes,Servicios Aéreos de los Andes
UL,ALK,SRILANKAN,SriLankan_Airlines,SriLankan Airlines
,CDL,CAROLINA,Sunbird_Airlines,Sunbird Airlines
,SCP,SCORPIO,Scorpio_Aviation,Scorpio Aviation
,SCQ,SCAVAC,OSM_Aviation_Academy,OSM Aviation Academy
,SIC,SICHART,SFS_Aviation,SFS Aviation
,SCR,SILVER CLOUD,Silver_Cloud_Air,Silver Cloud Air
,SCS,SOUTHERN CHARTERS,South_African_Non_Scheduled_Airways,South African Non Scheduled Airways
,SCT,SAAB-CRAFT,SAAB-Aircraft,SAAB-Aircraft
,SCV,SACSA,Servicios_AÃ©reos_Del_Centro,Servicios Aéreos Del Centro
SY,SCX,SUN COUNTRY,Sun_Country_Airlines,Sun Country Airlines
,SDA,SAINT ANDREWS,St._Andrews_Airways,St. Andrews Airways
,SDB,SU-CRAFT,Sukhoi_Design_Bureau_Company,Sukhoi Design Bureau Company
,SDC,SUNDANCE,Sunrise_Airlines,Sunrise Airlines
,SDD,SKY DANCE,Skymaster_Air_Taxi,Skymaster Air Taxi
,SDE,STAMPEDE,Air_Partners_Corp.,Air Partners Corp.
,SDF,SUNDORPH,Sundorph_Aeronautical_Corporation,Sundorph Aeronautical Corporation
,SDH,ARCOS,Servicio_De_Helicopteros,Servicio De Helicopteros
,SDK,SADELCA,SADELCA,SADELCA - Sociedad Aérea Del Caquetá
,SDL,SKYDRIFT,Skydrift,Skydrift
,SDN,BLUE NILE,Spirit_of_Africa_Airlines,Spirit of Africa Airlines
,SDU,SUD LINES,Sud_Airlines,Sud Airlines
,SDV,SELVA,Servicios_AÃ©reos_Del_Vaupes,Servicios Aéreos Del Vaupes
,SDX,SERVICIO TECNICO,Servicio_Tecnico_Aero_De_Mexico,Servicio Tecnico Aero De Mexico
,SDZ,SUDANA,Sudan_Pezetel_for_Aviation,Sudan Pezetel for Aviation
,SEA,SOUTHEAST AIR,Southeast_Air,Southeast Air
,SEB,SERVILUCE,Servicios_AÃ©reos_Luce,Servicios Aéreos Luce
,SED,SEDONA AIR,Sedona_Air_Center,Sedona Air Center
,SEE,SHAHEEN CARGO,Shaheen_Air_Cargo,Shaheen Air Cargo
,SEH,AIR CRETE,Sky_Express,Sky Express
SG,SEJ,SPICEJET,Spicejet,Spicejet
,SEK,SKALA,Skyjet_(Kazakhstani_airline),Skyjet
,SEL,SENTEL,Sentel_Corporation,Sentel Corporation
,SEO,SELCON AIR,Selcon_Airlines,Selcon Airlines
I6,SEQ,SKY EYES,Sky_Eyes,Sky Eyes
,SES,SERVISAL,Servicio_AÃ©reo_Saltillo,Servicio Aéreo Saltillo
EH,SET,SAETA,SAETA,SAETA
,SEV,CARGOPRESS,Serair_Transworld_Press,Serair Transworld Press
,SFA,SEFA,SEFA,SEFA
,SFC,SHUSWAP,Shuswap_Flight_Centre,Shuswap Flight Centre
,SFE,SEFOFANE,Sefofane_Air_Charters,Sefofane Air Charters
,SFF,SWIFTWING,Safewings_Aviation_Company,Safewings Aviation Company
,SFG,AERO GULF,Sun_Freight_Logistics,Sun Freight Logistics
7G,SFJ,STARFLYER,Star_Flyer,Star Flyer
,SFL,SOUTHFLIGHT,Southflight_Aviation,Southflight Aviation
,SFN,SAFIRAN,Safiran_Airlines,Safiran Airlines
,SFP,SAFE AIR,Safe_Air,Safe Air
FA,SFR,CARGO,Safair,Safair
,SFS,SOUTHERN FRONTIER,Southern_Frontier_Air_Transport,Southern Frontier Air Transport
,SFT,SKYFREIGHT,Skyfreight,Skyfreight
,SFU,SAINTS,Solent_Flight,Solent Flight
,SFX,SWAMP FOX,S.K._Logistics,S.K. Logistics
,SGB,SONGBIRD,"Sky_King,_Inc.","Sky King, Inc."
,SGC,SOUTHERNRIGHT,Southern_Right_Air_Charter,Southern Right Air Charter
,SGD,AIR BISHKEK,Sky_Gate_International_Aviation,Sky Gate International Aviation
,SGF,STAC,STAC_Swiss_Government_Flights,STAC Swiss Government Flights
,SGH,SERVISAIR,Servisair,Servisair
,SGI,SERAGRI,Servicios_AÃ©reos_AgrÃ­colas,Servicios Aéreos Agrícolas
,SGK,SKYWARD,Skyward_Aviation,Skyward Aviation
,SGM,SIGMA,Sky_Aircraft_Service,Sky Aircraft Service
,SGN,SIAM,SGA_Airlines,Siam GA
,SGP,SAGOLAIR,Sagolair_Transportes_Ejecutivos,Sagolair Transportes Ejecutivos
,SGS,SASKATCHEWAN,Saskatchewan_Government_Executive_Air_Service,Saskatchewan Government Executive Air Service
,SGT,SKYGATE,,Skygate
,SGU,RAUSHAN,Samgau,Samgau
,SGX,,Saga_Airlines,Saga Airlines
N5,SGY,SKAGWAY AIR,Skagway_Air_Service,Skagway Air Service
,SHB,SHABAIR,Shabair,Shabair
,SHC,SKY HARBOR CHEYENNE,Sky_Harbor_Air_Service,Sky Harbor Air Service
,SHD,,Sahara_Airlines_(Algeria),Sahara Airlines
,SHE,SHELL,Shell_Aircraft,Shell Aircraft
,SHG,SHOP AIR,Shoprite_Group,Shoprite Group
,SHJ,SHARJAH,Sharjah_Ruler's_Flight,Sharjah Ruler's Flight
,SHK,,Shorouk_Air,Shorouk Air
,SHL,SAMSON,Samson_Aviation,Samson Aviation
,SHM,SHELTAM,Sheltam_Aviation,Sheltam Aviation
,SHN,SUGAR ALFA,Shaheen_Airport_Services,Shaheen Airport Services
,SHO,,Sheremetyevo-Cargo,Sheremetyevo-Cargo
,SHP,SAF,Service_Aerien_Francais,Service Aerien Francais
,SHQ,SHANGHAI CARGO,Shanghai_Airlines_Cargo,Shanghai Airlines Cargo
,SHR,SHOOTER,Shooter_Air_Courier,Shooter Air Courier
,SHS,SHURA AIR,Shura_Air_Transport_Services,Shura Air Transport Services
HZ,SHU,SATAIR,Sakhalinskie_Aviatrassy,Sakhalinskie Aviatrassy (SAT)
SP,SAT,SATA,SATA_Air_Acores,SATA Air Acores
8S,,,Scorpio_Aviation,Scorpio Aviation
,SHV,SHAVANO,Shavano_Air,Shavano Air
,SHW,SHAWNEE,Shawnee_Airline,Shawnee Airline
,SHX,SLIM AIR,Slim_Aviation_Services,Slim Aviation Services
ZY,SHY,ANTALYA BIRD,Sky_Airlines,Sky Airlines
SQ,SIA,SINGAPORE,Singapore_Airlines,Singapore Airlines
5M,SIB,SIBAVIA,Sibaviatrans,Sibaviatrans
,SIE,SEREX,Sierra_Express,Sierra Express
SI,SIH,BLUEJET,Skynet_Airlines,Skynet Airlines
,SIJ,,Seco_International,Seco International
,SIL,SERVICIOS INTEGRALES,Servicios_AeronÃ¡uticos_Integrales,Servicios Aeronáuticos Integrales
,SIM,,Star_Air_(Sierra_Leone),Star Air
,SIO,SIRIO,Sirio,Sirio
,SIR,SALAIR,Salair_(airline),Salair
,SIS,,Saber_Airlines,Saber Airlines
XS,SIT,,SITA_(IT_company),SITA
,SIV,SLOVENIAN,Slovenian_Armed_Forces,Slovenian Armed Forces
,SIW,SIRIO EXECUTIVE,Sirio_Executive,Sirio Executive
,SJA,SERVICIOJAL,Servicios_AÃ©reos_Especiales_De_Jalisco,Servicios Aéreos Especiales De Jalisco
,SJC,SERVIEJECUTIVO,Servicios_Ejecutivos_Continental,Servicios Ejecutivos Continental
,SJE,SUNBIZ,Sunair_2001,Sunair 2001
,SJJ,SPIRIT JET,Spirit_Aviation,Spirit Aviation
,SJL,SERVICIOS JALISCO,Servicios_Especiales_Del_Pacifico_Jalisco,Servicios Especiales Del Pacifico Jalisco
,SJT,SWISS JET,Swiss_Jet,Swiss Jet
,SJY,SRIWIJAYA,Sriwijaya_Air,Sriwijaya Air
ZS,SMY,NAJIM,Sama_Airlines,Sama Airlines
,ALC,ACOM,"Southern_Jersey_Airways,_Inc.","Southern Jersey Airways, Inc."
,SPS,SALDUERO,SPASA,SPASA
,SPT,SPEED AVIATION,Speed_Aviation,Speed Aviation
,SPU,SPUTTER,Southeast_Airmotive,Southeast Airmotive
,SPV,SERVICIOS PRIVADOS,Servicios_Privados_De_AviaciÃ³n,Servicios Privados De Aviación
,SPW,SPEEDWING,Speedwings,Speedwings
,SPX,,Service_People_Gesellschaft_fÃ¼r_Charter_und_Service,Service People Gesellschaft für Charter und Service
,SQA,SLOVAK AEROCLUB,Slovak_National_Aeroclub,Slovak National Aeroclub
SQ,SQC,SINGCARGO,Singapore_Airlines_Cargo,Singapore Airlines Cargo
,SQF,SLOVAK AIRFORCE,Slovak_Air_Force,Slovak Air Force
,SQL,ALQUILER,Servicos_De_Alquiler,Servicos De Alquiler
,SRA,SAIR,Sair_Aviation,Sair Aviation
,SRC,SEARCA,Searca,Searca
FT,SRH,SIEMREAP AIR,Siem_Reap_Airways,Siem Reap Airways
SX,SRK,SKYFOX,Sky_Work_Airlines,Sky Work Airlines
SM,SRL,Starline,Swedline_Express,Swedline Express
,SRL,SERVICIOS PERSONAL,Servicios_AeronÃ¡uticos_Aero_Personal,Servicios Aeronáuticos Aero Personal
,SRN,SIRAIR,Sirair,Sirair
,SRO,SAEREO,Servicios_AÃ©reos_Ejecutivos_Saereo,Servicios Aéreos Ejecutivos Saereo
,SRQ,SEAIR,South_East_Asian_Airlines,South East Asian Airlines
S6,SRR,WHITESTAR,Star_Air_(Maersk_Air),Star Air
,SRS,PHOTO CHARLIE,Selkirk_Remote_Sensing,Selkirk Remote Sensing
,SRU,STAR-UP,Star_Up,Star Up
,SRW,SARIA,Sarit_Air_Lines,Sarit Airlines
,SRX,SIERRA EX,Sierra_Expressway_Airlines,Sierra Expressway Airlines
,SRZ,STRATO,Strato_Air_Services,Strato Air Services
,SSB,SASIR,Sasair,Sasair
,SSC,SOUTHERN SKIES,Southern_Seaplane,Southern Seaplane
,SSD,STAR SERVICE,Star_Service_International,Star Service International
,SSE,SUNSET,Servicios_AÃ©reos_Sunset,Servicios Aéreos Sunset
D2,SSF,SEVERSTAL,Severstal_Air_Company,Severstal Air Company
,SSG,SLOVAK GOVERNMENT,Slovak_Government_Flying_Service,Slovak Government Flying Service
VD,BBB,BLACKBIRD,SwedJet_Airways,SwedJet Airways
,SSK,SKYSTAR,Skystar_International,Skystar International
,SSO,DOPE,Special_Scope_Limited,Special Scope Limited
,SSP,STARSPEED,Starspeed,Starspeed
QF,SSQ,SUNSTATE,Sunstate_Airlines,Sunstate Airlines
,SSS,SAESA,SAESA,SAESA
,SST,SUNFLIGHT,Sunwest_Airlines,Sunwest Airlines
,SSU,SASCA,SASCA,SASCA
5G,SSV,SKYTOUR,Skyservice_Airlines,Skyservice Airlines
,SSW,STREAMLINE,Streamline_Aviation,Streamline Aviation
,SSY,SIERRA SKY,Sky_Aviation_(Sierra_Leone),Sky Aviation
,SSZ,SPECSAVERS,Specsavers_Aviation,Specsavers Aviation
,STA,STAR,Star_Aviation,Star Aviation
,STB,STATUS-ALPHA,Status-Alpha_Airline,Status-Alpha Airline
,STC,STADIUM,Stadium_City_Limited,Stadium City Limited
,STD,AERO AGUASCALINETES,Servicios_De_Aerotransportacion_De_Aguascalientes,Servicios De Aerotransportacion De Aguascalientes
,STE,SEMITRANS,Semitool_Europe,Semitool Europe
,STF,,SFT-Sudanese_Flight,SFT-Sudanese Flight
,STG,STAGE,"Sedalia,_Marshall,_Boonville_Stage_Line","Sedalia, Marshall, Boonville Stage Line"
,STH,,South-Airlines,South-Airlines
,STI,SONTAIR,Sontair,Sontair
,STJ,STELLAVIA,Sella_Aviation,Sella Aviation
RE,STK,STOBART,Stobart_Air,Stobart Air
,STL,STAPLEFORD,Stapleford_Flight_Centre,Stapleford Flight Centre
,STO,SLOPS,Streamline_Ops,Streamline Ops
,STQ,STERA,Star_Air_(Indonesia),Star Air
FS,STU,FUEGUINO,Servicios_de_Transportes_A%C3%A9reos_Fueguinos,Servicios de Transportes Aéreos Fueguinos
,STU,STARSOM,Star_African_Air,Star African Air
,SUU,SUNSTAR,Star_West_Aviation,Star West Aviation
,STV,SOUTHERN AVIATION,Southern_Aviation,Southern Aviation
,STW,SIERRA WHISKEY,South_West_Air_Corporation,South West Air Corporation
,STX,STARSAWAY,Stars_Away_Aviation,Stars Away Aviation
,STY,STYRIAN,Styrian_Airways,Styrian Airways
,SUA,AIR SILESIA,Silesia_Air,Silesia Air
,SUB,SUB AIR,,Suburban Air Freight
SD,SUD,SUDANAIR,Sudan_Airways,Sudan Airways
PI,SUF,SUNFLOWER,Fiji_Link,Sun Air (Fiji)
LW,FDY,FRIENDLY,Sun_Air_International,Sun Air International
,SUG,SUNU AIR,Sunu_Air,Sunu Air
,SUH,LIGHT AIR,Sun_Light,Sun Light
,SUI,SWISS AIR FORCE,Swiss_Air_Force,Swiss Air Force
,SUK,SKYCARGO,Superior_Aviation_Services,Superior Aviation Services
,SUM,SUMES,State_Unitary_Air_Enterprise,State Unitary Air Enterprise
,SUR,,Sun_Air_(Egypt),Sun Air
EZ,SUS,SUNSCAN,Sun_Air_of_Scandinavia,Sun Air of Scandinavia
,URF,SURF AIR,Surf_Air,Surf Air
,SUT,SISTEMAS AERONAUTICOS,Sistemas_Aeronauuticos_2000,Sistemas Aeronauuticos 2000
,SUV,DANCEAIR,Sundance_Air,Sundance Air
SV,SVA,SAUDIA,Saudia,Saudia
,SVD,GRENADINES,St._Vincent_Grenadines_Air_(1990),St. Vincent Grenadines Air (1990)
,SVF,SWEDEFORCE,Swedish_Armed_Forces,Swedish Armed Forces
,AWJ,SAHEL AIRLINES,Sahel_Airlines,Sahel Airlines
,SVH,SILVER,Sterling_Helicopters,Sterling Helicopters
,SVI,SETRA,Servicios_De_Transporte_AÃ©reo,Servicios De Transporte Aéreo
,SVJ,,Silver_Air_(Djibouti),Silver Air
,SVL,SEVAVIA,Sevastopol-Avia,Sevastopol-Avia
,SVN,SAVANAIR,Savanair_(Angola),Savanair (Angola)
,SVO,SERVIORIENTE,Servicios_Aeron%C3%A1uticos_De_Oriente,Servicios Aeronáuticos De Oriente
,SVS,AEREOS SAAR,Servicios_AÃ©reos_Saar,Servicios Aéreos Saar
,SVT,SIERRA SERVICES,Seven_Four_Eight_Air_Services,Seven Four Eight Air Services
,SVX,SECURITY AIR,Security_Aviation,Security Aviation
WN,SWA,SOUTHWEST,Southwest_Airlines,Southwest Airlines
,SWB,SWISSBOOGIE,Swissboogie_Parapro,Swissboogie Parapro
,SWC,SAINT CLAIR,South_West_Air,South West Air
A4,SWD,SOUTHERN WINDS,Southern_Winds_Airlines,Southern Winds Airlines
,SWE,SWEDELINE,Swedeways,Swedeways
,,AIR SEATTLE,Spurling_Aviation,Spurling Aviation
WG,SWG,SUNWING,Sunwing_Airlines,Sunwing Airlines
,SWI,SUNWORLD,Sunworld_Airlines,Sunworld Airlines
,SWJ,STATES,StatesWest_Airlines,StatesWest Airlines
,SWO,SIVA,Surinam_International_Victory_Airline,Surinam International Victory Airline
,SWP,STAR WORK,,Star Work Sky
,SWQ,SWIFTFLIGHT,Swift_Air,Swift Air (Interstate Equipment Leasing)
LX,SWR,SWISS,Swiss_International_Air_Lines,Swiss International Air Lines
SR,SWR,SWISSAIR,Swissair,Swissair
SR,SDR,SUNDAIR,Sundair,Sundair
,SWS,SUNNY WEST,Sunwest_Aviation,Sunwest Aviation (Lindquist Investment)
,SWT,SWIFT,Swiftair,Swiftair
LZ[45],SWU,EUROSWISS,Swiss_Global_Air_Lines,Swiss Global Air Lines
WV,SWV,FLYING SWEDE,Swe_Fly,Swe Fly
S8,SWW,WAY AERO,Shovkoviy_Shlyah,Shovkoviy Shlyah
Q4,SWX,SWAZI EXPRESS,Swazi_Express_Airways,Swazi Express Airways
WO,WSW,SWOOP,Swoop_(airline),Swoop
,SWY,SWISSLINK,Skyjet_(Swiss_airline),Sky Jet
,SWZ,SWISSBIRD,"Servair,_Private_Charter","Servair, Private Charter"
S8,SWZ,SKYWISE,Skywise_Airlines,Skywise Airline
,SXA,FERRY,Southern_Cross_Aviation,Southern Cross Aviation
,SXC,SKY EXEC,Sky_Exec_Aviation_Services,Sky Exec Aviation Services
,SXE,DOGWOOD EXPRESS,Southeast_Express_Airlines,Southeast Express Airlines
,SXM,SERVIMEX,Servicios_AÃ©reos_Especializados_Mexicanos,Servicios Aéreos Especializados Mexicanos
XQ,SXS,SUNEXPRESS,SunExpress,SunExpress
,SXT,SERTA,Servicios_De_Taxi_AÃ©reo,Servicios De Taxi Aéreo
,SXX,SATELLITE EXPRESS,Satellite_Aero,Satellite Aero
,SXY,SAFARI EXPRESS,Safari_Express_Cargo,Safari Express Cargo
,SYA,LINEAS CARDINAL,Skyways_(Argentinian_airline),Skyways
,SYC,SYSTEC,Systec_2000,Systec 2000
,SYE,,Sheba_Aviation,Sheba Aviation
,SYF,SKY FIRST,Sky_One_Express_Airlines,Sky One Express Airlines
,SYG,SYNERGY,Synergy_Aviation,Synergy Aviation
,SYI,,Sonalysts,Sonalysts
,SYJ,,Slate_Falls_Airways,Slate Falls Airways
,SYK,AEROCAB,Satsair,Satsair
,SYN,SYNCRUDE,Syncrude_Canada,Syncrude Canada
RB,SYR,SYRIANAIR,Syrian_Arab_Airlines,Syrian Arab Airlines
,SYS,SHAWBURY,RAF_Shawbury,Shawbury Flying Training Unit
,SYV,SPECIAL SYSTEM,Special_Aviation_Systems,Special Aviation Systems
AL,SYX,SKYWAY-EX,Skywalk_Airlines,Skywalk Airlines
ZP,AZQ,SILK LINE,Silk_Way_Airlines,Silk Way Airlines
7L,AZG,SILK WEST,Silk_Way_West_Airlines,Silk Way West Airlines
,SYY,SKY COACH,South_African_Historic_Flight,South African Historic Flight
,SZT,AERO ZEE,Servicios_AeronÃ¡uticos_Z,Servicios Aeronáuticos Z
,BHV,AVIASPEC,Specavia_Air_Company,Specavia Air Company
,BLY,BLARNEY,Starair,Starair
,BNC,BARNACLE AIR,Sundance_Air,Sundance Air
E5,BRZ,BERYOZA,Samara_Airlines,Samara Airlines
,CBN,CALIBRATION,Swedish_Civil_Aviation_Administration,Swedish Civil Aviation Administration
SC,CDG,SHANDONG,Shandong_Airlines,Shandong Airlines
,CDS,SPECDAS,Spectrem_Air,Spectrem Air
,CEE,CENTRA AEREOS,Servicios_AÃ©reos_Centrales,Servicios Aéreos Centrales
,CFL,SWEDISH,Swedish_Airlines,Swedish Airlines
,CGL,SEAGLE,Seagle_Air,Seagle Air
,CIG,SIRIUS AERO,Sirius-Aero,Sirius-Aero
,CNK,CHINOOK,Sunwest_Home_Aviation,Sunwest Home Aviation
SK,CNO,SCANOR,SAS_Braathens,SAS Braathens
9C,CQH,AIR SPRING,Spring_Airlines,Spring Airlines
3U,CSC,SI CHUAN,Sichuan_Airlines,Sichuan Airlines
FM,CSH,SHANGHAI AIR,Shanghai_Airlines,Shanghai Airlines
,CSY,SHUANGYANG,Shuangyang_General_Aviation,Shuangyang General Aviation
ZH,CSZ,SHENZHEN AIR,Shenzhen_Airlines,Shenzhen Airlines
8C,CXI,SHANXI,Shanxi_Airlines,Shanxi Airlines
,DKT,DAKOTA,Sioux_Falls_Aviation,Sioux Falls Aviation
,DKY,DAKOY,Servicios_AÃ©reos_Elite,Servicios Aéreos Elite
,DNI,AERO DENIM,Servicios_AÃ©reos_Denim,Servicios Aéreos Denim
,EAB,SWISS EAGLE,Swiss_Eagle,Swiss Eagle
,EAN,NIGERIA EXPRESS,Skypower_Express_Airways,Skypower Express Airways
,ENR,,Scenic_Air,Scenic Air
7L,ERO,ECHO ROMEO,Sun_D%27Or,Sun D'Or
NE,ESK,RELAX,SkyEurope,SkyEurope
CQ,EXL,,Sunshine_Express_Airlines,Sunshine Express Airlines
,EXY,EXPRESSWAYS,South_African_Express,South African Express
,FFD,FIRST FLIGHT,Stuttgarter_Flugdienst,Stuttgarter Flugdienst
,FFH,PEACE AIR,Shalom_Air_Services,Shalom Air Services
,FJE,ENVOY,Silverjet,Silverjet
,FLH,MILE HIGH,Sky_Bus,Sky Bus
,GAD,SOUTHCOAST,South_Coast_Aviation,South Coast Aviation
,GDE,GADEL,Servicios_AÃ©reos_Gadel,Servicios Aéreos Gadel
,GDG,GOLDEN GATE,S.P._Aviation,S.P. Aviation
,GIK,SEBA,Seba_Airlines,Seba Airlines
,GNA,SERVIGANA,Servicios_AÃ©reos_Gana,Servicios Aéreos Gana
,GSW,,Sky_Wings_Airlines,Sky Wings Airlines
,GXL,STARDUST,Star_XL_German_Airlines,Star XL German Airlines
,HAU,SKYHAUL,Skyhaul,Skyhaul
,HIP,STARSA,Starship,Starship
,HJE,GOSA,Servicios_Ejecutivos_Gosa,Servicios Ejecutivos Gosa
SO,HKA,SPEND AIR,Superior_Aviation_(airline),Superior Aviation
,HLO,HALO,Samaritan_Air_Service,Samaritan Air Service
,HRI,HELIRIM,Skyraidybos_Mokymo_Centras,Skyraidybos Mokymo Centras
,HSK,MATRA,Sky_Europe_Airlines,Sky Europe Airlines
,HSV,HIGHSWEDE,Svenska_Direktflyg,Svenska Direktflyg
,HSY,HELISKY,Sky_Helicopteros,Sky Helicopteros
TE,IGA,IGUANA,Skytaxi,Skytaxi
,IJS,,Silvair,Silvair
,ILS,SERVICIOS ILSA,Servicios_AÃ©reos_Ilsa,Servicios Aéreos Ilsa
,INK,SINCOM AVIA,Sincom-Avia,Sincom-Avia
,IRV,SAFAT AIR,Safat_Airlines,Safat Airlines
,IRZ,SAHA,Saha_Airlines,Saha Airlines Services
,JAM,SUNTRACK,Sunline_Express,Sunline Express
,JCM,SECUREAIR,Secure_Air_Charter,Secure Air Charter
,JIM,SARK,Sark_International_Airways,Sark International Airways
JK,JKK,SPANAIR,Spanair,Spanair
,KKS,KOKSHE,Salem_(airline),Salem
,KOP,COPTERS,Servicios_AÃ©reos_Copters,Servicios Aéreos Copters
,KSP,SAEP,Servicios_AÃ©reos_Expecializados_En_Transportes_Petroleros,Servicios Aéreos Expecializados En Transportes Petroleros
,KYB,SKY AIROPS,Skybridge_AirOps,Skybridge AirOps
,KYR,SKY AERONAUTICAL,Sky_Aeronautical_Services,Sky Aeronautical Services
,LGU,LAGUNA,Servicios_AÃ©reos_Ejecutivos_De_La_Laguna,Servicios Aéreos Ejecutivos De La Laguna
,LLA,LEO LOPOZ,Servico_Leo_Lopex,Servico Leo Lopex
,LLS,SERVIESTRELLA,Servicios_AÃ©reos_Estrella,Servicios Aéreos Estrella
,LMG,SOUTH AFRICAN,South_African_Air_Force,South African Air Force
,LMO,SKY HOLDINGS,Sky_One_Holdings,Sky One Holdings as Privaira
,LRS,,Sansa_Airlines,Sansa
,LSP,AIR TONY,Spectrum_Aviation_Incorporated,Spectrum Aviation Incorporated
,MCG,MEDICOPTER,SOS_Helikoptern_Gotland,SOS Helikoptern Gotland
,MDT,MIDNIGHT,Sundt_Air,Sundt Air
,MLO,MILENIO,Servicios_AÃ©reos_Milenio,Servicios Aéreos Milenio
,MMS,MUSAAD AIR,SAAD_(A320)_Limited,SAAD (A320) Limited
,MRI,MORITANI,Servicios_AÃ©reos_Moritani,Servicios Aéreos Moritani
2G*,MRR,MARINER,San_Juan_Airlines,San Juan Airlines
,MSG,SAR-REGIONAL,Servico_AÃ©reo_Regional,Servico Aéreo Regional
,MSP,SEGURIDAD,Servicio_De_Vigilancia_AÃ©rea_Del_Ministerio_De_Seguridad_PÃºblica,Servicio De Vigilancia Aérea Del Ministerio De Seguridad Pública
,MTG,,Servicios_AÃ©reos_MTT,Servicios Aéreos MTT
1Z,APD,,Sabre_(computer_system),Sabre Pacific
1S,,,Sabre_(computer_system),Sabre
1I,,,Sierra_Nevada_Airlines,Sierra Nevada Airlines
1H,,,Siren-Travel,Siren-Travel
1Q,,,Sirena_(airline),Sirena
,SBW,SNOWMAN,Snowbird_Airlines,Snowbird Airlines
1K,,,Southern_Cross_Distribution,Southern Cross Distribution
1K,,,Sutra_(air_company),Sutra
2C,,,SNCF,SNCF
2S,,,Star_Equatorial_Airlines,Star Equatorial Airlines
,NAD,SEULAWAH,Seulawah_Nad_Air,Seulawah Nad Air
,NAZ,NAZAS,Servicios_AÃ©reos_del_Nazas_S.A._de_C.V.,Servicios Aéreos del Nazas S.A. de C.V.
,NCS,COMMUTER-CANADA,Simpson_Air_Ltd,Simpson Air Ltd
NK,NKS,SPIRIT WINGS,Spirit_Airlines,Spirit Airlines
,NON,SERVICIOS LATINO,Servicios_AÃ©reos_Latinoamericanos,Servicios Aéreos Latinoamericanos
,NRZ,MONARREZ,Servicios_AÃ©reos_Monarrez,Servicios Aéreos Monarrez
,NSC,TRANS-SOCIETE,Societe_De_Transport_Aerien_De_Mauritanie,Societe De Transport Aerien De Mauritanie
,NSE,SATENA,SATENA,SATENA
,NTB,SERVINORTE,Servicios_AÃ©reos_Del_Norte,Servicios Aéreos Del Norte
,NTG,INTEGRALES,Servicios_Integrales_De_AviaciÃ³n,Servicios Integrales De Aviación
S0,OKS,SLOK GAMBIA,Slok_Air_Gambia,Slok Air Gambia
,OKT,SOKO AIR,Soko_Aviation,Soko Aviation
,OLC,SOLARCARGO,Solar_Cargo,Solar Cargo
,OLO,SOLO,Soloflight,Soloflight
,ONG,SONNIG,Sonnig_SA,Sonnig SA
SO,OSL,SOSOLISO,Sosoliso_Airlines,Sosoliso Airlines
,OSS,NOTICIOSOS,Servicios_AÃ©reos_Noticiosos,Servicios Aéreos Noticiosos
,OTL,SOUTHLINE,South_Airlines,South Airlines
VA,OZW,VELOCITY,Virgin_Australia_Regional_Airlines,Virgin Australia Regional Airlines
,PIV,AEROSOKOL,Sokol_(airline),Sokol
,PLT,PALMETTO,South_Carolina_Aeronautics_Commission,South Carolina Aeronautics Commission
,PMR,SERVICIOS PREMIER,Servicios_AÃ©reos_Premier,Servicios Aéreos Premier
,PNS,PENAS,Survey_Udara_(Penas),Survey Udara (Penas)
,POB,POBLANOS,Servicios_AÃ©reos_Poblanos,Servicios Aéreos Poblanos
,PSV,PROSERVICIOS,Servicios_A%C3%A9reos_Profesionales,Servicios Aéreos Profesionales
,PTM,POSTMAN,Southeastern_Airways,Southeastern Airways
,PUR,SPURWING,Spurwing_Airlines,Spurwing Airlines
1I,PZR,PHAZER,Sky_Trek_International_Airlines,Sky Trek International Airlines
,RBW,CAI HONG,Shandong_Airlines_Rainbow_Jet,Shandong Airlines Rainbow Jet
,REJ,REGIONAL LINK,SA_Airlink_Regional,SA Airlink Regional
,RER,REGAIR,Servicio_AÃ©reo_Regional_Regair,Servicio Aéreo Regional Regair
,RFT,ROMANIAN ACADEMY,Scoala_Superioara_De_Aviatie_Civila,Scoala Superioara De Aviatie Civila
,RGC,REGIOMONTANO,Servicios_AÃ©reos_Regiomontanos,Servicios Aéreos Regiomontanos
,RLS,S-AIRLINES,S-Air,S-Air
,RMP,SERAMSA,Servicios_De_Rampa_Y_Mostrador,Servicios De Rampa Y Mostrador
,RSE,RED SEA,SNAS_Aviation,SNAS Aviation
SX,SKB,SKYBUS,Skybus_Airlines,Skybus Airlines
,SKC,SKYMASTER AIR,Skymaster_Airlines,Skymaster Airlines
,SKD,SKY DAWG,Sky_Harbor_Air_Service,Sky Harbor Air Service
,SKE,SKYISLE,Sky_Tours,Sky Tours
,AZG,SAKSERVICE,Sakaviaservice,Sakaviaservice
,SKF,SKYCRAFT,Skycraft,Skycraft
,SKG,SKYCRAFT-CANADA,Skycraft_Air_Transport,Skycraft Air Transport
RU,SKI,SKYKING,SkyKing_Turks_and_Caicos_Airways,SkyKing Turks and Caicos Airways
,SKK,SKYLINK,Skylink_Aviation,Skylink Aviation
,SKL,SKYCHARTER,Skycharter_(Malton),Skycharter (Malton)
,SKN,SKYLINER,Skyline_Aviation_Services,Skyline Aviation Services
,SKO,SKYWORK,Scottish_Airways_Flyers,Scottish Airways Flyers
,SKR,SKYSCAPES,Skyscapes_Air_Charters,Skyscapes Air Charters
,SKS,SKY SERVICE,Sky_Service_(Belgium),Sky Service
,SKS,,Sky_Link_Aviation,Sky Link Aviation
S3,BBR,SANTA BARBARA,Santa_Barbara_Airlines,Santa Barbara Airlines
XT,SKT,SKY YOU,SkyStar_Airways,SkyStar Airways
H2,SKU,AEROSKY,Sky_Airline,Sky Airline
OO,SKW,SKYWEST,SkyWest_Airlines,SkyWest Airlines
JZ,SKX,SKY EXPRESS,Skyways_Express,Skyways Express
BC,SKY,SKYMARK,Skymark_Airlines,Skymark Airlines
,SKZ,SKYWAY-INC,Skyway_Enterprises,Skyway Enterprises
LJ,SLA,SELAIR,Sierra_National_Airlines,Sierra National Airlines
,SLB,SLOK AIR,Slok_Air,Slok Air
,SLD,SILVERLINE,Silver_Air_(Czech_Republic),Silver Air
,SLE,SLIPSTREAM,Safair,Streamline
,SLF,ELISTARFLY,Starfly,Starfly
,SLG,LIFEGUARD,Saskatchewan_Government,Saskatchewan Government
,SLH,SILVERHAWK,Silverhawk_Aviation,Silverhawk Aviation
,AGE,AEROANGEL,Servicios_AÃ©reos_de_Los_Ãngeles,Servicios Aéreos de Los Ángeles
MI,SLK,SILKAIR,SilkAir,SilkAir
6Q,SLL,SLOV LINE,Slovak_Airlines,Slovak Airlines
PY,SLM,SURINAM,Surinam_Airways,Surinam Airways
,SLN,SLOANE,Sloane_Aviation,Sloane Aviation
,SLP,SALPA,Salpa_Aviation,Salpa Aviation
,SLS,SERVICIOS SLAINTE,Servicios_AÃ©reos_Slainte,Servicios Aéreos Slainte
,SLV,AVISTELLA,Stella_Aviation,Stella Aviation
,SLW,SALMA AIR,Salama_Airlines_Nigeria,Salama Airlines Nigeria
,SLX,SETE,Sete_Linhas_AÃ©reas,Sete Linhas Aéreas
,SLY,SKYCO,Sky_Line_for_Air_Services,Sky Line for Air Services
,SLZ,LUZA,Super_Luza,Super Luza
,SMA,SESAME,SMA_Airlines,SMA Airlines
,SMC,SAMER,Sabang_Merauke_Raya_Air_Charter,Sabang Merauke Raya Air Charter
,SMD,SERVICIOS MARQUESA,Servicios_AÃ©reos_La_Marquesa,Servicios Aéreos La Marquesa
8D,,,Servant_Air,Servant Air
,SME,SEMICH,Semos,Semos
,SMF,GORDON,Smalandsflyg,Smalandsflyg
,SMH,SMITHAIR,Smithair,Smithair
,SMK,ERTIS,Semeyavia,Semeyavia
,SML,SMITH AIR,Smith_Air_(1976),Smith Air (1976)
,SMM,SUMMIT-AIR,Summit_Airlines,Summit Airlines
,SMQ,SAMAR AIR,Samar_Air,Samar Air
,SMR,SOMON AIR,Somon_Air,Somon Air
,SMT,SKYLIMIT,Skyline_(Nigeria),Skyline
,AOS,AEROSOL,"Servicios_AÃ©reos_Del_Sol,_S.A._de_C.V.","Servicios Aéreos Del Sol, S.A. de C.V."
,SNA,SENATOR,Senator_Aviation_Charter,Senator Aviation Charter
NB,SNB,STERLING,Sterling_Airlines,Sterling Airlines
,SNE,SANSA,Servicios_AÃ©reos_De_Nicaragua,Servicios Aéreos De Nicaragua (SANSA)
,SNF,SHANS AIR,Shans_Air,Shans Air
,SNH,SENSERVICE,Senair_Services,Senair Services
,SNI,SAVANAHLINE,Savanah_Airlines,Savanah Airlines
6J,SNJ,NEWSKY,Skynet_Asia_Airways,Skynet Asia Airways
,SNK,SUN KING,Southeast_Airlines,Southeast Airlines (Sun Jet International)
,SNL,SOONAIR,Soonair_Lines,Soonair Lines
,SNM,SERVIZI AEREI,Servizi_Aerei,Servizi Aerei
,SNP,SUN PACIFIC,Sun_Pacific_International,Sun Pacific International
,SNQ,EXECU-QUEST,Sun_Quest_Executive_Air_Charter,Sun Quest Executive Air Charter
,SNS,,Societe_Centrafricaine_De_Transport_Aerien,Societe Centrafricaine De Transport Aerien
,SNT,SUNCOAST,Suncoast_Aviation,Suncoast Aviation
,SNU,,Snunit_Aviation,Snunit Aviation
,SNV,SUDANESE,Sudanese_States_Aviation,Sudanese States Aviation
,SNW,SUN WEST,Sun_West_Airlines,Sun West Airlines
,SNX,SUNEX,Sun_Air_Aviation_Services,Sun Air Aviation Services
,SOB,STABO,Stabo_Freight,Stabo Freight
,SOC,,Southern_Cargo_Air_Lines,Southern Cargo Air Lines
,SOH,SOUTHERN OHIO,Southern_Ohio_Aviation_Sales,Southern Ohio Aviation Sales
,SOI,SOAVAIR,Southern_Aviation,Southern Aviation
IE,SOL,SOLOMON,Solomon_Airlines,Solomon Airlines
,SOM,SOMALAIR,Somali_Airlines,Somali Airlines
,SON,SUNSHINE TOURS,Sunshine_Air_Tours,Sunshine Air Tours
,SOO,SOUTHERN AIR,Southern_Air,Southern Air
,SOP,SOLINAIR,Solinair,Solinair
,SOR,SONAIR,Sonair_Servico_AÃ©reo,Sonair Servico Aéreo
,SOT,SOUTH COURIER,Southeast_Correct_Craft,Southeast Correct Craft
,SOU,SOUTHERN EXPRESS,Southern_Airways,Southern Airways
6W,SOV,SARATOV AIR,Saravia,Saratov Airlines Joint Stock Company
,SOW,SOWIND,Sowind_Air,Sowind Air
,SOX,SOLIDAIR,Solid_Air,Solid Air
HZ,SOZ,SATCO,Sat_Airlines,Sat Airlines
,SPA,SIERRA PACIFIC,Sierra_Pacific_Airlines,Sierra Pacific Airlines
,SPB,SPRING CLASSIC,Springbok_Classic_Air,Springbok Classic Air
,SPC,PORT,Skyworld_Airlines,Skyworld Airlines
,SPE,SPRAGUE,Sprague_Electric_Company,Sprague Electric Company
,SPF,SPACE WORLD,Space_World_Airline,Space World Airline
,SPG,SPRING AIR,Springdale_Air_Service,Springdale Air Service
,SPH,SAPPHIRE-CHARTER,Sapphire_Executive_Air,Sapphire Executive Air
,SPI,SOUTH PACIFIC,South_Pacific_Island_Airways,South Pacific Island Airways
,SPL,CORPORATIVOS LAGUNA,Servicios_Corporativos_AÃ©reos_De_La_Laguna,Servicios Corporativos Aéreos De La Laguna
,SPN,AIR SKORPIO,Skorpion_Air,Skorpion Air
,SPP,SAPPHIRE,Sapphire_Aviation,Sapphire Aviation
,SPQ,SERVICOS PALENQUE,Servicios_AÃ©reos_Palenque,Servicios Aéreos Palenque
,TBS,TRIBASA,Servicios_AÃ©reos_Tribasa,Servicios Aéreos Tribasa
S5,TCF,MERCURY,Shuttle_America,Shuttle America
,SVV,SALT,SALTAVIATION,SALTAVIATION
,TGT,TARGET,SAAB_Nyge_Aero,SAAB Nyge Aero
,THB,THAI SABAI,Spark_Air,Spark Air
,TIH,TIRIAC AIR,S_C_Ion_Tiriac,S C Ion Tiriac
,TRL,STARSTREAM,Starlite_Aviation,Starlite Aviation
,TRN,AEROTRON,Servicios_AÃ©reos_Corporativos,Servicios Aéreos Corporativos
,TTM,TOUT-AIR,Societe_Tout_Transport_Mauritanien,Societe Tout Transport Mauritanien
,TZU,TAMAZULA,Servicios_AÃ©reos_Tamazula,Servicios Aéreos Tamazula
,UGP,SHARINK,Shar_Ink,Shar Ink
,UKU,PYSHMA,Second_Sverdlovsk_Air_Enterprise,Second Sverdlovsk Air Enterprise
,UNT,UNIVERSITARIO,Servicios_AÃ©reos_Universitarios,Servicios Aéreos Universitarios
,USK,SKIF-AIR,Skif-Air,Skif-Air
,USN,SAMAS,Smarkand_Aero_Servise,Smarkand Aero Servise
,UZS,SOGDIANA,Samarkand_Airways,Samarkand Airways
,VDO,AVANDARO,Servicios_AÃ©reos_Avandaro,Servicios Aéreos Avandaro
,VGS,SMART,Stichting_Vliegschool_16Hoven,Stichting Vliegschool 16Hoven
,VRB,SILVERBACK,Silverback_Cargo_Freighters,Silverback Cargo Freighters
,VRS,VAIRSA,Sirvair,Sirvair
,VSV,VLASTA,Scat_Air,Scat Air
,VXN,VIXEN,Sunset_Aviation,Sunset Aviation
,TWY,TWILIGHT,"Sunset_Aviation,_LLC","Sunset Aviation, LLC"
,WCC,WEST COAST,Sport_Air_Travel,Sport Air Travel
,WFC,SWIFTCOPTERS,Swift_Copters,Swift Copters
,WLK,SKYWATCH,Skyrover_CC,Skyrover CC
,XLK,SAFARILINK,Safarilink_Aviation,Safarilink Aviation
,XMX,SENEAM,SENEAM,SENEAM
,XPG,,Southport_Air_Service,Southport Air Service
,XSA,,Spectrum_Air_Service,Spectrum Air Service
,XSN,,Stephenville_Aviation_Services,Stephenville Aviation Services
,XTA,TEXTRA,Servicios_AÃ©reos_Textra,Servicios Aéreos Textra
,XTR,EXTER,Sector_Airlines,Sector Airlines
,XXS,,Skyplan_Services,Skyplan Services
,YBE,YELLOW BIRD,Stewart_Aviation_Services,Stewart Aviation Services
R1,,,Sirin_(airline),Sirin
S6,,,Star_Air_(Denmark),Star Air
,FYA,FLYANT,Flyant,Servicios Aéreos Integrales / Flyant
,UFA,FLIGHT ACADEMY,The_State_Flight_Academy_of_Ukraine,State Flight Academy of Ukraine
,SAF,SINGA,Singapore_Air_Force,Singapore Air Force
Q4,TLK,STARLINK,Starlink_Aviation,Starlink Aviation
V9,HCW,STAR1,Star1_Airlines,Star1 Airlines
SO,AAS,,Sunshine_Airlines,Sunshine Airlines
O3,CSS,SHUN FENG,SF_Airlines,SF Airlines
,SXN,SAXONAIR,SaxonAir,SaxonAir
,RZ,YANGTZE RIVER,,Superna Airlines
5P,LLX,GERMANJET,Small_Planet_Airlines_(Germany),Small Planet Airlines
XG,SXD,SUNRISE,Sunexpress_Deutschland,Sunexpress Deutschland
,TEZ,,TezJet_Airlines,TezJet Airlines
,TXZ,EXPRESS AIR,Thai_Express_Air,Thai Express Air
WE,THD,THAI SMILE,Thai_Smile,Thai Smile Airways
VZ,TVJ,THAIVIET JET,Thai_Vietjet_Air,Thai Vietjet Air
,TQE,TAXAIR,Taxair_Mexiqienses,Taxair Mexiqienses
XJ,TAX,EXPRESS WING,Thai_AirAsia_X,Thai AirAsia X
SL,TLM,MENTARI,Thai_Lion_Air,Thai Lion Mentari
,TCB,AERO COLOMBIA,Transporte_Aereo_De_Colombia,Transporte Aereo De Colombia
,TDA,TREND AIR,Trend_Aviation,Trend Aviation
,THS,TUSAS,Turkish_Aerospace_Industries,Turkish Aerospace Industries
,TMD,TRANSMANDU,Transmandu,Transmandu
,TRK,TURKISH REPUBLIC,Turkish_Airlines,Turkish Airlines  General Aviation
,TKJ,TARKIM AVIATION,Tarkim_Aviation,Tarkim Aviation
,TCG,AFRICARGO,Transafricaine_Air_Cargo,Transafricaine Air Cargo
,TMC,TRAIL BLAZER,Travel_Management_Company,Travel Management Company
,SWD,SAWBLADE,Trifly,Trifly
,SWL,TRANSJET,Trans_Jet_Airways,Trans Jet Airways
,TAC,TURBOT,Turbot_Air_Cargo,Turbot Air Cargo
,TAD,TRANS DOMINICAN,Transporte_AÃ©reo_Dominicano,Transporte Aéreo Dominicano
EQ,TAE,TAME,TAME,TAME
,BAP,BIG APPLE,Trans_International_Express_Aviation,Trans International Express Aviation
,TAG,TAG U-S,TAG_Aviation_USA,TAG Aviation USA
,TAL,TALAIR,Talair,Talair
JJ,TAM,TAM,LATAM_Brasil,LATAM Brasil
,AUC,AUSCARGO,Transaustralian_Air_Express,Transaustralian Air Express
TP,TAP,AIR PORTUGAL,TAP_Portugal,TAP Portugal
TU,TAR,TUNAIR,Tunisair,Tunisair
,TAU,TRANSTAURO,Transportes_AÃ©reos_Tauro,Transportes Aéreos Tauro
,TAX,TRAVELAIR,Travel_Air,Travel Air
3V,TAY,QUALITY,TNT_Airways,TNT Airways
TR,TBA,TRANSBRASIL,Transbrasil,Transbrasil
,TBC,,Turbine_Air_Cargo_UK,Turbine Air Cargo UK
,TBD,ORCA,Thunderbird_Tours,Thunderbird Tours
M7,TBG,,Tropical_Airways,Tropical Airways
,TBH,,Trinity_Air_Bahamas,Trinity Air Bahamas
,TBI,TAB INTERNATIONAL,TAB_Express_International,TAB Express International
,TBM,TABAN AIR,Taban_Air_Lines,Taban Air Lines
,TBN,TEEBAH,Teebah_Airlines,Teebah Airlines
,TBR,TUBELAIR,Tubelair,Tubelair
,TBX,TABEX,Tobago_Express,Tobago Express
,TCA,TROPICANA,Tropican_Air_Services,Tropican Air Services
,TCB,TRANSCARIBE,Transporte_del_Caribe,Transporte del Caribe
,TCC,TRANSCAL,Trans_Continental_Airlines,Trans Continental Airlines
,TCD,TCHADLINES,Tchad_Airlines,Tchad Airlines
,TCE,TRANS-COLORADO,Trans-Colorado_Airlines,Trans-Colorado Airlines
T2,TCG,THAI CARGO,Thai_Air_Cargo,Thai Air Cargo
,TCH,TRANS GULF,Transcontinental_Air,Transcontinental Air
,TCM,TELEDYNE,Teledyne_Continental_Motors,Teledyne Continental Motors
,TCN,TRANSCON,Trans_Continental_Airlines,Trans Continental Airlines
,TCP,TRANSCORP,Transcorp_Airways,Transcorp Airways
,TCT,TRANS-CONT,Transcontinental_Sur,Transcontinental Sur
,TCU,TRANSGLOBAL,Transglobal_Airways_Corporation,Transglobal Airways Corporation
HQ,TCW,THOMAS COOK,Thomas_Cook_Airlines_Belgium,Thomas Cook Airlines
MT,TCX,THOMAS COOK,Thomas_Cook_Airlines,Thomas Cook Airlines
,TCY,TWIN CITY,Twin_Cities_Air_Service,Twin Cities Air Service
,TDC,TADAIR,Tadair,Tadair
,TDE,TELLURIDE,Flight_One,Tellavia / Flight One
,TDI,TRANSIXTLAN,Transportes_AÃ©reos_de_IxtlÃ¡n,Transportes Aéreos de Ixtlán
TQ,TDM,TANDEM,Tandem_Aero,Tandem Aero
,TDO,TRADO,TRADO,TRADO
,TDR,TRADEAIR,Trade_Air,Trade Air
,TDV,TAXI EVORA,Taxi_Aero_Nacional_Del_Evora,Taxi Aero Nacional Del Evora
,TDX,TRADEWINDS EXPRESS,Tradewinds_Airlines,Tradewinds Airlines
,TEB,TENIR AIR,Tenir_Airlines,Tenir Airlines
L9,TLW,Teamline,Teamline_Air,Teamline Air
,TEF,TECFOTO,Tecnicas_Fotograficas,Tecnicas Fotograficas
,TEH,TEMPELHOF,Tempelhof_Airways,Tempelhof Airways
,TEL,TELFORD,Telford_Aviation,Telford Aviation
,TEM,TECHMONT,Tech-Mont_Helicopter_Company,Tech-Mont Helicopter Company
,TEN,TENNESSEE,Tennessee_Airways,Tennessee Airways
,TER,TERRI-AIRE,Territorial_Airlines,Territorial Airlines
,TES,Tesaban,,Taespejo Portugal LDA
UE,TEP,TRANSEURLINE,Transeuropean_Airlines,Transeuropean Airlines
,TET,TEPAVIA,Tepavia-Trans_Airlines,Tepavia-Trans Airlines
,TFA,TRANS FLORIDA,Trans-Florida_Airlines,Trans-Florida Airlines
,TFB,ROYAL TEE-AIR,Tair_Airways,Tair Airways
,TFF,TALON FLIGHT,Talon_Air,Talon Air
,AIM,PIJO,Trabajos_AÃ©reos_Murcianos,Trabajos Aéreos Murcianos
,TFH,THAI HELICOPTER,Thai_Flying_Helicopter_Service,Thai Flying Helicopter Service
,TFI,,Transport_Facilitators,Transport Facilitators
,TFK,,Transafrik_International,Transafrik International
,TFO,TRANSPORTES PACIFICO,Transportes_AÃ©reos_del_PacÃ­fico,Transportes Aéreos del Pacífico
,TFT,THAI FLYING,Thai_Flying_Service,Thai Flying Service
,SRF,SAN RAFEAL,Transportes_AÃ©reos_San_Rafael,Transportes Aéreos San Rafael
,TFY,TAYSIDE,Tayside_Aviation,Tayside Aviation
,TGC,THANET,TG_Aviation,TG Aviation
,TGE,TASA,Trabajos_AÃ©reos,Trabajos Aéreos
,TGI,TRANSPORTE REGIONAL,Transportes_AÃ©reos_Regionales,Transportes Aéreos Regionales
,TGM,TAG ESPANA,TAG_Aviation_Espana,TAG Aviation Espana
,TGN,TRIGANA,Trigana_Air_Service,Trigana Air Service
2S,TGR,SATGURAIR,Satgur_Air_Transport,Satgur Air Transport
ZT,AWC,ZAP,Titan_Airways,Titan Airways
,TGO,TRANSPORT,Transport_Canada,Transport Canada
TR,TGW,GO CAT,Tiger_Airways,Tigerair Singapore
TT,TGG,TIGGOZ,Tigerair_Australia,Tigerair Australia
IT,TTW,SMART CAT,Tigerair_Taiwan,Tigerair Taiwan
,MDL,MANDALA,Tigerair_Mandala,Tigerair Mandala
DG,SRQ,BLUE JAY,Cebgo,Cebgo
,TGX,TRANSGABON,Transair_Gabon,Transair Gabon
,TGY,TRANS GUYANA,Trans_Guyana_Airways,Trans Guyana Airways
TG,THA,THAI,Thai_Airways_International,Thai Airways International
,THC,TARHEEL,Tar_Heel_Aviation,Tar Heel Aviation
,THE,TOUMAI AIR,Toumai_Air_Tchad,Toumai Air Tchad
,THF,TOURAINE HELICO,Touraine_Helicoptere,Touraine Helicoptere
,THG,THAI GLOBAL,Thai_Global_Airline,Thai Global Airline
,THJ,THAI JET,Thai_Jet_Intergroup,Thai Jet Intergroup
,THK,HUR KUS,Turk_Hava_Kurumu_Hava_Taksi_Isletmesi,Turk Hava Kurumu Hava Taksi Isletmesi
FD,AIQ,THAI ASIA,Thai_AirAsia,Thai AirAsia
,THO,LEMPIRA,TACA_De_Honduras,TACA De Honduras
,THR,TEHRAN AIR,Tehran_Airline,Tehran Airline
,THU,AIR THUNDER,Thunder_Airlines,Thunder Airlines
TK,THY,TURKISH,Turkish_Airlines,Turkish Airlines
,THZ,LYON HELIJET,Trans_Helicoptere_Service,Trans Helicoptere Service
,TIA,TRANS INTERNATIONAL,Trans_International_Airlines,Trans International Airlines
,TIC,TRAVEL INTERNATIONAL,Travel_International_Air_Charters,Travel International Air Charters
,TIE,TIME AIR,Time_Air,Time Air
,TIK,TICAIR,Tic_Air,Tic Air
,TIM,TEAM BRASIL,TEAM_Linhas_A%C3%A9reas,TEAM Linhas Aéreas
,TIN,TAINO,Taino_Tours,Taino Tours
,TIS,TESIS,Tesis,Tesis
,TIW,TIACA,Transcarga_Intl_Airways,Transcarga Intl Airways
,TJK,TAJIKAIR,Tajikair,Tajikair
,TJN,NERON,Tien-Shan_(airline),Tien-Shan
,TJS,TYROLJET,Tyrolean_Jet_Services,Tyrolean Jet Services
T7,TJT,TWINJET,Twin_Jet,Twin Jet
,TKC,TIKAL,Tikal_Jets_Airlines,Tikal Jets Airlines
,TKE,ISLAND BIRD,Take_Air_Line,Take Air Line
,TKX,TROPEXPRESS,Tropical_International_Airways,Tropical International Airways
9I,LLR,THAI SKY AIR,Thai_Sky_Airlines,Thai Sky Airlines
,TLA,TRANSLIFT,Translift_Airways,Translift Airways
,TLF,TRANS-LEONE,Transport_Africa,Transport Africa
,TLL,ATLANTIC LEONE,Trans_Atlantic_Airlines,Trans Atlantic Airlines
,TLO,TALON AIR,Eagle_Canyon_Airlines,Eagle Canyon Airlines
TD,TLP,TULIPAIR,Tulip_Air,Tulip Air
,TAJ,TUNISAVIA,Tunisavia,Tunisavia
,TLS,TEALSY,TLC_Air,TLC Air
,TLT,TURTLE,Turtle_Airways,Turtle Airways
,TLV,PAJAROS,Travelair_(Uruguayan_airline),Travelair
,TLX,TELESIS,Telesis_Transair,Telesis Transair
,TLY,TOPFLY,Top_Fly,Top Fly
TL,TMA,TANGO LIMA,Trans_Mediterranean_Airlines,Trans Mediterranean Airlines
,TMD,,Transmed_Airlines,Transmed Airlines
GY,TMG,TRILINES,Tri-MG_Intra_Asia_Airlines,Tri-MG Intra Asia Airlines
,TMH,TAXIMARAKAME,Taxis_Turisticos_Marakame,Taxis Turisticos Marakame
,TMI,TAMIRWAYS,Tamir_Airways,Tamir Airways
,TMK,TOMAHAWK,Tomahawk_Airways,Tomahawk Airways
OF,TML,TAM AIRLINE,Transports_et_Travaux_A%C3%A9riens_de_Madagascar,Transports et Travaux Aériens de Madagascar
,TMM,WILLOW RUN,TMC_Airlines,TMC Airlines
,TMQ,TRAM AIR,TRAM_AIR,TRAM
,TMR,TIMBER,Timberline_Air,Timberline Air
,TMS,TEMSCO,Temsco_Helicopters,Temsco Helicopters
,TMT,TRANS MIDWEST,Trans_Midwest_Airlines,Trans Midwest Airlines
,TMX,TRAMON,Tramon_Air,Tramon Air
,TMY,MUNDO MAYA,Transportes_AÃ©reos_del_Mundo_Maya,Transportes Aéreos del Mundo Maya
,TMZ,TRANS AMAZON,Transporte_Amazonair,Transporte Amazonair
,TNB,TRANS-BENIN,Trans_Air-Benin,Trans Air-Benin
,TNE,TAXINOROESTE,Taxis_AÃ©reos_del_Noroeste,Taxis Aéreos del Noroeste
,TNF,TRANSFAS,Transafricaine,Transafricaine
,TNG,,Tennessee_Air_National_Guard,Tennessee Air National Guard 164th Airlift Group
,TNI,TRANSINTER,Transair_International_Linhas_AÃ©reas,Transair International Linhas Aéreas
,TNL,SKY HORSE,Tengeriyn_Ulaach_Shine,Tengeriyn Ulaach Shine
3P,TNM,TIARA,Tiara_Air,Tiara Air
,TNP,TRANSPED,Transped_Aviation,Transped Aviation
,TNR,TAN AIR,Tanana_Air_Services,Tanana Air Services
,TNT,TRANS NORTH,Trans_North_Turbo_Air,Trans North Turbo Air
,TNV,TRANSNORTHERN,Transnorthern,Transnorthern
,TNW,TRANS-NATION,Trans_Nation_Airways,Trans Nation Airways
,TNX,TRAINER,Trener_Ltd,Trener Ltd
,TNY,TWINCAL,Twin_Town_Leasing_Company,Twin Town Leasing Company
7T,TOB,TOBRUK AIR,Tobruk_Air,Tobruk Air
,TOJ,TOJ AIRLINE,TOJ_Airlines,TOJ Airlines
TI,TOL,TOL AIR,Tol-Air_Services,Tol-Air Services
BY,TOM,TUI AIR[48],Thomson_Airways,Thomson Airways
,TOP,AIR TOP,Top_Air,Top Air
,TOR,TORONTAIR,Toronto_Airways,Toronto Airways
PM,TOS,TROPISER,Tropic_Air,Tropic Air
,TOT,,Totavia,Totavia
,TOW,TEE AIR,Tower_Air,Tower Air
,TOY,TOYOTA,Toyota_Canada_Inc.,Toyota Canada
QT,TPA,TAMPA,TAMPA_Cargo,TAMPA
,TPD,TOP SPEED,Top_Speed_(airline),Top Speed
,TPF,TAXIPACIFICO,Taxis_AÃ©reos_del_PacÃ­fico,Taxis Aéreos del Pacífico
,TPG,TRANSPEGASO,Transportes_AÃ©reos_Pegaso,Transportes Aéreos Pegaso
,TPL,INTERPILOT,TAR_Interpilot,TAR Interpilot
,TPM,TRANSPAIS,TranspaÃ­s_AÃ©reo,Transpaís Aéreo
,TPN,AEREA DELNORTE,TransportaciÃ³n_AÃ©rea_del_Norte,Transportación Aérea del Norte
,TPP,TRANS EXPRESS,Transpac_Express,Transpac Express
,TPR,TAXIS PARRAL,Taxis_AÃ©reos_De_Parral,Taxis Aéreos De Parral
,TPS,TAPSA,TAPSA,TAPSA Transportes Aéreos Petroleros
,TPT,TASSA,Transportes_AÃ©reo_del_Sureste,Transportes Aéreo del Sureste
,TPU,TRANS PERU,Trans_American_Airlines,Trans American Airlines (Trans Am)
,TPV,THAI PACIFIC,Thai_Pacific_Airlines_Business,Thai Pacific Airlines Business
,TPX,TRANSXALAPA,Transportes_AÃ©reos_De_Xalapa,Transportes Aéreos De Xalapa
,TPY,TRANS PROVINCIAL,Trans-Provincial_Airlines,Trans-Provincial Airlines
,TPZ,TRANSPAZ,Transportes_La_Paz,Transportes La Paz
K3,TQN,TAQUAN,Taquan_Air_Services,Taquan Air Services
,TQR,TRANSQUERETARO,TransportaciÃ³n_AÃ©rea_De_QuerÃ©taro,Transportación Aérea De Querétaro
GE,TNA,,TransAsia_Airways,TransAsia Airways
TO,TVF,FRANCE SOLEIL,Transavia_France,Transavia France
HV,TRA,TRANSAVIA,Transavia_Holland,Transavia Holland
,TRC,TRACKER,Trans_Air_Charter,Trans Air Charter
VR,TCV,CABOVERDE,TACV,TACV
,TRD,TRANS ISLAND,Trans_Island_Air,Trans Island Air
,TRF,TAXI JET,Taxi_Air_Fret,Taxi Air Fret
,TRG,,TRAGSA_(Medios_AÃ©reos),TRAGSA (Medios Aéreos)
,TRJ,HIGH TIDE,Trans_Euro_Air,Trans Euro Air
,TRM,SOTRANS,Transport_Aerien_de_Mauritanie,Transport Aerien de Mauritanie
,TRO,MOLOKAI,Tropic_Airlines-Air_Molokai,Tropic Airlines-Air Molokai
,TRQ,HUNTER,Turdus_Airways,Turdus Airways
,TRR,TRAMSON,Tramson_Limited,Tramson Limited
,TRT,TRANS ARABIAN,Trans_Arabian_Air_Transport,Trans Arabian Air Transport
,TRU,TRI AIR,Triangle_Airline_(Uganda),Triangle Airline (Uganda)
,TRW,TRANS-WEST,Transwestern_Airlines_of_Utah,Transwestern Airlines of Utah
,TRY,TRISTAR AIR,Tristar_Airlines,Tristar Airlines
T9,TRZ,TRANS-MERIDIAN,TransMeridian_Airlines,TransMeridian Airlines
,TSA,AIRTRAF,Transair_France,Transair France
,TSD,TAFI,TAF-Linhas_AÃ©reas,TAF-Linhas Aéreas
TH,TSE,TRANSMILE,Transmile_Air_Services,Transmile Air Services
,TSG,TRANS-CONGO,Trans-Air-Congo,Trans-Air-Congo
,TSI,TRANSPORTAIR,Transport'air,Transport'air
S5,TSJ,TRAST AERO,Trast_Aero,Trast Aero
,TSK,TOMSK-AVIA,Trast_Aero,Trast Aero
,TSL,THAI AVIATION,Thai_Aviation_Services,Thai Aviation Services
,SBT,TAFTAN,Taftan_Airlines,Taftan Airlines
9T,ABS,ATHABASKA,Transwest_Air,Transwest Air
,TSM,,Trans_Sayegh_Airport_Services,Trans Sayegh Airport Services
,TSN,AIR TRANS,Trans-Air_Services,Trans-Air Services
UN,TSO,TRANSOVIET,Transaero_Airlines,Transaero Airlines
,TSP,TRANSPO-INTER,Transportes_AÃ©reos_Inter,Transportes Aéreos Inter
,TSR,TRANS SERVICE,Trans_Service_Airlift,Trans Service Airlift
,TSS,TRI-STATE,Tri-State_Aero,Tri-State Aero
,TST,TRAST,TRAST,TRAST
,TSV,TROPIC,Tropair_Airservices,Tropair Airservices
,TSW,SWISSTRANS,Transwings,Transwings
T9,TSX,THAI STAR,Thai_Star_Airlines,Thai Star Airlines
,TSY,TRIPLE STAR,Tristar_Air,Tristar Air
,TTA,KANIMANBO,Sociedade_de_Transporte_e_Trabalho_AÃ©reo,TTA - Sociedade de Transporte e Trabalho Aéreo
,TTC,TRANSTECO,Transteco,Transteco
,TTH,,Tarhan_Tower_Airlines,Tarhan Tower Airlines
,TTL,TOTAL,Total_Linhas_A%C3%A9reas,Total Linhas Aéreas
,TTP,MIGHTY WING,Triple_O_Aviation,Triple O Aviation
,TTR,TRANSPORTACIONES,Transportaciones_Y_Servicios_AÃ©reos,Transportaciones Y Servicios Aéreos
,TTS,TECNICO,Transporte_AÃ©reo_TÃ©cnico_Ejecutivo,Transporte Aéreo Técnico Ejecutivo
T5,TUA,TURKMENISTAN,Turkmenistan_Airlines,Turkmenhovayollary
,TUC,TURICHILE,Turismo_AÃ©reo_de_Chile,Turismo Aéreo de Chile
UG,TUI,,Tuninter,Tuninter
,TUL,URSAL,Tulpar_Air,Tulpar Air
,TUM,TUMTEL,Tyumenspecavia,Tyumenspecavia
,TUO,TURISTICO,Taxi_AÃ©reo_TurÃ­stico,Taxi Aéreo Turístico
,TUX,TULPA,Tulpar_Air_Service,Tulpar Air Service
,TUZ,TUNA,Tuna_Aero,Tuna Aero
,TVA,TRANS-AMERICA,Trans_America_Airlines,Trans America Airlines
,TVH,TRAVASA,Trabajos_AÃ©reos_Vascongados,Trabajos Aéreos Vascongados
,TVI,TIRAMAVIA,Tiramavia,Tiramavia
7O[49],TVL,TRAVEL SERVICE,Travel_Service_(Hungary),Travel Service
,TVO,TRANS-BALLERIO,Transavio,Transavio
T6,TVR,TAVREY,Tavrey_Airlines,Tavrey Airlines
QS,TVS,SKYTRAVEL,Travel_Service_(airline),Travel Service
TW,TWB,TWAYAIR,T%27way_Air,T'way Air
,TWE,TRANSWEDE,Transwede_Airways,Transwede Airways
,TWJ,,Twinjet_Aircraft_Sales,Twinjet Aircraft Sales
,TWL,TRADEWINDS CANADA,Tradewinds_Aviation,Tradewinds Aviation
,TWM,,Transairways,Transairways
,TWO,COLIBRI,Twente_Airlines,Twente Airlines
,TWW,WELWITCHIA,Trans_Air_Welwitchia,Trans Air Welwitchia
,TXA,OKAY AIR,Texair_Charter,Texair Charter
AL,TXC,TRANSEXPORT,TransAVIAexport_Airlines,TransAVIAexport Airlines
,TXE,TRANSAIR EXPRESS,Transilvania_Express,Transilvania Express
,TXL,TAXI COZATL,Taxi_AÃ©reo_Cozatl,Taxi Aéreo Cozatl
,TXM,TAXIMEX,Taxi_AÃ©reo_de_MÃ©xico,Taxi Aéreo de México
,TXN,TEXAS NATIONAL,Texas_National_Airlines,Texas National Airlines
,TXO,TAXIS SINALOA,Taxis_AÃ©reos_de_Sinaloa,Taxis Aéreos de Sinaloa
,TXR,TAXIREY,Taxirey,Taxirey
,TXS,TEXAIR,Texas_Airlines,Texas Airlines
,TXT,TEXAS CHARTER,Texas_Air_Charters,Texas Air Charters
,TXZ,TEX STAR,Tex_Star_Air_Freight,Tex Star Air Freight
,TYF,TAYFLITE,Tayflite,Tayflite
,TYG,TRYGG,Trygg-Flyg,Trygg-Flyg
,TYW,TYROL AMBULANCE,Tyrol_Air_Ambulance,Tyrol Air Ambulance
,TZE,TRANSPORTE SAENZ,Transporte_AÃ©reo_Ernesto_Saenz,Transporte Aéreo Ernesto Saenz
,TZK,TAJIKSTAN,Tajikistan_International_Airlines,Tajikistan International Airlines
,BLI,BLUELINE,Thyssen_Krupp_AG,Thyssen Krupp AG
6B,BLX,BLUESCAN,TUIfly_Nordic,TUIfly Nordic
,BOL,BOL,Transportes_A%C3%A9reos_Bolivianos,Transportes Aéreos Bolivianos
,BOX,BOX,Tiphook_PLC,Tiphook PLC
,CHE,CHECK AIR,Top_Flight_Air_Service,Top Flight Air Service
,CLR,CLINTON AIRWAYS,Trans_International_Airlines,Trans America
,CLU,CAROLUS,Triple_Alpha,Triple Alpha
,CTP,CORTAS,Tashkent_Aircraft_Production_Corporation,Tashkent Aircraft Production Corporation
,CWT,TEXAS AIRWAYS,Texas_Airways,Texas Airways
,DCL,DON CARLOS,Transportes_AÃ©reos_Don_Carlos,Transportes Aéreos Don Carlos
,DOT,DOT TEL,Telnic_Limited,Telnic Limited
,DRC,TRITON AIR,Triton_Airlines,Triton Airlines
DT,DTA,DTA,TAAG_Angola_Airlines,TAAG Angola Airlines
SF,DTH,TASSILI AIR,Tassili_Airlines,Tassili Airlines
,EAR,EJECUTIVO-AEREO,Transporte_Ejecutivo_AÃ©reo,Transporte Ejecutivo Aéreo
,ELV,AEREOS SELVA,Transportes_AÃ©reos_Nacionales_De_Selva_Tans,Transportes Aéreos Nacionales De Selva Tans
,FBO,,TAG_Farnborough_Airport,TAG Farnborough Airport
,FNV,TRANSAVIASERVICE,Transaviaservice,Transaviaservice
,FPG,TAG AVIATION,TAG_Aviation,TAG Aviation
,GFN,GRIFFON,The_955_Preservation_Group,The 955 Preservation Group
,GJB,SKY TRUCK,Trans-Air-Link,Trans-Air-Link
TJ,GPD,GOODSPEED,Tradewind_Aviation,Tradewind Aviation
,HBA,HARBOR AIR,Trail_Lake_Flying_Service,Trail Lake Flying Service
,HET,HELITAF,TAF_Helicopters,TAF Helicopters
,HTO,HELI TANGO,Tango_Bravo,Tango Bravo
,HVK,TURKISH AIRFORCE,Turkish_Air_Force,Turkish Air Force
,IHS,,Thryluthjonustan,Thryluthjonustan
,TBA,TIBET,Tibet_Airlines,Tibet Airlines
,IRF,TA-AIR,TA-Air_Airline,TA-Air Airline
,IRR,TARAIR,Tara_Air_Line,Tara Air Line
,JCH,TRADING CARGO,Trading_Air_Cargo,Trading Air Cargo
,KCA,TRANS-KIEV,Trans-Kiev,Trans-Kiev
,JEL,JETEL,Tal_Air_Charters,Tal Air Charters
,KRA,REGATA,Transcontinental_Airlines,Transcontinental Airlines
,KTB,TRANSBALTIKA,Transaviabaltika,Transaviabaltika
,KTS,KOTAIR,Transair-Gyraintiee,Transair-Gyraintiee
PZ,LAP,PARAGUAYA,TAM_Mercosur,TAM Mercosur
,LCC,LANCAIR,The_Lancair_Company,The Lancair Company
,LEG,LEGACY,The_Army_Aviation_Heritage_Foundation,The Army Aviation Heritage Foundation
,LKW,TOPINTER,Top_Sky_International,Top Sky International
AX,LOF,WATERSKI,Trans_States_Airlines,Trans States Airlines
,LTA,LANTRA,Trans_Atlantis,Trans Atlantis
,MCT,TRANS CORTES,TransportaciÃ³n_AÃ©rea_Del_Mar_De_CortÃ©s,Transportación Aérea Del Mar De Cortés
,MGM,AERO EMM-GEE-EMM,Transporte_Aero_MGM,Transporte Aero MGM
,MOH,MOTH,Tigerfly,Tigerfly
,MPO,AMPARO,Transportes_AÃ©reos_Amparo,Transportes Aéreos Amparo
,MUI,MAUI,Trans_Air,Trans Air
,MXQ,MEXIQUENSES,Transportes_AÃ©reos_Mexiquenses,Transportes Aéreos Mexiquenses
1E,,,Travelsky_Technology,Travelsky Technology
2H,,,Thalys,Thalys
,NTR,NITRO,TNT_International_Aviation,TNT International Aviation
1L,OSY,OPEN SKIES,Treaty_on_Open_Skies,Open Skies Consultative Commission
,PCW,PACIFIC ORIENT,Trans-Pacific_Orient_Airways,Trans-Pacific Orient Airways
,PGS,,Tauranga_Aer_Club,Tauranga Aer Club
,PSS,PROGRESS,TSSKB-Progress,TSSKB-Progress
,RBD,RED BIRD,Trans_World_Express,Trans World Express
,REC,TRANS-RECO,Trans_Reco,Trans Reco
,RMS,TASS AIR,Tas_Aviation,Tas Aviation
RO,ROT,TAROM,Tarom,Tarom
,ROU,ROBINSON CRUSOE,Transportes_AÃ©reos_I._R._Crusoe,Transportes Aéreos I. R. Crusoe
,RRT,SIERRA ALTA,Transportes_AÃ©reos_Sierra,Transportes Aéreos Sierra
,RRY,AIRFERRY,Tbilisi_Aviation_University,Tbilisi Aviation University
,RTM,AERO TRANSAM,Trans_Am_Compania,Trans Am Compania
,SBJ,TRANS SAHARA,Trans_Sahara_Air,Trans Sahara Air
,SEI,TRANSPORTE SIERRA,Transportes_AÃ©reos_Sierra_Madre,Transportes Aéreos Sierra Madre
,SRT,TRASER,Trans_Asian_Airlines,Trans Asian Airlines
3T,URN,TURAN,Turan_Air,Turan Air
T4,TIB,TRIP,TRIP_Linhas_A%C3%A9reas,TRIP Linhas Aéreas
,USB,TUSHETI,Tusheti,Tusheti
,UTM,AVIATAPS,TAPC_Aviatrans_Aircompany,TAPC Aviatrans Aircompany
,UTN,TRANS-ULGII,Trans-Ulgii,Trans-Ulgii
,UTT,ARABIAN TRANSPORT,Transarabian_Transportation_Services,Transarabian Transportation Services
,VEN,TRANSAVEN AIRLINE,Transaven,Transaven
,VIP,SOVEREIGN,Tag_Aviation_UK,Tag Aviation UK
L6,VNZ,TBILAVIA,Tbilaviamsheni,Tbilaviamsheni
,VRC,VERACRUZ,Taxi_de_Veracruz,Taxi de Veracruz
XN,XAR,XPRESS,XpressAir,XpressAir
,XNR,TAXI NORTE,Taxi_Aero_Del_Norte,Taxi Aero Del Norte
VO,TYR,TYROLEAN,Tyrolean_Airways,Tyrolean Airways
TX,TAN,,Transportes_A%C3%A9reos_Nacionales,Transportes Aéreos Nacionales
OR,TFL,ORANGE,TUI_Airlines_Netherlands,TUI Airlines Netherlands
GS,GCR,BO HAI,Tianjin_Airlines,Tianjin Airlines
TZ,,,Tsaradia,Tsaradia
,TUM,UTAIR-CARGO,,UTAir
,UEU,UNITED EUROPEAN,United_European_Airlines,United European Airlines
,UCG,UNIWORLD,Uniworld_Air_Cargo,Uniworld Air Cargo
,CUH,LOULAN,Urumqi_Airlines,Urumqi Airlines
,DOD,,USAF,USAF Air Mobility Operations Control Center
,DOI,INTERIOR,U.S._Department_of_the_Interior,U.S. Department of the Interior
,CNV,CONVOY,,U.S. Navy Reserve Logistic Air Forces
,EXM,EXAM,United_Kingdom_Civil_Aviation_Authority,United Kingdom Civil Aviation Authority
,GIH,TRANSPORT AFRICAIN,Union_des_Transports_Africains_de_Guinee,Union des Transports Africains de Guinee
,GKA,GOLDEN KNIGHTS,US_Army_Parachute_Team,US Army Parachute Team
U5,GWY,GETAWAY,USA3000_Airlines,USA3000 Airlines
B7,UIA,GLORY,UNI_Air,UNI Air
,UAB,UNITED ARABIAN,United_Arabian_Airlines,United Arabian Airlines
UA,UAL,UNITED,United_Airlines,United Airlines
4H,UBD,UNITED BANGLADESH,United_Airways,United Airways
,UAC,UNITAIR,United_Air_Charters,United Air Charters
,UCS,UNITED CARRIERS,United_Carriers_Systems,United Carriers Systems
,UEA,UNITED EAGLE,United_Eagle_Airlines,United Eagle Airlines
U2,UFS,FEEDER EXPRESS,United_Express,United Feeder Service
,CFU,MINAIR,United_Kingdom_Civil_Aviation_Authority,United Kingdom Civil Aviation Authority
,KRF,KITTYHAWK,United_Kingdom_Royal_VIP_Flights,United Kingdom Royal VIP Flights
,KRH,SPARROWHAWK,United_Kingdom_Royal_VIP_Flight,United Kingdom Royal VIP Flight
,SDS,STANDARDS,United_Kingdom_Civil_Aviation_Authority,United Kingdom Civil Aviation Authority
,TQF,RAINBOW,United_Kingdom_Royal_VIP_Flights,United Kingdom Royal VIP Flights
,CGX,COASTGUARD AUXAIR,United_States_Coast_Guard_Auxiliary,United States Coast Guard Auxiliary
,AGR,AGRICULTURE,United_States_Department_Of_Agriculture,United States Department Of Agriculture
,UAD,,University_Air_Squadron,University Air Squadron
,UAJ,,University_Air_Squadron,University Air Squadron
,UAA,,University_Air_Squadron,University Air Squadron
,UAS,,University_Air_Squadron,University Air Squadron
,HBU,KHARKIV UNIVERSAL,Universal_Avia,Universal Avia
,HLE,HELIMED,UK_HEMS,UK HEMS
U7,JUS,JET USA,USA_Jet_Airlines,USA Jet Airlines
,LEA,LEADAIR,Unijet,Unijet
,MSH,MARSHALAIR,US_Marshals_Service,US Marshals Service
,NDU,SIOUX,University_of_North_Dakota,University of North Dakota
,PNA,PACIFIC NORTHERN,Universal_Airlines_(United_States),Universal Airlines
,,UPALI,Upali_Air,Upali Air
,RAU,UGANDA ROYAL,Uganda_Royal_Airways,Uganda Royal Airways
,SAU,UNISERVE,United_Aviation_Services,United Aviation Services
U6,SVR,SVERDLOVSK AIR,Ural_Airlines,Ural Airlines
,TRB,KIROVTRANS,Ukraine_Transavia,Ukraine Transavia
,UAF,UNIFORCE,United_Arab_Emirates_Air_Force,United Arab Emirates Air Force
,UAI,UNAIR,Union_Africaine_des_Transports,Union Africaine des Transports
,UCC,UGANDA CARGO,Uganda_Air_Cargo,Uganda Air Cargo
,UCH,US CHARTER,US_Airports_Air_Charter,US Airports Air Charter
,UCO,UCOAVIACION,Ucoaviacion,Ucoaviacion
,UES,AVIASYSTEM,Ues-Avia_Aircompany,Ues-Avia Aircompany
QU,UGA,UGANDA,Uganda_Airlines,Uganda Airlines
,UGC,URGEMER,Urgemer_Canarias,Urgemer Canarias
,UHL,UKRAINE COPTERS,Ukrainian_Helicopters,Ukrainian Helicopters
,UHS,PILOT AIR,Ulyanovsk_Higher_Civil_Aviation_School,Ulyanovsk Higher Civil Aviation School
,UJR,UNIVERSAL JET,Universal_Jet_Rental_de_Mexico,Universal Jet Rental de Mexico
,UJT,UNI-JET,Universal_Jet_Aviation,Universal Jet Aviation
,UKI,KHALIQ,UK_International_Airlines,UK International Airlines
,UKL,UKRAINE ALLIANCE,Ukraine_Air_Alliance,Ukraine Air Alliance
UF,UKM,UKRAINE MEDITERRANEE,UM_Airlines,UM Airlines
,UKN,ENTERPRISE UKRAINE,Ukraine_Air_Enterprise,Ukraine Air Enterprise
,UKP,POLICE,United_Kingdom_Home_Office,United Kingdom Home Office
6Z,UKS,CARGOTRANS,Ukrainian_Cargo_Airways,Ukrainian Cargo Airways
,ULT,ULTRAIR,Ultrair,Ultrair
,ULH,ULTIMATEHELI,Ultimate_HELI,Ultimate HELI
,ULR,VIPER,Ultimate_Air,Ultimate Air
,UNC,UNICOPTER,Uni-Fly,Uni-Fly
,UNF,UNION FLIGHTS,Union_Flights,Union Flights
,UNJ,PROJET,Universal_Jet,Universal Jet
,UNS,UNSPED,Unsped_Paket_Servisi,Unsped Paket Servisi
,UNU,UNIEURO,Unifly_Servizi_Aerei,Unifly Servizi Aerei
,UPL,PILOT SCHOOL,Ukrainian_Pilot_School,Ukrainian Pilot School
5X,UPS,UPS,United_Parcel_Service,United Parcel Service
,URV,URAI,Uraiavia,Uraiavia
US,AWE,CACTUS,US_Airways,US Airways
,USF,AFRICA EXPRESS,USAfrica_Airways,USAfrica Airways
UH,USH,US-HELI,US_Helicopter,US Helicopter
,USJ,USJET,US_Jet,US Jet
,USX,AIR EXPRESS,,US Express
UT,UTA,UTAIR,UTair_Aviation,UTair Aviation
,UTG,UTAGE,UTAGE,UTAGE
,UTR,AIRUT,Utair_South_Africa,Utair South Africa
,UTS,AIRRUH,Ukrainian_State_Air_Traffic_Service_Enterprise,Ukrainian State Air Traffic Service Enterprise
,UTU,,Urartu-Air,Urartu-Air
,UVA,UNIVERSAL,Universal_Airways,Universal Airways
,UVG,GUYANA JET,Universal_Airlines_(Guyana),Universal Airlines
,UVM,UVAVEMEX,Uvavemex,Uvavemex
,AIO,AIR CHIEF,United_States_Air_Force,United States Air Force
,UVN,UNITED AVIATION,United_Aviation,United Aviation
HY,UZB,UZBEK,Uzbekistan_Airways,Uzbekistan Airways
PS,AUI,UKRAINE INTERNATIONAL,Ukraine_International_Airlines,Ukraine International Airlines
,WEC,AIRGO,Universal_Airlines_(United_States),Universal Airlines
US,,,Unavia_Suisse,Unavia Suisse
US,USA,CACTUS,US_Airways,US Airways
,QID,QUID,100th_Air_Refueling_Wing,USAF 100th Air Refueling Wing
,UIT,ARCTIC,University_of_TromsÃ¸_School_of_Aviation,University of Tromsø School of Aviation
JW,VNL,VANILLA,Vanilla_Air,Vanilla Air
,VAI,,Volant_Aviation_International,Volant Aviation International
,VAR,VECA,Veca_Airlines,Veca Airlines
,VLR,VOLAX,Volare_22_X,Volare 22 X
,VDR,VOLDIR,Voldirect,Voldirect
,VVV,VALAIRJET,Valair_AviaÃ§Ã£o_Lda,Valair Aviação Lda
VB,VIV,AEROENLACES,Vivaaerobus.com,VIVA Aerobus
,VIL,TURTLE DOVE,V_I_Airlink,V I Airlink
VA,VOZ,VELOCITY,Virgin_Australia,Virgin Australia Airlines
,VBA,VEEBEE,V_Bird_Airlines_Netherlands,V Bird Airlines Netherlands
,WIW,VEE-AVIA,V-avia_Airline,V-avia Airline
,VBD,VEEBIRD-AVIA,V-Berd-Avia,V-Berd-Avia
,VAC,VACATIONAIR,Vacationair,Vacationair
,RDW,ROADWATCH,Valair_AG_(Helicoptere),Valair AG (Helicoptere)
,VLA,NALAU,Valan_International_Cargo_Charter,Valan International Cargo Charter
,VLN,VALAN,Valan_Limited,Valan Limited
,EHR,ROTOR,Valfell-Verkflug,Valfell-Verkflug
VF,VLU,VALUAIR,Valuair,Valuair
J7,VJA,CRITTER,ValuJet_Airlines,ValuJet Airlines
,VAA,EUROVAN,Van_Air_Europe,Van Air Europe
,VGC,VANGUARDIA COLIMA,Vanguardia_en_AviaciÃ³n_en_Colima,Vanguardia en Aviación en Colima
,VGD,VANGUARD AIR,Vanguard_Airlines,Vanguard Airlines
0V,VFC,VASCO AIR,Vietnam_Air_Services_Company,Vietnam Air Services Company (VASCO)
,VAG,SEGA,Vega,Vega
,WGA,WEGA FRANKO,Vega_Air_Company,Vega Air Company
,WEL,VELES,"Veles,_Ukrainian_Aviation_Company","Veles, Ukrainian Aviation Company"
,VTX,VERATAXIS,Verataxis,Verataxis
,BTP,NET RAIL,Veritair,Veritair Ltd
VC,VAL,VOYAGEUR,Voyageur_Airways,Voyageur Airways
,GRV,NIGHT RIDER,Vernicos_Aviation,Vernicos Aviation
VN,HVN,VIET NAM AIRLINES,Vietnam_Airlines,Vietnam Airlines
,KWA,VOZAIR,Vozdushnaya_Academy,Vozdushnaya Academy
NN,MOV,MOV AIR,VIM_Airlines,VIM Airlines
,MVY,,VIM-Aviaservice,VIM-Aviaservice
2R,,,Via_Rail_Canada,Via Rail Canada
,ENV,ENDEAVOUR,Victoria_Aviation,Victoria Aviation
,VCT,VISCOUNT AIR,Viscount_Air_Service,Viscount Air Service
,SSI,SUPER JET,Vision_Airlines_(SSI),Vision Airlines
,FXF,FOX FLIGHT,VIP_Air_Charter,VIP Air Charter
,PAV,NICOL,VIP_Avia_(PAV),VIP Avia
,PRX,PAREX,VIP_Avia_(PRX),VIP Avia
,VAT,VISIONAIR,Visionair,Visionair
VA,,,Viasa,Viasa
,VCA,VICA,VICA_-_Viacao_Charter_AÃ©reos,VICA - Viacao Charter Aéreos
,VCM,CARMEN,Volare_Air_Charter_Company,Volare Air Charter Company
Y4,VOI,VOLARIS,Volaris,Volaris
VI,VDA,VOLGA,Volga-Dnepr_Airlines,Volga-Dnepr Airlines
,VEA,VEGA AIRLINES,Vega_Airlines,Vega Airlines
,VEC,VECAR,Venescar_Internacional,Venescar Internacional
,VEE,VICTOR ECHO,Victor_Echo,Victor Echo
,VEI,GREEN ISLE,Virgin_Express_Ireland,Virgin Express Ireland
VX,VRD,REDWOOD,Virgin_America,Virgin America
VJ,VJC,VIETJET,Vietjet_Air,Vietjet Air
V4,VES,VIEQUES,Vieques_Air_Link,Vieques Air Link
TV,VEX,VIRGIN EXPRESS,Virgin_Express,Virgin Express
,VFT,ZETA FLIGHTS,VZ_Flights,VZ Flights
VK,VGN,VIRGIN NIGERIA,Virgin_Nigeria_Airways,Virgin Nigeria Airways
,VGV,VOLOGDA AIR,Vologda_State_Air_Enterprise,Vologda State Air Enterprise
,VHA,AIR V-H,VH-Air_Industrie,VH-Air Industrie
,VHM,EARLY BIRD,VHM_Schul-und-Charterflug,VHM Schul-und-Charterflug
,VIB,VITUS,Vibroair_Flugservice,Vibroair Flugservice
,VIC,VIP-EJECUTIVO,VIP_Servicios_AÃ©reos_Ejecutivos,VIP Servicios Aéreos Ejecutivos
,VIE,VIP EMPRESARIAL,VIP_Empresarial,VIP Empresarial
,VIF,VIENNA FLIGHT,VIF_Luftahrt,VIF Luftahrt
,VIG,VEGA AVIATION,Vega_Aviation,Vega Aviation
,VIH,VICHI,Vichi,Vichi
,VIK,SWEDJET,Viking_Airlines,Viking Airlines
,VIN,VINAIR,Vinair_AeroserviÃ§os,Vinair Aeroserviços
VS,VIR,VIRGIN,Virgin_Atlantic_Airways,Virgin Atlantic Airways
,VJM,VIAJES MEXICANOS,Viajes_Ejecutivos_Mexicanos,Viajes Ejecutivos Mexicanos
,VJT,VISTA,Vistajet,Vistajet
,VJT,VISTA MALTA,Vistajet,Vistajet
ZG,VVM,JACKPOT,Viva_Macau,Viva Macau
VE,VLE,VOLA,C.A.I._Second,C.A.I. Second
VY,VLG,VUELING,Vueling_Airlines,Vueling Airlines
XF,VLK,VLADAIR,Vladivostok_Air,Vladivostok Air
LC,VLO,VELOG,Varig_Log%C3%ADstica,Varig Logística
,VLT,VERTICAL,Vertical-T_Air_Company,Vertical-T Air Company
,VMA,VERO MONMOUTH,Vero_Monmouth_Airlines,Vero Monmouth Airlines
,VNK,,Vipport,Vipport Joint Stock Company
VM,VOA,VIAGGIO,Viaggio_Air,Viaggio Air
,VOG,VOYAGER AIR,Voyager_Airlines,Voyager Airlines
9V,VPA,VIAIR,Vipair_Airlines,Vipair Airlines
,VPB,VETERAN,Veteran_Air,Veteran Air
,VPV,VIP AVIA,VIP-Avia,VIP-Avia
,VRA,VERITAIR,Vertair,Vertair
,VRE,UKRAINE VOLARE,Volare_Airlines_(Ukraine),Volare Airlines
,VRL,VOAR LINHAS,Voar_Lda,Voar Lda
RG,VRN,VARIG,VRG_Linhas_A%C3%A9reas,VRG Linhas Aéreas
FC,VVC,VivaColombia,VivaColombia,VivaColombia
,VSB,VICKERS,Vickers_Limited,Vickers Limited
,VSN,VISION,Vision_Airways_Corporation,Vision Airways Corporation
,VSO,VASO,Voronezh_Aircraft_Manufacturing_Society,Voronezh Aircraft Manufacturing Society
VP,VSP,VASP,VASP,VASP
,VSS,WATERBIRD,Virign_Islands_Seaplane_Shuttle,Virign Islands Seaplane Shuttle
,VTC,VUELOS TOLLOCAN,Vuelos_Especializados_Tollocan,Vuelos Especializados Tollocan
,VTH,VUELOS TEHUACAN,Vuelos_Corporativos_de_Tehuacan,Vuelos Corporativos de Tehuacan
V7,VOE,VOLOTEA,Volotea,Volotea
,VTK,VOSTOK,Vostok_Airlines,Vostok Airlines
,VTL,VITALA,Victor_Tagle_Larrain,Victor Tagle Larrain
,VTV,VOINTEH,Vointeh,Vointeh
,VUR,VIPEC,Vuelos_Internos_Privados_VIP,Vuelos Internos Privados VIP
,VUS,VUELA BUS,Vuela_Bus,Vuela Bus
,VZL,VZLYET,Vzlyet,Vzlyet
VG,VLM,RUBENS,VLM_Airlines,VLM Airlines
,WCY,TITAN AIR,Viking_Express,Viking Express
,WEV,VICTORIA UGANDA,Victoria_International_Airways,Victoria International Airways
G6,WLG,GOUMRAK,Air_Volga,Air Volga
,VNR,VIENNAIR,Viennair,Viennair
UK,VTI,Vistara,Vistara,Vistara
,WGN,WESTERN GLOBAL,Western_Global_Airlines,Western Global Airlines
,WCC,WEST COAST,West_Coast_Charters,West Coast Charters
WD,WDL,WDL,WDL_Aviation,WDL Aviation
,WRR,WRAP AIR,WRA_Inc,WRA Inc
,XWS,,WSI_Corporation,WSI Corporation
,CGG,CHARGE,Walmart_Aviation,Walmart Aviation
,WAS,WALSTEN,Walsten_Air_Services,Walsten Air Services
,GOT,GOTHIC,WaltAir,WaltAir
,WPT,WAPITI,Wapiti_Aviation,Wapiti Aviation
,WAV,WARBELOW,Warbelow%27s_Air_Ventures,Warbelow's Air Ventures
,ATX,AIRTAX,Warwickshire_Aerocentre_Ltd.,Warwickshire Aerocentre Ltd.
WT,WSG,WASAYA,Wasaya_Airways,Wasaya Airways
,,WAYRAPERÚ,Wayraper%C3%BA,Wayraperú
,WTC,WATCO,Weasua_Air_Transport_Company,Weasua Air Transport Company
WH,WEB,WEB-BRASIL,WebJet_Linhas_A%C3%A9reas,WebJet Linhas Aéreas
,TDB,THUNDER BAY,Welch_Aviation,Welch Aviation
2W,WLC,WELCOMEAIR,Welcome_Air,Welcome Air
,BLW,BLUESTAR,Wermlandsflyg_AB,Wermlandsflyg AB
,WCB,KILO YANKEE,West_Africa_Airlines,West Africa Airlines
,WTF,WESTAF AIRTRANS,West_African_Air_Transport,West African Air Transport
WZ,WSF,,West_African_Airlines,West African Airlines
,WAC,WESTAF CARGO,West_African_Cargo_Airlines,West African Cargo Airlines
,WLX,WEST LUX,West_Air_Luxembourg,West Air Luxembourg
PT,SWN,AIR SWEDEN,West_Air_Sweden,West Air Sweden
YH,WCW,WEST,West_Caribbean_Airways,West Caribbean Airways
,WCR,WEST CARIBBEAN,West_Caribbean_Costa_Rica,West Caribbean Costa Rica
8O,YWZ,COAST AIR,West_Coast_Air,West Coast Air
,,,West_Coast_Airlines,West Coast Airlines
,WCG,WHISKY INDIA,West_Coast_Airlines,West Coast Airlines
,WCA,WEST-LEONE,West_Coast_Airways,West Coast Airways
,TEE,TEEBIRD,West_Freugh_DTEO,West Freugh DTEO
,WEW,WESTWIND,West_Wind_Aviation,West Wind Aviation
WS,WJA,WESTJET,WestJet,WestJet
,WAA,WESTAIR WINGS,Westair_Aviation_(Namibia),Westair Aviation
,WSC,WESTCAR,Westair_Cargo_Airlines,Westair Cargo Airlines
,PCM,PAC VALLEY,Westair_Industries,Westair Industries
,BLK,BLUE FLAME,Westcoast_Energy,Westcoast Energy
XP,CXP,RUBY MOUNTAIN,Western_(airline),Western
,WST,WESTERN BAHAMAS,Western_Air,Western Air
,NPC,NORPAC,Western_Air_Couriers,Western Air Couriers
,WAE,WESTERN EXPRESS,Western_Air_Express,Western Air Express
WA,WAL,WESTERN,Western_Airlines,Western Airlines
WA,KLC,CITY,KLM_Cityhopper,KLM Cityhopper
,WAL,WESTERN ARCTIC,Western_Arctic_Air,Western Arctic Air
,WTV,WESTAVIA,Western_Aviators,Western Aviators
,WES,WEST EX,Western_Express_Air_Lines,Western Express Air Lines
W7,KMR,KOMSTAR,Western_Pacific_Airlines,Western Pacific Airlines
,WPA,WESTPAC,Western_Pacific_Airservice,Western Pacific Airservice
,WSL,WEST LINE,Westflight_Aviation,Westflight Aviation
,WSA,WESTATES,Westgates_Airlines,Westgates Airlines
,WHE,WESTLAND,Westland_Helicopters,Westland Helicopters
,WTP,WESTPOINT,Westpoint_Air,Westpoint Air
CN,WWD,WESTWARD,Westward_Airways_(Nebraska),Westward Airways
,WHT,WHITEJET,White_(airline),White
,WEA,WHITE EAGLE,White_Eagle_Aviation,White Eagle Aviation
,WRA,,White_River_Air_Services,White River Air Services
,WWL,,Whyalla_Airlines,Whyalla Airlines
WF,WIF,WIDEROE,Wider%C3%B8e,Widerøe
,WIG,WIGGINS AIRWAYS,Wiggins_Airways,Wiggins Airways
,WHS,WEEKING,Wiking_Helikopter_Service,Wiking Helikopter Service
,WFO,WILBURS,Wilbur's_Flight_Operations,Wilbur's Flight Operations
,WLS,WALES,Air_Wales_Virtual,Air Wales Virtual
,WGP,GRAND PRIX,Williams_Grand_Prix_Engineering,Williams Grand Prix Engineering
,WDA,WIMBI DIRA,Wimbi_Dira_Airways,Wimbi Dira Airways
,WNA,WINAIR,Winair,Winair
IV,JET,GHIBLI,Wind_Jet,Wind Jet
,WSI,WIND SPIRIT,Wind_Spirit_Air,Wind Spirit Air
7W,QGA,QUADRIGA,Windrose_Air,Windrose Air
,WIA,WINDWARD,Windward_Islands_Airways_International,Windward Islands Airways International
IW,WON,WINGS ABADI,Wings_Air,Wings Air
,WAT,,Wings_Air_Transport,Wings Air Transport
,WAW,WING SHUTTLE,Wings_Airways,Wings Airways
,WOL,WINGJET,Wings_Aviation,Wings Aviation
,WEX,WINGS EXPRESS,Wings_Express,Wings Express
,WLB,WING LEBANON,Wings_of_Lebanon_Aviation,Wings of Lebanon Aviation
,WIN,WINLINK,Winlink,Winlink
,WAG,,Wisconsin_Air_National_Guard,Wisconsin Air National Guard
,WSM,WISMAN,Wisman_Aviation,Wisman Aviation
8Z,WVL,WIZZBUL,Wizz_Air_Bulgaria,Wizz Air Bulgaria
W6,WZZ,WIZZAIR,Wizz_Air,Wizz Air
,WNR,WONDAIR,Wondair_on_Demand_Aviation,Wondair on Demand Aviation
,CWY,CAUSEWAY,Woodgate_Aviation,Woodgate Aviation
WO,WOA,WORLD,World_Airways,World Airways
,XWW,,World_Weatherwatch,World Weatherwatch
,WWM,MANAS WING,World_Wing_Aviation,World Wing Aviation
1P,,,Worldspan,Worldspan
UI,CSW,SILKITALIA,SW_Italia,SW Italia
,WWS,,Worldwide_Aviation_Services,Worldwide Aviation Services
,WWI,WORLDWIDE,Worldwide_Jet_Charter,Worldwide Jet Charter
WW,WOW,WOW air,WOW_air,WOW air
,WRT,WRIGHT-AIR,Wright_Airlines,Wright Airlines
8V,WRF,WRIGHT FLYER,Wright_Air_Service,Wright Air Service
,CWU,WUHAN AIR,Wuhan_Airlines,Wuhan Airlines
,WYC,WYCOMBE,Wycombe_Air_Centre,Wycombe Air Centre
,WYG,WYOMING,Wyoming_Airlines,Wyoming Airlines
KW,WAN,WATANIYA,Wataniya_Airways,Wataniya Airways
3W*,VNR,WANAIR,Wan_Air,Wan Air
WR,WEN,ENCORE,WestJet_Encore,WestJet Encore
,XJE,,X-Jet,X-Jet
,XAB,AERO XABRE,Xabre_Aerolineas,Xabre Aerolineas
,XAE,AURA,Xair,Xair
,XJC,EXCLUSIVE JET,XJC,XJC Limited
,XER,XEROX,Xerox_Corporation,Xerox Corporation
7A,XRC,TUNISIA CARGO,Express_Air_Cargo,Express Air Cargo
MF,CXA,XIAMEN AIR,Xiamen_Airlines,Xiamen Airlines
,CXJ,XINJIANG,Xinjiang_Airlines,Xinjiang Airlines
,XJT,XRAY,Xjet,Xjet Limited
SE,SEU,STARWAY,XL_Airways_France,XL Airways France
,GXL,STARDUST,XL_Airways_Germany,XL Airways Germany
,XOJ,EXOJET,XOJet,XOJet
,XPS,XP PARCEL,XP_International,XP International
,,BIGSPLASH,Xpedite,Xpedite
,RAG,RAGLAN,Xstrata_Nickel,Xstrata Nickel (Raglan Mine)
XP,CXP,CASINO EXPRESS,Xtra_Airways,Xtra Airways
,DGA,YELLOW RIVER,Yellow_River_Delta_General_Aviation,Yellow River Delta General Aviation
,YRG,YAKAIR GEORGIA,Yak_Air,Yak Air
,AKY,YAK-SERVICE,Yak-Service,Yak-Service
,YAK,YAK AVIA,Yakolev,Yakolev
YL,LLM,YAMAL,Yamal_Airlines,Yamal Airlines
,YAK,YAK AVIA,Yakolev,Yakolev
R3,SYL,AIR YAKUTIA,Yakutia_Airlines,Yakutia Airlines
,CYG,VICAIR,Yana_Airlines,Yana Airlines
,AYG,AIR YANGON,Yangon_Airways,Yangon Airways
Y8,YZR,YANGTZE RIVER,Yangtze_River_Express,Yangtze River Express
,LYH,HELIGUYANE,Yankee_Lima_Helicopteres,Yankee Lima Helicopteres
,MHD,YAS AIR,Yas_Air_Kish,Yas Air Kish
Y0,EMJ,,Yellow_Air_Taxi/Friendship_Airways,Yellow Air Taxi/Friendship Airways
,ELW,YELLOW WINGS,Yellow_Wings_Air_Services,Yellow Wings Air Services
IY,IYE,YEMENI,Yemenia,Yemenia
,ERV,YEREVAN-AVIA,Yerevan-Avia,Yerevan-Avia
,NYT,YETI AIRLINES,Yeti_Airlines,Yeti Airlines
,YFS,YOUNG AIR,Young_Flying_Service,Young Flying Service
,AYU,,Yuhi_Air_Lines,Yuhi Air Lines
,AYE,AIR YING AN,Yunnan_Yingan_Airlines,Yunnan Yingan Airlines
4Y,UYA,,Yute_Air_Alaska,Yute Air Alaska
,UGN,PLUTON,Yuzhnaya_Aircompany,Yuzhnaya Aircompany
2N,UMK,YUZMASH,Yuzhmashavia,Yuzhmashavia
,BZE,ZENSTAR,Zenith_Aviation,Zenith Aviation
,AZB,ZAAB AIR,Zaab_Air,Zaab Air
,AZJ,,Zas_Air,Zas Air
,AZR,ZENAIR,Zenith_Air,Zenith Air
,CDC,HUALONG,Zhejiang_Loong_Airlines,Zhejiang Loong Airlines
,CIT,ZANE,Zanesville_Aviation,Zanesville Aviation
,EMR,ZENMOUR,Zenmour_Airlines,Zenmour Airlines
,EZD,ZEST AIRWAYS,Zest_Airways,Zest Airways
,GZQ,ZAGROS,Zagros_Air,Zagros Air
C4,IMX,ZIMEX,Zimex_Aviation,Zimex Aviation
,IZG,ZAGROS,Zagros_Airlines,Zagros Airlines
,JTU,ZHETYSU,Zhetysu,Zhetysu
,KVZ,,Z-Aero_Airlines,Z-Aero Airlines
,KZH,,Zhez_Air,Zhez Air
,MBG,CHALGROVE,Zephyr_Aviation,Zephyr Aviation
,MLU,MALI LOSINJ,Zracno_Pristaniste_Mali_Losinj,Zracno Pristaniste Mali Losinj
,MZE,,Zenith_Aviation_(Malta),Zenith Aviation (Malta)
,PZY,ZAPOLYARYE,Zapolyarye_Airline_Company,Zapolyarye Airline Company
,RZR,RECOVERY,Zephyr_Express,Zephyr Express
,RZU,ZHERSU AVIA,Zhersu_Avia,Zhersu Avia
,SYZ,ZIL AIR,Zil_Air,Zil Air
,TAN,ZANAIR,Zanair,Zanair
,ZAI,ZASAIR,Zaire_Aero_Service,Zaire Aero Service
,ZAK,ZAMBIA SKIES,Zambia_Skyways,Zambia Skyways
,ZAR,ZAIREAN,Zairean_Airlines,Zairean Airlines
,ZAV,ZETAVIA,Zetavia,Zetavia
,ZAW,ZED AIR,Zoom_Airways,Zoom Airways
,ZMA,ZAMBEZI WINGS,Zambezi_Airlines,Zambezi Airlines
,RZV,ZEDAVIA,Z-Avia,Z-Avia
Q3,MBN,ZAMBIANA,Zambian_Airways,Zambian Airways
,ZAN,ZANTOP,Zantop_International_Airlines,Zantop International Airlines
,ZAS,ZAS AIRLINES,ZAS_Airlines_of_Egypt,ZAS Airlines of Egypt
,CJG,ZHEJIANG,Zhejiang_Airlines,Zhejiang Airlines
,CFZ,ZHONGFEI,Zhongfei_General_Aviation,Zhongfei General Aviation
,CYN,ZHONGYUAN,Zhongyuan_Airlines,Zhongyuan Aviation
3J,WZP,ZIPPER,Zip_(airline),Zip
Z4,OOM,ZOOM,Zoom_Airlines,Zoom Airlines
,ORZ,ZOREX,Zorex,Zorex`
