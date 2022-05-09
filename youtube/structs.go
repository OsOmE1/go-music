package youtube

type streamingData struct {
	ExpiresInSeconds string `json:"expiresInSeconds"`
	Formats          []struct {
		Itag             int    `json:"itag"`
		MimeType         string `json:"mimeType"`
		Bitrate          int    `json:"bitrate"`
		Width            int    `json:"width"`
		Height           int    `json:"height"`
		LastModified     string `json:"lastModified"`
		ContentLength    string `json:"contentLength,omitempty"`
		Quality          string `json:"quality"`
		Fps              int    `json:"fps"`
		QualityLabel     string `json:"qualityLabel"`
		ProjectionType   string `json:"projectionType"`
		AverageBitrate   int    `json:"averageBitrate,omitempty"`
		AudioQuality     string `json:"audioQuality"`
		ApproxDurationMs string `json:"approxDurationMs"`
		AudioSampleRate  string `json:"audioSampleRate"`
		AudioChannels    int    `json:"audioChannels"`
		SignatureCipher  string `json:"signatureCipher"`
	} `json:"formats"`
	AdaptiveFormats []struct {
		Itag      int    `json:"itag"`
		MimeType  string `json:"mimeType"`
		Bitrate   int    `json:"bitrate"`
		Width     int    `json:"width,omitempty"`
		Height    int    `json:"height,omitempty"`
		InitRange struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"initRange"`
		IndexRange struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"indexRange"`
		LastModified     string `json:"lastModified"`
		ContentLength    string `json:"contentLength"`
		Quality          string `json:"quality"`
		Fps              int    `json:"fps,omitempty"`
		QualityLabel     string `json:"qualityLabel,omitempty"`
		ProjectionType   string `json:"projectionType"`
		AverageBitrate   int    `json:"averageBitrate"`
		ApproxDurationMs string `json:"approxDurationMs"`
		SignatureCipher  string `json:"signatureCipher"`
		Url              string `json:"url"`
		ColorInfo        struct {
			Primaries               string `json:"primaries"`
			TransferCharacteristics string `json:"transferCharacteristics"`
			MatrixCoefficients      string `json:"matrixCoefficients"`
		} `json:"colorInfo,omitempty"`
		HighReplication bool    `json:"highReplication,omitempty"`
		AudioQuality    string  `json:"audioQuality,omitempty"`
		AudioSampleRate string  `json:"audioSampleRate,omitempty"`
		AudioChannels   int     `json:"audioChannels,omitempty"`
		LoudnessDb      float64 `json:"loudnessDb,omitempty"`
	} `json:"adaptiveFormats"`
}

