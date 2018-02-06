package main

import (
	"dolphin/api"
	"dolphin/api/bolt"
	"dolphin/api/cli"
	"dolphin/api/cron"
	"dolphin/api/crypto"
	"dolphin/api/file"
	"dolphin/api/http"
	"dolphin/api/jwt"
	"dolphin/api/ldap"
	"dolphin/api/libcompose"

	"log"

)

func initCLI() *dockm.CLIFlags {
	var cli dockm.CLIService = &cli.Service{}
	flags, err := cli.ParseFlags(dockm.APIVersion)
	if err != nil {
		log.Fatal(err)
	}

	err = cli.ValidateFlags(flags)
	if err != nil {
		log.Fatal(err)
	}
	return flags
}

func initFileService(dataStorePath string) dockm.FileService {
	fileService, err := file.NewService(dataStorePath, "")
	if err != nil {
		log.Fatal(err)
	}
	return fileService
}

func initStore(dataStorePath string) *bolt.Store {
	store, err := bolt.NewStore(dataStorePath)
	if err != nil {
		log.Fatal(err)
	}

	err = store.Open()
	if err != nil {
		log.Fatal(err)
	}

	err = store.MigrateData()
	if err != nil {
		log.Fatal(err)
	}
	return store
}

func initStackManager() dockm.StackManager {
	return libcompose.NewStackManager()
}

func initJWTService(authenticationEnabled bool) dockm.JWTService {
	if authenticationEnabled {
		jwtService, err := jwt.NewService()
		if err != nil {
			log.Fatal(err)
		}
		return jwtService
	}
	return nil
}

func initCryptoService() dockm.CryptoService {
	return &crypto.Service{}
}

func initLDAPService() dockm.LDAPService {
	return &ldap.Service{}
}

func initEndpointWatcher(endpointService dockm.EndpointService, externalEnpointFile string, syncInterval string) bool {
	authorizeEndpointMgmt := true
	if externalEnpointFile != "" {
		authorizeEndpointMgmt = false
		log.Println("Using external endpoint definition. Endpoint management via the API will be disabled.")
		endpointWatcher := cron.NewWatcher(endpointService, syncInterval)
		err := endpointWatcher.WatchEndpointFile(externalEnpointFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	return authorizeEndpointMgmt
}

func initStatus(authorizeEndpointMgmt bool, flags *dockm.CLIFlags) *dockm.Status {
	return &dockm.Status{
		Analytics:          !*flags.NoAnalytics,
		Authentication:     !*flags.NoAuth,
		EndpointManagement: authorizeEndpointMgmt,
		Version:            dockm.APIVersion,
	}
}

func initDockerHub(dockerHubService dockm.DockerHubService) error {
	_, err := dockerHubService.DockerHub()
	if err == dockm.ErrDockerHubNotFound {
		dockerhub := &dockm.DockerHub{
			Authentication: false,
			Username:       "",
			Password:       "",
		}
		return dockerHubService.StoreDockerHub(dockerhub)
	} else if err != nil {
		return err
	}

	return nil
}

func initSettings(settingsService dockm.SettingsService, flags *dockm.CLIFlags) error {
	_, err := settingsService.Settings()
	if err == dockm.ErrSettingsNotFound {
		settings := &dockm.Settings{
			LogoURL:                     *flags.Logo,
			DisplayExternalContributors: true,
			AuthenticationMethod:        dockm.AuthenticationInternal,
			LDAPSettings: dockm.LDAPSettings{
				TLSConfig: dockm.TLSConfiguration{},
				SearchSettings: []dockm.LDAPSearchSettings{
					dockm.LDAPSearchSettings{},
				},
			},
		}

		if *flags.Templates != "" {
			settings.TemplatesURL = *flags.Templates
		} else {
			settings.TemplatesURL = dockm.DefaultTemplatesURL
		}

		if *flags.Labels != nil {
			settings.BlackListedLabels = *flags.Labels
		} else {
			settings.BlackListedLabels = make([]dockm.Pair, 0)
		}

		return settingsService.StoreSettings(settings)
	} else if err != nil {
		return err
	}

	return nil
}

func retrieveFirstEndpointFromDatabase(endpointService dockm.EndpointService) *dockm.Endpoint {
	endpoints, err := endpointService.Endpoints()
	if err != nil {
		log.Fatal(err)
	}
	return &endpoints[0]
}

func main() {
	flags := initCLI()

	fileService := initFileService(*flags.Data)

	store := initStore(*flags.Data)
	defer store.Close()

	stackManager := initStackManager()

	jwtService := initJWTService(!*flags.NoAuth)

	cryptoService := initCryptoService()

	ldapService := initLDAPService()

	authorizeEndpointMgmt := initEndpointWatcher(store.EndpointService, *flags.ExternalEndpoints, *flags.SyncInterval)

	err := initSettings(store.SettingsService, flags)
	if err != nil {
		log.Fatal(err)
	}

	err = initDockerHub(store.DockerHubService)
	if err != nil {
		log.Fatal(err)
	}

	applicationStatus := initStatus(authorizeEndpointMgmt, flags)

	if *flags.Endpoint != "" {
		var endpoints []dockm.Endpoint
		endpoints, err := store.EndpointService.Endpoints()
		if err != nil {
			log.Fatal(err)
		}
		if len(endpoints) == 0 {
			endpoint := &dockm.Endpoint{
				Name: "primary",
				URL:  *flags.Endpoint,
				TLSConfig: dockm.TLSConfiguration{
					TLS:           *flags.TLSVerify,
					TLSSkipVerify: false,
					TLSCACertPath: *flags.TLSCacert,
					TLSCertPath:   *flags.TLSCert,
					TLSKeyPath:    *flags.TLSKey,
				},
				AuthorizedUsers: []dockm.UserID{},
				AuthorizedTeams: []dockm.TeamID{},
			}
			err = store.EndpointService.CreateEndpoint(endpoint)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Println("Instance already has defined endpoints. Skipping the endpoint defined via CLI.")
		}
	}

	if *flags.AdminPassword != "" {
		log.Printf("Creating admin user with password hash %s", *flags.AdminPassword)
		user := &dockm.User{
			Username: "admin",
			Role:     dockm.AdministratorRole,
			Password: *flags.AdminPassword,
		}
		err := store.UserService.CreateUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}

	var server dockm.Server = &http.Server{
		Status:                 applicationStatus,
		BindAddress:            *flags.Addr,
		AssetsPath:             *flags.Assets,
		AuthDisabled:           *flags.NoAuth,
		EndpointManagement:     authorizeEndpointMgmt,
		UserService:            store.UserService,
		TeamService:            store.TeamService,
		TeamMembershipService:  store.TeamMembershipService,
		EndpointService:        store.EndpointService,
		AppToContainerService:  store.AppToContainerService,//click2cloud-apptocontainer
		OrchestrationService:  	store.OrchestrationService,//click2cloud-orchestration
		ResourceControlService: store.ResourceControlService,
		SettingsService:        store.SettingsService,
		RegistryService:        store.RegistryService,
		DockerHubService:       store.DockerHubService,
		StackService:           store.StackService,
		StackManager:           stackManager,
		CryptoService:          cryptoService,
		JWTService:             jwtService,
		FileService:            fileService,
		LDAPService:            ldapService,
		SSL:                    *flags.SSL,
		SSLCert:                *flags.SSLCert,
		SSLKey:                 *flags.SSLKey,
	}

	log.Printf("Starting dockm %s on %s",dockm.APIVersion, *flags.Addr)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