type youtubePlayerResponse struct {
	ResponseContext struct {
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MainAppWebResponseContext struct {
			LoggedOut bool `json:"loggedOut"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	TrackingParams    string `json:"trackingParams"`
	PlayabilityStatus struct {
		Status          string `json:"status"`
		PlayableInEmbed bool   `json:"playableInEmbed"`
		Miniplayer      struct {
			MiniplayerRenderer struct {
				PlaybackMode string `json:"playbackMode"`
			} `json:"miniplayerRenderer"`
		} `json:"miniplayer"`
		ContextParams string `json:"contextParams"`
	} `json:"playabilityStatus"`
	StreamingData    streamingData `json:"streamingData"`
	PlaybackTracking struct {
		VideostatsPlaybackURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"videostatsPlaybackUrl"`
		VideostatsDelayplayURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"videostatsDelayplayUrl"`
		VideostatsWatchtimeURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"videostatsWatchtimeUrl"`
		PtrackingURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"ptrackingUrl"`
		QoeURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"qoeUrl"`
		AtrURL struct {
			BaseURL                 string `json:"baseUrl"`
			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
		} `json:"atrUrl"`
		VideostatsScheduledFlushWalltimeSeconds []int `json:"videostatsScheduledFlushWalltimeSeconds"`
		VideostatsDefaultFlushIntervalSeconds   int   `json:"videostatsDefaultFlushIntervalSeconds"`
	} `json:"playbackTracking"`
	Captions struct {
		PlayerCaptionsRenderer struct {
			BaseURL    string `json:"baseUrl"`
			Visibility string `json:"visibility"`
		} `json:"playerCaptionsRenderer"`
		PlayerCaptionsTracklistRenderer struct {
			CaptionTracks []struct {
				BaseURL string `json:"baseUrl"`
				Name    struct {
					SimpleText string `json:"simpleText"`
				} `json:"name"`
				VssID          string `json:"vssId"`
				LanguageCode   string `json:"languageCode"`
				Kind           string `json:"kind"`
				IsTranslatable bool   `json:"isTranslatable"`
			} `json:"captionTracks"`
			AudioTracks []struct {
				CaptionTrackIndices []int `json:"captionTrackIndices"`
			} `json:"audioTracks"`
			TranslationLanguages []struct {
				LanguageCode string `json:"languageCode"`
				LanguageName struct {
					SimpleText string `json:"simpleText"`
				} `json:"languageName"`
			} `json:"translationLanguages"`
			DefaultAudioTrackIndex int `json:"defaultAudioTrackIndex"`
		} `json:"playerCaptionsTracklistRenderer"`
	} `json:"captions"`
	VideoDetails struct {
		VideoID          string   `json:"videoId"`
		Title            string   `json:"title"`
		LengthSeconds    string   `json:"lengthSeconds"`
		Keywords         []string `json:"keywords"`
		ChannelID        string   `json:"channelId"`
		IsOwnerViewing   bool     `json:"isOwnerViewing"`
		ShortDescription string   `json:"shortDescription"`
		IsCrawlable      bool     `json:"isCrawlable"`
		Thumbnail        struct {
			Thumbnails []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"thumbnails"`
		} `json:"thumbnail"`
		AverageRating     float64 `json:"averageRating"`
		AllowRatings      bool    `json:"allowRatings"`
		ViewCount         string  `json:"viewCount"`
		Author            string  `json:"author"`
		IsPrivate         bool    `json:"isPrivate"`
		IsUnpluggedCorpus bool    `json:"isUnpluggedCorpus"`
		IsLiveContent     bool    `json:"isLiveContent"`
	} `json:"videoDetails"`
	PlayerConfig struct {
		AudioConfig struct {
			LoudnessDb              float64 `json:"loudnessDb"`
			PerceptualLoudnessDb    float64 `json:"perceptualLoudnessDb"`
			EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
		} `json:"audioConfig"`
		StreamSelectionConfig struct {
			MaxBitrate string `json:"maxBitrate"`
		} `json:"streamSelectionConfig"`
		MediaCommonConfig struct {
			DynamicReadaheadConfig struct {
				MaxReadAheadMediaTimeMs int `json:"maxReadAheadMediaTimeMs"`
				MinReadAheadMediaTimeMs int `json:"minReadAheadMediaTimeMs"`
				ReadAheadGrowthRateMs   int `json:"readAheadGrowthRateMs"`
			} `json:"dynamicReadaheadConfig"`
		} `json:"mediaCommonConfig"`
		WebPlayerConfig struct {
			WebPlayerActionsPorting struct {
				GetSharePanelCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					WebPlayerShareEntityServiceEndpoint struct {
						SerializedShareEntity string `json:"serializedShareEntity"`
					} `json:"webPlayerShareEntityServiceEndpoint"`
				} `json:"getSharePanelCommand"`
				SubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					SubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"subscribeEndpoint"`
				} `json:"subscribeCommand"`
				UnsubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					UnsubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"unsubscribeEndpoint"`
				} `json:"unsubscribeCommand"`
				AddToWatchLaterCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					PlaylistEditEndpoint struct {
						PlaylistID string `json:"playlistId"`
						Actions    []struct {
							AddedVideoID string `json:"addedVideoId"`
							Action       string `json:"action"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"addToWatchLaterCommand"`
				RemoveFromWatchLaterCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					PlaylistEditEndpoint struct {
						PlaylistID string `json:"playlistId"`
						Actions    []struct {
							Action         string `json:"action"`
							RemovedVideoID string `json:"removedVideoId"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"removeFromWatchLaterCommand"`
			} `json:"webPlayerActionsPorting"`
		} `json:"webPlayerConfig"`
	} `json:"playerConfig"`
	Storyboards struct {
		PlayerStoryboardSpecRenderer struct {
			Spec string `json:"spec"`
		} `json:"playerStoryboardSpecRenderer"`
	} `json:"storyboards"`
	Microformat struct {
		PlayerMicroformatRenderer struct {
			Thumbnail struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			Embed struct {
				IframeURL      string `json:"iframeUrl"`
				FlashURL       string `json:"flashUrl"`
				Width          int    `json:"width"`
				Height         int    `json:"height"`
				FlashSecureURL string `json:"flashSecureUrl"`
			} `json:"embed"`
			Title struct {
				SimpleText string `json:"simpleText"`
			} `json:"title"`
			Description struct {
				SimpleText string `json:"simpleText"`
			} `json:"description"`
			LengthSeconds      string   `json:"lengthSeconds"`
			OwnerProfileURL    string   `json:"ownerProfileUrl"`
			ExternalChannelID  string   `json:"externalChannelId"`
			IsFamilySafe       bool     `json:"isFamilySafe"`
			AvailableCountries []string `json:"availableCountries"`
			IsUnlisted         bool     `json:"isUnlisted"`
			HasYpcMetadata     bool     `json:"hasYpcMetadata"`
			ViewCount          string   `json:"viewCount"`
			Category           string   `json:"category"`
			PublishDate        string   `json:"publishDate"`
			OwnerChannelName   string   `json:"ownerChannelName"`
			UploadDate         string   `json:"uploadDate"`
		} `json:"playerMicroformatRenderer"`
	} `json:"microformat"`
	Cards struct {
		CardCollectionRenderer struct {
			Cards []struct {
				CardRenderer struct {
					Teaser struct {
						SimpleCardTeaserRenderer struct {
							Message struct {
								SimpleText string `json:"simpleText"`
							} `json:"message"`
							TrackingParams       string `json:"trackingParams"`
							Prominent            bool   `json:"prominent"`
							LogVisibilityUpdates bool   `json:"logVisibilityUpdates"`
						} `json:"simpleCardTeaserRenderer"`
					} `json:"teaser"`
					Content struct {
						VideoInfoCardContentRenderer struct {
							VideoThumbnail struct {
								Thumbnails []struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
							} `json:"videoThumbnail"`
							LengthString struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"lengthString"`
							VideoTitle struct {
								SimpleText string `json:"simpleText"`
							} `json:"videoTitle"`
							ChannelName struct {
								SimpleText string `json:"simpleText"`
							} `json:"channelName"`
							ViewCountText struct {
								SimpleText string `json:"simpleText"`
							} `json:"viewCountText"`
							Action struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoID                            string `json:"videoId"`
									WatchEndpointSupportedOnesieConfig struct {
										HTML5PlaybackOnesieConfig struct {
											CommonConfig struct {
												URL string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
								} `json:"watchEndpoint"`
							} `json:"action"`
							TrackingParams string `json:"trackingParams"`
							CustomMessage  struct {
								SimpleText string `json:"simpleText"`
							} `json:"customMessage"`
						} `json:"videoInfoCardContentRenderer"`
					} `json:"content"`
					CueRanges []struct {
						StartCardActiveMs string `json:"startCardActiveMs"`
						EndCardActiveMs   string `json:"endCardActiveMs"`
						TeaserDurationMs  string `json:"teaserDurationMs"`
						IconAfterTeaserMs string `json:"iconAfterTeaserMs"`
					} `json:"cueRanges"`
					Icon struct {
						InfoCardIconRenderer struct {
							TrackingParams string `json:"trackingParams"`
						} `json:"infoCardIconRenderer"`
					} `json:"icon"`
					TrackingParams string `json:"trackingParams"`
					CardID         string `json:"cardId"`
					Feature        string `json:"feature"`
				} `json:"cardRenderer"`
			} `json:"cards"`
			HeaderText struct {
				SimpleText string `json:"simpleText"`
			} `json:"headerText"`
			Icon struct {
				InfoCardIconRenderer struct {
					TrackingParams string `json:"trackingParams"`
				} `json:"infoCardIconRenderer"`
			} `json:"icon"`
			CloseButton struct {
				InfoCardIconRenderer struct {
					TrackingParams string `json:"trackingParams"`
				} `json:"infoCardIconRenderer"`
			} `json:"closeButton"`
			TrackingParams           string `json:"trackingParams"`
			AllowTeaserDismiss       bool   `json:"allowTeaserDismiss"`
			LogIconVisibilityUpdates bool   `json:"logIconVisibilityUpdates"`
		} `json:"cardCollectionRenderer"`
	} `json:"cards"`
	Attestation struct {
		PlayerAttestationRenderer struct {
			Challenge    string `json:"challenge"`
			BotguardData struct {
				Program            string `json:"program"`
				InterpreterSafeURL struct {
					PrivateDoNotAccessOrElseTrustedResourceURLWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
				} `json:"interpreterSafeUrl"`
				ServerEnvironment int `json:"serverEnvironment"`
			} `json:"botguardData"`
		} `json:"playerAttestationRenderer"`
	} `json:"attestation"`
	Messages []struct {
		MealbarPromoRenderer struct {
			Icon struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"icon"`
			MessageTexts []struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"messageTexts"`
			ActionButton struct {
				ButtonRenderer struct {
					Style string `json:"style"`
					Size  string `json:"size"`
					Text  struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams    string `json:"clickTrackingParams"`
						CommandExecutorCommand struct {
							Commands []struct {
								ClickTrackingParams string `json:"clickTrackingParams,omitempty"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
										APIURL      string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								BrowseEndpoint struct {
									BrowseID string `json:"browseId"`
									Params   string `json:"params"`
								} `json:"browseEndpoint,omitempty"`
								FeedbackEndpoint struct {
									FeedbackToken string `json:"feedbackToken"`
									UIActions     struct {
										HideEnclosingContainer bool `json:"hideEnclosingContainer"`
									} `json:"uiActions"`
								} `json:"feedbackEndpoint,omitempty"`
							} `json:"commands"`
						} `json:"commandExecutorCommand"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"actionButton"`
			DismissButton struct {
				ButtonRenderer struct {
					Style string `json:"style"`
					Size  string `json:"size"`
					Text  struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams    string `json:"clickTrackingParams"`
						CommandExecutorCommand struct {
							Commands []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										APIURL   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								FeedbackEndpoint struct {
									FeedbackToken string `json:"feedbackToken"`
									UIActions     struct {
										HideEnclosingContainer bool `json:"hideEnclosingContainer"`
									} `json:"uiActions"`
								} `json:"feedbackEndpoint"`
							} `json:"commands"`
						} `json:"commandExecutorCommand"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"dismissButton"`
			TriggerCondition    string `json:"triggerCondition"`
			Style               string `json:"style"`
			TrackingParams      string `json:"trackingParams"`
			ImpressionEndpoints []struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					WebCommandMetadata struct {
						SendPost bool   `json:"sendPost"`
						APIURL   string `json:"apiUrl"`
					} `json:"webCommandMetadata"`
				} `json:"commandMetadata"`
				FeedbackEndpoint struct {
					FeedbackToken string `json:"feedbackToken"`
					UIActions     struct {
						HideEnclosingContainer bool `json:"hideEnclosingContainer"`
					} `json:"uiActions"`
				} `json:"feedbackEndpoint"`
			} `json:"impressionEndpoints"`
			IsVisible    bool `json:"isVisible"`
			MessageTitle struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"messageTitle"`
		} `json:"mealbarPromoRenderer"`
	} `json:"messages"`
	Endscreen struct {
		EndscreenRenderer struct {
			Elements []struct {
				EndscreenElementRenderer struct {
					Style string `json:"style"`
					Image struct {
						Thumbnails []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"thumbnails"`
					} `json:"image"`
					Icon struct {
						Thumbnails []struct {
							URL string `json:"url"`
						} `json:"thumbnails"`
					} `json:"icon"`
					Left        float64 `json:"left"`
					Width       float64 `json:"width"`
					Top         float64 `json:"top"`
					AspectRatio float64 `json:"aspectRatio"`
					StartMs     string  `json:"startMs"`
					EndMs       string  `json:"endMs"`
					Title       struct {
						Accessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
						SimpleText string `json:"simpleText"`
					} `json:"title"`
					Metadata struct {
						SimpleText string `json:"simpleText"`
					} `json:"metadata"`
					CallToAction struct {
						SimpleText string `json:"simpleText"`
					} `json:"callToAction"`
					Dismiss struct {
						SimpleText string `json:"simpleText"`
					} `json:"dismiss"`
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								APIURL      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseID string `json:"browseId"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					HovercardButton struct {
						SubscribeButtonRenderer struct {
							ButtonText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"buttonText"`
							Subscribed           bool   `json:"subscribed"`
							Enabled              bool   `json:"enabled"`
							Type                 string `json:"type"`
							ChannelID            string `json:"channelId"`
							ShowPreferences      bool   `json:"showPreferences"`
							SubscribedButtonText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"subscribedButtonText"`
							UnsubscribedButtonText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"unsubscribedButtonText"`
							TrackingParams        string `json:"trackingParams"`
							UnsubscribeButtonText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"unsubscribeButtonText"`
							ServiceEndpoints []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										APIURL   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								SubscribeEndpoint struct {
									ChannelIds []string `json:"channelIds"`
									Params     string   `json:"params"`
								} `json:"subscribeEndpoint,omitempty"`
								SignalServiceEndpoint struct {
									Signal  string `json:"signal"`
									Actions []struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										OpenPopupAction     struct {
											Popup struct {
												ConfirmDialogRenderer struct {
													TrackingParams string `json:"trackingParams"`
													DialogMessages []struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"dialogMessages"`
													ConfirmButton struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Text       struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
															ServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		APIURL   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																UnsubscribeEndpoint struct {
																	ChannelIds []string `json:"channelIds"`
																	Params     string   `json:"params"`
																} `json:"unsubscribeEndpoint"`
															} `json:"serviceEndpoint"`
															Accessibility struct {
																Label string `json:"label"`
															} `json:"accessibility"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"confirmButton"`
													CancelButton struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Text       struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
															Accessibility struct {
																Label string `json:"label"`
															} `json:"accessibility"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"cancelButton"`
													PrimaryIsCancel bool `json:"primaryIsCancel"`
												} `json:"confirmDialogRenderer"`
											} `json:"popup"`
											PopupType string `json:"popupType"`
										} `json:"openPopupAction"`
									} `json:"actions"`
								} `json:"signalServiceEndpoint,omitempty"`
							} `json:"serviceEndpoints"`
							SubscribeAccessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"subscribeAccessibility"`
							UnsubscribeAccessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"unsubscribeAccessibility"`
							SignInEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL string `json:"url"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
							} `json:"signInEndpoint"`
						} `json:"subscribeButtonRenderer"`
					} `json:"hovercardButton"`
					TrackingParams string `json:"trackingParams"`
					IsSubscribe    bool   `json:"isSubscribe"`
					ID             string `json:"id"`
				} `json:"endscreenElementRenderer"`
			} `json:"elements"`
			StartMs        string `json:"startMs"`
			TrackingParams string `json:"trackingParams"`
		} `json:"endscreenRenderer"`
	} `json:"endscreen"`
	FrameworkUpdates struct {
		EntityBatchUpdate struct {
			Mutations []struct {
				EntityKey string `json:"entityKey"`
				Type      string `json:"type"`
				Payload   struct {
					OfflineabilityEntity struct {
						Key                     string `json:"key"`
						AccessState             string `json:"accessState"`
						AddToOfflineButtonState string `json:"addToOfflineButtonState"`
					} `json:"offlineabilityEntity"`
				} `json:"payload"`
			} `json:"mutations"`
			Timestamp struct {
				Seconds string `json:"seconds"`
				Nanos   int    `json:"nanos"`
			} `json:"timestamp"`
		} `json:"entityBatchUpdate"`
	} `json:"frameworkUpdates"`
}

type videoRenderer struct {
	VideoId   string `json:"videoId"`
	Thumbnail struct {
		Thumbnails []struct {
			Url    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"thumbnails"`
	} `json:"thumbnail"`
	Title struct {
		Runs []struct {
			Text string `json:"text"`
		} `json:"runs"`
		Accessibility struct {
			AccessibilityData struct {
				Label string `json:"label"`
			} `json:"accessibilityData"`
		} `json:"accessibility"`
	} `json:"title"`
	LongBylineText struct {
		Runs []struct {
			Text               string `json:"text"`
			NavigationEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					WebCommandMetadata struct {
						Url         string `json:"url"`
						WebPageType string `json:"webPageType"`
						RootVe      int    `json:"rootVe"`
						ApiUrl      string `json:"apiUrl"`
					} `json:"webCommandMetadata"`
				} `json:"commandMetadata"`
				BrowseEndpoint struct {
					BrowseId         string `json:"browseId"`
					CanonicalBaseUrl string `json:"canonicalBaseUrl,omitempty"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint"`
		} `json:"runs"`
	} `json:"longBylineText"`
	PublishedTimeText struct {
		SimpleText string `json:"simpleText"`
	} `json:"publishedTimeText,omitempty"`
	LengthText struct {
		Accessibility struct {
			AccessibilityData struct {
				Label string `json:"label"`
			} `json:"accessibilityData"`
		} `json:"accessibility"`
		SimpleText string `json:"simpleText"`
	} `json:"lengthText,omitempty"`
	ViewCountText struct {
		SimpleText string `json:"simpleText,omitempty"`
		Runs       []struct {
			Text string `json:"text"`
		} `json:"runs,omitempty"`
	} `json:"viewCountText"`
	NavigationEndpoint struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
		CommandMetadata     struct {
			WebCommandMetadata struct {
				Url         string `json:"url"`
				WebPageType string `json:"webPageType"`
				RootVe      int    `json:"rootVe"`
			} `json:"webCommandMetadata"`
		} `json:"commandMetadata"`
		WatchEndpoint struct {
			VideoId                            string `json:"videoId"`
			Params                             string `json:"params"`
			WatchEndpointSupportedOnesieConfig struct {
				Html5PlaybackOnesieConfig struct {
					CommonConfig struct {
						Url string `json:"url"`
					} `json:"commonConfig"`
				} `json:"html5PlaybackOnesieConfig"`
			} `json:"watchEndpointSupportedOnesieConfig"`
		} `json:"watchEndpoint"`
	} `json:"navigationEndpoint"`
	OwnerBadges []struct {
		MetadataBadgeRenderer struct {
			Icon struct {
				IconType string `json:"iconType"`
			} `json:"icon"`
			Style             string `json:"style"`
			Tooltip           string `json:"tooltip"`
			TrackingParams    string `json:"trackingParams"`
			AccessibilityData struct {
				Label string `json:"label"`
			} `json:"accessibilityData"`
		} `json:"metadataBadgeRenderer"`
	} `json:"ownerBadges,omitempty"`
	OwnerText struct {
		Runs []struct {
			Text               string `json:"text"`
			NavigationEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					WebCommandMetadata struct {
						Url         string `json:"url"`
						WebPageType string `json:"webPageType"`
						RootVe      int    `json:"rootVe"`
						ApiUrl      string `json:"apiUrl"`
					} `json:"webCommandMetadata"`
				} `json:"commandMetadata"`
				BrowseEndpoint struct {
					BrowseId         string `json:"browseId"`
					CanonicalBaseUrl string `json:"canonicalBaseUrl,omitempty"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint"`
		} `json:"runs"`
	} `json:"ownerText"`
	ShortBylineText struct {
		Runs []struct {
			Text               string `json:"text"`
			NavigationEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					WebCommandMetadata struct {
						Url         string `json:"url"`
						WebPageType string `json:"webPageType"`
						RootVe      int    `json:"rootVe"`
						ApiUrl      string `json:"apiUrl"`
					} `json:"webCommandMetadata"`
				} `json:"commandMetadata"`
				BrowseEndpoint struct {
					BrowseId         string `json:"browseId"`
					CanonicalBaseUrl string `json:"canonicalBaseUrl,omitempty"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint"`
		} `json:"runs"`
	} `json:"shortBylineText"`
	TrackingParams     string `json:"trackingParams"`
	ShowActionMenu     bool   `json:"showActionMenu"`
	ShortViewCountText struct {
		Accessibility struct {
			AccessibilityData struct {
				Label string `json:"label"`
			} `json:"accessibilityData"`
		} `json:"accessibility,omitempty"`
		SimpleText string `json:"simpleText,omitempty"`
		Runs       []struct {
			Text string `json:"text"`
		} `json:"runs,omitempty"`
	} `json:"shortViewCountText"`
	Menu struct {
		MenuRenderer struct {
			Items []struct {
				MenuServiceItemRenderer struct {
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					ServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams  string `json:"clickTrackingParams"`
								AddToPlaylistCommand struct {
									OpenMiniplayer      bool   `json:"openMiniplayer"`
									VideoId             string `json:"videoId"`
									ListType            string `json:"listType"`
									OnCreateListCommand struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												ApiUrl   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										CreatePlaylistServiceEndpoint struct {
											VideoIds []string `json:"videoIds"`
											Params   string   `json:"params"`
										} `json:"createPlaylistServiceEndpoint"`
									} `json:"onCreateListCommand"`
									VideoIds []string `json:"videoIds"`
								} `json:"addToPlaylistCommand"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"serviceEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"menuServiceItemRenderer"`
			} `json:"items"`
			TrackingParams string `json:"trackingParams"`
			Accessibility  struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibility"`
		} `json:"menuRenderer"`
	} `json:"menu"`
	ChannelThumbnailSupportedRenderers struct {
		ChannelThumbnailWithLinkRenderer struct {
			Thumbnail struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			NavigationEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					WebCommandMetadata struct {
						Url         string `json:"url"`
						WebPageType string `json:"webPageType"`
						RootVe      int    `json:"rootVe"`
						ApiUrl      string `json:"apiUrl"`
					} `json:"webCommandMetadata"`
				} `json:"commandMetadata"`
				BrowseEndpoint struct {
					BrowseId         string `json:"browseId"`
					CanonicalBaseUrl string `json:"canonicalBaseUrl,omitempty"`
				} `json:"browseEndpoint"`
			} `json:"navigationEndpoint"`
			Accessibility struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibility"`
		} `json:"channelThumbnailWithLinkRenderer"`
	} `json:"channelThumbnailSupportedRenderers"`
	ThumbnailOverlays []struct {
		ThumbnailOverlayTimeStatusRenderer struct {
			Text struct {
				Accessibility struct {
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"accessibility"`
				SimpleText string `json:"simpleText"`
			} `json:"text"`
			Style string `json:"style"`
		} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
		ThumbnailOverlayToggleButtonRenderer struct {
			IsToggled     bool `json:"isToggled,omitempty"`
			UntoggledIcon struct {
				IconType string `json:"iconType"`
			} `json:"untoggledIcon"`
			ToggledIcon struct {
				IconType string `json:"iconType"`
			} `json:"toggledIcon"`
			UntoggledTooltip         string `json:"untoggledTooltip"`
			ToggledTooltip           string `json:"toggledTooltip"`
			UntoggledServiceEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					WebCommandMetadata struct {
						SendPost bool   `json:"sendPost"`
						ApiUrl   string `json:"apiUrl,omitempty"`
					} `json:"webCommandMetadata"`
				} `json:"commandMetadata"`
				PlaylistEditEndpoint struct {
					PlaylistId string `json:"playlistId"`
					Actions    []struct {
						AddedVideoId string `json:"addedVideoId"`
						Action       string `json:"action"`
					} `json:"actions"`
				} `json:"playlistEditEndpoint,omitempty"`
				SignalServiceEndpoint struct {
					Signal  string `json:"signal"`
					Actions []struct {
						ClickTrackingParams  string `json:"clickTrackingParams"`
						AddToPlaylistCommand struct {
							OpenMiniplayer      bool   `json:"openMiniplayer"`
							VideoId             string `json:"videoId"`
							ListType            string `json:"listType"`
							OnCreateListCommand struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										ApiUrl   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								CreatePlaylistServiceEndpoint struct {
									VideoIds []string `json:"videoIds"`
									Params   string   `json:"params"`
								} `json:"createPlaylistServiceEndpoint"`
							} `json:"onCreateListCommand"`
							VideoIds []string `json:"videoIds"`
						} `json:"addToPlaylistCommand"`
					} `json:"actions"`
				} `json:"signalServiceEndpoint,omitempty"`
			} `json:"untoggledServiceEndpoint"`
			ToggledServiceEndpoint struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					WebCommandMetadata struct {
						SendPost bool   `json:"sendPost"`
						ApiUrl   string `json:"apiUrl"`
					} `json:"webCommandMetadata"`
				} `json:"commandMetadata"`
				PlaylistEditEndpoint struct {
					PlaylistId string `json:"playlistId"`
					Actions    []struct {
						Action         string `json:"action"`
						RemovedVideoId string `json:"removedVideoId"`
					} `json:"actions"`
				} `json:"playlistEditEndpoint"`
			} `json:"toggledServiceEndpoint,omitempty"`
			UntoggledAccessibility struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"untoggledAccessibility"`
			ToggledAccessibility struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"toggledAccessibility"`
			TrackingParams string `json:"trackingParams"`
		} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
		ThumbnailOverlayNowPlayingRenderer struct {
			Text struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"text"`
		} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
	} `json:"thumbnailOverlays"`
	DetailedMetadataSnippets []struct {
		SnippetText struct {
			Runs []struct {
				Text string `json:"text"`
				Bold bool   `json:"bold,omitempty"`
			} `json:"runs"`
		} `json:"snippetText"`
		SnippetHoverText struct {
			Runs []struct {
				Text string `json:"text"`
			} `json:"runs"`
		} `json:"snippetHoverText"`
		MaxOneLine bool `json:"maxOneLine"`
	} `json:"detailedMetadataSnippets"`
	Badges []struct {
		MetadataBadgeRenderer struct {
			Style             string `json:"style"`
			Label             string `json:"label"`
			TrackingParams    string `json:"trackingParams"`
			AccessibilityData struct {
				Label string `json:"label"`
			} `json:"accessibilityData,omitempty"`
		} `json:"metadataBadgeRenderer"`
	} `json:"badges,omitempty"`
}

type youtubeSearchResponse struct {
	ResponseContext struct {
		VisitorData           string `json:"visitorData"`
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MainAppWebResponseContext struct {
			LoggedOut bool `json:"loggedOut"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	EstimatedResults string `json:"estimatedResults"`
	Contents         struct {
		TwoColumnSearchResultsRenderer struct {
			PrimaryContents struct {
				SectionListRenderer struct {
					Contents []struct {
						ItemSectionRenderer struct {
							Contents []struct {
								VideoRenderer videoRenderer `json:"videoRenderer,omitempty"`
								RadioRenderer struct {
									PlaylistId string `json:"playlistId"`
									Title      struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									Thumbnail struct {
										Thumbnails []struct {
											Url    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
									} `json:"thumbnail"`
									VideoCountText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"videoCountText"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										WatchEndpoint struct {
											VideoId          string `json:"videoId"`
											PlaylistId       string `json:"playlistId"`
											Params           string `json:"params"`
											ContinuePlayback bool   `json:"continuePlayback"`
											LoggingContext   struct {
												VssLoggingContext struct {
													SerializedContextData string `json:"serializedContextData"`
												} `json:"vssLoggingContext"`
											} `json:"loggingContext"`
											WatchEndpointSupportedOnesieConfig struct {
												Html5PlaybackOnesieConfig struct {
													CommonConfig struct {
														Url string `json:"url"`
													} `json:"commonConfig"`
												} `json:"html5PlaybackOnesieConfig"`
											} `json:"watchEndpointSupportedOnesieConfig"`
										} `json:"watchEndpoint"`
									} `json:"navigationEndpoint"`
									ShortBylineText struct {
										SimpleText string `json:"simpleText"`
									} `json:"shortBylineText"`
									TrackingParams string `json:"trackingParams"`
									Videos         []struct {
										ChildVideoRenderer struct {
											Title struct {
												SimpleText string `json:"simpleText"`
											} `json:"title"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												WatchEndpoint struct {
													VideoId        string `json:"videoId"`
													PlaylistId     string `json:"playlistId"`
													Params         string `json:"params"`
													LoggingContext struct {
														VssLoggingContext struct {
															SerializedContextData string `json:"serializedContextData"`
														} `json:"vssLoggingContext"`
													} `json:"loggingContext"`
													WatchEndpointSupportedOnesieConfig struct {
														Html5PlaybackOnesieConfig struct {
															CommonConfig struct {
																Url string `json:"url"`
															} `json:"commonConfig"`
														} `json:"html5PlaybackOnesieConfig"`
													} `json:"watchEndpointSupportedOnesieConfig"`
												} `json:"watchEndpoint"`
											} `json:"navigationEndpoint"`
											LengthText struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
												SimpleText string `json:"simpleText"`
											} `json:"lengthText"`
											VideoId string `json:"videoId"`
										} `json:"childVideoRenderer"`
									} `json:"videos"`
									ThumbnailText struct {
										Runs []struct {
											Text string `json:"text"`
											Bold bool   `json:"bold,omitempty"`
										} `json:"runs"`
									} `json:"thumbnailText"`
									LongBylineText struct {
										SimpleText string `json:"simpleText"`
									} `json:"longBylineText"`
									ThumbnailOverlays []struct {
										ThumbnailOverlayBottomPanelRenderer struct {
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
										} `json:"thumbnailOverlayBottomPanelRenderer,omitempty"`
										ThumbnailOverlayHoverTextRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
										} `json:"thumbnailOverlayHoverTextRenderer,omitempty"`
										ThumbnailOverlayNowPlayingRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
										} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
									} `json:"thumbnailOverlays"`
									VideoCountShortText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"videoCountShortText"`
								} `json:"radioRenderer,omitempty"`
								PlaylistRenderer struct {
									PlaylistId string `json:"playlistId"`
									Title      struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									Thumbnails []struct {
										Thumbnails []struct {
											Url    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
									} `json:"thumbnails"`
									VideoCount         string `json:"videoCount"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										WatchEndpoint struct {
											VideoId        string `json:"videoId"`
											PlaylistId     string `json:"playlistId"`
											Params         string `json:"params"`
											LoggingContext struct {
												VssLoggingContext struct {
													SerializedContextData string `json:"serializedContextData"`
												} `json:"vssLoggingContext"`
											} `json:"loggingContext"`
											WatchEndpointSupportedOnesieConfig struct {
												Html5PlaybackOnesieConfig struct {
													CommonConfig struct {
														Url string `json:"url"`
													} `json:"commonConfig"`
												} `json:"html5PlaybackOnesieConfig"`
											} `json:"watchEndpointSupportedOnesieConfig"`
										} `json:"watchEndpoint"`
									} `json:"navigationEndpoint"`
									ViewPlaylistText struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														ApiUrl      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseId string `json:"browseId"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"viewPlaylistText"`
									ShortBylineText struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														ApiUrl      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseId string `json:"browseId"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"shortBylineText"`
									Videos []struct {
										ChildVideoRenderer struct {
											Title struct {
												SimpleText string `json:"simpleText"`
											} `json:"title"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												WatchEndpoint struct {
													VideoId        string `json:"videoId"`
													PlaylistId     string `json:"playlistId"`
													LoggingContext struct {
														VssLoggingContext struct {
															SerializedContextData string `json:"serializedContextData"`
														} `json:"vssLoggingContext"`
													} `json:"loggingContext"`
													WatchEndpointSupportedOnesieConfig struct {
														Html5PlaybackOnesieConfig struct {
															CommonConfig struct {
																Url string `json:"url"`
															} `json:"commonConfig"`
														} `json:"html5PlaybackOnesieConfig"`
													} `json:"watchEndpointSupportedOnesieConfig"`
												} `json:"watchEndpoint"`
											} `json:"navigationEndpoint"`
											LengthText struct {
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
												SimpleText string `json:"simpleText"`
											} `json:"lengthText"`
											VideoId string `json:"videoId"`
										} `json:"childVideoRenderer"`
									} `json:"videos"`
									VideoCountText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"videoCountText"`
									TrackingParams string `json:"trackingParams"`
									ThumbnailText  struct {
										Runs []struct {
											Text string `json:"text"`
											Bold bool   `json:"bold,omitempty"`
										} `json:"runs"`
									} `json:"thumbnailText"`
									LongBylineText struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														ApiUrl      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseId string `json:"browseId"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"longBylineText"`
									ThumbnailRenderer struct {
										PlaylistVideoThumbnailRenderer struct {
											Thumbnail struct {
												Thumbnails []struct {
													Url    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
											} `json:"thumbnail"`
										} `json:"playlistVideoThumbnailRenderer"`
									} `json:"thumbnailRenderer"`
									ThumbnailOverlays []struct {
										ThumbnailOverlaySidePanelRenderer struct {
											Text struct {
												SimpleText string `json:"simpleText"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
										} `json:"thumbnailOverlaySidePanelRenderer,omitempty"`
										ThumbnailOverlayHoverTextRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
										} `json:"thumbnailOverlayHoverTextRenderer,omitempty"`
										ThumbnailOverlayNowPlayingRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
										} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
									} `json:"thumbnailOverlays"`
								} `json:"playlistRenderer,omitempty"`
								ShelfRenderer struct {
									Title struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									Content struct {
										VerticalListRenderer struct {
											Items []struct {
												VideoRenderer struct {
													VideoId   string `json:"videoId"`
													Thumbnail struct {
														Thumbnails []struct {
															Url    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
													} `json:"thumbnail"`
													Title struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
													} `json:"title"`
													LongBylineText struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		ApiUrl      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseId         string `json:"browseId"`
																	CanonicalBaseUrl string `json:"canonicalBaseUrl,omitempty"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"longBylineText"`
													PublishedTimeText struct {
														SimpleText string `json:"simpleText"`
													} `json:"publishedTimeText"`
													LengthText struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"lengthText"`
													ViewCountText struct {
														SimpleText string `json:"simpleText"`
													} `json:"viewCountText"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														WatchEndpoint struct {
															VideoId                            string `json:"videoId"`
															WatchEndpointSupportedOnesieConfig struct {
																Html5PlaybackOnesieConfig struct {
																	CommonConfig struct {
																		Url string `json:"url"`
																	} `json:"commonConfig"`
																} `json:"html5PlaybackOnesieConfig"`
															} `json:"watchEndpointSupportedOnesieConfig"`
														} `json:"watchEndpoint"`
													} `json:"navigationEndpoint"`
													OwnerBadges []struct {
														MetadataBadgeRenderer struct {
															Icon struct {
																IconType string `json:"iconType"`
															} `json:"icon"`
															Style             string `json:"style"`
															Tooltip           string `json:"tooltip"`
															TrackingParams    string `json:"trackingParams"`
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"metadataBadgeRenderer"`
													} `json:"ownerBadges,omitempty"`
													OwnerText struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		ApiUrl      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseId         string `json:"browseId"`
																	CanonicalBaseUrl string `json:"canonicalBaseUrl,omitempty"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"ownerText"`
													ShortBylineText struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		ApiUrl      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseId         string `json:"browseId"`
																	CanonicalBaseUrl string `json:"canonicalBaseUrl,omitempty"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"shortBylineText"`
													TrackingParams     string `json:"trackingParams"`
													ShowActionMenu     bool   `json:"showActionMenu"`
													ShortViewCountText struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"shortViewCountText"`
													Menu struct {
														MenuRenderer struct {
															Items []struct {
																MenuServiceItemRenderer struct {
																	Text struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																	Icon struct {
																		IconType string `json:"iconType"`
																	} `json:"icon"`
																	ServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool `json:"sendPost"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		SignalServiceEndpoint struct {
																			Signal  string `json:"signal"`
																			Actions []struct {
																				ClickTrackingParams  string `json:"clickTrackingParams"`
																				AddToPlaylistCommand struct {
																					OpenMiniplayer      bool   `json:"openMiniplayer"`
																					VideoId             string `json:"videoId"`
																					ListType            string `json:"listType"`
																					OnCreateListCommand struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								SendPost bool   `json:"sendPost"`
																								ApiUrl   string `json:"apiUrl"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						CreatePlaylistServiceEndpoint struct {
																							VideoIds []string `json:"videoIds"`
																							Params   string   `json:"params"`
																						} `json:"createPlaylistServiceEndpoint"`
																					} `json:"onCreateListCommand"`
																					VideoIds []string `json:"videoIds"`
																				} `json:"addToPlaylistCommand"`
																			} `json:"actions"`
																		} `json:"signalServiceEndpoint"`
																	} `json:"serviceEndpoint"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"menuServiceItemRenderer"`
															} `json:"items"`
															TrackingParams string `json:"trackingParams"`
															Accessibility  struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibility"`
														} `json:"menuRenderer"`
													} `json:"menu"`
													ChannelThumbnailSupportedRenderers struct {
														ChannelThumbnailWithLinkRenderer struct {
															Thumbnail struct {
																Thumbnails []struct {
																	Url    string `json:"url"`
																	Width  int    `json:"width"`
																	Height int    `json:"height"`
																} `json:"thumbnails"`
															} `json:"thumbnail"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		ApiUrl      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseId         string `json:"browseId"`
																	CanonicalBaseUrl string `json:"canonicalBaseUrl,omitempty"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
															Accessibility struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibility"`
														} `json:"channelThumbnailWithLinkRenderer"`
													} `json:"channelThumbnailSupportedRenderers"`
													ThumbnailOverlays []struct {
														ThumbnailOverlayTimeStatusRenderer struct {
															Text struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"text"`
															Style string `json:"style"`
														} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
														ThumbnailOverlayToggleButtonRenderer struct {
															IsToggled     bool `json:"isToggled,omitempty"`
															UntoggledIcon struct {
																IconType string `json:"iconType"`
															} `json:"untoggledIcon"`
															ToggledIcon struct {
																IconType string `json:"iconType"`
															} `json:"toggledIcon"`
															UntoggledTooltip         string `json:"untoggledTooltip"`
															ToggledTooltip           string `json:"toggledTooltip"`
															UntoggledServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		ApiUrl   string `json:"apiUrl,omitempty"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																PlaylistEditEndpoint struct {
																	PlaylistId string `json:"playlistId"`
																	Actions    []struct {
																		AddedVideoId string `json:"addedVideoId"`
																		Action       string `json:"action"`
																	} `json:"actions"`
																} `json:"playlistEditEndpoint,omitempty"`
																SignalServiceEndpoint struct {
																	Signal  string `json:"signal"`
																	Actions []struct {
																		ClickTrackingParams  string `json:"clickTrackingParams"`
																		AddToPlaylistCommand struct {
																			OpenMiniplayer      bool   `json:"openMiniplayer"`
																			VideoId             string `json:"videoId"`
																			ListType            string `json:"listType"`
																			OnCreateListCommand struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool   `json:"sendPost"`
																						ApiUrl   string `json:"apiUrl"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				CreatePlaylistServiceEndpoint struct {
																					VideoIds []string `json:"videoIds"`
																					Params   string   `json:"params"`
																				} `json:"createPlaylistServiceEndpoint"`
																			} `json:"onCreateListCommand"`
																			VideoIds []string `json:"videoIds"`
																		} `json:"addToPlaylistCommand"`
																	} `json:"actions"`
																} `json:"signalServiceEndpoint,omitempty"`
															} `json:"untoggledServiceEndpoint"`
															ToggledServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		ApiUrl   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																PlaylistEditEndpoint struct {
																	PlaylistId string `json:"playlistId"`
																	Actions    []struct {
																		Action         string `json:"action"`
																		RemovedVideoId string `json:"removedVideoId"`
																	} `json:"actions"`
																} `json:"playlistEditEndpoint"`
															} `json:"toggledServiceEndpoint,omitempty"`
															UntoggledAccessibility struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"untoggledAccessibility"`
															ToggledAccessibility struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"toggledAccessibility"`
															TrackingParams string `json:"trackingParams"`
														} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
														ThumbnailOverlayNowPlayingRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
														} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
													} `json:"thumbnailOverlays"`
													DetailedMetadataSnippets []struct {
														SnippetText struct {
															Runs []struct {
																Text string `json:"text"`
																Bold bool   `json:"bold,omitempty"`
															} `json:"runs"`
														} `json:"snippetText"`
														SnippetHoverText struct {
															Runs []struct {
																Text string `json:"text"`
															} `json:"runs"`
														} `json:"snippetHoverText"`
														MaxOneLine bool `json:"maxOneLine"`
													} `json:"detailedMetadataSnippets"`
												} `json:"videoRenderer"`
											} `json:"items"`
											CollapsedItemCount       int `json:"collapsedItemCount"`
											CollapsedStateButtonText struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
												Accessibility struct {
													AccessibilityData struct {
														Label string `json:"label"`
													} `json:"accessibilityData"`
												} `json:"accessibility"`
											} `json:"collapsedStateButtonText"`
											TrackingParams string `json:"trackingParams"`
										} `json:"verticalListRenderer"`
									} `json:"content"`
									TrackingParams string `json:"trackingParams"`
								} `json:"shelfRenderer,omitempty"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
						} `json:"itemSectionRenderer,omitempty"`
						ContinuationItemRenderer struct {
							Trigger              string `json:"trigger"`
							ContinuationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										ApiUrl   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								ContinuationCommand struct {
									Token   string `json:"token"`
									Request string `json:"request"`
								} `json:"continuationCommand"`
							} `json:"continuationEndpoint"`
						} `json:"continuationItemRenderer,omitempty"`
					} `json:"contents"`
					TrackingParams string `json:"trackingParams"`
					SubMenu        struct {
						SearchSubMenuRenderer struct {
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"title"`
							Groups []struct {
								SearchFilterGroupRenderer struct {
									Title struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									Filters []struct {
										SearchFilterRenderer struct {
											Label struct {
												SimpleText string `json:"simpleText"`
											} `json:"label"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												SearchEndpoint struct {
													Query  string `json:"query"`
													Params string `json:"params"`
												} `json:"searchEndpoint"`
											} `json:"navigationEndpoint,omitempty"`
											Tooltip        string `json:"tooltip"`
											TrackingParams string `json:"trackingParams"`
											Status         string `json:"status,omitempty"`
										} `json:"searchFilterRenderer"`
									} `json:"filters"`
									TrackingParams string `json:"trackingParams"`
								} `json:"searchFilterGroupRenderer"`
							} `json:"groups"`
							TrackingParams string `json:"trackingParams"`
							Button         struct {
								ToggleButtonRenderer struct {
									Style struct {
										StyleType string `json:"styleType"`
									} `json:"style"`
									IsToggled   bool `json:"isToggled"`
									IsDisabled  bool `json:"isDisabled"`
									DefaultIcon struct {
										IconType string `json:"iconType"`
									} `json:"defaultIcon"`
									DefaultText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"defaultText"`
									Accessibility struct {
										Label string `json:"label"`
									} `json:"accessibility"`
									TrackingParams string `json:"trackingParams"`
									DefaultTooltip string `json:"defaultTooltip"`
									ToggledTooltip string `json:"toggledTooltip"`
									ToggledStyle   struct {
										StyleType string `json:"styleType"`
									} `json:"toggledStyle"`
									AccessibilityData struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibilityData"`
								} `json:"toggleButtonRenderer"`
							} `json:"button"`
						} `json:"searchSubMenuRenderer"`
					} `json:"subMenu"`
					HideBottomSeparator bool   `json:"hideBottomSeparator"`
					TargetId            string `json:"targetId"`
				} `json:"sectionListRenderer"`
			} `json:"primaryContents"`
			SecondaryContents struct {
				SecondarySearchContainerRenderer struct {
					Contents []struct {
						UniversalWatchCardRenderer struct {
							Header struct {
								WatchCardRichHeaderRenderer struct {
									Title struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									TitleNavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId string `json:"browseId"`
										} `json:"browseEndpoint"`
									} `json:"titleNavigationEndpoint"`
									Subtitle struct {
										SimpleText string `json:"simpleText"`
									} `json:"subtitle"`
									Avatar struct {
										Thumbnails []struct {
											Url    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
										PlaceholderColor int64 `json:"placeholderColor"`
									} `json:"avatar"`
									ColorSupportedDatas struct {
										BasicColorPaletteData struct {
											BackgroundColor      int64 `json:"backgroundColor"`
											ForegroundTitleColor int64 `json:"foregroundTitleColor"`
											ForegroundBodyColor  int64 `json:"foregroundBodyColor"`
										} `json:"basicColorPaletteData"`
									} `json:"colorSupportedDatas"`
									TrackingParams               string `json:"trackingParams"`
									DarkThemeColorSupportedDatas struct {
										BasicColorPaletteData struct {
											BackgroundColor      int64 `json:"backgroundColor"`
											ForegroundTitleColor int64 `json:"foregroundTitleColor"`
											ForegroundBodyColor  int64 `json:"foregroundBodyColor"`
										} `json:"basicColorPaletteData"`
									} `json:"darkThemeColorSupportedDatas"`
									Style string `json:"style"`
								} `json:"watchCardRichHeaderRenderer"`
							} `json:"header"`
							CallToAction struct {
								WatchCardHeroVideoRenderer struct {
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										WatchEndpoint struct {
											VideoId        string `json:"videoId"`
											PlaylistId     string `json:"playlistId"`
											Params         string `json:"params"`
											LoggingContext struct {
												VssLoggingContext struct {
													SerializedContextData string `json:"serializedContextData"`
												} `json:"vssLoggingContext"`
											} `json:"loggingContext"`
											WatchEndpointSupportedOnesieConfig struct {
												Html5PlaybackOnesieConfig struct {
													CommonConfig struct {
														Url string `json:"url"`
													} `json:"commonConfig"`
												} `json:"html5PlaybackOnesieConfig"`
											} `json:"watchEndpointSupportedOnesieConfig"`
										} `json:"watchEndpoint"`
									} `json:"navigationEndpoint"`
									TrackingParams     string `json:"trackingParams"`
									CallToActionButton struct {
										CallToActionButtonRenderer struct {
											Label struct {
												SimpleText string `json:"simpleText"`
											} `json:"label"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Style string `json:"style"`
										} `json:"callToActionButtonRenderer"`
									} `json:"callToActionButton"`
									HeroImage struct {
										CollageHeroImageRenderer struct {
											LeftThumbnail struct {
												Thumbnails []struct {
													Url    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
											} `json:"leftThumbnail"`
											TopRightThumbnail struct {
												Thumbnails []struct {
													Url    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
											} `json:"topRightThumbnail"`
											BottomRightThumbnail struct {
												Thumbnails []struct {
													Url    string `json:"url"`
													Width  int    `json:"width"`
													Height int    `json:"height"`
												} `json:"thumbnails"`
											} `json:"bottomRightThumbnail"`
										} `json:"collageHeroImageRenderer"`
									} `json:"heroImage"`
									Accessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
								} `json:"watchCardHeroVideoRenderer"`
							} `json:"callToAction"`
							Sections []struct {
								WatchCardSectionSequenceRenderer struct {
									Lists []struct {
										VerticalWatchCardListRenderer struct {
											Items []struct {
												WatchCardCompactVideoRenderer struct {
													Title struct {
														SimpleText string `json:"simpleText"`
													} `json:"title"`
													Subtitle struct {
														SimpleText string `json:"simpleText"`
													} `json:"subtitle"`
													LengthText struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"lengthText"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														WatchEndpoint struct {
															VideoId        string `json:"videoId"`
															PlaylistId     string `json:"playlistId"`
															LoggingContext struct {
																VssLoggingContext struct {
																	SerializedContextData string `json:"serializedContextData"`
																} `json:"vssLoggingContext"`
															} `json:"loggingContext"`
															WatchEndpointSupportedOnesieConfig struct {
																Html5PlaybackOnesieConfig struct {
																	CommonConfig struct {
																		Url string `json:"url"`
																	} `json:"commonConfig"`
																} `json:"html5PlaybackOnesieConfig"`
															} `json:"watchEndpointSupportedOnesieConfig"`
														} `json:"watchEndpoint"`
													} `json:"navigationEndpoint"`
													TrackingParams string `json:"trackingParams"`
													Style          string `json:"style"`
													Byline         struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		ApiUrl      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseId string `json:"browseId"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"byline"`
												} `json:"watchCardCompactVideoRenderer"`
											} `json:"items"`
											TrackingParams string `json:"trackingParams"`
										} `json:"verticalWatchCardListRenderer,omitempty"`
										HorizontalCardListRenderer struct {
											Cards []struct {
												SearchRefinementCardRenderer struct {
													Thumbnail struct {
														Thumbnails []struct {
															Url    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
													} `json:"thumbnail"`
													Query struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"query"`
													SearchEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														WatchPlaylistEndpoint struct {
															PlaylistId string `json:"playlistId"`
														} `json:"watchPlaylistEndpoint"`
													} `json:"searchEndpoint"`
													TrackingParams                    string `json:"trackingParams"`
													SearchRefinementCardRendererStyle struct {
														Value string `json:"value"`
													} `json:"searchRefinementCardRendererStyle"`
												} `json:"searchRefinementCardRenderer"`
											} `json:"cards"`
											TrackingParams string `json:"trackingParams"`
											Header         struct {
												TitleAndButtonListHeaderRenderer struct {
													Title struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"title"`
													TrackingParams string `json:"trackingParams"`
												} `json:"titleAndButtonListHeaderRenderer"`
											} `json:"header"`
											PreviousButton struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Icon       struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													TrackingParams string `json:"trackingParams"`
												} `json:"buttonRenderer"`
											} `json:"previousButton"`
											NextButton struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Icon       struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													TrackingParams string `json:"trackingParams"`
												} `json:"buttonRenderer"`
											} `json:"nextButton"`
										} `json:"horizontalCardListRenderer,omitempty"`
									} `json:"lists"`
									TrackingParams string `json:"trackingParams"`
									ListTitles     []struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"listTitles,omitempty"`
								} `json:"watchCardSectionSequenceRenderer"`
							} `json:"sections"`
							CollapsedLabel struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"collapsedLabel"`
							TrackingParams string `json:"trackingParams"`
						} `json:"universalWatchCardRenderer"`
					} `json:"contents"`
					TrackingParams string `json:"trackingParams"`
				} `json:"secondarySearchContainerRenderer"`
			} `json:"secondaryContents"`
		} `json:"twoColumnSearchResultsRenderer"`
	} `json:"contents"`
	TrackingParams string   `json:"trackingParams"`
	Refinements    []string `json:"refinements"`
}

type playlistAPIResponse struct {
	ResponseContext struct {
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MainAppWebResponseContext struct {
			LoggedOut bool `json:"loggedOut"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			YtConfigData struct {
				VisitorData           string `json:"visitorData"`
				RootVisualElementType int    `json:"rootVisualElementType"`
			} `json:"ytConfigData"`
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Selected bool `json:"selected"`
					Content  struct {
						SectionListRenderer struct {
							Contents []struct {
								ItemSectionRenderer struct {
									Contents []struct {
										PlaylistVideoListRenderer struct {
											Contents []struct {
												PlaylistVideoRenderer struct {
													VideoId   string `json:"videoId"`
													Thumbnail struct {
														Thumbnails []struct {
															Url    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
													} `json:"thumbnail"`
													Title struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
													} `json:"title"`
													Index struct {
														SimpleText string `json:"simpleText"`
													} `json:"index"`
													ShortBylineText struct {
														Runs []struct {
															Text               string `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		ApiUrl      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseId         string `json:"browseId"`
																	CanonicalBaseUrl string `json:"canonicalBaseUrl"`
																} `json:"browseEndpoint"`
															} `json:"navigationEndpoint"`
														} `json:"runs"`
													} `json:"shortBylineText"`
													LengthText struct {
														Accessibility struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibility"`
														SimpleText string `json:"simpleText"`
													} `json:"lengthText"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														WatchEndpoint struct {
															VideoId        string `json:"videoId"`
															PlaylistId     string `json:"playlistId"`
															Index          int    `json:"index"`
															Params         string `json:"params"`
															LoggingContext struct {
																VssLoggingContext struct {
																	SerializedContextData string `json:"serializedContextData"`
																} `json:"vssLoggingContext"`
															} `json:"loggingContext"`
															WatchEndpointSupportedOnesieConfig struct {
																Html5PlaybackOnesieConfig struct {
																	CommonConfig struct {
																		Url string `json:"url"`
																	} `json:"commonConfig"`
																} `json:"html5PlaybackOnesieConfig"`
															} `json:"watchEndpointSupportedOnesieConfig"`
														} `json:"watchEndpoint"`
													} `json:"navigationEndpoint"`
													LengthSeconds  string `json:"lengthSeconds"`
													TrackingParams string `json:"trackingParams"`
													IsPlayable     bool   `json:"isPlayable"`
													Menu           struct {
														MenuRenderer struct {
															Items []struct {
																MenuServiceItemRenderer struct {
																	Text struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																	Icon struct {
																		IconType string `json:"iconType"`
																	} `json:"icon"`
																	ServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool `json:"sendPost"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		SignalServiceEndpoint struct {
																			Signal  string `json:"signal"`
																			Actions []struct {
																				ClickTrackingParams  string `json:"clickTrackingParams"`
																				AddToPlaylistCommand struct {
																					OpenMiniplayer      bool   `json:"openMiniplayer"`
																					VideoId             string `json:"videoId"`
																					ListType            string `json:"listType"`
																					OnCreateListCommand struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								SendPost bool   `json:"sendPost"`
																								ApiUrl   string `json:"apiUrl"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						CreatePlaylistServiceEndpoint struct {
																							VideoIds []string `json:"videoIds"`
																							Params   string   `json:"params"`
																						} `json:"createPlaylistServiceEndpoint"`
																					} `json:"onCreateListCommand"`
																					VideoIds []string `json:"videoIds"`
																				} `json:"addToPlaylistCommand"`
																			} `json:"actions"`
																		} `json:"signalServiceEndpoint"`
																	} `json:"serviceEndpoint"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"menuServiceItemRenderer"`
															} `json:"items"`
															TrackingParams string `json:"trackingParams"`
															Accessibility  struct {
																AccessibilityData struct {
																	Label string `json:"label"`
																} `json:"accessibilityData"`
															} `json:"accessibility"`
														} `json:"menuRenderer"`
													} `json:"menu"`
													ThumbnailOverlays []struct {
														ThumbnailOverlayTimeStatusRenderer struct {
															Text struct {
																Accessibility struct {
																	AccessibilityData struct {
																		Label string `json:"label"`
																	} `json:"accessibilityData"`
																} `json:"accessibility"`
																SimpleText string `json:"simpleText"`
															} `json:"text"`
															Style string `json:"style"`
														} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
														ThumbnailOverlayNowPlayingRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
														} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
													} `json:"thumbnailOverlays"`
													Badges []struct {
														MetadataBadgeRenderer struct {
															Style          string `json:"style"`
															Label          string `json:"label"`
															TrackingParams string `json:"trackingParams"`
														} `json:"metadataBadgeRenderer"`
													} `json:"badges,omitempty"`
												} `json:"playlistVideoRenderer,omitempty"`
												ContinuationItemRenderer struct {
													Trigger              string `json:"trigger"`
													ContinuationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																SendPost bool   `json:"sendPost"`
																ApiUrl   string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														ContinuationCommand struct {
															Token   string `json:"token"`
															Request string `json:"request"`
														} `json:"continuationCommand"`
													} `json:"continuationEndpoint"`
												} `json:"continuationItemRenderer,omitempty"`
											} `json:"contents"`
											PlaylistId     string `json:"playlistId"`
											IsEditable     bool   `json:"isEditable"`
											CanReorder     bool   `json:"canReorder"`
											TrackingParams string `json:"trackingParams"`
											TargetId       string `json:"targetId"`
										} `json:"playlistVideoListRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
								} `json:"itemSectionRenderer"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
					TrackingParams string `json:"trackingParams"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"twoColumnBrowseResultsRenderer"`
	} `json:"contents"`
	Metadata struct {
		PlaylistMetadataRenderer struct {
			Title                  string `json:"title"`
			Description            string `json:"description"`
			AndroidAppindexingLink string `json:"androidAppindexingLink"`
			IosAppindexingLink     string `json:"iosAppindexingLink"`
		} `json:"playlistMetadataRenderer"`
	} `json:"metadata"`
	TrackingParams string `json:"trackingParams"`
	Topbar         struct {
		DesktopTopbarRenderer struct {
			Logo struct {
				TopbarLogoRenderer struct {
					IconImage struct {
						IconType string `json:"iconType"`
					} `json:"iconImage"`
					TooltipText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"tooltipText"`
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								ApiUrl      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseId string `json:"browseId"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					TrackingParams    string `json:"trackingParams"`
					OverrideEntityKey string `json:"overrideEntityKey"`
				} `json:"topbarLogoRenderer"`
			} `json:"logo"`
			Searchbox struct {
				FusionSearchboxRenderer struct {
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					PlaceholderText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"placeholderText"`
					Config struct {
						WebSearchboxConfig struct {
							RequestLanguage     string `json:"requestLanguage"`
							RequestDomain       string `json:"requestDomain"`
							HasOnscreenKeyboard bool   `json:"hasOnscreenKeyboard"`
							FocusSearchbox      bool   `json:"focusSearchbox"`
						} `json:"webSearchboxConfig"`
					} `json:"config"`
					TrackingParams string `json:"trackingParams"`
					SearchEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SearchEndpoint struct {
							Query string `json:"query"`
						} `json:"searchEndpoint"`
					} `json:"searchEndpoint"`
					ClearButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Icon       struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
						} `json:"buttonRenderer"`
					} `json:"clearButton"`
				} `json:"fusionSearchboxRenderer"`
			} `json:"searchbox"`
			TrackingParams string `json:"trackingParams"`
			CountryCode    string `json:"countryCode"`
			TopbarButtons  []struct {
				TopbarMenuButtonRenderer struct {
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					MenuRenderer struct {
						MultiPageMenuRenderer struct {
							Sections []struct {
								MultiPageMenuSectionRenderer struct {
									Items []struct {
										CompactLinkRenderer struct {
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Title struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"title"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														Url         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												UrlEndpoint struct {
													Url    string `json:"url"`
													Target string `json:"target"`
												} `json:"urlEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"compactLinkRenderer"`
									} `json:"items"`
									TrackingParams string `json:"trackingParams"`
								} `json:"multiPageMenuSectionRenderer"`
							} `json:"sections"`
							TrackingParams string `json:"trackingParams"`
							Style          string `json:"style"`
						} `json:"multiPageMenuRenderer"`
					} `json:"menuRenderer,omitempty"`
					TrackingParams string `json:"trackingParams"`
					Accessibility  struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
					Tooltip     string `json:"tooltip"`
					Style       string `json:"style"`
					TargetId    string `json:"targetId,omitempty"`
					MenuRequest struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								ApiUrl   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										MultiPageMenuRenderer struct {
											TrackingParams     string `json:"trackingParams"`
											Style              string `json:"style"`
											ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
										} `json:"multiPageMenuRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
									BeReused  bool   `json:"beReused"`
								} `json:"openPopupAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"menuRequest,omitempty"`
				} `json:"topbarMenuButtonRenderer,omitempty"`
				ButtonRenderer struct {
					Style string `json:"style"`
					Size  string `json:"size"`
					Text  struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignInEndpoint struct {
							IdamTag string `json:"idamTag"`
						} `json:"signInEndpoint"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
					TargetId       string `json:"targetId"`
				} `json:"buttonRenderer,omitempty"`
			} `json:"topbarButtons"`
			HotkeyDialog struct {
				HotkeyDialogRenderer struct {
					Title struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"title"`
					Sections []struct {
						HotkeyDialogSectionRenderer struct {
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"title"`
							Options []struct {
								HotkeyDialogSectionOptionRenderer struct {
									Label struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"label"`
									Hotkey                   string `json:"hotkey"`
									HotkeyAccessibilityLabel struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"hotkeyAccessibilityLabel,omitempty"`
								} `json:"hotkeyDialogSectionOptionRenderer"`
							} `json:"options"`
						} `json:"hotkeyDialogSectionRenderer"`
					} `json:"sections"`
					DismissButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"text"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"dismissButton"`
					TrackingParams string `json:"trackingParams"`
				} `json:"hotkeyDialogRenderer"`
			} `json:"hotkeyDialog"`
			BackButton struct {
				ButtonRenderer struct {
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"backButton"`
			ForwardButton struct {
				ButtonRenderer struct {
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"forwardButton"`
			A11YSkipNavigationButton struct {
				ButtonRenderer struct {
					Style      string `json:"style"`
					Size       string `json:"size"`
					IsDisabled bool   `json:"isDisabled"`
					Text       struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"a11ySkipNavigationButton"`
			VoiceSearchButton struct {
				ButtonRenderer struct {
					Style           string `json:"style"`
					Size            string `json:"size"`
					IsDisabled      bool   `json:"isDisabled"`
					ServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										VoiceSearchDialogRenderer struct {
											PlaceholderHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"placeholderHeader"`
											PromptHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"promptHeader"`
											ExampleQuery1 struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"exampleQuery1"`
											ExampleQuery2 struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"exampleQuery2"`
											PromptMicrophoneLabel struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"promptMicrophoneLabel"`
											LoadingHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"loadingHeader"`
											ConnectionErrorHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"connectionErrorHeader"`
											ConnectionErrorMicrophoneLabel struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"connectionErrorMicrophoneLabel"`
											PermissionsHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"permissionsHeader"`
											PermissionsSubtext struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"permissionsSubtext"`
											DisabledHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"disabledHeader"`
											DisabledSubtext struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"disabledSubtext"`
											MicrophoneButtonAriaLabel struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"microphoneButtonAriaLabel"`
											ExitButton struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Icon       struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													TrackingParams    string `json:"trackingParams"`
													AccessibilityData struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibilityData"`
												} `json:"buttonRenderer"`
											} `json:"exitButton"`
											TrackingParams            string `json:"trackingParams"`
											MicrophoneOffPromptHeader struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"microphoneOffPromptHeader"`
										} `json:"voiceSearchDialogRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
								} `json:"openPopupAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"serviceEndpoint"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					Tooltip           string `json:"tooltip"`
					TrackingParams    string `json:"trackingParams"`
					AccessibilityData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibilityData"`
				} `json:"buttonRenderer"`
			} `json:"voiceSearchButton"`
		} `json:"desktopTopbarRenderer"`
	} `json:"topbar"`
	Microformat struct {
		MicroformatDataRenderer struct {
			UrlCanonical string `json:"urlCanonical"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			Thumbnail    struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			SiteName           string `json:"siteName"`
			AppName            string `json:"appName"`
			AndroidPackage     string `json:"androidPackage"`
			IosAppStoreId      string `json:"iosAppStoreId"`
			IosAppArguments    string `json:"iosAppArguments"`
			OgType             string `json:"ogType"`
			UrlApplinksWeb     string `json:"urlApplinksWeb"`
			UrlApplinksIos     string `json:"urlApplinksIos"`
			UrlApplinksAndroid string `json:"urlApplinksAndroid"`
			UrlTwitterIos      string `json:"urlTwitterIos"`
			UrlTwitterAndroid  string `json:"urlTwitterAndroid"`
			TwitterCardType    string `json:"twitterCardType"`
			TwitterSiteHandle  string `json:"twitterSiteHandle"`
			SchemaDotOrgType   string `json:"schemaDotOrgType"`
			Noindex            bool   `json:"noindex"`
			Unlisted           bool   `json:"unlisted"`
			LinkAlternates     []struct {
				HrefUrl string `json:"hrefUrl"`
			} `json:"linkAlternates"`
		} `json:"microformatDataRenderer"`
	} `json:"microformat"`
	Sidebar struct {
		PlaylistSidebarRenderer struct {
			Items []struct {
				PlaylistSidebarPrimaryInfoRenderer struct {
					ThumbnailRenderer struct {
						PlaylistVideoThumbnailRenderer struct {
							Thumbnail struct {
								Thumbnails []struct {
									Url    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
							} `json:"thumbnail"`
						} `json:"playlistVideoThumbnailRenderer"`
					} `json:"thumbnailRenderer"`
					Title struct {
						Runs []struct {
							Text               string `json:"text"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoId        string `json:"videoId"`
									PlaylistId     string `json:"playlistId"`
									LoggingContext struct {
										VssLoggingContext struct {
											SerializedContextData string `json:"serializedContextData"`
										} `json:"vssLoggingContext"`
									} `json:"loggingContext"`
									WatchEndpointSupportedOnesieConfig struct {
										Html5PlaybackOnesieConfig struct {
											CommonConfig struct {
												Url string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
								} `json:"watchEndpoint"`
							} `json:"navigationEndpoint"`
						} `json:"runs"`
					} `json:"title"`
					Stats []struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs,omitempty"`
						SimpleText string `json:"simpleText,omitempty"`
					} `json:"stats"`
					Menu struct {
						MenuRenderer struct {
							Items []struct {
								MenuNavigationItemRenderer struct {
									Text struct {
										SimpleText string `json:"simpleText"`
									} `json:"text"`
									Icon struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												IgnoreNavigation bool `json:"ignoreNavigation"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ModalEndpoint struct {
											Modal struct {
												ModalWithTitleAndButtonRenderer struct {
													Title struct {
														SimpleText string `json:"simpleText"`
													} `json:"title"`
													Content struct {
														SimpleText string `json:"simpleText"`
													} `json:"content"`
													Button struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Text       struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																SignInEndpoint struct {
																	NextEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				Url         string `json:"url"`
																				WebPageType string `json:"webPageType"`
																				RootVe      int    `json:"rootVe"`
																				ApiUrl      string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		BrowseEndpoint struct {
																			BrowseId string `json:"browseId"`
																		} `json:"browseEndpoint"`
																	} `json:"nextEndpoint"`
																} `json:"signInEndpoint"`
															} `json:"navigationEndpoint"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"button"`
												} `json:"modalWithTitleAndButtonRenderer"`
											} `json:"modal"`
										} `json:"modalEndpoint"`
									} `json:"navigationEndpoint"`
									TrackingParams string `json:"trackingParams"`
								} `json:"menuNavigationItemRenderer"`
							} `json:"items"`
							TrackingParams  string `json:"trackingParams"`
							TopLevelButtons []struct {
								ToggleButtonRenderer struct {
									Style struct {
										StyleType string `json:"styleType"`
									} `json:"style"`
									Size struct {
										SizeType string `json:"sizeType"`
									} `json:"size"`
									IsToggled   bool `json:"isToggled"`
									IsDisabled  bool `json:"isDisabled"`
									DefaultIcon struct {
										IconType string `json:"iconType"`
									} `json:"defaultIcon"`
									ToggledIcon struct {
										IconType string `json:"iconType"`
									} `json:"toggledIcon"`
									TrackingParams            string `json:"trackingParams"`
									DefaultTooltip            string `json:"defaultTooltip"`
									ToggledTooltip            string `json:"toggledTooltip"`
									DefaultNavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												IgnoreNavigation bool `json:"ignoreNavigation"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ModalEndpoint struct {
											Modal struct {
												ModalWithTitleAndButtonRenderer struct {
													Title struct {
														SimpleText string `json:"simpleText"`
													} `json:"title"`
													Content struct {
														SimpleText string `json:"simpleText"`
													} `json:"content"`
													Button struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Text       struct {
																SimpleText string `json:"simpleText"`
															} `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																SignInEndpoint struct {
																	NextEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				Url         string `json:"url"`
																				WebPageType string `json:"webPageType"`
																				RootVe      int    `json:"rootVe"`
																				ApiUrl      string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		BrowseEndpoint struct {
																			BrowseId string `json:"browseId"`
																		} `json:"browseEndpoint"`
																	} `json:"nextEndpoint"`
																	IdamTag string `json:"idamTag"`
																} `json:"signInEndpoint"`
															} `json:"navigationEndpoint"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"button"`
												} `json:"modalWithTitleAndButtonRenderer"`
											} `json:"modal"`
										} `json:"modalEndpoint"`
									} `json:"defaultNavigationEndpoint"`
									AccessibilityData struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibilityData"`
									ToggledAccessibilityData struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"toggledAccessibilityData"`
								} `json:"toggleButtonRenderer,omitempty"`
								ButtonRenderer struct {
									Style      string `json:"style"`
									Size       string `json:"size"`
									IsDisabled bool   `json:"isDisabled"`
									Icon       struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										WatchEndpoint struct {
											VideoId        string `json:"videoId"`
											PlaylistId     string `json:"playlistId"`
											Params         string `json:"params"`
											LoggingContext struct {
												VssLoggingContext struct {
													SerializedContextData string `json:"serializedContextData"`
												} `json:"vssLoggingContext"`
											} `json:"loggingContext"`
											WatchEndpointSupportedOnesieConfig struct {
												Html5PlaybackOnesieConfig struct {
													CommonConfig struct {
														Url string `json:"url"`
													} `json:"commonConfig"`
												} `json:"html5PlaybackOnesieConfig"`
											} `json:"watchEndpointSupportedOnesieConfig"`
										} `json:"watchEndpoint"`
									} `json:"navigationEndpoint,omitempty"`
									Accessibility struct {
										Label string `json:"label"`
									} `json:"accessibility"`
									Tooltip         string `json:"tooltip"`
									TrackingParams  string `json:"trackingParams"`
									ServiceEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												ApiUrl   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ShareEntityServiceEndpoint struct {
											SerializedShareEntity string `json:"serializedShareEntity"`
											Commands              []struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												OpenPopupAction     struct {
													Popup struct {
														UnifiedSharePanelRenderer struct {
															TrackingParams     string `json:"trackingParams"`
															ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
														} `json:"unifiedSharePanelRenderer"`
													} `json:"popup"`
													PopupType string `json:"popupType"`
													BeReused  bool   `json:"beReused"`
												} `json:"openPopupAction"`
											} `json:"commands"`
										} `json:"shareEntityServiceEndpoint"`
									} `json:"serviceEndpoint,omitempty"`
								} `json:"buttonRenderer,omitempty"`
							} `json:"topLevelButtons"`
							Accessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
						} `json:"menuRenderer"`
					} `json:"menu"`
					ThumbnailOverlays []struct {
						ThumbnailOverlaySidePanelRenderer struct {
							Text struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Icon struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
						} `json:"thumbnailOverlaySidePanelRenderer"`
					} `json:"thumbnailOverlays"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								Url         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						WatchEndpoint struct {
							VideoId        string `json:"videoId"`
							PlaylistId     string `json:"playlistId"`
							LoggingContext struct {
								VssLoggingContext struct {
									SerializedContextData string `json:"serializedContextData"`
								} `json:"vssLoggingContext"`
							} `json:"loggingContext"`
							WatchEndpointSupportedOnesieConfig struct {
								Html5PlaybackOnesieConfig struct {
									CommonConfig struct {
										Url string `json:"url"`
									} `json:"commonConfig"`
								} `json:"html5PlaybackOnesieConfig"`
							} `json:"watchEndpointSupportedOnesieConfig"`
						} `json:"watchEndpoint"`
					} `json:"navigationEndpoint"`
					Description struct {
						Runs []struct {
							Text               string `json:"text"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								UrlEndpoint struct {
									Url      string `json:"url"`
									Target   string `json:"target"`
									Nofollow bool   `json:"nofollow"`
								} `json:"urlEndpoint"`
							} `json:"navigationEndpoint,omitempty"`
						} `json:"runs"`
					} `json:"description"`
					ShowMoreText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"showMoreText"`
				} `json:"playlistSidebarPrimaryInfoRenderer,omitempty"`
				PlaylistSidebarSecondaryInfoRenderer struct {
					VideoOwner struct {
						VideoOwnerRenderer struct {
							Thumbnail struct {
								Thumbnails []struct {
									Url    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
							} `json:"thumbnail"`
							Title struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												Url         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												ApiUrl      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseId         string `json:"browseId"`
											CanonicalBaseUrl string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
								} `json:"runs"`
							} `json:"title"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										Url         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
										ApiUrl      string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								BrowseEndpoint struct {
									BrowseId         string `json:"browseId"`
									CanonicalBaseUrl string `json:"canonicalBaseUrl"`
								} `json:"browseEndpoint"`
							} `json:"navigationEndpoint"`
							TrackingParams string `json:"trackingParams"`
						} `json:"videoOwnerRenderer"`
					} `json:"videoOwner"`
					Button struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"text"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										IgnoreNavigation bool `json:"ignoreNavigation"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								ModalEndpoint struct {
									Modal struct {
										ModalWithTitleAndButtonRenderer struct {
											Title struct {
												SimpleText string `json:"simpleText"`
											} `json:"title"`
											Content struct {
												SimpleText string `json:"simpleText"`
											} `json:"content"`
											Button struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Text       struct {
														SimpleText string `json:"simpleText"`
													} `json:"text"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																Url         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														SignInEndpoint struct {
															NextEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		Url         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																		ApiUrl      string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																BrowseEndpoint struct {
																	BrowseId string `json:"browseId"`
																} `json:"browseEndpoint"`
															} `json:"nextEndpoint"`
															ContinueAction string `json:"continueAction"`
															IdamTag        string `json:"idamTag"`
														} `json:"signInEndpoint"`
													} `json:"navigationEndpoint"`
													TrackingParams string `json:"trackingParams"`
												} `json:"buttonRenderer"`
											} `json:"button"`
										} `json:"modalWithTitleAndButtonRenderer"`
									} `json:"modal"`
								} `json:"modalEndpoint"`
							} `json:"navigationEndpoint"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"button"`
				} `json:"playlistSidebarSecondaryInfoRenderer,omitempty"`
			} `json:"items"`
			TrackingParams string `json:"trackingParams"`
		} `json:"playlistSidebarRenderer"`
	} `json:"sidebar"`
}

type videoAPIResponse struct {
	ResponseContext struct {
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MaxAgeSeconds             int `json:"maxAgeSeconds"`
		MainAppWebResponseContext struct {
			DatasyncId string `json:"datasyncId"`
			LoggedOut  bool   `json:"loggedOut"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	PlayabilityStatus struct {
		Status          string `json:"status"`
		PlayableInEmbed bool   `json:"playableInEmbed"`
		Miniplayer      struct {
			MiniplayerRenderer struct {
				PlaybackMode string `json:"playbackMode"`
			} `json:"miniplayerRenderer"`
		} `json:"miniplayer"`
		ContextParams string `json:"contextParams"`
	} `json:"playabilityStatus"`
	StreamingData streamingData `json:"streamingData"`
	PlayerAds     []struct {
		PlayerLegacyDesktopWatchAdsRenderer struct {
			PlayerAdParams struct {
				ShowContentThumbnail bool   `json:"showContentThumbnail"`
				EnabledEngageTypes   string `json:"enabledEngageTypes"`
			} `json:"playerAdParams"`
			GutParams struct {
				Tag string `json:"tag"`
			} `json:"gutParams"`
			ShowCompanion bool `json:"showCompanion"`
			ShowInstream  bool `json:"showInstream"`
			UseGut        bool `json:"useGut"`
		} `json:"playerLegacyDesktopWatchAdsRenderer"`
	} `json:"playerAds"`
	PlaybackTracking struct {
		VideostatsPlaybackUrl struct {
			BaseUrl string `json:"baseUrl"`
		} `json:"videostatsPlaybackUrl"`
		VideostatsDelayplayUrl struct {
			BaseUrl string `json:"baseUrl"`
		} `json:"videostatsDelayplayUrl"`
		VideostatsWatchtimeUrl struct {
			BaseUrl string `json:"baseUrl"`
		} `json:"videostatsWatchtimeUrl"`
		PtrackingUrl struct {
			BaseUrl string `json:"baseUrl"`
		} `json:"ptrackingUrl"`
		QoeUrl struct {
			BaseUrl string `json:"baseUrl"`
		} `json:"qoeUrl"`
		AtrUrl struct {
			BaseUrl                 string `json:"baseUrl"`
			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
		} `json:"atrUrl"`
		VideostatsScheduledFlushWalltimeSeconds []int `json:"videostatsScheduledFlushWalltimeSeconds"`
		VideostatsDefaultFlushIntervalSeconds   int   `json:"videostatsDefaultFlushIntervalSeconds"`
		YoutubeRemarketingUrl                   struct {
			BaseUrl                 string `json:"baseUrl"`
			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
		} `json:"youtubeRemarketingUrl"`
	} `json:"playbackTracking"`
	VideoDetails struct {
		VideoId          string   `json:"videoId"`
		Title            string   `json:"title"`
		LengthSeconds    string   `json:"lengthSeconds"`
		Keywords         []string `json:"keywords"`
		ChannelId        string   `json:"channelId"`
		IsOwnerViewing   bool     `json:"isOwnerViewing"`
		ShortDescription string   `json:"shortDescription"`
		IsCrawlable      bool     `json:"isCrawlable"`
		Thumbnail        struct {
			Thumbnails []struct {
				Url    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"thumbnails"`
		} `json:"thumbnail"`
		AverageRating     float64 `json:"averageRating"`
		AllowRatings      bool    `json:"allowRatings"`
		ViewCount         string  `json:"viewCount"`
		Author            string  `json:"author"`
		IsPrivate         bool    `json:"isPrivate"`
		IsUnpluggedCorpus bool    `json:"isUnpluggedCorpus"`
		IsLiveContent     bool    `json:"isLiveContent"`
	} `json:"videoDetails"`
	Annotations []struct {
		PlayerAnnotationsExpandedRenderer struct {
			FeaturedChannel struct {
				StartTimeMs string `json:"startTimeMs"`
				EndTimeMs   string `json:"endTimeMs"`
				Watermark   struct {
					Thumbnails []struct {
						Url    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"thumbnails"`
				} `json:"watermark"`
				TrackingParams     string `json:"trackingParams"`
				NavigationEndpoint struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							Url         string `json:"url"`
							WebPageType string `json:"webPageType"`
							RootVe      int    `json:"rootVe"`
							ApiUrl      string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					BrowseEndpoint struct {
						BrowseId string `json:"browseId"`
					} `json:"browseEndpoint"`
				} `json:"navigationEndpoint"`
				ChannelName     string `json:"channelName"`
				SubscribeButton struct {
					SubscribeButtonRenderer struct {
						ButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"buttonText"`
						Subscribed           bool   `json:"subscribed"`
						Enabled              bool   `json:"enabled"`
						Type                 string `json:"type"`
						ChannelId            string `json:"channelId"`
						ShowPreferences      bool   `json:"showPreferences"`
						SubscribedButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"subscribedButtonText"`
						UnsubscribedButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"unsubscribedButtonText"`
						TrackingParams        string `json:"trackingParams"`
						UnsubscribeButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"unsubscribeButtonText"`
						ServiceEndpoints []struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									SendPost bool   `json:"sendPost"`
									ApiUrl   string `json:"apiUrl,omitempty"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							SubscribeEndpoint struct {
								ChannelIds []string `json:"channelIds"`
								Params     string   `json:"params"`
							} `json:"subscribeEndpoint,omitempty"`
							SignalServiceEndpoint struct {
								Signal  string `json:"signal"`
								Actions []struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									OpenPopupAction     struct {
										Popup struct {
											ConfirmDialogRenderer struct {
												TrackingParams string `json:"trackingParams"`
												DialogMessages []struct {
													Runs []struct {
														Text string `json:"text"`
													} `json:"runs"`
												} `json:"dialogMessages"`
												ConfirmButton struct {
													ButtonRenderer struct {
														Style      string `json:"style"`
														Size       string `json:"size"`
														IsDisabled bool   `json:"isDisabled"`
														Text       struct {
															Runs []struct {
																Text string `json:"text"`
															} `json:"runs"`
														} `json:"text"`
														ServiceEndpoint struct {
															ClickTrackingParams string `json:"clickTrackingParams"`
															CommandMetadata     struct {
																WebCommandMetadata struct {
																	SendPost bool   `json:"sendPost"`
																	ApiUrl   string `json:"apiUrl"`
																} `json:"webCommandMetadata"`
															} `json:"commandMetadata"`
															UnsubscribeEndpoint struct {
																ChannelIds []string `json:"channelIds"`
																Params     string   `json:"params"`
															} `json:"unsubscribeEndpoint"`
														} `json:"serviceEndpoint"`
														Accessibility struct {
															Label string `json:"label"`
														} `json:"accessibility"`
														TrackingParams string `json:"trackingParams"`
													} `json:"buttonRenderer"`
												} `json:"confirmButton"`
												CancelButton struct {
													ButtonRenderer struct {
														Style      string `json:"style"`
														Size       string `json:"size"`
														IsDisabled bool   `json:"isDisabled"`
														Text       struct {
															Runs []struct {
																Text string `json:"text"`
															} `json:"runs"`
														} `json:"text"`
														Accessibility struct {
															Label string `json:"label"`
														} `json:"accessibility"`
														TrackingParams string `json:"trackingParams"`
													} `json:"buttonRenderer"`
												} `json:"cancelButton"`
												PrimaryIsCancel bool `json:"primaryIsCancel"`
											} `json:"confirmDialogRenderer"`
										} `json:"popup"`
										PopupType string `json:"popupType"`
									} `json:"openPopupAction"`
								} `json:"actions"`
							} `json:"signalServiceEndpoint,omitempty"`
						} `json:"serviceEndpoints"`
						SubscribeAccessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"subscribeAccessibility"`
						UnsubscribeAccessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"unsubscribeAccessibility"`
					} `json:"subscribeButtonRenderer"`
				} `json:"subscribeButton"`
			} `json:"featuredChannel"`
			AllowSwipeDismiss bool   `json:"allowSwipeDismiss"`
			AnnotationId      string `json:"annotationId"`
		} `json:"playerAnnotationsExpandedRenderer"`
	} `json:"annotations"`
	PlayerConfig struct {
		AudioConfig struct {
			LoudnessDb              float64 `json:"loudnessDb"`
			PerceptualLoudnessDb    float64 `json:"perceptualLoudnessDb"`
			EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
		} `json:"audioConfig"`
		StreamSelectionConfig struct {
			MaxBitrate string `json:"maxBitrate"`
		} `json:"streamSelectionConfig"`
		MediaCommonConfig struct {
			DynamicReadaheadConfig struct {
				MaxReadAheadMediaTimeMs int `json:"maxReadAheadMediaTimeMs"`
				MinReadAheadMediaTimeMs int `json:"minReadAheadMediaTimeMs"`
				ReadAheadGrowthRateMs   int `json:"readAheadGrowthRateMs"`
			} `json:"dynamicReadaheadConfig"`
		} `json:"mediaCommonConfig"`
		WebPlayerConfig struct {
			WebPlayerActionsPorting struct {
				GetSharePanelCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							ApiUrl   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					WebPlayerShareEntityServiceEndpoint struct {
						SerializedShareEntity string `json:"serializedShareEntity"`
					} `json:"webPlayerShareEntityServiceEndpoint"`
				} `json:"getSharePanelCommand"`
				SubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							ApiUrl   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					SubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"subscribeEndpoint"`
				} `json:"subscribeCommand"`
				UnsubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							ApiUrl   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					UnsubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"unsubscribeEndpoint"`
				} `json:"unsubscribeCommand"`
				AddToWatchLaterCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							ApiUrl   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					PlaylistEditEndpoint struct {
						PlaylistId string `json:"playlistId"`
						Actions    []struct {
							AddedVideoId string `json:"addedVideoId"`
							Action       string `json:"action"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"addToWatchLaterCommand"`
				RemoveFromWatchLaterCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							ApiUrl   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					PlaylistEditEndpoint struct {
						PlaylistId string `json:"playlistId"`
						Actions    []struct {
							Action         string `json:"action"`
							RemovedVideoId string `json:"removedVideoId"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"removeFromWatchLaterCommand"`
			} `json:"webPlayerActionsPorting"`
		} `json:"webPlayerConfig"`
	} `json:"playerConfig"`
	Storyboards struct {
		PlayerStoryboardSpecRenderer struct {
			Spec string `json:"spec"`
		} `json:"playerStoryboardSpecRenderer"`
	} `json:"storyboards"`
	Microformat struct {
		PlayerMicroformatRenderer struct {
			Thumbnail struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			Embed struct {
				IframeUrl      string `json:"iframeUrl"`
				FlashUrl       string `json:"flashUrl"`
				Width          int    `json:"width"`
				Height         int    `json:"height"`
				FlashSecureUrl string `json:"flashSecureUrl"`
			} `json:"embed"`
			Title struct {
				SimpleText string `json:"simpleText"`
			} `json:"title"`
			Description struct {
				SimpleText string `json:"simpleText"`
			} `json:"description"`
			LengthSeconds      string   `json:"lengthSeconds"`
			OwnerProfileUrl    string   `json:"ownerProfileUrl"`
			ExternalChannelId  string   `json:"externalChannelId"`
			IsFamilySafe       bool     `json:"isFamilySafe"`
			AvailableCountries []string `json:"availableCountries"`
			IsUnlisted         bool     `json:"isUnlisted"`
			HasYpcMetadata     bool     `json:"hasYpcMetadata"`
			ViewCount          string   `json:"viewCount"`
			Category           string   `json:"category"`
			PublishDate        string   `json:"publishDate"`
			OwnerChannelName   string   `json:"ownerChannelName"`
			UploadDate         string   `json:"uploadDate"`
		} `json:"playerMicroformatRenderer"`
	} `json:"microformat"`
	TrackingParams string `json:"trackingParams"`
	Attestation    struct {
		PlayerAttestationRenderer struct {
			Challenge    string `json:"challenge"`
			BotguardData struct {
				Program            string `json:"program"`
				InterpreterSafeUrl struct {
					PrivateDoNotAccessOrElseTrustedResourceUrlWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
				} `json:"interpreterSafeUrl"`
				ServerEnvironment int `json:"serverEnvironment"`
			} `json:"botguardData"`
		} `json:"playerAttestationRenderer"`
	} `json:"attestation"`
	AdPlacements []struct {
		AdPlacementRenderer struct {
			Config struct {
				AdPlacementConfig struct {
					Kind         string `json:"kind"`
					AdTimeOffset struct {
						OffsetStartMilliseconds string `json:"offsetStartMilliseconds"`
						OffsetEndMilliseconds   string `json:"offsetEndMilliseconds"`
					} `json:"adTimeOffset"`
					HideCueRangeMarker bool `json:"hideCueRangeMarker"`
				} `json:"adPlacementConfig"`
			} `json:"config"`
			Renderer struct {
				AdBreakServiceRenderer struct {
					PrefetchMilliseconds string `json:"prefetchMilliseconds"`
					GetAdBreakUrl        string `json:"getAdBreakUrl"`
				} `json:"adBreakServiceRenderer"`
			} `json:"renderer"`
		} `json:"adPlacementRenderer"`
	} `json:"adPlacements"`
	FrameworkUpdates struct {
		EntityBatchUpdate struct {
			Mutations []struct {
				EntityKey string `json:"entityKey"`
				Type      string `json:"type"`
				Payload   struct {
					OfflineabilityEntity struct {
						Key                     string `json:"key"`
						AccessState             string `json:"accessState"`
						AddToOfflineButtonState string `json:"addToOfflineButtonState"`
					} `json:"offlineabilityEntity"`
				} `json:"payload"`
			} `json:"mutations"`
			Timestamp struct {
				Seconds string `json:"seconds"`
				Nanos   int    `json:"nanos"`
			} `json:"timestamp"`
		} `json:"entityBatchUpdate"`
	} `json:"frameworkUpdates"`
}

type youtubeConfigSet struct {
	ClientCanaryState        string `json:"CLIENT_CANARY_STATE"`
	Device                   string `json:"DEVICE"`
	DisableYTImgDelayLoading bool   `json:"DISABLE_YT_IMG_DELAY_LOADING"`
	ElementPoolDefaultCap    int    `json:"ELEMENT_POOL_DEFAULT_CAP"`
	EventID                  string `json:"EVENT_ID"`
	ExperimentFlags          struct {
		AllowHttpsStreamingForAll                                      bool          `json:"allow_https_streaming_for_all"`
		AllowSkipNetworkless                                           bool          `json:"allow_skip_networkless"`
		AutoescapeTempdataUrl                                          bool          `json:"autoescape_tempdata_url"`
		BotguardPeriodicRefresh                                        bool          `json:"botguard_periodic_refresh"`
		BrowseNextContinuationsMigrationPlaylist                       bool          `json:"browse_next_continuations_migration_playlist"`
		CacheUtcOffsetMinutesInPrefCookie                              bool          `json:"cache_utc_offset_minutes_in_pref_cookie"`
		CancelPendingNavs                                              bool          `json:"cancel_pending_navs"`
		CheckUserLactAtPromptShownTimeOnWeb                            bool          `json:"check_user_lact_at_prompt_shown_time_on_web"`
		ClearUserPartitionedLs                                         bool          `json:"clear_user_partitioned_ls"`
		ClipsEnableSpeedLines                                          bool          `json:"clips_enable_speed_lines"`
		ColdMissingHistory                                             bool          `json:"cold_missing_history"`
		CsiOnGel                                                       bool          `json:"csi_on_gel"`
		DecorateAutoplayRenderer                                       bool          `json:"decorate_autoplay_renderer"`
		DeferMenus                                                     bool          `json:"defer_menus"`
		DeferOverlays                                                  bool          `json:"defer_overlays"`
		DeferRenderingOutsideVisibleArea                               bool          `json:"defer_rendering_outside_visible_area"`
		DeprecatePairServletEnabled                                    bool          `json:"deprecate_pair_servlet_enabled"`
		DesktopAddToPlaylistRendererDialogPopup                        bool          `json:"desktop_add_to_playlist_renderer_dialog_popup"`
		DesktopAdjustTouchTarget                                       bool          `json:"desktop_adjust_touch_target"`
		DesktopAnimateMiniplayer                                       bool          `json:"desktop_animate_miniplayer"`
		DesktopClientRelease                                           bool          `json:"desktop_client_release"`
		DesktopEnableWcrMultiStageCanary                               bool          `json:"desktop_enable_wcr_multi_stage_canary"`
		DesktopFixCarouselVideoTimeout                                 bool          `json:"desktop_fix_carousel_video_timeout"`
		DesktopKeyboardCaptureKeydownKillswitch                        bool          `json:"desktop_keyboard_capture_keydown_killswitch"`
		DesktopMixUseSampledColorForBottomBar                          bool          `json:"desktop_mix_use_sampled_color_for_bottom_bar"`
		DesktopMixUseSampledColorForBottomBarSearch                    bool          `json:"desktop_mix_use_sampled_color_for_bottom_bar_search"`
		DesktopMixUseSampledColorForBottomBarWatchNext                 bool          `json:"desktop_mix_use_sampled_color_for_bottom_bar_watch_next"`
		DesktopNotificationHighPriorityIgnorePush                      bool          `json:"desktop_notification_high_priority_ignore_push"`
		DesktopNotificationSetTitleBar                                 bool          `json:"desktop_notification_set_title_bar"`
		DesktopPersistentMenu                                          bool          `json:"desktop_persistent_menu"`
		DesktopPlayerTouchGestures                                     bool          `json:"desktop_player_touch_gestures"`
		DesktopSearchProminentThumbs                                   bool          `json:"desktop_search_prominent_thumbs"`
		DesktopSparklesLightCtaButton                                  bool          `json:"desktop_sparkles_light_cta_button"`
		DesktopSwipeableGuide                                          bool          `json:"desktop_swipeable_guide"`
		DesktopTextAdsGrayVisurl                                       bool          `json:"desktop_text_ads_gray_visurl"`
		DesktopThemeableVulcan                                         bool          `json:"desktop_themeable_vulcan"`
		DesktopTouchGesturesUsageLog                                   bool          `json:"desktop_touch_gestures_usage_log"`
		DisableChildNodeAutoFormattedStrings                           bool          `json:"disable_child_node_auto_formatted_strings"`
		DisableDependencyInjection                                     bool          `json:"disable_dependency_injection"`
		DisableFeaturesForSupex                                        bool          `json:"disable_features_for_supex"`
		DisableLegacyDesktopRemoteQueue                                bool          `json:"disable_legacy_desktop_remote_queue"`
		DisableSimpleMixedDirectionFormattedStrings                    bool          `json:"disable_simple_mixed_direction_formatted_strings"`
		DisableThumbnailPreloading                                     bool          `json:"disable_thumbnail_preloading"`
		ElementPoolPopulatorAutoAbort                                  bool          `json:"element_pool_populator_auto_abort"`
		EnableBrowserCookieStatusMonitoring                            bool          `json:"enable_browser_cookie_status_monitoring"`
		EnableBusinessEmailRevealServletMigration                      bool          `json:"enable_business_email_reveal_servlet_migration"`
		EnableButtonBehaviorReuse                                      bool          `json:"enable_button_behavior_reuse"`
		EnableCallToActionClarificationRendererBottomSectionConditions bool          `json:"enable_call_to_action_clarification_renderer_bottom_section_conditions"`
		EnableChannelCreationAvatarEditing                             bool          `json:"enable_channel_creation_avatar_editing"`
		EnableClientSliLogging                                         bool          `json:"enable_client_sli_logging"`
		EnableClientStreamzWeb                                         bool          `json:"enable_client_streamz_web"`
		EnableCommentsContinuationCommandForWeb                        bool          `json:"enable_comments_continuation_command_for_web"`
		EnableDockedChatMessages                                       bool          `json:"enable_docked_chat_messages"`
		EnableDownloadsQualitySelector                                 bool          `json:"enable_downloads_quality_selector"`
		EnableFullyExpandedClipRangeInProgressBar                      bool          `json:"enable_fully_expanded_clip_range_in_progress_bar"`
		EnableGelLogCommands                                           bool          `json:"enable_gel_log_commands"`
		EnableGetAccountSwitcherEndpointOnWebfe                        bool          `json:"enable_get_account_switcher_endpoint_on_webfe"`
		EnableGrayVisurl                                               bool          `json:"enable_gray_visurl"`
		EnableGuideDownloadsEntryRenderer                              bool          `json:"enable_guide_downloads_entry_renderer"`
		EnableMainAppClientSliLogging                                  bool          `json:"enable_main_app_client_sli_logging"`
		EnableMastheadQuartilePingFix                                  bool          `json:"enable_masthead_quartile_ping_fix"`
		EnableMembershipsAndPurchases                                  bool          `json:"enable_memberships_and_purchases"`
		EnableMentionsInReposts                                        bool          `json:"enable_mentions_in_reposts"`
		EnableMicroformatData                                          bool          `json:"enable_microformat_data"`
		EnableMixedDirectionFormattedStrings                           bool          `json:"enable_mixed_direction_formatted_strings"`
		EnableMultiImagePostCreation                                   bool          `json:"enable_multi_image_post_creation"`
		EnableNamesHandlesAccountSwitcher                              bool          `json:"enable_names_handles_account_switcher"`
		EnableOfferSuppression                                         bool          `json:"enable_offer_suppression"`
		EnablePollChoiceBorderOnWeb                                    bool          `json:"enable_poll_choice_border_on_web"`
		EnablePolymerResin                                             bool          `json:"enable_polymer_resin"`
		EnablePolymerResinMigration                                    bool          `json:"enable_polymer_resin_migration"`
		EnablePostCctLinks                                             bool          `json:"enable_post_cct_links"`
		EnablePostScheduling                                           bool          `json:"enable_post_scheduling"`
		EnablePremiumVoluntaryPause                                    bool          `json:"enable_premium_voluntary_pause"`
		EnableProgrammedPlaylistRedesign                               bool          `json:"enable_programmed_playlist_redesign"`
		EnablePurchaseActivityInPaidMemberships                        bool          `json:"enable_purchase_activity_in_paid_memberships"`
		EnableReelWatchSequence                                        bool          `json:"enable_reel_watch_sequence"`
		EnableServiceAjaxCsn                                           bool          `json:"enable_service_ajax_csn"`
		EnableServletErrorsStreamz                                     bool          `json:"enable_servlet_errors_streamz"`
		EnableServletStreamz                                           bool          `json:"enable_servlet_streamz"`
		EnableSharePanelPageAsScreenLayer                              bool          `json:"enable_share_panel_page_as_screen_layer"`
		EnableShortsResolveUrlWithZeroFrame                            bool          `json:"enable_shorts_resolve_url_with_zero_frame"`
		EnableSignals                                                  bool          `json:"enable_signals"`
		EnableSliFlush                                                 bool          `json:"enable_sli_flush"`
		EnableStreamlineRepostFlow                                     bool          `json:"enable_streamline_repost_flow"`
		EnableTopsoilWtaForHalftimeLiveInfra                           bool          `json:"enable_topsoil_wta_for_halftime_live_infra"`
		EnableUnavailableVideosWatchPage                               bool          `json:"enable_unavailable_videos_watch_page"`
		EnableUnifiedShowPageWebClient                                 bool          `json:"enable_unified_show_page_web_client"`
		EnableWatchNextPauseAutoplayLact                               bool          `json:"enable_watch_next_pause_autoplay_lact"`
		EnableWebKetchupHeroAnimation                                  bool          `json:"enable_web_ketchup_hero_animation"`
		EnableWebPosterHoverAnimation                                  bool          `json:"enable_web_poster_hover_animation"`
		EnableYoodle                                                   bool          `json:"enable_yoodle"`
		EnableYpcSpinners                                              bool          `json:"enable_ypc_spinners"`
		EnableYtcRefundsSubmitFormSignalAction                         bool          `json:"enable_ytc_refunds_submit_form_signal_action"`
		EnableYtcSelfServeRefunds                                      bool          `json:"enable_ytc_self_serve_refunds"`
		EndpointHandlerLoggingCleanupKillswitch                        bool          `json:"endpoint_handler_logging_cleanup_killswitch"`
		ExportNetworklessOptions                                       bool          `json:"export_networkless_options"`
		ExternalFullscreen                                             bool          `json:"external_fullscreen"`
		ExternalFullscreenWithEdu                                      bool          `json:"external_fullscreen_with_edu"`
		FillWebPlayerContextConfig                                     bool          `json:"fill_web_player_context_config"`
		ForwardDomainAdminStateOnEmbeds                                bool          `json:"forward_domain_admin_state_on_embeds"`
		GfeedbackForSignedOutUsersEnabled                              bool          `json:"gfeedback_for_signed_out_users_enabled"`
		GlobalSpacebarPause                                            bool          `json:"global_spacebar_pause"`
		Html5EnableSingleVideoVodIvarOnPacf                            bool          `json:"html5_enable_single_video_vod_ivar_on_pacf"`
		Html5EnableVideoOverlayOnInplayerSlotForTv                     bool          `json:"html5_enable_video_overlay_on_inplayer_slot_for_tv"`
		Html5PacfEnableDai                                             bool          `json:"html5_pacf_enable_dai"`
		Html5UserPartitionedLs                                         bool          `json:"html5_user_partitioned_ls"`
		IncludeAutoplayCountInPlaylists                                bool          `json:"include_autoplay_count_in_playlists"`
		InlinePlaybackDisableEnsureHover                               bool          `json:"inline_playback_disable_ensure_hover"`
		IsBrowserSupportForWebcamStreaming                             bool          `json:"is_browser_support_for_webcam_streaming"`
		IsPartOfAnyUserEngagementExperiment                            bool          `json:"is_part_of_any_user_engagement_experiment"`
		KevlarActionRouterNodeRemoveSenderBehavior                     bool          `json:"kevlar_action_router_node_remove_sender_behavior"`
		KevlarAllowPlaylistReorder                                     bool          `json:"kevlar_allow_playlist_reorder"`
		KevlarAllowQueueReorder                                        bool          `json:"kevlar_allow_queue_reorder"`
		KevlarAppInitializerCleanup                                    bool          `json:"kevlar_app_initializer_cleanup"`
		KevlarAppShortcuts                                             bool          `json:"kevlar_app_shortcuts"`
		KevlarAppshellServiceWorker                                    bool          `json:"kevlar_appshell_service_worker"`
		KevlarAutofocusMenuOnKeyboardNav                               bool          `json:"kevlar_autofocus_menu_on_keyboard_nav"`
		KevlarAutonavMiniplayerFix                                     bool          `json:"kevlar_autonav_miniplayer_fix"`
		KevlarAutonavPopupFiltering                                    bool          `json:"kevlar_autonav_popup_filtering"`
		KevlarBackgroundColorUpdate                                    bool          `json:"kevlar_background_color_update"`
		KevlarCacheColdLoadResponse                                    bool          `json:"kevlar_cache_cold_load_response"`
		KevlarCacheInitialData                                         bool          `json:"kevlar_cache_initial_data"`
		KevlarCacheOnTtlBrowse                                         bool          `json:"kevlar_cache_on_ttl_browse"`
		KevlarCacheOnTtlPlayer                                         bool          `json:"kevlar_cache_on_ttl_player"`
		KevlarCacheOnTtlSearch                                         bool          `json:"kevlar_cache_on_ttl_search"`
		KevlarCalculateGridCollapsible                                 bool          `json:"kevlar_calculate_grid_collapsible"`
		KevlarCancelScheduledCommentJobsOnNavigate                     bool          `json:"kevlar_cancel_scheduled_comment_jobs_on_navigate"`
		KevlarCenterSearchResults                                      bool          `json:"kevlar_center_search_results"`
		KevlarChannelTrailerMultiAttach                                bool          `json:"kevlar_channel_trailer_multi_attach"`
		KevlarChannelsPlayerHandleMissingSwfconfig                     bool          `json:"kevlar_channels_player_handle_missing_swfconfig"`
		KevlarChaptersListViewSeekByChapter                            bool          `json:"kevlar_chapters_list_view_seek_by_chapter"`
		KevlarCleanUp                                                  bool          `json:"kevlar_clean_up"`
		KevlarClearNonDisplayableUrlParams                             bool          `json:"kevlar_clear_non_displayable_url_params"`
		KevlarClientSaveSubsPreferences                                bool          `json:"kevlar_client_save_subs_preferences"`
		KevlarClientSideFillerData                                     bool          `json:"kevlar_client_side_filler_data"`
		KevlarClientSideScreens                                        bool          `json:"kevlar_client_side_screens"`
		KevlarCollectBatteryNetworkStatus                              bool          `json:"kevlar_collect_battery_network_status"`
		KevlarCollectHoverTouchSupport                                 bool          `json:"kevlar_collect_hover_touch_support"`
		KevlarCommandHandler                                           bool          `json:"kevlar_command_handler"`
		KevlarCommandHandlerClicks                                     bool          `json:"kevlar_command_handler_clicks"`
		KevlarCommandHandlerInitPlugin                                 bool          `json:"kevlar_command_handler_init_plugin"`
		KevlarCommandHandlerToggleButtons                              bool          `json:"kevlar_command_handler_toggle_buttons"`
		KevlarCommandUrl                                               bool          `json:"kevlar_command_url"`
		KevlarContinuePlaybackWithoutPlayerResponse                    bool          `json:"kevlar_continue_playback_without_player_response"`
		KevlarCopyPlaylist                                             bool          `json:"kevlar_copy_playlist"`
		KevlarCtrlTapFix                                               bool          `json:"kevlar_ctrl_tap_fix"`
		KevlarDecorateEndpointWithOnesieConfig                         bool          `json:"kevlar_decorate_endpoint_with_onesie_config"`
		KevlarDelayWatchInitialData                                    bool          `json:"kevlar_delay_watch_initial_data"`
		KevlarDisableBackgroundPrefetch                                bool          `json:"kevlar_disable_background_prefetch"`
		KevlarDisableFormattedStringEmptyStyleClass                    bool          `json:"kevlar_disable_formatted_string_empty_style_class"`
		KevlarDisableHtmlImports                                       bool          `json:"kevlar_disable_html_imports"`
		KevlarDragdropFastScroll                                       bool          `json:"kevlar_dragdrop_fast_scroll"`
		KevlarDropdownFix                                              bool          `json:"kevlar_dropdown_fix"`
		KevlarDroppablePrefetchableRequests                            bool          `json:"kevlar_droppable_prefetchable_requests"`
		KevlarEagerShellBootViaOnePlatform                             bool          `json:"kevlar_eager_shell_boot_via_one_platform"`
		KevlarEnableEditablePlaylists                                  bool          `json:"kevlar_enable_editable_playlists"`
		KevlarEnableReorderablePlaylists                               bool          `json:"kevlar_enable_reorderable_playlists"`
		KevlarEnableSlis                                               bool          `json:"kevlar_enable_slis"`
		KevlarEntitiesProcessor                                        bool          `json:"kevlar_entities_processor"`
		KevlarExitFullscreenLeavingWatch                               bool          `json:"kevlar_exit_fullscreen_leaving_watch"`
		KevlarFallbackToPageDataRootVe                                 bool          `json:"kevlar_fallback_to_page_data_root_ve"`
		KevlarFetchNetworklessSupport                                  bool          `json:"kevlar_fetch_networkless_support"`
		KevlarFixMiniplayerLogging                                     bool          `json:"kevlar_fix_miniplayer_logging"`
		KevlarFixPlaylistContinuation                                  bool          `json:"kevlar_fix_playlist_continuation"`
		KevlarFlexibleMenu                                             bool          `json:"kevlar_flexible_menu"`
		KevlarFluidTouchScroll                                         bool          `json:"kevlar_fluid_touch_scroll"`
		KevlarFrontendQueueRecover                                     bool          `json:"kevlar_frontend_queue_recover"`
		KevlarFrontendVideoListActions                                 bool          `json:"kevlar_frontend_video_list_actions"`
		KevlarGelErrorRouting                                          bool          `json:"kevlar_gel_error_routing"`
		KevlarGuideAjaxMigration                                       bool          `json:"kevlar_guide_ajax_migration"`
		KevlarGuideLockScrollbar                                       bool          `json:"kevlar_guide_lock_scrollbar"`
		KevlarGuideRefresh                                             bool          `json:"kevlar_guide_refresh"`
		KevlarGuideStore                                               bool          `json:"kevlar_guide_store"`
		KevlarHelpUseLocale                                            bool          `json:"kevlar_help_use_locale"`
		KevlarHideDownloadButtonOnClient                               bool          `json:"kevlar_hide_download_button_on_client"`
		KevlarHidePlaylistPlaybackStatus                               bool          `json:"kevlar_hide_playlist_playback_status"`
		KevlarHidePpUrlParam                                           bool          `json:"kevlar_hide_pp_url_param"`
		KevlarHideTimeContinueUrlParam                                 bool          `json:"kevlar_hide_time_continue_url_param"`
		KevlarHomeSkeleton                                             bool          `json:"kevlar_home_skeleton"`
		KevlarHomeSkeletonHideLater                                    bool          `json:"kevlar_home_skeleton_hide_later"`
		KevlarIncludeQueryInSearchEndpoint                             bool          `json:"kevlar_include_query_in_search_endpoint"`
		KevlarInjector                                                 bool          `json:"kevlar_injector"`
		KevlarInlinedHtmlTemplatesPolymerFlags                         bool          `json:"kevlar_inlined_html_templates_polymer_flags"`
		KevlarJsFixes                                                  bool          `json:"kevlar_js_fixes"`
		KevlarKeyboardButtonFocus                                      bool          `json:"kevlar_keyboard_button_focus"`
		KevlarLargerThreeDotTap                                        bool          `json:"kevlar_larger_three_dot_tap"`
		KevlarLazyListResumeForAutofill                                bool          `json:"kevlar_lazy_list_resume_for_autofill"`
		KevlarLocalInnertubeResponse                                   bool          `json:"kevlar_local_innertube_response"`
		KevlarLogInitialScreen                                         bool          `json:"kevlar_log_initial_screen"`
		KevlarLoggedOutTopbarMenuMigration                             bool          `json:"kevlar_logged_out_topbar_menu_migration"`
		KevlarMacroMarkersKeyboardShortcut                             bool          `json:"kevlar_macro_markers_keyboard_shortcut"`
		KevlarMastheadStore                                            bool          `json:"kevlar_masthead_store"`
		KevlarMealbarAbovePlayer                                       bool          `json:"kevlar_mealbar_above_player"`
		KevlarMiniplayer                                               bool          `json:"kevlar_miniplayer"`
		KevlarMiniplayerExpandTop                                      bool          `json:"kevlar_miniplayer_expand_top"`
		KevlarMiniplayerNoUpdateOnDeactivate                           bool          `json:"kevlar_miniplayer_no_update_on_deactivate"`
		KevlarMiniplayerPlayPauseOnScrim                               bool          `json:"kevlar_miniplayer_play_pause_on_scrim"`
		KevlarMiniplayerQueueUserActivation                            bool          `json:"kevlar_miniplayer_queue_user_activation"`
		KevlarMiniplayerSetElementEarly                                bool          `json:"kevlar_miniplayer_set_element_early"`
		KevlarMiniplayerSetWatchNext                                   bool          `json:"kevlar_miniplayer_set_watch_next"`
		KevlarMixHandleFirstEndpointDifferent                          bool          `json:"kevlar_mix_handle_first_endpoint_different"`
		KevlarNextColdOnAuthChangeDetected                             bool          `json:"kevlar_next_cold_on_auth_change_detected"`
		KevlarNitrateDrivenTooltips                                    bool          `json:"kevlar_nitrate_driven_tooltips"`
		KevlarNoAutoscrollOnPlaylistHover                              bool          `json:"kevlar_no_autoscroll_on_playlist_hover"`
		KevlarNoRedirectToClassicKs                                    bool          `json:"kevlar_no_redirect_to_classic_ks"`
		KevlarNoUrlParams                                              bool          `json:"kevlar_no_url_params"`
		KevlarOnePickAddVideoToPlaylist                                bool          `json:"kevlar_one_pick_add_video_to_playlist"`
		KevlarOpInfra                                                  bool          `json:"kevlar_op_infra"`
		KevlarOpWarmPages                                              bool          `json:"kevlar_op_warm_pages"`
		KevlarOpWatchOfflineFallback                                   bool          `json:"kevlar_op_watch_offline_fallback"`
		KevlarPandownPolyfill                                          bool          `json:"kevlar_pandown_polyfill"`
		KevlarPassiveEventListeners                                    bool          `json:"kevlar_passive_event_listeners"`
		KevlarPersistentResponseStore                                  bool          `json:"kevlar_persistent_response_store"`
		KevlarPickerAjaxMigration                                      bool          `json:"kevlar_picker_ajax_migration"`
		KevlarPlaybackAssociatedQueue                                  bool          `json:"kevlar_playback_associated_queue"`
		KevlarPlayerCachedLoadConfig                                   bool          `json:"kevlar_player_cached_load_config"`
		KevlarPlayerDisableRvsUpdate                                   bool          `json:"kevlar_player_disable_rvs_update"`
		KevlarPlayerLoadPlayerNoOp                                     bool          `json:"kevlar_player_load_player_no_op"`
		KevlarPlayerMigrateAssetLoading                                bool          `json:"kevlar_player_migrate_asset_loading"`
		KevlarPlayerPlaylistUseLocalIndex                              bool          `json:"kevlar_player_playlist_use_local_index"`
		KevlarPlayerResponseSwfConfigWrapperKillswitch                 bool          `json:"kevlar_player_response_swf_config_wrapper_killswitch"`
		KevlarPlayerWatchEndpointNavigation                            bool          `json:"kevlar_player_watch_endpoint_navigation"`
		KevlarPlaylistAutonavLoopFix                                   bool          `json:"kevlar_playlist_autonav_loop_fix"`
		KevlarPlaylistDragHandles                                      bool          `json:"kevlar_playlist_drag_handles"`
		KevlarPlaylistResponsive                                       bool          `json:"kevlar_playlist_responsive"`
		KevlarPopulateCommandOnDownloadButton                          bool          `json:"kevlar_populate_command_on_download_button"`
		KevlarPrefetch                                                 bool          `json:"kevlar_prefetch"`
		KevlarPrefetchRedirects                                        bool          `json:"kevlar_prefetch_redirects"`
		KevlarPreparePlayerOnMiniplayerActivation                      bool          `json:"kevlar_prepare_player_on_miniplayer_activation"`
		KevlarPreventPolymerDynamicFontLoad                            bool          `json:"kevlar_prevent_polymer_dynamic_font_load"`
		KevlarQueueUseDedicatedListType                                bool          `json:"kevlar_queue_use_dedicated_list_type"`
		KevlarQueueUseUpdateApi                                        bool          `json:"kevlar_queue_use_update_api"`
		KevlarRemovePrepopulator                                       bool          `json:"kevlar_remove_prepopulator"`
		KevlarRendererstamperEventListener                             bool          `json:"kevlar_rendererstamper_event_listener"`
		KevlarResolvePlaylistEndpointAsWatchEndpoint                   bool          `json:"kevlar_resolve_playlist_endpoint_as_watch_endpoint"`
		KevlarSaveQueue                                                bool          `json:"kevlar_save_queue"`
		KevlarScrollChipsOnTouch                                       bool          `json:"kevlar_scroll_chips_on_touch"`
		KevlarScrollbarRework                                          bool          `json:"kevlar_scrollbar_rework"`
		KevlarServiceCommandCheck                                      bool          `json:"kevlar_service_command_check"`
		KevlarSetInternalPlayerSize                                    bool          `json:"kevlar_set_internal_player_size"`
		KevlarShellForDownloadsPage                                    bool          `json:"kevlar_shell_for_downloads_page"`
		KevlarShortsPageEnabled                                        bool          `json:"kevlar_shorts_page_enabled"`
		KevlarShouldMaintainStableList                                 bool          `json:"kevlar_should_maintain_stable_list"`
		KevlarSnapStateRefresh                                         bool          `json:"kevlar_snap_state_refresh"`
		KevlarStandardScrollbarColor                                   bool          `json:"kevlar_standard_scrollbar_color"`
		KevlarStartupLifecycle                                         bool          `json:"kevlar_startup_lifecycle"`
		KevlarSystemIcons                                              bool          `json:"kevlar_system_icons"`
		KevlarThemedStandardizedScrollbar                              bool          `json:"kevlar_themed_standardized_scrollbar"`
		KevlarThreeDotInk                                              bool          `json:"kevlar_three_dot_ink"`
		KevlarThumbnailFluid                                           bool          `json:"kevlar_thumbnail_fluid"`
		KevlarToastManager                                             bool          `json:"kevlar_toast_manager"`
		KevlarTopbarLogoFallbackHome                                   bool          `json:"kevlar_topbar_logo_fallback_home"`
		KevlarTouchFeedback                                            bool          `json:"kevlar_touch_feedback"`
		KevlarTouchFeedbackLockups                                     bool          `json:"kevlar_touch_feedback_lockups"`
		KevlarTouchGestureVes                                          bool          `json:"kevlar_touch_gesture_ves"`
		KevlarTranscriptEngagementPanel                                bool          `json:"kevlar_transcript_engagement_panel"`
		KevlarTranscriptPanelRefreshedStyles                           bool          `json:"kevlar_transcript_panel_refreshed_styles"`
		KevlarTunerRunDefaultCommentsDelay                             bool          `json:"kevlar_tuner_run_default_comments_delay"`
		KevlarTunerShouldDeferDetach                                   bool          `json:"kevlar_tuner_should_defer_detach"`
		KevlarTypographySpacingUpdate                                  bool          `json:"kevlar_typography_spacing_update"`
		KevlarTypographyUpdate                                         bool          `json:"kevlar_typography_update"`
		KevlarUnavailableVideoErrorUiClient                            bool          `json:"kevlar_unavailable_video_error_ui_client"`
		KevlarUnifiedServerContractClient                              bool          `json:"kevlar_unified_server_contract_client"`
		KevlarUpdateYoutubeSans                                        bool          `json:"kevlar_update_youtube_sans"`
		KevlarUpdatedLogoIcons                                         bool          `json:"kevlar_updated_logo_icons"`
		KevlarUseEndpointForChannelCreationForm                        bool          `json:"kevlar_use_endpoint_for_channel_creation_form"`
		KevlarUseOnePlatformForQueueRefresh                            bool          `json:"kevlar_use_one_platform_for_queue_refresh"`
		KevlarUsePageCommandUrl                                        bool          `json:"kevlar_use_page_command_url"`
		KevlarUsePageDataWillUpdate                                    bool          `json:"kevlar_use_page_data_will_update"`
		KevlarUsePlayerResponseForUpdates                              bool          `json:"kevlar_use_player_response_for_updates"`
		KevlarUseResponseTtlToInvalidateCache                          bool          `json:"kevlar_use_response_ttl_to_invalidate_cache"`
		KevlarUseYtdPlayer                                             bool          `json:"kevlar_use_ytd_player"`
		KevlarUserPartitionedDownloadQuality                           bool          `json:"kevlar_user_partitioned_download_quality"`
		KevlarVariableYoutubeSans                                      bool          `json:"kevlar_variable_youtube_sans"`
		KevlarVoiceLoggingFix                                          bool          `json:"kevlar_voice_logging_fix"`
		KevlarVoiceSearch                                              bool          `json:"kevlar_voice_search"`
		KevlarWatchColorUpdate                                         bool          `json:"kevlar_watch_color_update"`
		KevlarWatchDragHandles                                         bool          `json:"kevlar_watch_drag_handles"`
		KevlarWatchFocusOnEngagementPanels                             bool          `json:"kevlar_watch_focus_on_engagement_panels"`
		KevlarWatchGesturePandown                                      bool          `json:"kevlar_watch_gesture_pandown"`
		KevlarWatchIncreasedWidthThreshold                             bool          `json:"kevlar_watch_increased_width_threshold"`
		KevlarWatchJsPanelHeight                                       bool          `json:"kevlar_watch_js_panel_height"`
		KevlarWatchMetadataRefreshForLiveKillswitch                    bool          `json:"kevlar_watch_metadata_refresh_for_live_killswitch"`
		KevlarWatchMetadataRefreshLeftAlignedVideoActions              bool          `json:"kevlar_watch_metadata_refresh_left_aligned_video_actions"`
		KevlarWatchMetadataRefreshLowerCaseVideoActions                bool          `json:"kevlar_watch_metadata_refresh_lower_case_video_actions"`
		KevlarWatchMetadataRefreshNarrowerItemWrap                     bool          `json:"kevlar_watch_metadata_refresh_narrower_item_wrap"`
		KevlarWatchMetadataRefreshNormalDescriptionPreamble            bool          `json:"kevlar_watch_metadata_refresh_normal_description_preamble"`
		KevlarWatchNavigationClearAutoplayCountSessionData             bool          `json:"kevlar_watch_navigation_clear_autoplay_count_session_data"`
		KevlarWatchNextChipsContinuationsMigration                     bool          `json:"kevlar_watch_next_chips_continuations_migration"`
		KevlarWatchPanelHeightMatchesPlayer                            bool          `json:"kevlar_watch_panel_height_matches_player"`
		KevlarWatchSkeleton                                            bool          `json:"kevlar_watch_skeleton"`
		KevlarWebResponseContextYtConfigDeprecation                    bool          `json:"kevlar_web_response_context_yt_config_deprecation"`
		KevlarWoffle                                                   bool          `json:"kevlar_woffle"`
		KevlarWoffleDlManager                                          bool          `json:"kevlar_woffle_dl_manager"`
		KevlarWoffleFallbackImage                                      bool          `json:"kevlar_woffle_fallback_image"`
		KevlarWoffleSettings                                           bool          `json:"kevlar_woffle_settings"`
		KevlarWoffleUseOfflineabilityUtil                              bool          `json:"kevlar_woffle_use_offlineability_util"`
		LiveChatIncreasedMinHeight                                     bool          `json:"live_chat_increased_min_height"`
		LiveChatInitFetchNetworkManager                                bool          `json:"live_chat_init_fetch_network_manager"`
		LiveChatOverPlaylist                                           bool          `json:"live_chat_over_playlist"`
		LogHeartbeatWithLifecycles                                     bool          `json:"log_heartbeat_with_lifecycles"`
		LogVisOnTabChange                                              bool          `json:"log_vis_on_tab_change"`
		LogWebEndpointToLayer                                          bool          `json:"log_web_endpoint_to_layer"`
		MdxEnablePrivacyDisclosureUi                                   bool          `json:"mdx_enable_privacy_disclosure_ui"`
		MdxLoadCastApiBootstrapScript                                  bool          `json:"mdx_load_cast_api_bootstrap_script"`
		MigrateToUserPartitionedDownloadQuality                        bool          `json:"migrate_to_user_partitioned_download_quality"`
		NetworklessGel                                                 bool          `json:"networkless_gel"`
		NetworklessLogging                                             bool          `json:"networkless_logging"`
		NoSubCountOnSubButton                                          bool          `json:"no_sub_count_on_sub_button"`
		NwlSendFastOnUnload                                            bool          `json:"nwl_send_fast_on_unload"`
		NwlSwHealthPayloads                                            bool          `json:"nwl_sw_health_payloads"`
		NwlThrottlingRaceFix                                           bool          `json:"nwl_throttling_race_fix"`
		NwlTriggerThrottleAfterReset                                   bool          `json:"nwl_trigger_throttle_after_reset"`
		OfflineErrorHandling                                           bool          `json:"offline_error_handling"`
		OmitInnertubeApiKeyForBearerAuthHeader                         bool          `json:"omit_innertube_api_key_for_bearer_auth_header"`
		PageidAsHeaderWeb                                              bool          `json:"pageid_as_header_web"`
		PdgDesktopSuperThanksHeaderUpdate                              bool          `json:"pdg_desktop_super_thanks_header_update"`
		PdgDisableWebSuperVodExplicitClickLogging                      bool          `json:"pdg_disable_web_super_vod_explicit_click_logging"`
		PesAesAll                                                      bool          `json:"pes_aes_all"`
		PesMigrateAssociationData                                      bool          `json:"pes_migrate_association_data"`
		PlayerAllowAutonavAfterPlaylist                                bool          `json:"player_allow_autonav_after_playlist"`
		PlayerBootstrapMethod                                          bool          `json:"player_bootstrap_method"`
		PlayerDoubletapToSeek                                          bool          `json:"player_doubletap_to_seek"`
		PlayerEnablePlaybackPlaylistChange                             bool          `json:"player_enable_playback_playlist_change"`
		PlayerEndscreenEllipsisFix                                     bool          `json:"player_endscreen_ellipsis_fix"`
		Polymer2ElementPoolProperties                                  bool          `json:"polymer2_element_pool_properties"`
		Polymer2PolyfillManualFlush                                    bool          `json:"polymer2_polyfill_manual_flush"`
		PolymerBadBuildLabels                                          bool          `json:"polymer_bad_build_labels"`
		PolymerTaskManagerProxiedPromise                               bool          `json:"polymer_task_manager_proxied_promise"`
		PolymerVerifiyAppState                                         bool          `json:"polymer_verifiy_app_state"`
		PolymerVideoRendererDeferMenu                                  bool          `json:"polymer_video_renderer_defer_menu"`
		PolymerWarmThumbnailPreload                                    bool          `json:"polymer_warm_thumbnail_preload"`
		QoeSendAndWrite                                                bool          `json:"qoe_send_and_write"`
		RecordAppCrashedWeb                                            bool          `json:"record_app_crashed_web"`
		ReloadWithoutPolymerInnertube                                  bool          `json:"reload_without_polymer_innertube"`
		RemoveYtSimpleEndpointFromDesktopDisplayAdTitle                bool          `json:"remove_yt_simple_endpoint_from_desktop_display_ad_title"`
		RichGridContentVisibilityOptimization                          bool          `json:"rich_grid_content_visibility_optimization"`
		RichGridEnableEdgeToEdge                                       bool          `json:"rich_grid_enable_edge_to_edge"`
		RichGridMiniMode                                               bool          `json:"rich_grid_mini_mode"`
		ScreenManagerLogServletEi                                      bool          `json:"screen_manager_log_servlet_ei"`
		SearchUiEnablePveBuyButton                                     bool          `json:"search_ui_enable_pve_buy_button"`
		SearchUiOfficialCardsEnablePaidVirtualEventBuyButton           bool          `json:"search_ui_official_cards_enable_paid_virtual_event_buy_button"`
		SearchboxReporting                                             bool          `json:"searchbox_reporting"`
		ServePdpAtCanonicalUrl                                         bool          `json:"serve_pdp_at_canonical_url"`
		ServiceWorkerEnabled                                           bool          `json:"service_worker_enabled"`
		ServiceWorkerPushEnabled                                       bool          `json:"service_worker_push_enabled"`
		ServiceWorkerPushHomePagePrompt                                bool          `json:"service_worker_push_home_page_prompt"`
		ServiceWorkerPushWatchPagePrompt                               bool          `json:"service_worker_push_watch_page_prompt"`
		ServiceWorkerSubscribeWithVapidKey                             bool          `json:"service_worker_subscribe_with_vapid_key"`
		ShouldClearVideoDataOnPlayerCuedUnstarted                      bool          `json:"should_clear_video_data_on_player_cued_unstarted"`
		SkipEndpointParamComparison                                    bool          `json:"skip_endpoint_param_comparison"`
		SkipLsGelRetry                                                 bool          `json:"skip_ls_gel_retry"`
		SpfKevlarAssumeChunked                                         bool          `json:"spf_kevlar_assume_chunked"`
		StateMachineDynamicEventsLifecycles                            bool          `json:"state_machine_dynamic_events_lifecycles"`
		SuppressError204Logging                                        bool          `json:"suppress_error_204_logging"`
		TopbarPersistentStoreFallback                                  bool          `json:"topbar_persistent_store_fallback"`
		UseBetterPostDismissals                                        bool          `json:"use_better_post_dismissals"`
		UseDocumentLifecycles                                          bool          `json:"use_document_lifecycles"`
		UseOneplatformForVideoPreview                                  bool          `json:"use_oneplatform_for_video_preview"`
		UseScreenManagerUtil                                           bool          `json:"use_screen_manager_util"`
		UseSourceElementIfPresentForActions                            bool          `json:"use_source_element_if_present_for_actions"`
		UseUndefinedCsnAnyLayer                                        bool          `json:"use_undefined_csn_any_layer"`
		UseWatchFragments2                                             bool          `json:"use_watch_fragments2"`
		VssFinalPingSendAndWrite                                       bool          `json:"vss_final_ping_send_and_write"`
		VssPlaybackUseSendAndWrite                                     bool          `json:"vss_playback_use_send_and_write"`
		WarmLoadNavStartWeb                                            bool          `json:"warm_load_nav_start_web"`
		WarmOpCsnCleanup                                               bool          `json:"warm_op_csn_cleanup"`
		WebAlwaysLoadChatSupport                                       bool          `json:"web_always_load_chat_support"`
		WebApiUrl                                                      bool          `json:"web_api_url"`
		WebAppshellPurgeTrigger                                        bool          `json:"web_appshell_purge_trigger"`
		WebAppshellRefreshTrigger                                      bool          `json:"web_appshell_refresh_trigger"`
		WebAutonavAllowOffByDefault                                    bool          `json:"web_autonav_allow_off_by_default"`
		WebBroadcastEocsWithLifecycle                                  bool          `json:"web_broadcast_eocs_with_lifecycle"`
		WebCleanSwLogsStore                                            bool          `json:"web_clean_sw_logs_store"`
		WebCompactVideoSingleLine                                      bool          `json:"web_compact_video_single_line"`
		WebDedupeVeGrafting                                            bool          `json:"web_dedupe_ve_grafting"`
		WebDeprecateServiceAjaxMapDependency                           bool          `json:"web_deprecate_service_ajax_map_dependency"`
		WebDontCancelPendingNavigationSameUrl                          bool          `json:"web_dont_cancel_pending_navigation_same_url"`
		WebEnableAdSignalsInItContext                                  bool          `json:"web_enable_ad_signals_in_it_context"`
		WebEnableHistoryCacheMap                                       bool          `json:"web_enable_history_cache_map"`
		WebEpChevronTapTargetSize                                      bool          `json:"web_ep_chevron_tap_target_size"`
		WebFaviconImageUpdate                                          bool          `json:"web_favicon_image_update"`
		WebForwardCommandOnPbj                                         bool          `json:"web_forward_command_on_pbj"`
		WebGelTimeoutCap                                               bool          `json:"web_gel_timeout_cap"`
		WebHideAutonavHeadline                                         bool          `json:"web_hide_autonav_headline"`
		WebHideAutonavKeyline                                          bool          `json:"web_hide_autonav_keyline"`
		WebLogMemoryTotalKbytes                                        bool          `json:"web_log_memory_total_kbytes"`
		WebLogPlayerWatchNextTicks                                     bool          `json:"web_log_player_watch_next_ticks"`
		WebLogReelsTicks                                               bool          `json:"web_log_reels_ticks"`
		WebMoveAutoplayVideoUnderChip                                  bool          `json:"web_move_autoplay_video_under_chip"`
		WebOfflinePromoViaGetPlayer                                    bool          `json:"web_offline_promo_via_get_player"`
		WebOpenGuideItemsInNewTab                                      bool          `json:"web_open_guide_items_in_new_tab"`
		WebPlayerEnableIpp                                             bool          `json:"web_player_enable_ipp"`
		WebPlayerMoveAutonavToggle                                     bool          `json:"web_player_move_autonav_toggle"`
		WebPlayerTouchModeImprovements                                 bool          `json:"web_player_touch_mode_improvements"`
		WebPlayerWatchNextResponse                                     bool          `json:"web_player_watch_next_response"`
		WebPlaylistWatchPanelOverflowWithAddTo                         bool          `json:"web_playlist_watch_panel_overflow_with_add_to"`
		WebPrsTestingModeKillswitch                                    bool          `json:"web_prs_testing_mode_killswitch"`
		WebResponseProcessorSupport                                    bool          `json:"web_response_processor_support"`
		WebShortsPageEnabled                                           bool          `json:"web_shorts_page_enabled"`
		WebUseCacheForImageFallback                                    bool          `json:"web_use_cache_for_image_fallback"`
		WebUseOverflowMenuForPlaylistWatchPanel                        bool          `json:"web_use_overflow_menu_for_playlist_watch_panel"`
		WebYtConfigContext                                             bool          `json:"web_yt_config_context"`
		WoffleOrchestration                                            bool          `json:"woffle_orchestration"`
		WofflePromoViaGda                                              bool          `json:"woffle_promo_via_gda"`
		YourDataEntrypoint                                             bool          `json:"your_data_entrypoint"`
		YtidbFetchDatasyncIdsForDataCleanup                            bool          `json:"ytidb_fetch_datasync_ids_for_data_cleanup"`
		YtidbIsSupportedCacheSuccessResult                             bool          `json:"ytidb_is_supported_cache_success_result"`
		YtidbStopTransactionCommit                                     bool          `json:"ytidb_stop_transaction_commit"`
		AddtoAjaxLogWarningFraction                                    float64       `json:"addto_ajax_log_warning_fraction"`
		AutoplayPauseByLactSamplingFraction                            float64       `json:"autoplay_pause_by_lact_sampling_fraction"`
		BrowseAjaxLogWarningFraction                                   float64       `json:"browse_ajax_log_warning_fraction"`
		KevlarAsyncStampProbability                                    float64       `json:"kevlar_async_stamp_probability"`
		KevlarTunerClampDevicePixelRatio                               float64       `json:"kevlar_tuner_clamp_device_pixel_ratio"`
		KevlarTunerThumbnailFactor                                     float64       `json:"kevlar_tuner_thumbnail_factor"`
		LogWindowOnerrorFraction                                       float64       `json:"log_window_onerror_fraction"`
		PolymerReportClientUrlRequestedRate                            float64       `json:"polymer_report_client_url_requested_rate"`
		PolymerReportMissingWebNavigationEndpointRate                  float64       `json:"polymer_report_missing_web_navigation_endpoint_rate"`
		WebSystemHealthFraction                                        float64       `json:"web_system_health_fraction"`
		YtidbTransactionEndedEventRateLimit                            float64       `json:"ytidb_transaction_ended_event_rate_limit"`
		AutoplayPauseByLactSec                                         int           `json:"autoplay_pause_by_lact_sec"`
		AutoplayTime                                                   int           `json:"autoplay_time"`
		AutoplayTimeForFullscreen                                      int           `json:"autoplay_time_for_fullscreen"`
		AutoplayTimeForMusicContent                                    int           `json:"autoplay_time_for_music_content"`
		AutoplayTimeForMusicContentAfterAutoplayedVideo                int           `json:"autoplay_time_for_music_content_after_autoplayed_video"`
		BotguardAsyncSnapshotTimeoutMs                                 int           `json:"botguard_async_snapshot_timeout_ms"`
		CheckNavigatorAccuracyTimeoutMs                                int           `json:"check_navigator_accuracy_timeout_ms"`
		ClientStreamzWebFlushCount                                     int           `json:"client_streamz_web_flush_count"`
		ClientStreamzWebFlushIntervalSeconds                           int           `json:"client_streamz_web_flush_interval_seconds"`
		DesktopSearchSuggestionTapTarget                               int           `json:"desktop_search_suggestion_tap_target"`
		ExternalFullscreenButtonClickThreshold                         int           `json:"external_fullscreen_button_click_threshold"`
		ExternalFullscreenButtonShownThreshold                         int           `json:"external_fullscreen_button_shown_threshold"`
		GetAsyncTimeoutMs                                              int           `json:"get_async_timeout_ms"`
		HighPriorityFlyoutFrequency                                    int           `json:"high_priority_flyout_frequency"`
		InitialGelBatchTimeout                                         int           `json:"initial_gel_batch_timeout"`
		KevlarAsyncStampDelay                                          int           `json:"kevlar_async_stamp_delay"`
		KevlarMiniGuideWidthThreshold                                  int           `json:"kevlar_mini_guide_width_threshold"`
		KevlarPersistentGuideWidthThreshold                            int           `json:"kevlar_persistent_guide_width_threshold"`
		KevlarTimeCachingEndThreshold                                  int           `json:"kevlar_time_caching_end_threshold"`
		KevlarTimeCachingStartThreshold                                int           `json:"kevlar_time_caching_start_threshold"`
		KevlarTooltipImpressionCap                                     int           `json:"kevlar_tooltip_impression_cap"`
		KevlarTunerDefaultCommentsDelay                                int           `json:"kevlar_tuner_default_comments_delay"`
		KevlarTunerSchedulerSoftStateTimerMs                           int           `json:"kevlar_tuner_scheduler_soft_state_timer_ms"`
		KevlarTunerVisibilityTimeBetweenJobsMs                         int           `json:"kevlar_tuner_visibility_time_between_jobs_ms"`
		KevlarWatchFlexyMetadataHeight                                 int           `json:"kevlar_watch_flexy_metadata_height"`
		KevlarWatchMetadataRefreshDescriptionLines                     int           `json:"kevlar_watch_metadata_refresh_description_lines"`
		LeaderElectionCheckInterval                                    int           `json:"leader_election_check_interval"`
		LeaderElectionLeaseTtl                                         int           `json:"leader_election_lease_ttl"`
		LeaderElectionRenewalInterval                                  int           `json:"leader_election_renewal_interval"`
		LogWebMetaIntervalMs                                           int           `json:"log_web_meta_interval_ms"`
		MaxDurationToConsiderMouseoverAsHover                          int           `json:"max_duration_to_consider_mouseover_as_hover"`
		MinMouseStillDuration                                          int           `json:"min_mouse_still_duration"`
		MinimumDurationToConsiderMouseoverAsHover                      int           `json:"minimum_duration_to_consider_mouseover_as_hover"`
		NetworkPollingInterval                                         int           `json:"network_polling_interval"`
		PbjNavigateLimit                                               int           `json:"pbj_navigate_limit"`
		PostTypeIconsRearrange                                         int           `json:"post_type_icons_rearrange"`
		PrefetchCommentsMsAfterVideo                                   int           `json:"prefetch_comments_ms_after_video"`
		ServiceWorkerPushLoggedOutPromptWatches                        int           `json:"service_worker_push_logged_out_prompt_watches"`
		ServiceWorkerPushPromptCap                                     int           `json:"service_worker_push_prompt_cap"`
		ServiceWorkerPushPromptDelayMicroseconds                       int64         `json:"service_worker_push_prompt_delay_microseconds"`
		UserEngagementExperimentsRateLimitMs                           int           `json:"user_engagement_experiments_rate_limit_ms"`
		UserMentionSuggestionsEduImpressionCap                         int           `json:"user_mention_suggestions_edu_impression_cap"`
		ViewportLoadCollectionWaitTime                                 int           `json:"viewport_load_collection_wait_time"`
		VisibilityTimeBetweenJobsMs                                    int           `json:"visibility_time_between_jobs_ms"`
		WatchNextPauseAutoplayLactSec                                  int           `json:"watch_next_pause_autoplay_lact_sec"`
		WebEmulatedIdleCallbackDelay                                   int           `json:"web_emulated_idle_callback_delay"`
		WebForegroundHeartbeatIntervalMs                               int           `json:"web_foreground_heartbeat_interval_ms"`
		WebGelDebounceMs                                               int           `json:"web_gel_debounce_ms"`
		WebInlinePlayerTriggeringDelay                                 int           `json:"web_inline_player_triggering_delay"`
		WebLoggingMaxBatch                                             int           `json:"web_logging_max_batch"`
		YoodleEndTimeUtc                                               int           `json:"yoodle_end_time_utc"`
		YoodleStartTimeUtc                                             int           `json:"yoodle_start_time_utc"`
		YtidbRemakeDbRetries                                           int           `json:"ytidb_remake_db_retries"`
		WebClientReleaseProcessCriticalYoutubeWebClientVersionOverride string        `json:"WebClientReleaseProcessCritical__youtube_web_client_version_override"`
		CbV2Uxe                                                        string        `json:"cb_v2_uxe"`
		DebugForcedInternalcountrycode                                 string        `json:"debug_forced_internalcountrycode"`
		DesktopSboxIcon                                                string        `json:"desktop_sbox_icon"`
		DesktopSearchProminentThumbsStyle                              string        `json:"desktop_search_prominent_thumbs_style"`
		DesktopSearchbarStyle                                          string        `json:"desktop_searchbar_style"`
		DesktopSuggestionBoxStyle                                      string        `json:"desktop_suggestion_box_style"`
		DesktopWebClientVersionOverride                                string        `json:"desktop_web_client_version_override"`
		EmbedsWebSynthChHeadersBannedUrlsRegex                         string        `json:"embeds_web_synth_ch_headers_banned_urls_regex"`
		KevlarLinkCapturingMode                                        string        `json:"kevlar_link_capturing_mode"`
		KevlarNextUpNextEduEmoji                                       string        `json:"kevlar_next_up_next_edu_emoji"`
		LiveChatUnicodeEmojiJsonUrl                                    string        `json:"live_chat_unicode_emoji_json_url"`
		PolymerTaskManagerStatus                                       string        `json:"polymer_task_manager_status"`
		ServiceWorkerPushForceNotificationPromptTag                    string        `json:"service_worker_push_force_notification_prompt_tag"`
		ServiceWorkerScope                                             string        `json:"service_worker_scope"`
		WebClientVersionOverride                                       string        `json:"web_client_version_override"`
		YoodleBaseUrl                                                  string        `json:"yoodle_base_url"`
		YoodleWebpBaseUrl                                              string        `json:"yoodle_webp_base_url"`
		ConditionalLabIds                                              []interface{} `json:"conditional_lab_ids"`
		GuideBusinessInfoCountries                                     []string      `json:"guide_business_info_countries"`
		GuideLegalFooterEnabledCountries                               []string      `json:"guide_legal_footer_enabled_countries"`
		KevlarCommandHandlerCommandBanlist                             []interface{} `json:"kevlar_command_handler_command_banlist"`
		KevlarOpBrowseSampledPrefixIds                                 []interface{} `json:"kevlar_op_browse_sampled_prefix_ids"`
		KevlarPageServiceUrlPrefixCarveouts                            []interface{} `json:"kevlar_page_service_url_prefix_carveouts"`
		TenVideoReordering                                             []int         `json:"ten_video_reordering"`
		TwelveVideoReordering                                          []int         `json:"twelve_video_reordering"`
		WebOpContinuationTypeBanlist                                   []interface{} `json:"web_op_continuation_type_banlist"`
		WebOpEndpointBanlist                                           []interface{} `json:"web_op_endpoint_banlist"`
		WebOpSignalTypeBanlist                                         []interface{} `json:"web_op_signal_type_banlist"`
	} `json:"EXPERIMENT_FLAGS"`
	GAPIHintParams            string `json:"GAPI_HINT_PARAMS"`
	GAPIHost                  string `json:"GAPI_HOST"`
	GAPILocale                string `json:"GAPI_LOCALE"`
	GL                        string `json:"GL"`
	GoogleFeedbackProductId   string `json:"GOOGLE_FEEDBACK_PRODUCT_ID"`
	GoogleFeedbackProductData struct {
		Polymer        string `json:"polymer"`
		Polymer2       string `json:"polymer2"`
		AcceptLanguage string `json:"accept_language"`
	} `json:"GOOGLE_FEEDBACK_PRODUCT_DATA"`
	HL                     string `json:"HL"`
	HTMLDir                string `json:"HTML_DIR"`
	HTMLLang               string `json:"HTML_LANG"`
	InnerTubeApiKey        string `json:"INNERTUBE_API_KEY"`
	InnerTubeApiVersion    string `json:"INNERTUBE_API_VERSION"`
	InnerTubeClientName    string `json:"INNERTUBE_CLIENT_NAME"`
	InnerTubeClientVersion string `json:"INNERTUBE_CLIENT_VERSION"`
	InnerTubeContext       struct {
		Client struct {
			Hl               string `json:"hl"`
			Gl               string `json:"gl"`
			RemoteHost       string `json:"remoteHost"`
			DeviceMake       string `json:"deviceMake"`
			DeviceModel      string `json:"deviceModel"`
			VisitorData      string `json:"visitorData"`
			UserAgent        string `json:"userAgent"`
			ClientName       string `json:"clientName"`
			ClientVersion    string `json:"clientVersion"`
			OsName           string `json:"osName"`
			OsVersion        string `json:"osVersion"`
			OriginalUrl      string `json:"originalUrl"`
			Platform         string `json:"platform"`
			ClientFormFactor string `json:"clientFormFactor"`
			ConfigInfo       struct {
				AppInstallData string `json:"appInstallData"`
			} `json:"configInfo"`
			BrowserName    string `json:"browserName"`
			BrowserVersion string `json:"browserVersion"`
		} `json:"client"`
		User struct {
			LockedSafetyMode bool `json:"lockedSafetyMode"`
		} `json:"user"`
		Request struct {
			UseSsl bool `json:"useSsl"`
		} `json:"request"`
		ClickTracking struct {
			ClickTrackingParams string `json:"clickTrackingParams"`
		} `json:"clickTracking"`
	} `json:"INNERTUBE_CONTEXT"`
	InnerTubeContextClientName          int    `json:"INNERTUBE_CONTEXT_CLIENT_NAME"`
	InnerTubeContextClientVersion       string `json:"INNERTUBE_CONTEXT_CLIENT_VERSION"`
	InnerTubeContextGL                  string `json:"INNERTUBE_CONTEXT_GL"`
	InnerTubeContextHL                  string `json:"INNERTUBE_CONTEXT_HL"`
	LatestECatcherServiceTrackingParams struct {
		ClientName string `json:"client.name"`
	} `json:"LATEST_ECATCHER_SERVICE_TRACKING_PARAMS"`
	LoggedId       bool   `json:"LOGGED_IN"`
	PageBuildLabel string `json:"PAGE_BUILD_LABEL"`
	PageCL         int    `json:"PAGE_CL"`
	Scheduler      struct {
		UseRaf  bool `json:"useRaf"`
		Timeout int  `json:"timeout"`
	} `json:"scheduler"`
	ServerName              string `json:"SERVER_NAME"`
	SessionIndex            string `json:"SESSION_INDEX"`
	SigninUrl               string `json:"SIGNIN_URL"`
	VisitorData             string `json:"VISITOR_DATA"`
	WebPlayerContextConfigs struct {
		WebPlayerContextConfigIDKevlarWatch struct {
			TransparentBackground         bool   `json:"transparentBackground"`
			ShowMiniplayerButton          bool   `json:"showMiniplayerButton"`
			ExternalFullscreen            bool   `json:"externalFullscreen"`
			ShowMiniplayerUiWhenMinimized bool   `json:"showMiniplayerUiWhenMinimized"`
			RootElementId                 string `json:"rootElementId"`
			JsUrl                         string `json:"jsUrl"`
			CssUrl                        string `json:"cssUrl"`
			ContextId                     string `json:"contextId"`
			EventLabel                    string `json:"eventLabel"`
			ContentRegion                 string `json:"contentRegion"`
			Hl                            string `json:"hl"`
			HostLanguage                  string `json:"hostLanguage"`
			PlayerStyle                   string `json:"playerStyle"`
			InnertubeApiKey               string `json:"innertubeApiKey"`
			InnertubeApiVersion           string `json:"innertubeApiVersion"`
			InnertubeContextClientVersion string `json:"innertubeContextClientVersion"`
			Device                        struct {
				Brand            string `json:"brand"`
				Model            string `json:"model"`
				Browser          string `json:"browser"`
				BrowserVersion   string `json:"browserVersion"`
				Os               string `json:"os"`
				OsVersion        string `json:"osVersion"`
				Platform         string `json:"platform"`
				InterfaceName    string `json:"interfaceName"`
				InterfaceVersion string `json:"interfaceVersion"`
			} `json:"device"`
			SerializedExperimentIds   string `json:"serializedExperimentIds"`
			SerializedExperimentFlags string `json:"serializedExperimentFlags"`
			CspNonce                  string `json:"cspNonce"`
			CanaryState               string `json:"canaryState"`
			EnableCsiLogging          bool   `json:"enableCsiLogging"`
			CsiPageType               string `json:"csiPageType"`
			DatasyncId                string `json:"datasyncId"`
			AllowWoffleManagement     bool   `json:"allowWoffleManagement"`
		} `json:"WEB_PLAYER_CONTEXT_CONFIG_ID_KEVLAR_WATCH"`
		WebPlayerContextConfigIDKevlarChannelTrailer struct {
			RootElementId                 string `json:"rootElementId"`
			JsUrl                         string `json:"jsUrl"`
			CssUrl                        string `json:"cssUrl"`
			ContextId                     string `json:"contextId"`
			EventLabel                    string `json:"eventLabel"`
			ContentRegion                 string `json:"contentRegion"`
			Hl                            string `json:"hl"`
			HostLanguage                  string `json:"hostLanguage"`
			PlayerStyle                   string `json:"playerStyle"`
			InnertubeApiKey               string `json:"innertubeApiKey"`
			InnertubeApiVersion           string `json:"innertubeApiVersion"`
			InnertubeContextClientVersion string `json:"innertubeContextClientVersion"`
			Device                        struct {
				Brand            string `json:"brand"`
				Model            string `json:"model"`
				Browser          string `json:"browser"`
				BrowserVersion   string `json:"browserVersion"`
				Os               string `json:"os"`
				OsVersion        string `json:"osVersion"`
				Platform         string `json:"platform"`
				InterfaceName    string `json:"interfaceName"`
				InterfaceVersion string `json:"interfaceVersion"`
			} `json:"device"`
			SerializedExperimentIds   string `json:"serializedExperimentIds"`
			SerializedExperimentFlags string `json:"serializedExperimentFlags"`
			CspNonce                  string `json:"cspNonce"`
			CanaryState               string `json:"canaryState"`
			EnableCsiLogging          bool   `json:"enableCsiLogging"`
			CsiPageType               string `json:"csiPageType"`
			DatasyncId                string `json:"datasyncId"`
		} `json:"WEB_PLAYER_CONTEXT_CONFIG_ID_KEVLAR_CHANNEL_TRAILER"`
		WebPlayerContextConfigIDKevlarPlaylistOverview struct {
			RootElementId                 string `json:"rootElementId"`
			JsUrl                         string `json:"jsUrl"`
			CssUrl                        string `json:"cssUrl"`
			ContextId                     string `json:"contextId"`
			EventLabel                    string `json:"eventLabel"`
			ContentRegion                 string `json:"contentRegion"`
			Hl                            string `json:"hl"`
			HostLanguage                  string `json:"hostLanguage"`
			PlayerStyle                   string `json:"playerStyle"`
			InnertubeApiKey               string `json:"innertubeApiKey"`
			InnertubeApiVersion           string `json:"innertubeApiVersion"`
			InnertubeContextClientVersion string `json:"innertubeContextClientVersion"`
			Device                        struct {
				Brand            string `json:"brand"`
				Model            string `json:"model"`
				Browser          string `json:"browser"`
				BrowserVersion   string `json:"browserVersion"`
				Os               string `json:"os"`
				OsVersion        string `json:"osVersion"`
				Platform         string `json:"platform"`
				InterfaceName    string `json:"interfaceName"`
				InterfaceVersion string `json:"interfaceVersion"`
			} `json:"device"`
			SerializedExperimentIds   string `json:"serializedExperimentIds"`
			SerializedExperimentFlags string `json:"serializedExperimentFlags"`
			DisableSharing            bool   `json:"disableSharing"`
			HideInfo                  bool   `json:"hideInfo"`
			DisableWatchLater         bool   `json:"disableWatchLater"`
			CspNonce                  string `json:"cspNonce"`
			CanaryState               string `json:"canaryState"`
			EnableCsiLogging          bool   `json:"enableCsiLogging"`
			CsiPageType               string `json:"csiPageType"`
			DatasyncId                string `json:"datasyncId"`
		} `json:"WEB_PLAYER_CONTEXT_CONFIG_ID_KEVLAR_PLAYLIST_OVERVIEW"`
		WebPlayerContextConfigIDKevlarVerticalLandingPagePromo struct {
			RootElementId                 string `json:"rootElementId"`
			JsUrl                         string `json:"jsUrl"`
			CssUrl                        string `json:"cssUrl"`
			ContextId                     string `json:"contextId"`
			EventLabel                    string `json:"eventLabel"`
			ContentRegion                 string `json:"contentRegion"`
			Hl                            string `json:"hl"`
			HostLanguage                  string `json:"hostLanguage"`
			PlayerStyle                   string `json:"playerStyle"`
			InnertubeApiKey               string `json:"innertubeApiKey"`
			InnertubeApiVersion           string `json:"innertubeApiVersion"`
			InnertubeContextClientVersion string `json:"innertubeContextClientVersion"`
			ControlsType                  int    `json:"controlsType"`
			DisableRelatedVideos          bool   `json:"disableRelatedVideos"`
			AnnotationsLoadPolicy         int    `json:"annotationsLoadPolicy"`
			Device                        struct {
				Brand            string `json:"brand"`
				Model            string `json:"model"`
				Browser          string `json:"browser"`
				BrowserVersion   string `json:"browserVersion"`
				Os               string `json:"os"`
				OsVersion        string `json:"osVersion"`
				Platform         string `json:"platform"`
				InterfaceName    string `json:"interfaceName"`
				InterfaceVersion string `json:"interfaceVersion"`
			} `json:"device"`
			SerializedExperimentIds   string `json:"serializedExperimentIds"`
			SerializedExperimentFlags string `json:"serializedExperimentFlags"`
			HideInfo                  bool   `json:"hideInfo"`
			StartMuted                bool   `json:"startMuted"`
			EnableMutedAutoplay       bool   `json:"enableMutedAutoplay"`
			CspNonce                  string `json:"cspNonce"`
			CanaryState               string `json:"canaryState"`
			EnableCsiLogging          bool   `json:"enableCsiLogging"`
			CsiPageType               string `json:"csiPageType"`
			DatasyncId                string `json:"datasyncId"`
		} `json:"WEB_PLAYER_CONTEXT_CONFIG_ID_KEVLAR_VERTICAL_LANDING_PAGE_PROMO"`
		WebPlayerContextConfigIDKevlarSponsorShipsOffer struct {
			RootElementId                 string `json:"rootElementId"`
			JsUrl                         string `json:"jsUrl"`
			CssUrl                        string `json:"cssUrl"`
			ContextId                     string `json:"contextId"`
			EventLabel                    string `json:"eventLabel"`
			ContentRegion                 string `json:"contentRegion"`
			Hl                            string `json:"hl"`
			HostLanguage                  string `json:"hostLanguage"`
			PlayerStyle                   string `json:"playerStyle"`
			InnertubeApiKey               string `json:"innertubeApiKey"`
			InnertubeApiVersion           string `json:"innertubeApiVersion"`
			InnertubeContextClientVersion string `json:"innertubeContextClientVersion"`
			DisableRelatedVideos          bool   `json:"disableRelatedVideos"`
			AnnotationsLoadPolicy         int    `json:"annotationsLoadPolicy"`
			Device                        struct {
				Brand            string `json:"brand"`
				Model            string `json:"model"`
				Browser          string `json:"browser"`
				BrowserVersion   string `json:"browserVersion"`
				Os               string `json:"os"`
				OsVersion        string `json:"osVersion"`
				Platform         string `json:"platform"`
				InterfaceName    string `json:"interfaceName"`
				InterfaceVersion string `json:"interfaceVersion"`
			} `json:"device"`
			SerializedExperimentIds   string `json:"serializedExperimentIds"`
			SerializedExperimentFlags string `json:"serializedExperimentFlags"`
			DisableFullscreen         bool   `json:"disableFullscreen"`
			CspNonce                  string `json:"cspNonce"`
			CanaryState               string `json:"canaryState"`
			DatasyncId                string `json:"datasyncId"`
		} `json:"WEB_PLAYER_CONTEXT_CONFIG_ID_KEVLAR_SPONSORSHIPS_OFFER"`
		WebPlayerContextConfigIDKevlarShorts struct {
			RootElementId                 string `json:"rootElementId"`
			JsUrl                         string `json:"jsUrl"`
			CssUrl                        string `json:"cssUrl"`
			ContextId                     string `json:"contextId"`
			EventLabel                    string `json:"eventLabel"`
			ContentRegion                 string `json:"contentRegion"`
			Hl                            string `json:"hl"`
			HostLanguage                  string `json:"hostLanguage"`
			PlayerStyle                   string `json:"playerStyle"`
			InnertubeApiKey               string `json:"innertubeApiKey"`
			InnertubeApiVersion           string `json:"innertubeApiVersion"`
			InnertubeContextClientVersion string `json:"innertubeContextClientVersion"`
			ControlsType                  int    `json:"controlsType"`
			DisableKeyboardControls       bool   `json:"disableKeyboardControls"`
			DisableRelatedVideos          bool   `json:"disableRelatedVideos"`
			AnnotationsLoadPolicy         int    `json:"annotationsLoadPolicy"`
			Device                        struct {
				Brand            string `json:"brand"`
				Model            string `json:"model"`
				Browser          string `json:"browser"`
				BrowserVersion   string `json:"browserVersion"`
				Os               string `json:"os"`
				OsVersion        string `json:"osVersion"`
				Platform         string `json:"platform"`
				InterfaceName    string `json:"interfaceName"`
				InterfaceVersion string `json:"interfaceVersion"`
			} `json:"device"`
			SerializedExperimentIds   string `json:"serializedExperimentIds"`
			SerializedExperimentFlags string `json:"serializedExperimentFlags"`
			HideInfo                  bool   `json:"hideInfo"`
			DisableFullscreen         bool   `json:"disableFullscreen"`
			CspNonce                  string `json:"cspNonce"`
			CanaryState               string `json:"canaryState"`
			EnableCsiLogging          bool   `json:"enableCsiLogging"`
			DatasyncId                string `json:"datasyncId"`
			StoreUserVolume           bool   `json:"storeUserVolume"`
			DisableSeek               bool   `json:"disableSeek"`
			DisablePaidContentOverlay bool   `json:"disablePaidContentOverlay"`
		} `json:"WEB_PLAYER_CONTEXT_CONFIG_ID_KEVLAR_SHORTS"`
	} `json:"WEB_PLAYER_CONTEXT_CONFIGS"`
	XSRFFieldName              string `json:"XSRF_FIELD_NAME"`
	XSRFToken                  string `json:"XSRF_TOKEN"`
	YPCMBUrl                   string `json:"YPC_MB_URL"`
	YTRFamilyCreationUrl       string `json:"YTR_FAMILY_CREATION_URL"`
	ServerVersion              string `json:"SERVER_VERSION"`
	ReuseComponents            bool   `json:"REUSE_COMPONENTS"`
	StampersStableList         bool   `json:"STAMPER_STABLE_LIST"`
	DataSyncID                 string `json:"DATASYNC_ID"`
	SerializedClientConfigData string `json:"SERIALIZED_CLIENT_CONFIG_DATA"`
	LiveChatBaseTangoConfig    struct {
		ApiKey            string `json:"apiKey"`
		ChannelUri        string `json:"channelUri"`
		ClientName        string `json:"clientName"`
		RequiresAuthToken bool   `json:"requiresAuthToken"`
		SenderUri         string `json:"senderUri"`
		UseNewTango       bool   `json:"useNewTango"`
	} `json:"LIVE_CHAT_BASE_TANGO_CONFIG"`
	FexpExperiments                  []int  `json:"FEXP_EXPERIMENTS"`
	LiveChatSendMessageAction        string `json:"LIVE_CHAT_SEND_MESSAGE_ACTION"`
	RooteVEType                      int    `json:"ROOT_VE_TYPE"`
	ClientProtocol                   string `json:"CLIENT_PROTOCOL"`
	ClientTransport                  string `json:"CLIENT_TRANSPORT"`
	VisibilityTimeBetweenJobsMs      int    `json:"VISIBILITY_TIME_BETWEEN_JOBS_MS"`
	StartInTheaterMode               bool   `json:"START_IN_THEATER_MODE"`
	StartInFullWindowMode            bool   `json:"START_IN_FULL_WINDOW_MODE"`
	ServiceWorkerPromptNotifications bool   `json:"SERVICE_WORKER_PROMPT_NOTIFICATIONS"`
	SboxLabels                       struct {
		SuggestionDismissLabel   string `json:"SUGGESTION_DISMISS_LABEL"`
		SuggestionDismissedLabel string `json:"SUGGESTION_DISMISSED_LABEL"`
	} `json:"SBOX_LABELS"`
	OnePickUrl          string   `json:"ONE_PICK_URL"`
	NoEmptyDataImg      bool     `json:"NO_EMPTY_DATA_IMG"`
	MentionsEduHelpLink string   `json:"MENTIONS_EDU_HELP_LINK"`
	IsHomePageCold      bool     `json:"IS_HOMEPAGE_COLD"`
	DeferredDetach      bool     `json:"DEFERRED_DETACH"`
	ZwieBackPingUrls    []string `json:"ZWIEBACK_PING_URLS"`
	VozAPIKey           string   `json:"VOZ_API_KEY"`
	STS                 int      `json:"STS"`
	SBoxSettings        struct {
		HasOnScreenKeyBoard                         bool   `json:"HAS_ON_SCREEN_KEYBOARD"`
		IsFusion                                    bool   `json:"IS_FUSION"`
		IsPolymer                                   bool   `json:"IS_POLYMER"`
		RequestDomain                               string `json:"REQUEST_DOMAIN"`
		RequestLanguage                             string `json:"REQUEST_LANGUAGE"`
		SendVisitorData                             bool   `json:"SEND_VISITOR_DATA"`
		SearchBoxBehaviorExperiment                 string `json:"SEARCHBOX_BEHAVIOR_EXPERIMENT"`
		SearchBoxEnableRefineMentSuggest            bool   `json:"SEARCHBOX_ENABLE_REFINEMENT_SUGGEST"`
		SearchBoxTapTargetExperiment                int    `json:"SEARCHBOX_TAP_TARGET_EXPERIMENT"`
		SearchBoxZeroTypingSuggestUseRegularSuggest string `json:"SEARCHBOX_ZERO_TYPING_SUGGEST_USE_REGULAR_SUGGEST"`
		SuggExpID                                   string `json:"SUGG_EXP_ID"`
		VisitorData                                 string `json:"VISITOR_DATA"`
		SearchBoxHostOverride                       string `json:"SEARCHBOX_HOST_OVERRIDE"`
		HideRemoveLink                              bool   `json:"HIDE_REMOVE_LINK"`
	} `json:"SBOX_SETTINGS"`
	SBoxJsUrl          string `json:"SBOX_JS_URL"`
	RecaptchaV3Sitekey string `json:"RECAPTCHA_V3_SITEKEY"`
	PlayerJsUrl        string `json:"PLAYER_JS_URL"`
	PlayerCssUrl       string `json:"PLAYER_CSS_URL"`
	LinkGalDomain      string `json:"LINK_GAL_DOMAIN"`
	LinkOISDomain      string `json:"LINK_OIS_DOMAIN"`
	IsTablet           bool   `json:"IS_TABLET"`
	LinkApiKey         string `json:"LINK_API_KEY"`
	DisableWarmLoads   bool   `json:"DISABLE_WARM_LOADS"`
}

type musicConfigSet struct {
	ClientCanaryState     string `json:"CLIENT_CANARY_STATE"`
	Device                string `json:"DEVICE"`
	ElementPoolDefaultCap int    `json:"ELEMENT_POOL_DEFAULT_CAP"`
	EventID               string `json:"EVENT_ID"`
	ExperimentFlags       struct {
		AllowSkipNetworkless                                    bool          `json:"allow_skip_networkless"`
		ClearUserPartitionedLs                                  bool          `json:"clear_user_partitioned_ls"`
		CsiOnGel                                                bool          `json:"csi_on_gel"`
		DesktopTextAdsGrayVisurl                                bool          `json:"desktop_text_ads_gray_visurl"`
		DisableChildNodeAutoFormattedStrings                    bool          `json:"disable_child_node_auto_formatted_strings"`
		DisableSimpleMixedDirectionFormattedStrings             bool          `json:"disable_simple_mixed_direction_formatted_strings"`
		DisableThumbnailPreloading                              bool          `json:"disable_thumbnail_preloading"`
		EnableCastOnMusicWeb                                    bool          `json:"enable_cast_on_music_web"`
		EnableClientSliLogging                                  bool          `json:"enable_client_sli_logging"`
		EnableClientStreamzWeb                                  bool          `json:"enable_client_streamz_web"`
		EnableGetAccountSwitcherEndpointOnWebfe                 bool          `json:"enable_get_account_switcher_endpoint_on_webfe"`
		EnableGrayVisurl                                        bool          `json:"enable_gray_visurl"`
		EnableMembershipsAndPurchases                           bool          `json:"enable_memberships_and_purchases"`
		EnableMixedDirectionFormattedStrings                    bool          `json:"enable_mixed_direction_formatted_strings"`
		EnableModularPlayerPageServer                           bool          `json:"enable_modular_player_page_server"`
		EnableNamesHandlesAccountSwitcher                       bool          `json:"enable_names_handles_account_switcher"`
		EnablePolymerResin                                      bool          `json:"enable_polymer_resin"`
		EnablePremiumVoluntaryPause                             bool          `json:"enable_premium_voluntary_pause"`
		EnableSharePanelPageAsScreenLayer                       bool          `json:"enable_share_panel_page_as_screen_layer"`
		EnableSliFlush                                          bool          `json:"enable_sli_flush"`
		EnableWebMediaSessionMetadataFix                        bool          `json:"enable_web_media_session_metadata_fix"`
		EnableWebUploadSupport                                  bool          `json:"enable_web_upload_support"`
		EnableYpcSpinners                                       bool          `json:"enable_ypc_spinners"`
		ExportNetworklessOptions                                bool          `json:"export_networkless_options"`
		ForwardDomainAdminStateOnEmbeds                         bool          `json:"forward_domain_admin_state_on_embeds"`
		HandleMusicResolveUrlRedirectResponse                   bool          `json:"handle_music_resolve_url_redirect_response"`
		Html5EnableSingleVideoVodIvarOnPacf                     bool          `json:"html5_enable_single_video_vod_ivar_on_pacf"`
		Html5EnableVideoOverlayOnInplayerSlotForTv              bool          `json:"html5_enable_video_overlay_on_inplayer_slot_for_tv"`
		Html5PacfEnableDai                                      bool          `json:"html5_pacf_enable_dai"`
		KevlarAttachVimioBehavior                               bool          `json:"kevlar_attach_vimio_behavior"`
		KevlarDropdownFix                                       bool          `json:"kevlar_dropdown_fix"`
		KevlarEnableVimioCallback                               bool          `json:"kevlar_enable_vimio_callback"`
		KevlarEnableVimioLogging                                bool          `json:"kevlar_enable_vimio_logging"`
		KevlarGelErrorRouting                                   bool          `json:"kevlar_gel_error_routing"`
		KevlarImportVimioBehavior                               bool          `json:"kevlar_import_vimio_behavior"`
		KevlarUseVimioBehavior                                  bool          `json:"kevlar_use_vimio_behavior"`
		LogHeartbeatWithLifecycles                              bool          `json:"log_heartbeat_with_lifecycles"`
		LogWebEndpointToLayer                                   bool          `json:"log_web_endpoint_to_layer"`
		MusicEnableExploreTabOnWeb                              bool          `json:"music_enable_explore_tab_on_web"`
		MusicEnableImproveYourRecommendationsSetting            bool          `json:"music_enable_improve_your_recommendations_setting"`
		MusicWebCanary                                          bool          `json:"music_web_canary"`
		MusicWebEnableAddToPlaylistV2CompleteUi                 bool          `json:"music_web_enable_add_to_playlist_v2_complete_ui"`
		MusicWebEnableClientSidePlaybackScreens                 bool          `json:"music_web_enable_client_side_playback_screens"`
		MusicWebEnableDragDropUpload                            bool          `json:"music_web_enable_drag_drop_upload"`
		MusicWebEnablePlayerBarVeLoggingFixes                   bool          `json:"music_web_enable_player_bar_ve_logging_fixes"`
		MusicWebEnableServiceWorker                             bool          `json:"music_web_enable_service_worker"`
		MusicWebEnableStateSubscriptionStatus                   bool          `json:"music_web_enable_state_subscription_status"`
		MusicWebEnableWugApplicationSettingsEndpoint            bool          `json:"music_web_enable_wug_application_settings_endpoint"`
		MusicWebEnableWugBrowseEndpoint                         bool          `json:"music_web_enable_wug_browse_endpoint"`
		MusicWebEnableWugChannelCreationFormEndpoint            bool          `json:"music_web_enable_wug_channel_creation_form_endpoint"`
		MusicWebEnableWugFeedbackEndpoint                       bool          `json:"music_web_enable_wug_feedback_endpoint"`
		MusicWebEnableWugFetchDataBrowse                        bool          `json:"music_web_enable_wug_fetch_data_browse"`
		MusicWebEnableWugFetchGetSearchSuggestions              bool          `json:"music_web_enable_wug_fetch_get_search_suggestions"`
		MusicWebEnableWugFetchNext                              bool          `json:"music_web_enable_wug_fetch_next"`
		MusicWebEnableWugGetAccountMenuEndpoint                 bool          `json:"music_web_enable_wug_get_account_menu_endpoint"`
		MusicWebEnableWugMusicDeletePrivatelyOwnedEntityCommand bool          `json:"music_web_enable_wug_music_delete_privately_owned_entity_command"`
		MusicWebEnableWugPlaylistEditEndpoint                   bool          `json:"music_web_enable_wug_playlist_edit_endpoint"`
		MusicWebEnableWugQueueAddEndpoint                       bool          `json:"music_web_enable_wug_queue_add_endpoint"`
		MusicWebEnableWugSearchEndpoint                         bool          `json:"music_web_enable_wug_search_endpoint"`
		MusicWebEnableWugSetSettingEndpoint                     bool          `json:"music_web_enable_wug_set_setting_endpoint"`
		MusicWebEnableWugSubscribeEndpoint                      bool          `json:"music_web_enable_wug_subscribe_endpoint"`
		MusicWebEnableWugYpcEndpoints                           bool          `json:"music_web_enable_wug_ypc_endpoints"`
		MusicWebEnableWugYtNextContinuationBehavior             bool          `json:"music_web_enable_wug_yt_next_continuation_behavior"`
		MusicWebIsCanary                                        bool          `json:"music_web_is_canary"`
		MusicWebMarkRootVisible                                 bool          `json:"music_web_mark_root_visible"`
		MusicWebPingAdStart                                     bool          `json:"music_web_ping_ad_start"`
		MusicWebPlayerContextConfig                             bool          `json:"music_web_player_context_config"`
		NetworklessLogging                                      bool          `json:"networkless_logging"`
		NwlThrottlingRaceFix                                    bool          `json:"nwl_throttling_race_fix"`
		NwlTriggerThrottleAfterReset                            bool          `json:"nwl_trigger_throttle_after_reset"`
		OmitInnertubeApiKeyForBearerAuthHeader                  bool          `json:"omit_innertube_api_key_for_bearer_auth_header"`
		PageidAsHeaderWeb                                       bool          `json:"pageid_as_header_web"`
		PesMigrateAssociationData                               bool          `json:"pes_migrate_association_data"`
		PolymerBadBuildLabels                                   bool          `json:"polymer_bad_build_labels"`
		PolymerVerifiyAppState                                  bool          `json:"polymer_verifiy_app_state"`
		QoeSendAndWrite                                         bool          `json:"qoe_send_and_write"`
		ScreenManagerLogServletEi                               bool          `json:"screen_manager_log_servlet_ei"`
		SkipLsGelRetry                                          bool          `json:"skip_ls_gel_retry"`
		StateMachineDynamicEventsLifecycles                     bool          `json:"state_machine_dynamic_events_lifecycles"`
		SuppressError204Logging                                 bool          `json:"suppress_error_204_logging"`
		UseDocumentLifecycles                                   bool          `json:"use_document_lifecycles"`
		UseScreenManagerUtil                                    bool          `json:"use_screen_manager_util"`
		UseUndefinedCsnAnyLayer                                 bool          `json:"use_undefined_csn_any_layer"`
		VssFinalPingSendAndWrite                                bool          `json:"vss_final_ping_send_and_write"`
		VssPlaybackUseSendAndWrite                              bool          `json:"vss_playback_use_send_and_write"`
		WebApiUrl                                               bool          `json:"web_api_url"`
		WebDedupeVeGrafting                                     bool          `json:"web_dedupe_ve_grafting"`
		WebEnableAdSignalsInItContext                           bool          `json:"web_enable_ad_signals_in_it_context"`
		WebGelTimeoutCap                                        bool          `json:"web_gel_timeout_cap"`
		WebPlaybackAssociatedVe                                 bool          `json:"web_playback_associated_ve"`
		WebRemixNwl                                             bool          `json:"web_remix_nwl"`
		YtImgShadowTriggerShowOnVisible                         bool          `json:"yt_img_shadow_trigger_show_on_visible"`
		YtidbIsSupportedCacheSuccessResult                      bool          `json:"ytidb_is_supported_cache_success_result"`
		YtidbStopTransactionCommit                              bool          `json:"ytidb_stop_transaction_commit"`
		AddtoAjaxLogWarningFraction                             float64       `json:"addto_ajax_log_warning_fraction"`
		LogWindowOnerrorFraction                                float64       `json:"log_window_onerror_fraction"`
		YtidbTransactionEndedEventRateLimit                     float64       `json:"ytidb_transaction_ended_event_rate_limit"`
		BotguardAsyncSnapshotTimeoutMs                          int           `json:"botguard_async_snapshot_timeout_ms"`
		CheckNavigatorAccuracyTimeoutMs                         int           `json:"check_navigator_accuracy_timeout_ms"`
		ClientStreamzWebFlushCount                              int           `json:"client_streamz_web_flush_count"`
		ClientStreamzWebFlushIntervalSeconds                    int           `json:"client_streamz_web_flush_interval_seconds"`
		InitialGelBatchTimeout                                  int           `json:"initial_gel_batch_timeout"`
		LeaderElectionCheckInterval                             int           `json:"leader_election_check_interval"`
		LeaderElectionLeaseTtl                                  int           `json:"leader_election_lease_ttl"`
		LeaderElectionRenewalInterval                           int           `json:"leader_election_renewal_interval"`
		MusicWebCanaryStage                                     int           `json:"music_web_canary_stage"`
		MusicWebDelayWatchNextMs                                int           `json:"music_web_delay_watch_next_ms"`
		MusicWebListContinuationPrescanHeightPx                 int           `json:"music_web_list_continuation_prescan_height_px"`
		MusicWebSessionCheckIntervalMillis                      int           `json:"music_web_session_check_interval_millis"`
		NetworkPollingInterval                                  int           `json:"network_polling_interval"`
		WebForegroundHeartbeatIntervalMs                        int           `json:"web_foreground_heartbeat_interval_ms"`
		WebGelDebounceMs                                        int           `json:"web_gel_debounce_ms"`
		WebLoggingMaxBatch                                      int           `json:"web_logging_max_batch"`
		YtidbRemakeDbRetries                                    int           `json:"ytidb_remake_db_retries"`
		CbV2Uxe                                                 string        `json:"cb_v2_uxe"`
		DebugForcedPromoId                                      string        `json:"debug_forced_promo_id"`
		MusicWebBodyLineHeight                                  string        `json:"music_web_body_line_height"`
		MusicWebTitleLineHeight                                 string        `json:"music_web_title_line_height"`
		UserPreferenceCollectionInitialBrowseId                 string        `json:"user_preference_collection_initial_browse_id"`
		WebClientVersionOverride                                string        `json:"web_client_version_override"`
		KevlarCommandHandlerCommandBanlist                      []interface{} `json:"kevlar_command_handler_command_banlist"`
		WebOpContinuationTypeBanlist                            []interface{} `json:"web_op_continuation_type_banlist"`
		WebOpEndpointBanlist                                    []interface{} `json:"web_op_endpoint_banlist"`
		WebOpSignalTypeBanlist                                  []interface{} `json:"web_op_signal_type_banlist"`
	} `json:"EXPERIMENT_FLAGS"`
	GAPIHintParams         string `json:"GAPI_HINT_PARAMS"`
	GAPIHost               string `json:"GAPI_HOST"`
	GAPILocale             string `json:"GAPI_LOCALE"`
	GL                     string `json:"GL"`
	HL                     string `json:"HL"`
	HTMLDir                string `json:"HTML_DIR"`
	HTMLKang               string `json:"HTML_LANG"`
	InnerTubeApiKey        string `json:"INNERTUBE_API_KEY"`
	InnerTubeApiVersion    string `json:"INNERTUBE_API_VERSION"`
	InnerTubeClientName    string `json:"INNERTUBE_CLIENT_NAME"`
	InnerTubeClientVersion string `json:"INNERTUBE_CLIENT_VERSION"`
	InnerTubeContext       struct {
		Client struct {
			Hl               string `json:"hl"`
			Gl               string `json:"gl"`
			RemoteHost       string `json:"remoteHost"`
			DeviceMake       string `json:"deviceMake"`
			DeviceModel      string `json:"deviceModel"`
			VisitorData      string `json:"visitorData"`
			UserAgent        string `json:"userAgent"`
			ClientName       string `json:"clientName"`
			ClientVersion    string `json:"clientVersion"`
			OsName           string `json:"osName"`
			OsVersion        string `json:"osVersion"`
			OriginalUrl      string `json:"originalUrl"`
			Platform         string `json:"platform"`
			ClientFormFactor string `json:"clientFormFactor"`
			ConfigInfo       struct {
				AppInstallData string `json:"appInstallData"`
			} `json:"configInfo"`
			BrowserName    string `json:"browserName"`
			BrowserVersion string `json:"browserVersion"`
		} `json:"client"`
		User struct {
			LockedSafetyMode bool `json:"lockedSafetyMode"`
		} `json:"user"`
		Request struct {
			UseSsl bool `json:"useSsl"`
		} `json:"request"`
		ClickTracking struct {
			ClickTrackingParams string `json:"clickTrackingParams"`
		} `json:"clickTracking"`
	} `json:"INNERTUBE_CONTEXT"`
	InnerTubeContextClientName          int    `json:"INNERTUBE_CONTEXT_CLIENT_NAME"`
	InnerTubeContextClientVersion       string `json:"INNERTUBE_CONTEXT_CLIENT_VERSION"`
	InnerTubeContextGL                  string `json:"INNERTUBE_CONTEXT_GL"`
	InnerTubeContextHL                  string `json:"INNERTUBE_CONTEXT_HL"`
	LatestECatcherServiceTrackingParams struct {
		ClientName string `json:"client.name"`
	} `json:"LATEST_ECATCHER_SERVICE_TRACKING_PARAMS"`
	LoggedIn                bool   `json:"LOGGED_IN"`
	PageBuildLabel          string `json:"PAGE_BUILD_LABEL"`
	PageCL                  int    `json:"PAGE_CL"`
	ServerName              string `json:"SERVER_NAME"`
	SigninUrl               string `json:"SIGNIN_URL"`
	VisitorData             string `json:"VISITOR_DATA"`
	WebPlayerContextConfigs struct {
		WebPlayerContextConfigIDMusicWatch struct {
			RootElementId                 string `json:"rootElementId"`
			JsUrl                         string `json:"jsUrl"`
			CssUrl                        string `json:"cssUrl"`
			ContextId                     string `json:"contextId"`
			EventLabel                    string `json:"eventLabel"`
			ContentRegion                 string `json:"contentRegion"`
			Hl                            string `json:"hl"`
			HostLanguage                  string `json:"hostLanguage"`
			InnertubeApiKey               string `json:"innertubeApiKey"`
			InnertubeApiVersion           string `json:"innertubeApiVersion"`
			InnertubeContextClientVersion string `json:"innertubeContextClientVersion"`
			ControlsType                  int    `json:"controlsType"`
			DisableKeyboardControls       bool   `json:"disableKeyboardControls"`
			DisableRelatedVideos          bool   `json:"disableRelatedVideos"`
			AnnotationsLoadPolicy         int    `json:"annotationsLoadPolicy"`
			Device                        struct {
				Brand            string `json:"brand"`
				Model            string `json:"model"`
				Browser          string `json:"browser"`
				BrowserVersion   string `json:"browserVersion"`
				Os               string `json:"os"`
				OsVersion        string `json:"osVersion"`
				Platform         string `json:"platform"`
				InterfaceName    string `json:"interfaceName"`
				InterfaceVersion string `json:"interfaceVersion"`
			} `json:"device"`
			SerializedExperimentIds            string `json:"serializedExperimentIds"`
			SerializedExperimentFlags          string `json:"serializedExperimentFlags"`
			DisableSharing                     bool   `json:"disableSharing"`
			HideInfo                           bool   `json:"hideInfo"`
			DisableWatchLater                  bool   `json:"disableWatchLater"`
			MobileIphoneSupportsInlinePlayback bool   `json:"mobileIphoneSupportsInlinePlayback"`
			IsMobileDevice                     bool   `json:"isMobileDevice"`
			CspNonce                           string `json:"cspNonce"`
			CanaryState                        string `json:"canaryState"`
			EnableCsiLogging                   bool   `json:"enableCsiLogging"`
			DatasyncId                         string `json:"datasyncId"`
		} `json:"WEB_PLAYER_CONTEXT_CONFIG_ID_MUSIC_WATCH"`
	} `json:"WEB_PLAYER_CONTEXT_CONFIGS"`
	XSRFFieldName              string `json:"XSRF_FIELD_NAME"`
	XSRFToken                  string `json:"XSRF_TOKEN"`
	YPCMBUrl                   string `json:"YPC_MB_URL"`
	YTRFamilyCreationUrl       string `json:"YTR_FAMILY_CREATION_URL"`
	ServerVersion              string `json:"SERVER_VERSION"`
	Locale                     string `json:"LOCALE"`
	ReuseComponents            bool   `json:"REUSE_COMPONENTS"`
	StamperStableList          bool   `json:"STAMPER_STABLE_LIST"`
	DataSyncID                 string `json:"DATASYNC_ID"`
	SerializedClientConfigData string `json:"SERIALIZED_CLIENT_CONFIG_DATA"`
	MxdConfig                  struct {
		Device                        string   `json:"device"`
		App                           string   `json:"app"`
		AppId                         string   `json:"appId"`
		DisableDial                   bool     `json:"disableDial"`
		Theme                         string   `json:"theme"`
		LoadCastApiSetupScript        bool     `json:"loadCastApiSetupScript"`
		Capabilities                  []string `json:"capabilities"`
		DisableAutomaticScreenCache   bool     `json:"disableAutomaticScreenCache"`
		ForceMirroring                bool     `json:"forceMirroring"`
		EnableConnectWithInitialState bool     `json:"enableConnectWithInitialState"`
	} `json:"MDX_CONFIG"`
	ClientProtocol           string `json:"CLIENT_PROTOCOL"`
	ClientTransport          string `json:"CLIENT_TRANSPORT"`
	UseEmbeddedInnertubeData bool   `json:"USE_EMBEDDED_INNERTUBE_DATA"`
	VisibilityRoot           string `json:"VISIBILITY_ROOT"`
	YTMusicIconSrc           string `json:"YTMUSIC_ICON_SRC"`
	YTMusicLogoSrc           string `json:"YTMUSIC_LOGO_SRC"`
	UploadUrl                string `json:"UPLOAD_URL"`
	TransferPageSigninUrl    string `json:"TRANSFER_PAGE_SIGNIN_URL"`
	LogoutUrl                string `json:"LOGOUT_URL"`
	IsSubscriber             bool   `json:"IS_SUBSCRIBER"`
	IsMobileWeb              bool   `json:"IS_MOBILE_WEB"`
	InitialEndpoint          string `json:"INITIAL_ENDPOINT"`
	HotKeyDialog             struct {
		Title struct {
			Runs []struct {
				Text string `json:"text"`
			} `json:"runs"`
		} `json:"title"`
		Sections []struct {
			HotkeyDialogSectionRenderer struct {
				Title struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
				} `json:"title"`
				Options []struct {
					HotkeyDialogSectionOptionRenderer struct {
						Label struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"label"`
						Hotkey string `json:"hotkey"`
					} `json:"hotkeyDialogSectionOptionRenderer"`
				} `json:"options"`
			} `json:"hotkeyDialogSectionRenderer"`
		} `json:"sections"`
	} `json:"HOTKEY_DIALOG"`
	DefaultAlbumImageSrc          string `json:"DEFAULT_ALBUM_IMAGE_SRC"`
	AudioQuality                  string `json:"AUDIO_QUALITY"`
	AddScraperAttributes          bool   `json:"ADD_SCRAPER_ATTRIBUTES"`
	ActiveAccountIsMadisonAccount bool   `json:"ACTIVE_ACCOUNT_IS_MADISON_ACCOUNT"`
	YTMusicWhiteIconSrc           string `json:"YTMUSIC_WHITE_ICON_SRC"`
	YTmusicWhiteLogoSrc           string `json:"YTMUSIC_WHITE_LOGO_SRC"`
	ActiveAccountCanUpload        bool   `json:"ACTIVE_ACCOUNT_CAN_UPLOAD"`
}

type playerQueryResponse struct {
	ResponseContext struct {
		VisitorData           string `json:"visitorData"`
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
	} `json:"responseContext"`
	TrackingParams    string `json:"trackingParams"`
	PlayabilityStatus struct {
		Status               string `json:"status"`
		PlayableInEmbed      bool   `json:"playableInEmbed"`
		AudioOnlyPlayability struct {
			AudioOnlyPlayabilityRenderer struct {
				TrackingParams        string `json:"trackingParams"`
				AudioOnlyAvailability string `json:"audioOnlyAvailability"`
			} `json:"audioOnlyPlayabilityRenderer"`
		} `json:"audioOnlyPlayability"`
		Miniplayer struct {
			MiniplayerRenderer struct {
				PlaybackMode string `json:"playbackMode"`
			} `json:"miniplayerRenderer"`
		} `json:"miniplayer"`
		ContextParams string `json:"contextParams"`
	} `json:"playabilityStatus"`
	StreamingData streamingData `json:"streamingData"`
	PlayerAds     []struct {
		PlayerLegacyDesktopWatchAdsRenderer struct {
			PlayerAdParams struct {
				ShowContentThumbnail bool   `json:"showContentThumbnail"`
				EnabledEngageTypes   string `json:"enabledEngageTypes"`
			} `json:"playerAdParams"`
			GutParams struct {
				Tag string `json:"tag"`
			} `json:"gutParams"`
			ShowCompanion bool `json:"showCompanion"`
			ShowInstream  bool `json:"showInstream"`
			UseGut        bool `json:"useGut"`
		} `json:"playerLegacyDesktopWatchAdsRenderer"`
	} `json:"playerAds"`
	PlaybackTracking struct {
		VideostatsPlaybackUrl struct {
			BaseUrl string `json:"baseUrl"`
			Headers []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"videostatsPlaybackUrl"`
		VideostatsDelayplayUrl struct {
			BaseUrl string `json:"baseUrl"`
			Headers []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"videostatsDelayplayUrl"`
		VideostatsWatchtimeUrl struct {
			BaseUrl string `json:"baseUrl"`
			Headers []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"videostatsWatchtimeUrl"`
		PtrackingUrl struct {
			BaseUrl string `json:"baseUrl"`
			Headers []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"ptrackingUrl"`
		QoeUrl struct {
			BaseUrl string `json:"baseUrl"`
			Headers []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"qoeUrl"`
		AtrUrl struct {
			BaseUrl                 string `json:"baseUrl"`
			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
			Headers                 []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"atrUrl"`
		VideostatsScheduledFlushWalltimeSeconds []int `json:"videostatsScheduledFlushWalltimeSeconds"`
		VideostatsDefaultFlushIntervalSeconds   int   `json:"videostatsDefaultFlushIntervalSeconds"`
	} `json:"playbackTracking"`
	VideoDetails struct {
		VideoId        string `json:"videoId"`
		Title          string `json:"title"`
		LengthSeconds  string `json:"lengthSeconds"`
		ChannelId      string `json:"channelId"`
		IsOwnerViewing bool   `json:"isOwnerViewing"`
		IsCrawlable    bool   `json:"isCrawlable"`
		Thumbnail      struct {
			Thumbnails []struct {
				Url    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"thumbnails"`
		} `json:"thumbnail"`
		AverageRating     float64 `json:"averageRating"`
		AllowRatings      bool    `json:"allowRatings"`
		ViewCount         string  `json:"viewCount"`
		Author            string  `json:"author"`
		IsPrivate         bool    `json:"isPrivate"`
		IsUnpluggedCorpus bool    `json:"isUnpluggedCorpus"`
		MusicVideoType    string  `json:"musicVideoType"`
		IsLiveContent     bool    `json:"isLiveContent"`
	} `json:"videoDetails"`
	PlayerConfig struct {
		AudioConfig struct {
			LoudnessDb              float64 `json:"loudnessDb"`
			PerceptualLoudnessDb    float64 `json:"perceptualLoudnessDb"`
			EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
		} `json:"audioConfig"`
		StreamSelectionConfig struct {
			MaxBitrate string `json:"maxBitrate"`
		} `json:"streamSelectionConfig"`
		MediaCommonConfig struct {
			DynamicReadaheadConfig struct {
				MaxReadAheadMediaTimeMs int `json:"maxReadAheadMediaTimeMs"`
				MinReadAheadMediaTimeMs int `json:"minReadAheadMediaTimeMs"`
				ReadAheadGrowthRateMs   int `json:"readAheadGrowthRateMs"`
			} `json:"dynamicReadaheadConfig"`
		} `json:"mediaCommonConfig"`
		WebPlayerConfig struct {
			WebPlayerActionsPorting struct {
				SubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					SubscribeEndpoint   struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"subscribeEndpoint"`
				} `json:"subscribeCommand"`
				UnsubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					UnsubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"unsubscribeEndpoint"`
				} `json:"unsubscribeCommand"`
				AddToWatchLaterCommand struct {
					ClickTrackingParams  string `json:"clickTrackingParams"`
					PlaylistEditEndpoint struct {
						PlaylistId string `json:"playlistId"`
						Actions    []struct {
							AddedVideoId string `json:"addedVideoId"`
							Action       string `json:"action"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"addToWatchLaterCommand"`
				RemoveFromWatchLaterCommand struct {
					ClickTrackingParams  string `json:"clickTrackingParams"`
					PlaylistEditEndpoint struct {
						PlaylistId string `json:"playlistId"`
						Actions    []struct {
							Action         string `json:"action"`
							RemovedVideoId string `json:"removedVideoId"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"removeFromWatchLaterCommand"`
			} `json:"webPlayerActionsPorting"`
		} `json:"webPlayerConfig"`
	} `json:"playerConfig"`
	Storyboards struct {
		PlayerStoryboardSpecRenderer struct {
			Spec string `json:"spec"`
		} `json:"playerStoryboardSpecRenderer"`
	} `json:"storyboards"`
	Microformat struct {
		MicroformatDataRenderer struct {
			UrlCanonical string `json:"urlCanonical"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			Thumbnail    struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			SiteName           string   `json:"siteName"`
			AppName            string   `json:"appName"`
			AndroidPackage     string   `json:"androidPackage"`
			IosAppStoreId      string   `json:"iosAppStoreId"`
			IosAppArguments    string   `json:"iosAppArguments"`
			OgType             string   `json:"ogType"`
			UrlApplinksIos     string   `json:"urlApplinksIos"`
			UrlApplinksAndroid string   `json:"urlApplinksAndroid"`
			UrlTwitterIos      string   `json:"urlTwitterIos"`
			UrlTwitterAndroid  string   `json:"urlTwitterAndroid"`
			TwitterCardType    string   `json:"twitterCardType"`
			TwitterSiteHandle  string   `json:"twitterSiteHandle"`
			SchemaDotOrgType   string   `json:"schemaDotOrgType"`
			Noindex            bool     `json:"noindex"`
			Unlisted           bool     `json:"unlisted"`
			Paid               bool     `json:"paid"`
			FamilySafe         bool     `json:"familySafe"`
			Tags               []string `json:"tags"`
			AvailableCountries []string `json:"availableCountries"`
			PageOwnerDetails   struct {
				Name              string `json:"name"`
				ExternalChannelId string `json:"externalChannelId"`
				YoutubeProfileUrl string `json:"youtubeProfileUrl"`
			} `json:"pageOwnerDetails"`
			VideoDetails struct {
				ExternalVideoId string `json:"externalVideoId"`
				DurationSeconds string `json:"durationSeconds"`
				DurationIso8601 string `json:"durationIso8601"`
			} `json:"videoDetails"`
			LinkAlternates []struct {
				HrefUrl       string `json:"hrefUrl"`
				Title         string `json:"title,omitempty"`
				AlternateType string `json:"alternateType,omitempty"`
			} `json:"linkAlternates"`
			ViewCount   string `json:"viewCount"`
			PublishDate string `json:"publishDate"`
			Category    string `json:"category"`
			UploadDate  string `json:"uploadDate"`
		} `json:"microformatDataRenderer"`
	} `json:"microformat"`
	Attestation struct {
		PlayerAttestationRenderer struct {
			Challenge    string `json:"challenge"`
			BotguardData struct {
				Program            string `json:"program"`
				InterpreterSafeUrl struct {
					PrivateDoNotAccessOrElseTrustedResourceUrlWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
				} `json:"interpreterSafeUrl"`
				ServerEnvironment int `json:"serverEnvironment"`
			} `json:"botguardData"`
		} `json:"playerAttestationRenderer"`
	} `json:"attestation"`
	Messages []struct {
		MealbarPromoRenderer struct {
			Icon struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"icon"`
			MessageTexts []struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"messageTexts"`
			ActionButton struct {
				ButtonRenderer struct {
					Size string `json:"size"`
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					ServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						FeedbackEndpoint    struct {
							FeedbackToken string `json:"feedbackToken"`
							UiActions     struct {
								HideEnclosingContainer bool `json:"hideEnclosingContainer"`
							} `json:"uiActions"`
						} `json:"feedbackEndpoint"`
					} `json:"serviceEndpoint"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						BrowseEndpoint      struct {
							BrowseId string `json:"browseId"`
							Params   string `json:"params"`
						} `json:"browseEndpoint"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"buttonRenderer"`
			} `json:"actionButton"`
			DismissButton struct {
				ButtonRenderer struct {
					Size string `json:"size"`
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					ServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						FeedbackEndpoint    struct {
							FeedbackToken string `json:"feedbackToken"`
							UiActions     struct {
								HideEnclosingContainer bool `json:"hideEnclosingContainer"`
							} `json:"uiActions"`
						} `json:"feedbackEndpoint"`
					} `json:"serviceEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"buttonRenderer"`
			} `json:"dismissButton"`
			TriggerCondition    string `json:"triggerCondition"`
			Style               string `json:"style"`
			TrackingParams      string `json:"trackingParams"`
			ImpressionEndpoints []struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				FeedbackEndpoint    struct {
					FeedbackToken string `json:"feedbackToken"`
					UiActions     struct {
						HideEnclosingContainer bool `json:"hideEnclosingContainer"`
					} `json:"uiActions"`
				} `json:"feedbackEndpoint"`
			} `json:"impressionEndpoints"`
			IsVisible    bool `json:"isVisible"`
			MessageTitle struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"messageTitle"`
		} `json:"mealbarPromoRenderer"`
	} `json:"messages"`
	AdPlacements []struct {
		AdPlacementRenderer struct {
			Config struct {
				AdPlacementConfig struct {
					Kind         string `json:"kind"`
					AdTimeOffset struct {
						OffsetStartMilliseconds string `json:"offsetStartMilliseconds"`
						OffsetEndMilliseconds   string `json:"offsetEndMilliseconds"`
					} `json:"adTimeOffset"`
					HideCueRangeMarker bool `json:"hideCueRangeMarker"`
				} `json:"adPlacementConfig"`
			} `json:"config"`
			Renderer struct {
				InstreamVideoAdRenderer struct {
					SkipOffsetMilliseconds int `json:"skipOffsetMilliseconds"`
					Pings                  struct {
						ImpressionPings []struct {
							BaseUrl string `json:"baseUrl"`
							Headers []struct {
								HeaderType string `json:"headerType"`
							} `json:"headers,omitempty"`
						} `json:"impressionPings"`
						ErrorPings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"errorPings"`
						MutePings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"mutePings"`
						UnmutePings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"unmutePings"`
						PausePings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"pausePings"`
						RewindPings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"rewindPings"`
						ResumePings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"resumePings"`
						SkipPings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"skipPings"`
						ClosePings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"closePings"`
						ProgressPings []struct {
							BaseUrl            string `json:"baseUrl"`
							OffsetMilliseconds int    `json:"offsetMilliseconds"`
						} `json:"progressPings"`
						FullscreenPings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"fullscreenPings"`
						ActiveViewViewablePings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"activeViewViewablePings"`
						EndFullscreenPings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"endFullscreenPings"`
						ActiveViewMeasurablePings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"activeViewMeasurablePings"`
						AbandonPings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"abandonPings"`
						ActiveViewFullyViewableAudibleHalfDurationPings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"activeViewFullyViewableAudibleHalfDurationPings"`
						CompletePings []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"completePings"`
					} `json:"pings"`
					ClickthroughEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						LoggingUrls         []struct {
							BaseUrl string `json:"baseUrl"`
						} `json:"loggingUrls"`
						UrlEndpoint struct {
							Url    string `json:"url"`
							Target string `json:"target"`
						} `json:"urlEndpoint"`
					} `json:"clickthroughEndpoint"`
					CsiParameters []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"csiParameters"`
					PlayerVars    string `json:"playerVars"`
					PlayerOverlay struct {
						InstreamAdPlayerOverlayRenderer struct {
							SkipOrPreviewRenderer struct {
								SkipAdRenderer struct {
									PreskipRenderer struct {
										AdPreviewRenderer struct {
											Thumbnail struct {
												Thumbnail struct {
													Thumbnails []struct {
														Url    string `json:"url"`
														Width  int    `json:"width"`
														Height int    `json:"height"`
													} `json:"thumbnails"`
												} `json:"thumbnail"`
												TrackingParams string `json:"trackingParams"`
											} `json:"thumbnail"`
											TrackingParams     string `json:"trackingParams"`
											TemplatedCountdown struct {
												TemplatedAdText struct {
													Text           string `json:"text"`
													IsTemplated    bool   `json:"isTemplated"`
													TrackingParams string `json:"trackingParams"`
												} `json:"templatedAdText"`
											} `json:"templatedCountdown"`
											DurationMilliseconds int `json:"durationMilliseconds"`
										} `json:"adPreviewRenderer"`
									} `json:"preskipRenderer"`
									SkippableRenderer struct {
										SkipButtonRenderer struct {
											Message struct {
												Text           string `json:"text"`
												IsTemplated    bool   `json:"isTemplated"`
												TrackingParams string `json:"trackingParams"`
											} `json:"message"`
											TrackingParams string `json:"trackingParams"`
										} `json:"skipButtonRenderer"`
									} `json:"skippableRenderer"`
									TrackingParams         string `json:"trackingParams"`
									SkipOffsetMilliseconds int    `json:"skipOffsetMilliseconds"`
								} `json:"skipAdRenderer"`
							} `json:"skipOrPreviewRenderer"`
							TrackingParams          string `json:"trackingParams"`
							VisitAdvertiserRenderer struct {
								ButtonRenderer struct {
									Style string `json:"style"`
									Text  struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"text"`
									Icon struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										LoggingUrls         []struct {
											BaseUrl string `json:"baseUrl"`
										} `json:"loggingUrls"`
										UrlEndpoint struct {
											Url    string `json:"url"`
											Target string `json:"target"`
										} `json:"urlEndpoint"`
									} `json:"navigationEndpoint"`
									TrackingParams string `json:"trackingParams"`
								} `json:"buttonRenderer"`
							} `json:"visitAdvertiserRenderer"`
							AdBadgeRenderer struct {
								SimpleAdBadgeRenderer struct {
									Text struct {
										Text           string `json:"text"`
										IsTemplated    bool   `json:"isTemplated"`
										TrackingParams string `json:"trackingParams"`
									} `json:"text"`
									TrackingParams string `json:"trackingParams"`
								} `json:"simpleAdBadgeRenderer"`
							} `json:"adBadgeRenderer"`
							AdDurationRemaining struct {
								AdDurationRemainingRenderer struct {
									TemplatedCountdown struct {
										TemplatedAdText struct {
											Text           string `json:"text"`
											IsTemplated    bool   `json:"isTemplated"`
											TrackingParams string `json:"trackingParams"`
										} `json:"templatedAdText"`
									} `json:"templatedCountdown"`
									TrackingParams string `json:"trackingParams"`
								} `json:"adDurationRemainingRenderer"`
							} `json:"adDurationRemaining"`
							AdInfoRenderer struct {
								AdHoverTextButtonRenderer struct {
									Button struct {
										ButtonRenderer struct {
											Style           string `json:"style"`
											Size            string `json:"size"`
											IsDisabled      bool   `json:"isDisabled"`
											ServiceEndpoint struct {
												ClickTrackingParams  string `json:"clickTrackingParams"`
												AdInfoDialogEndpoint struct {
													Dialog struct {
														AdInfoDialogRenderer struct {
															DialogMessage struct {
																Runs []struct {
																	Text               string `json:"text"`
																	Bold               bool   `json:"bold,omitempty"`
																	NavigationEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		UrlEndpoint         struct {
																			Url    string `json:"url"`
																			Target string `json:"target"`
																		} `json:"urlEndpoint"`
																		LoggingUrls []struct {
																			BaseUrl string `json:"baseUrl"`
																		} `json:"loggingUrls,omitempty"`
																	} `json:"navigationEndpoint,omitempty"`
																} `json:"runs"`
															} `json:"dialogMessage"`
															MuteAdRenderer struct {
																ButtonRenderer struct {
																	Style string `json:"style"`
																	Text  struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																	Icon struct {
																		IconType string `json:"iconType"`
																	} `json:"icon"`
																	NavigationEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		LoggingUrls         []struct {
																			BaseUrl string `json:"baseUrl"`
																		} `json:"loggingUrls"`
																		AdFeedbackEndpoint struct {
																			Content struct {
																				AdFeedbackRenderer struct {
																					Title struct {
																						Runs []struct {
																							Text string `json:"text"`
																						} `json:"runs"`
																					} `json:"title"`
																					ConfirmLabel struct {
																						Runs []struct {
																							Text string `json:"text"`
																						} `json:"runs"`
																					} `json:"confirmLabel"`
																					CancelLabel struct {
																						Runs []struct {
																							Text string `json:"text"`
																						} `json:"runs"`
																					} `json:"cancelLabel"`
																					Reasons []struct {
																						Reason struct {
																							Runs []struct {
																								Text string `json:"text"`
																							} `json:"runs"`
																						} `json:"reason"`
																						Endpoint struct {
																							ClickTrackingParams string `json:"clickTrackingParams"`
																							LoggingUrls         []struct {
																								BaseUrl string `json:"baseUrl"`
																							} `json:"loggingUrls"`
																							MuteAdEndpoint struct {
																								Type string `json:"type"`
																							} `json:"muteAdEndpoint"`
																						} `json:"endpoint"`
																					} `json:"reasons"`
																					CompletionMessage struct {
																						Runs []struct {
																							Text string `json:"text"`
																						} `json:"runs"`
																					} `json:"completionMessage"`
																					TrackingParams string `json:"trackingParams"`
																					CancelEndpoint struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						MuteAdEndpoint      struct {
																							Type string `json:"type"`
																						} `json:"muteAdEndpoint"`
																					} `json:"cancelEndpoint"`
																					ImpressionEndpoint struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						MuteAdEndpoint      struct {
																							Type string `json:"type"`
																						} `json:"muteAdEndpoint"`
																					} `json:"impressionEndpoint"`
																					CancelRenderer struct {
																						ButtonRenderer struct {
																							Style           string `json:"style"`
																							Size            string `json:"size"`
																							IsDisabled      bool   `json:"isDisabled"`
																							ServiceEndpoint struct {
																								ClickTrackingParams string `json:"clickTrackingParams"`
																								MuteAdEndpoint      struct {
																									Type string `json:"type"`
																								} `json:"muteAdEndpoint"`
																							} `json:"serviceEndpoint"`
																							Icon struct {
																								IconType string `json:"iconType"`
																							} `json:"icon"`
																							TrackingParams string `json:"trackingParams"`
																						} `json:"buttonRenderer"`
																					} `json:"cancelRenderer"`
																					UndoRenderer struct {
																						ButtonRenderer struct {
																							Style string `json:"style"`
																							Size  string `json:"size"`
																							Text  struct {
																								Runs []struct {
																									Text string `json:"text"`
																								} `json:"runs"`
																							} `json:"text"`
																							ServiceEndpoint struct {
																								ClickTrackingParams string `json:"clickTrackingParams"`
																								MuteAdEndpoint      struct {
																									Type string `json:"type"`
																								} `json:"muteAdEndpoint"`
																							} `json:"serviceEndpoint"`
																							TrackingParams string `json:"trackingParams"`
																						} `json:"buttonRenderer"`
																					} `json:"undoRenderer"`
																					ReasonsTitle struct {
																						Runs []struct {
																							Text string `json:"text"`
																						} `json:"runs"`
																					} `json:"reasonsTitle"`
																				} `json:"adFeedbackRenderer"`
																			} `json:"content"`
																		} `json:"adFeedbackEndpoint"`
																	} `json:"navigationEndpoint"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"buttonRenderer"`
															} `json:"muteAdRenderer"`
															ConfirmLabel struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"confirmLabel"`
															ConfirmServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																LoggingUrls         []struct {
																	BaseUrl string `json:"baseUrl"`
																} `json:"loggingUrls"`
																PingingEndpoint struct {
																	Hack bool `json:"hack"`
																} `json:"pingingEndpoint"`
															} `json:"confirmServiceEndpoint"`
															TrackingParams       string `json:"trackingParams"`
															CloseOverlayRenderer struct {
																ButtonRenderer struct {
																	Style           string `json:"style"`
																	Size            string `json:"size"`
																	IsDisabled      bool   `json:"isDisabled"`
																	ServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		LoggingUrls         []struct {
																			BaseUrl string `json:"baseUrl"`
																		} `json:"loggingUrls"`
																		PingingEndpoint struct {
																			Hack bool `json:"hack"`
																		} `json:"pingingEndpoint"`
																	} `json:"serviceEndpoint"`
																	Icon struct {
																		IconType string `json:"iconType"`
																	} `json:"icon"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"buttonRenderer"`
															} `json:"closeOverlayRenderer"`
															Title struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"title"`
															AdReasons []struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"adReasons"`
															ImpressionEndpoints []struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																LoggingUrls         []struct {
																	BaseUrl string `json:"baseUrl"`
																} `json:"loggingUrls"`
																PingingEndpoint struct {
																	Hack bool `json:"hack"`
																} `json:"pingingEndpoint"`
															} `json:"impressionEndpoints"`
														} `json:"adInfoDialogRenderer"`
													} `json:"dialog"`
												} `json:"adInfoDialogEndpoint"`
											} `json:"serviceEndpoint"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
									HoverText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"hoverText"`
									TrackingParams string `json:"trackingParams"`
								} `json:"adHoverTextButtonRenderer"`
							} `json:"adInfoRenderer"`
							AdLayoutLoggingData struct {
								SerializedAdServingDataEntry string `json:"serializedAdServingDataEntry"`
							} `json:"adLayoutLoggingData"`
							ElementId string `json:"elementId"`
						} `json:"instreamAdPlayerOverlayRenderer"`
					} `json:"playerOverlay"`
					ElementId                   string `json:"elementId"`
					TrackingParams              string `json:"trackingParams"`
					LegacyInfoCardVastExtension string `json:"legacyInfoCardVastExtension"`
					SodarExtensionData          struct {
						Siub string `json:"siub"`
						Bgub string `json:"bgub"`
						Scs  string `json:"scs"`
						Bgp  string `json:"bgp"`
					} `json:"sodarExtensionData"`
					ExternalVideoId     string `json:"externalVideoId"`
					AdLayoutLoggingData struct {
						SerializedAdServingDataEntry string `json:"serializedAdServingDataEntry"`
					} `json:"adLayoutLoggingData"`
				} `json:"instreamVideoAdRenderer"`
			} `json:"renderer"`
		} `json:"adPlacementRenderer"`
	} `json:"adPlacements"`
}

type musicItemRenderer struct {
	TrackingParams string `json:"trackingParams"`
	Thumbnail      struct {
		MusicThumbnailRenderer struct {
			Thumbnail struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			ThumbnailCrop  string `json:"thumbnailCrop"`
			ThumbnailScale string `json:"thumbnailScale"`
			TrackingParams string `json:"trackingParams"`
		} `json:"musicThumbnailRenderer"`
	} `json:"thumbnail"`
	Overlay struct {
		MusicItemThumbnailOverlayRenderer struct {
			Background struct {
				VerticalGradient struct {
					GradientLayerColors []string `json:"gradientLayerColors"`
				} `json:"verticalGradient"`
			} `json:"background"`
			Content struct {
				MusicPlayButtonRenderer struct {
					PlayNavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						WatchEndpoint       struct {
							VideoId                            string `json:"videoId"`
							WatchEndpointMusicSupportedConfigs struct {
								WatchEndpointMusicConfig struct {
									MusicVideoType string `json:"musicVideoType"`
								} `json:"watchEndpointMusicConfig"`
							} `json:"watchEndpointMusicSupportedConfigs"`
						} `json:"watchEndpoint,omitempty"`
						WatchPlaylistEndpoint struct {
							PlaylistId string `json:"playlistId"`
							Params     string `json:"params,omitempty"`
						} `json:"watchPlaylistEndpoint,omitempty"`
					} `json:"playNavigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
					PlayIcon       struct {
						IconType string `json:"iconType"`
					} `json:"playIcon"`
					PauseIcon struct {
						IconType string `json:"iconType"`
					} `json:"pauseIcon"`
					IconColor             int64 `json:"iconColor"`
					BackgroundColor       int   `json:"backgroundColor"`
					ActiveBackgroundColor int   `json:"activeBackgroundColor"`
					LoadingIndicatorColor int64 `json:"loadingIndicatorColor"`
					PlayingIcon           struct {
						IconType string `json:"iconType"`
					} `json:"playingIcon"`
					IconLoadingColor      int    `json:"iconLoadingColor"`
					ActiveScaleFactor     int    `json:"activeScaleFactor"`
					ButtonSize            string `json:"buttonSize"`
					RippleTarget          string `json:"rippleTarget"`
					AccessibilityPlayData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibilityPlayData"`
					AccessibilityPauseData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibilityPauseData"`
				} `json:"musicPlayButtonRenderer"`
			} `json:"content"`
			ContentPosition string `json:"contentPosition"`
			DisplayStyle    string `json:"displayStyle"`
		} `json:"musicItemThumbnailOverlayRenderer"`
	} `json:"overlay,omitempty"`
	FlexColumns []struct {
		MusicResponsiveListItemFlexColumnRenderer struct {
			Text struct {
				Runs []struct {
					Text               string `json:"text"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						WatchEndpoint       struct {
							VideoId                            string `json:"videoId"`
							WatchEndpointMusicSupportedConfigs struct {
								WatchEndpointMusicConfig struct {
									MusicVideoType string `json:"musicVideoType"`
								} `json:"watchEndpointMusicConfig"`
							} `json:"watchEndpointMusicSupportedConfigs"`
						} `json:"watchEndpoint,omitempty"`
						BrowseEndpoint struct {
							BrowseId                              string `json:"browseId"`
							BrowseEndpointContextSupportedConfigs struct {
								BrowseEndpointContextMusicConfig struct {
									PageType string `json:"pageType"`
								} `json:"browseEndpointContextMusicConfig"`
							} `json:"browseEndpointContextSupportedConfigs"`
						} `json:"browseEndpoint,omitempty"`
					} `json:"navigationEndpoint,omitempty"`
				} `json:"runs"`
			} `json:"text"`
			DisplayPriority string `json:"displayPriority"`
		} `json:"musicResponsiveListItemFlexColumnRenderer"`
	} `json:"flexColumns"`
	Menu struct {
		MenuRenderer struct {
			Items []struct {
				MenuNavigationItemRenderer struct {
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						WatchEndpoint       struct {
							VideoId        string `json:"videoId"`
							PlaylistId     string `json:"playlistId"`
							Params         string `json:"params"`
							LoggingContext struct {
								VssLoggingContext struct {
									SerializedContextData string `json:"serializedContextData"`
								} `json:"vssLoggingContext"`
							} `json:"loggingContext"`
							WatchEndpointMusicSupportedConfigs struct {
								WatchEndpointMusicConfig struct {
									MusicVideoType string `json:"musicVideoType"`
								} `json:"watchEndpointMusicConfig"`
							} `json:"watchEndpointMusicSupportedConfigs"`
						} `json:"watchEndpoint,omitempty"`
						ModalEndpoint struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint,omitempty"`
						BrowseEndpoint struct {
							BrowseId                              string `json:"browseId"`
							BrowseEndpointContextSupportedConfigs struct {
								BrowseEndpointContextMusicConfig struct {
									PageType string `json:"pageType"`
								} `json:"browseEndpointContextMusicConfig"`
							} `json:"browseEndpointContextSupportedConfigs"`
						} `json:"browseEndpoint,omitempty"`
						ShareEntityEndpoint struct {
							SerializedShareEntity string `json:"serializedShareEntity"`
							SharePanelType        string `json:"sharePanelType"`
						} `json:"shareEntityEndpoint,omitempty"`
						WatchPlaylistEndpoint struct {
							PlaylistId string `json:"playlistId"`
							Params     string `json:"params"`
						} `json:"watchPlaylistEndpoint,omitempty"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"menuNavigationItemRenderer,omitempty"`
				MenuServiceItemRenderer struct {
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					ServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						QueueAddEndpoint    struct {
							QueueTarget struct {
								VideoId    string `json:"videoId,omitempty"`
								PlaylistId string `json:"playlistId,omitempty"`
							} `json:"queueTarget"`
							QueueInsertPosition string `json:"queueInsertPosition"`
							Commands            []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								AddToToastAction    struct {
									Item struct {
										NotificationTextRenderer struct {
											SuccessResponseText struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"successResponseText"`
											TrackingParams string `json:"trackingParams"`
										} `json:"notificationTextRenderer"`
									} `json:"item"`
								} `json:"addToToastAction"`
							} `json:"commands"`
						} `json:"queueAddEndpoint"`
					} `json:"serviceEndpoint"`
					TrackingParams string `json:"trackingParams"`
				} `json:"menuServiceItemRenderer,omitempty"`
				ToggleMenuServiceItemRenderer struct {
					DefaultText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"defaultText"`
					DefaultIcon struct {
						IconType string `json:"iconType"`
					} `json:"defaultIcon"`
					DefaultServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						ModalEndpoint       struct {
							Modal struct {
								ModalWithTitleAndButtonRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Content struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"content"`
									Button struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											IsDisabled bool   `json:"isDisabled"`
											Text       struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignInEndpoint      struct {
													Hack bool `json:"hack"`
												} `json:"signInEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"button"`
								} `json:"modalWithTitleAndButtonRenderer"`
							} `json:"modal"`
						} `json:"modalEndpoint"`
					} `json:"defaultServiceEndpoint"`
					ToggledText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"toggledText"`
					ToggledIcon struct {
						IconType string `json:"iconType"`
					} `json:"toggledIcon"`
					TrackingParams         string `json:"trackingParams"`
					ToggledServiceEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						LikeEndpoint        struct {
							Status string `json:"status"`
							Target struct {
								PlaylistId string `json:"playlistId"`
							} `json:"target"`
						} `json:"likeEndpoint"`
					} `json:"toggledServiceEndpoint,omitempty"`
				} `json:"toggleMenuServiceItemRenderer,omitempty"`
			} `json:"items"`
			TrackingParams string `json:"trackingParams"`
			Accessibility  struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibility"`
		} `json:"menuRenderer"`
	} `json:"menu"`
	PlaylistItemData struct {
		VideoId string `json:"videoId"`
	} `json:"playlistItemData,omitempty"`
	FlexColumnDisplayStyle string `json:"flexColumnDisplayStyle"`
	ItemHeight             string `json:"itemHeight"`
	Badges                 []struct {
		MusicInlineBadgeRenderer struct {
			TrackingParams string `json:"trackingParams"`
			Icon           struct {
				IconType string `json:"iconType"`
			} `json:"icon"`
			AccessibilityData struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibilityData"`
		} `json:"musicInlineBadgeRenderer"`
	} `json:"badges,omitempty"`
	NavigationEndpoint struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
		BrowseEndpoint      struct {
			BrowseId                              string `json:"browseId"`
			BrowseEndpointContextSupportedConfigs struct {
				BrowseEndpointContextMusicConfig struct {
					PageType string `json:"pageType"`
				} `json:"browseEndpointContextMusicConfig"`
			} `json:"browseEndpointContextSupportedConfigs"`
		} `json:"browseEndpoint"`
	} `json:"navigationEndpoint,omitempty"`
}

type musicSearchResponse struct {
	ResponseContext struct {
		VisitorData           string `json:"visitorData"`
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MaxAgeSeconds int `json:"maxAgeSeconds"`
	} `json:"responseContext"`
	Contents struct {
		TabbedSearchResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Title    string `json:"title"`
					Selected bool   `json:"selected"`
					Content  struct {
						SectionListRenderer struct {
							Contents []struct {
								MusicShelfRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`
									Contents []struct {
										MusicResponsiveListItemRenderer musicItemRenderer `json:"musicResponsiveListItemRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
									ShelfDivider   struct {
										MusicShelfDividerRenderer struct {
											Hidden bool `json:"hidden"`
										} `json:"musicShelfDividerRenderer"`
									} `json:"shelfDivider"`
									BottomText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"bottomText,omitempty"`
									BottomEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										SearchEndpoint      struct {
											Query  string `json:"query"`
											Params string `json:"params"`
										} `json:"searchEndpoint"`
									} `json:"bottomEndpoint,omitempty"`
								} `json:"musicShelfRenderer"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
							Header         struct {
								ChipCloudRenderer struct {
									Chips []struct {
										ChipCloudChipRenderer struct {
											Style struct {
												StyleType string `json:"styleType"`
											} `json:"style"`
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SearchEndpoint      struct {
													Query  string `json:"query"`
													Params string `json:"params"`
												} `json:"searchEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibilityData"`
											IsSelected bool `json:"isSelected"`
										} `json:"chipCloudChipRenderer"`
									} `json:"chips"`
									CollapsedRowCount    int    `json:"collapsedRowCount"`
									TrackingParams       string `json:"trackingParams"`
									HorizontalScrollable bool   `json:"horizontalScrollable"`
								} `json:"chipCloudRenderer"`
							} `json:"header"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
					TabIdentifier  string `json:"tabIdentifier"`
					TrackingParams string `json:"trackingParams"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"tabbedSearchResultsRenderer"`
	} `json:"contents"`
	TrackingParams string `json:"trackingParams"`
}

type musicPlayerParseResponse struct {
	ResponseContext struct {
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
	} `json:"responseContext"`
	PlayabilityStatus struct {
		Status               string `json:"status"`
		PlayableInEmbed      bool   `json:"playableInEmbed"`
		AudioOnlyPlayability struct {
			AudioOnlyPlayabilityRenderer struct {
				AudioOnlyPlayability  bool   `json:"audioOnlyPlayability"`
				TrackingParams        string `json:"trackingParams"`
				AudioOnlyAvailability string `json:"audioOnlyAvailability"`
				PromoRenderer         struct {
					MealbarPromoRenderer struct {
						MessageTexts []struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"messageTexts"`
						ActionButton struct {
							ButtonRenderer struct {
								Size string `json:"size"`
								Text struct {
									Runs []struct {
										Text string `json:"text"`
									} `json:"runs"`
								} `json:"text"`
								ServiceEndpoint struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									FeedbackEndpoint    struct {
										FeedbackToken string `json:"feedbackToken"`
										UiActions     struct {
											HideEnclosingContainer bool `json:"hideEnclosingContainer"`
										} `json:"uiActions"`
									} `json:"feedbackEndpoint"`
								} `json:"serviceEndpoint"`
								NavigationEndpoint struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									BrowseEndpoint      struct {
										BrowseId string `json:"browseId"`
										Params   string `json:"params"`
									} `json:"browseEndpoint"`
								} `json:"navigationEndpoint"`
								TrackingParams string `json:"trackingParams"`
							} `json:"buttonRenderer"`
						} `json:"actionButton"`
						DismissButton struct {
							ButtonRenderer struct {
								Size string `json:"size"`
								Text struct {
									Runs []struct {
										Text string `json:"text"`
									} `json:"runs"`
								} `json:"text"`
								ServiceEndpoint struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									FeedbackEndpoint    struct {
										FeedbackToken string `json:"feedbackToken"`
										UiActions     struct {
											HideEnclosingContainer bool `json:"hideEnclosingContainer"`
										} `json:"uiActions"`
									} `json:"feedbackEndpoint"`
								} `json:"serviceEndpoint"`
								TrackingParams string `json:"trackingParams"`
							} `json:"buttonRenderer"`
						} `json:"dismissButton"`
						TriggerCondition    string `json:"triggerCondition"`
						Style               string `json:"style"`
						TrackingParams      string `json:"trackingParams"`
						ImpressionEndpoints []struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							FeedbackEndpoint    struct {
								FeedbackToken string `json:"feedbackToken"`
								UiActions     struct {
									HideEnclosingContainer bool `json:"hideEnclosingContainer"`
								} `json:"uiActions"`
							} `json:"feedbackEndpoint"`
						} `json:"impressionEndpoints"`
						IsVisible    bool `json:"isVisible"`
						MessageTitle struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"messageTitle"`
					} `json:"mealbarPromoRenderer"`
				} `json:"promoRenderer"`
			} `json:"audioOnlyPlayabilityRenderer"`
		} `json:"audioOnlyPlayability"`
		Miniplayer struct {
			MiniplayerRenderer struct {
				PlaybackMode string `json:"playbackMode"`
			} `json:"miniplayerRenderer"`
		} `json:"miniplayer"`
		ContextParams string `json:"contextParams"`
	} `json:"playabilityStatus"`
	StreamingData struct {
		ExpiresInSeconds string `json:"expiresInSeconds"`
		Formats          []struct {
			Itag             int    `json:"itag"`
			MimeType         string `json:"mimeType"`
			Bitrate          int    `json:"bitrate"`
			Width            int    `json:"width"`
			Height           int    `json:"height"`
			LastModified     string `json:"lastModified"`
			ContentLength    string `json:"contentLength"`
			Quality          string `json:"quality"`
			Fps              int    `json:"fps"`
			QualityLabel     string `json:"qualityLabel"`
			ProjectionType   string `json:"projectionType"`
			AverageBitrate   int    `json:"averageBitrate"`
			AudioQuality     string `json:"audioQuality"`
			ApproxDurationMs string `json:"approxDurationMs"`
			AudioSampleRate  string `json:"audioSampleRate"`
			AudioChannels    int    `json:"audioChannels"`
			SignatureCipher  string `json:"signatureCipher"`
		} `json:"formats"`
		AdaptiveFormats []struct {
			Itag      int    `json:"itag"`
			MimeType  string `json:"mimeType"`
			Bitrate   int    `json:"bitrate"`
			Width     int    `json:"width,omitempty"`
			Height    int    `json:"height,omitempty"`
			InitRange struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"initRange"`
			IndexRange struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"indexRange"`
			LastModified     string `json:"lastModified"`
			ContentLength    string `json:"contentLength"`
			Quality          string `json:"quality"`
			Fps              int    `json:"fps,omitempty"`
			QualityLabel     string `json:"qualityLabel,omitempty"`
			ProjectionType   string `json:"projectionType"`
			AverageBitrate   int    `json:"averageBitrate"`
			ApproxDurationMs string `json:"approxDurationMs"`
			SignatureCipher  string `json:"signatureCipher"`
			ColorInfo        struct {
				Primaries               string `json:"primaries"`
				TransferCharacteristics string `json:"transferCharacteristics"`
				MatrixCoefficients      string `json:"matrixCoefficients"`
			} `json:"colorInfo,omitempty"`
			HighReplication bool    `json:"highReplication,omitempty"`
			AudioQuality    string  `json:"audioQuality,omitempty"`
			AudioSampleRate string  `json:"audioSampleRate,omitempty"`
			AudioChannels   int     `json:"audioChannels,omitempty"`
			LoudnessDb      float64 `json:"loudnessDb,omitempty"`
		} `json:"adaptiveFormats"`
	} `json:"streamingData"`
	PlayerAds []struct {
		PlayerLegacyDesktopWatchAdsRenderer struct {
			PlayerAdParams struct {
				ShowContentThumbnail bool   `json:"showContentThumbnail"`
				EnabledEngageTypes   string `json:"enabledEngageTypes"`
			} `json:"playerAdParams"`
			GutParams struct {
				Tag string `json:"tag"`
			} `json:"gutParams"`
			ShowCompanion bool `json:"showCompanion"`
			ShowInstream  bool `json:"showInstream"`
			UseGut        bool `json:"useGut"`
		} `json:"playerLegacyDesktopWatchAdsRenderer"`
	} `json:"playerAds"`
	PlaybackTracking struct {
		VideostatsPlaybackUrl struct {
			BaseUrl string `json:"baseUrl"`
			Headers []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"videostatsPlaybackUrl"`
		VideostatsDelayplayUrl struct {
			BaseUrl string `json:"baseUrl"`
			Headers []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"videostatsDelayplayUrl"`
		VideostatsWatchtimeUrl struct {
			BaseUrl string `json:"baseUrl"`
			Headers []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"videostatsWatchtimeUrl"`
		PtrackingUrl struct {
			BaseUrl string `json:"baseUrl"`
			Headers []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"ptrackingUrl"`
		QoeUrl struct {
			BaseUrl string `json:"baseUrl"`
			Headers []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"qoeUrl"`
		AtrUrl struct {
			BaseUrl                 string `json:"baseUrl"`
			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
			Headers                 []struct {
				HeaderType string `json:"headerType"`
			} `json:"headers"`
		} `json:"atrUrl"`
		VideostatsScheduledFlushWalltimeSeconds []int `json:"videostatsScheduledFlushWalltimeSeconds"`
		VideostatsDefaultFlushIntervalSeconds   int   `json:"videostatsDefaultFlushIntervalSeconds"`
	} `json:"playbackTracking"`
	VideoDetails struct {
		VideoId        string `json:"videoId"`
		Title          string `json:"title"`
		LengthSeconds  string `json:"lengthSeconds"`
		ChannelId      string `json:"channelId"`
		IsOwnerViewing bool   `json:"isOwnerViewing"`
		IsCrawlable    bool   `json:"isCrawlable"`
		Thumbnail      struct {
			Thumbnails []struct {
				Url    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"thumbnails"`
		} `json:"thumbnail"`
		AllowRatings      bool   `json:"allowRatings"`
		ViewCount         string `json:"viewCount"`
		Author            string `json:"author"`
		IsPrivate         bool   `json:"isPrivate"`
		IsUnpluggedCorpus bool   `json:"isUnpluggedCorpus"`
		MusicVideoType    string `json:"musicVideoType"`
		IsLiveContent     bool   `json:"isLiveContent"`
	} `json:"videoDetails"`
	PlayerConfig struct {
		AudioConfig struct {
			LoudnessDb              float64 `json:"loudnessDb"`
			PerceptualLoudnessDb    float64 `json:"perceptualLoudnessDb"`
			EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
		} `json:"audioConfig"`
		StreamSelectionConfig struct {
			MaxBitrate string `json:"maxBitrate"`
		} `json:"streamSelectionConfig"`
		MediaCommonConfig struct {
			DynamicReadaheadConfig struct {
				MaxReadAheadMediaTimeMs int `json:"maxReadAheadMediaTimeMs"`
				MinReadAheadMediaTimeMs int `json:"minReadAheadMediaTimeMs"`
				ReadAheadGrowthRateMs   int `json:"readAheadGrowthRateMs"`
			} `json:"dynamicReadaheadConfig"`
		} `json:"mediaCommonConfig"`
		WebPlayerConfig struct {
			WebPlayerActionsPorting struct {
				SubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					SubscribeEndpoint   struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"subscribeEndpoint"`
				} `json:"subscribeCommand"`
				UnsubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					UnsubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"unsubscribeEndpoint"`
				} `json:"unsubscribeCommand"`
				AddToWatchLaterCommand struct {
					ClickTrackingParams  string `json:"clickTrackingParams"`
					PlaylistEditEndpoint struct {
						PlaylistId string `json:"playlistId"`
						Actions    []struct {
							AddedVideoId string `json:"addedVideoId"`
							Action       string `json:"action"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"addToWatchLaterCommand"`
				RemoveFromWatchLaterCommand struct {
					ClickTrackingParams  string `json:"clickTrackingParams"`
					PlaylistEditEndpoint struct {
						PlaylistId string `json:"playlistId"`
						Actions    []struct {
							Action         string `json:"action"`
							RemovedVideoId string `json:"removedVideoId"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"removeFromWatchLaterCommand"`
			} `json:"webPlayerActionsPorting"`
		} `json:"webPlayerConfig"`
	} `json:"playerConfig"`
	Storyboards struct {
		PlayerStoryboardSpecRenderer struct {
			Spec string `json:"spec"`
		} `json:"playerStoryboardSpecRenderer"`
	} `json:"storyboards"`
	Microformat struct {
		MicroformatDataRenderer struct {
			UrlCanonical string `json:"urlCanonical"`
			Title        string `json:"title"`
			Description  string `json:"description"`
			Thumbnail    struct {
				Thumbnails []struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			SiteName           string   `json:"siteName"`
			AppName            string   `json:"appName"`
			AndroidPackage     string   `json:"androidPackage"`
			IosAppStoreId      string   `json:"iosAppStoreId"`
			IosAppArguments    string   `json:"iosAppArguments"`
			OgType             string   `json:"ogType"`
			UrlApplinksIos     string   `json:"urlApplinksIos"`
			UrlApplinksAndroid string   `json:"urlApplinksAndroid"`
			UrlTwitterIos      string   `json:"urlTwitterIos"`
			UrlTwitterAndroid  string   `json:"urlTwitterAndroid"`
			TwitterCardType    string   `json:"twitterCardType"`
			TwitterSiteHandle  string   `json:"twitterSiteHandle"`
			SchemaDotOrgType   string   `json:"schemaDotOrgType"`
			Noindex            bool     `json:"noindex"`
			Unlisted           bool     `json:"unlisted"`
			Paid               bool     `json:"paid"`
			FamilySafe         bool     `json:"familySafe"`
			Tags               []string `json:"tags"`
			AvailableCountries []string `json:"availableCountries"`
			PageOwnerDetails   struct {
				Name              string `json:"name"`
				ExternalChannelId string `json:"externalChannelId"`
				YoutubeProfileUrl string `json:"youtubeProfileUrl"`
			} `json:"pageOwnerDetails"`
			VideoDetails struct {
				ExternalVideoId string `json:"externalVideoId"`
				DurationSeconds string `json:"durationSeconds"`
				DurationIso8601 string `json:"durationIso8601"`
			} `json:"videoDetails"`
			LinkAlternates []struct {
				HrefUrl       string `json:"hrefUrl"`
				Title         string `json:"title,omitempty"`
				AlternateType string `json:"alternateType,omitempty"`
			} `json:"linkAlternates"`
			ViewCount   string `json:"viewCount"`
			PublishDate string `json:"publishDate"`
			Category    string `json:"category"`
			UploadDate  string `json:"uploadDate"`
		} `json:"microformatDataRenderer"`
	} `json:"microformat"`
	TrackingParams string `json:"trackingParams"`
	Attestation    struct {
		PlayerAttestationRenderer struct {
			Challenge    string `json:"challenge"`
			BotguardData struct {
				Program            string `json:"program"`
				InterpreterSafeUrl struct {
					PrivateDoNotAccessOrElseTrustedResourceUrlWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
				} `json:"interpreterSafeUrl"`
				ServerEnvironment int `json:"serverEnvironment"`
			} `json:"botguardData"`
		} `json:"playerAttestationRenderer"`
	} `json:"attestation"`
	AdPlacements []struct {
		AdPlacementRenderer struct {
			Config struct {
				AdPlacementConfig struct {
					Kind         string `json:"kind"`
					AdTimeOffset struct {
						OffsetStartMilliseconds string `json:"offsetStartMilliseconds"`
						OffsetEndMilliseconds   string `json:"offsetEndMilliseconds"`
					} `json:"adTimeOffset"`
					HideCueRangeMarker bool `json:"hideCueRangeMarker"`
				} `json:"adPlacementConfig"`
			} `json:"config"`
			Renderer struct {
				ClientForecastingAdRenderer struct {
					ImpressionUrls []struct {
						BaseUrl string `json:"baseUrl"`
						Headers []struct {
							HeaderType string `json:"headerType"`
						} `json:"headers,omitempty"`
					} `json:"impressionUrls"`
				} `json:"clientForecastingAdRenderer,omitempty"`
				AdBreakServiceRenderer struct {
					PrefetchMilliseconds string `json:"prefetchMilliseconds"`
					GetAdBreakUrl        string `json:"getAdBreakUrl"`
				} `json:"adBreakServiceRenderer,omitempty"`
			} `json:"renderer"`
			AdSlotLoggingData struct {
				SerializedSlotAdServingDataEntry string `json:"serializedSlotAdServingDataEntry"`
			} `json:"adSlotLoggingData,omitempty"`
		} `json:"adPlacementRenderer"`
	} `json:"adPlacements"`
}
