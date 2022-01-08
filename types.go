package youtube_live_chat

type SimpleLiveChatMessage struct {
	Msg string
}

type LiveChat struct {
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
	TrackingParams       string `json:"trackingParams"`
	ContinuationContents struct {
		LiveChatContinuation struct {
			Continuations []struct {
				TimedContinuationData struct {
					TimeoutMs           int    `json:"timeoutMs"`
					Continuation        string `json:"continuation"`
					ClickTrackingParams string `json:"clickTrackingParams"`
				} `json:"timedContinuationData"`
				InvalidationContinuationData struct {
					InvalidationID struct {
						ObjectSource             int    `json:"objectSource"`
						ObjectID                 string `json:"objectId"`
						Topic                    string `json:"topic"`
						SubscribeToGcmTopics     bool   `json:"subscribeToGcmTopics"`
						ProtoCreationTimestampMs string `json:"protoCreationTimestampMs"`
					} `json:"invalidationId"`
					TimeoutMs           int    `json:"timeoutMs"`
					Continuation        string `json:"continuation"`
					ClickTrackingParams string `json:"clickTrackingParams"`
				} `json:"invalidationContinuationData"`
			} `json:"continuations"`
			Actions []struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				AddChatItemAction   struct {
					Item struct {
						LiveChatTextMessageRenderer struct {
							Message struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"message"`
							AuthorName struct {
								SimpleText string `json:"simpleText"`
							} `json:"authorName"`
							AuthorPhoto struct {
								Thumbnails []struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
							} `json:"authorPhoto"`
							ContextMenuEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										IgnoreNavigation bool `json:"ignoreNavigation"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								LiveChatItemContextMenuEndpoint struct {
									Params string `json:"params"`
								} `json:"liveChatItemContextMenuEndpoint"`
							} `json:"contextMenuEndpoint"`
							ID                       string `json:"id"`
							TimestampUsec            string `json:"timestampUsec"`
							AuthorExternalChannelID  string `json:"authorExternalChannelId"`
							ContextMenuAccessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"contextMenuAccessibility"`
						} `json:"liveChatTextMessageRenderer"`
					} `json:"item"`
					ClientID string `json:"clientId"`
				} `json:"addChatItemAction,omitempty"`
				MarkChatItemAsDeletedAction struct {
					DeletedStateMessage struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"deletedStateMessage"`
					TargetItemID string `json:"targetItemId"`
				} `json:"markChatItemAsDeletedAction,omitempty"`
			} `json:"actions"`
			ActionPanel struct {
				LiveChatMessageInputRenderer struct {
					InputField struct {
						LiveChatTextInputFieldRenderer struct {
							Placeholder struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"placeholder"`
							MaxCharacterLimit   int `json:"maxCharacterLimit"`
							EmojiCharacterCount int `json:"emojiCharacterCount"`
						} `json:"liveChatTextInputFieldRenderer"`
					} `json:"inputField"`
					SendButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Icon       struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
						} `json:"buttonRenderer"`
					} `json:"sendButton"`
					Pickers []struct {
						EmojiPickerRenderer struct {
							ID         string `json:"id"`
							Categories []struct {
								EmojiPickerCategoryRenderer struct {
									CategoryID string `json:"categoryId"`
									Title      struct {
										SimpleText string `json:"simpleText"`
									} `json:"title"`
									EmojiIds       []string `json:"emojiIds"`
									TrackingParams string   `json:"trackingParams"`
								} `json:"emojiPickerCategoryRenderer"`
							} `json:"categories"`
							CategoryButtons []struct {
								EmojiPickerCategoryButtonRenderer struct {
									CategoryID string `json:"categoryId"`
									Icon       struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									Tooltip       string `json:"tooltip"`
									Accessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
								} `json:"emojiPickerCategoryButtonRenderer"`
							} `json:"categoryButtons"`
							SearchPlaceholderText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"searchPlaceholderText"`
							SearchNoResultsText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"searchNoResultsText"`
							PickSkinToneText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"pickSkinToneText"`
							TrackingParams           string `json:"trackingParams"`
							ClearSearchLabel         string `json:"clearSearchLabel"`
							SkinToneGenericLabel     string `json:"skinToneGenericLabel"`
							SkinToneLightLabel       string `json:"skinToneLightLabel"`
							SkinToneMediumLightLabel string `json:"skinToneMediumLightLabel"`
							SkinToneMediumLabel      string `json:"skinToneMediumLabel"`
							SkinToneMediumDarkLabel  string `json:"skinToneMediumDarkLabel"`
							SkinToneDarkLabel        string `json:"skinToneDarkLabel"`
						} `json:"emojiPickerRenderer"`
					} `json:"pickers"`
					PickerButtons []struct {
						LiveChatIconToggleButtonRenderer struct {
							TargetID string `json:"targetId"`
							Icon     struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							Tooltip       string `json:"tooltip"`
							Accessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
							ToggledIcon struct {
								IconType string `json:"iconType"`
							} `json:"toggledIcon"`
							TrackingParams string `json:"trackingParams"`
						} `json:"liveChatIconToggleButtonRenderer"`
					} `json:"pickerButtons"`
					InteractionMessage struct {
						MessageRenderer struct {
							TrackingParams string `json:"trackingParams"`
							Button         struct {
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
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										SignInEndpoint struct {
											NextEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												WatchEndpoint struct {
													VideoID string `json:"videoId"`
												} `json:"watchEndpoint"`
											} `json:"nextEndpoint"`
										} `json:"signInEndpoint"`
									} `json:"navigationEndpoint"`
									Accessibility struct {
										Label string `json:"label"`
									} `json:"accessibility"`
									TrackingParams    string `json:"trackingParams"`
									AccessibilityData struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibilityData"`
								} `json:"buttonRenderer"`
							} `json:"button"`
							Subtext struct {
								MessageSubtextRenderer struct {
									Text struct {
										SimpleText string `json:"simpleText"`
									} `json:"text"`
								} `json:"messageSubtextRenderer"`
							} `json:"subtext"`
						} `json:"messageRenderer"`
					} `json:"interactionMessage"`
					TargetID string `json:"targetId"`
				} `json:"liveChatMessageInputRenderer"`
			} `json:"actionPanel"`
			ItemList struct {
				LiveChatItemListRenderer struct {
					MaxItemsToDisplay       int `json:"maxItemsToDisplay"`
					MoreCommentsBelowButton struct {
						ButtonRenderer struct {
							Style string `json:"style"`
							Icon  struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
						} `json:"buttonRenderer"`
					} `json:"moreCommentsBelowButton"`
					EnablePauseChatKeyboardShortcuts bool `json:"enablePauseChatKeyboardShortcuts"`
				} `json:"liveChatItemListRenderer"`
			} `json:"itemList"`
			Header struct {
				LiveChatHeaderRenderer struct {
					OverflowMenu struct {
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
										ClickTrackingParams              string `json:"clickTrackingParams"`
										ShowLiveChatParticipantsEndpoint struct {
											Hack bool `json:"hack"`
										} `json:"showLiveChatParticipantsEndpoint"`
									} `json:"serviceEndpoint"`
									TrackingParams string `json:"trackingParams"`
								} `json:"menuServiceItemRenderer,omitempty"`
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
										CommandMetadata     struct {
											WebCommandMetadata struct {
												IgnoreNavigation bool `json:"ignoreNavigation"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										UserFeedbackEndpoint struct {
											Hack             bool   `json:"hack"`
											BucketIdentifier string `json:"bucketIdentifier"`
										} `json:"userFeedbackEndpoint"`
									} `json:"navigationEndpoint"`
									TrackingParams string `json:"trackingParams"`
								} `json:"menuNavigationItemRenderer,omitempty"`
							} `json:"items"`
							TrackingParams string `json:"trackingParams"`
							Accessibility  struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
						} `json:"menuRenderer"`
					} `json:"overflowMenu"`
					CollapseButton struct {
						ButtonRenderer struct {
							Style         string `json:"style"`
							Size          string `json:"size"`
							IsDisabled    bool   `json:"isDisabled"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"collapseButton"`
					ViewSelector struct {
						SortFilterSubMenuRenderer struct {
							SubMenuItems []struct {
								Title        string `json:"title"`
								Selected     bool   `json:"selected"`
								Continuation struct {
									ReloadContinuationData struct {
										Continuation        string `json:"continuation"`
										ClickTrackingParams string `json:"clickTrackingParams"`
									} `json:"reloadContinuationData"`
								} `json:"continuation"`
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								Subtitle       string `json:"subtitle"`
								TrackingParams string `json:"trackingParams"`
							} `json:"subMenuItems"`
							Accessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
							TrackingParams string `json:"trackingParams"`
						} `json:"sortFilterSubMenuRenderer"`
					} `json:"viewSelector"`
				} `json:"liveChatHeaderRenderer"`
			} `json:"header"`
			Ticker struct {
				LiveChatTickerRenderer struct {
					Sentinel bool `json:"sentinel"`
				} `json:"liveChatTickerRenderer"`
			} `json:"ticker"`
			TrackingParams   string `json:"trackingParams"`
			ParticipantsList struct {
				LiveChatParticipantsListRenderer struct {
					Title struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"title"`
					BackButton struct {
						ButtonRenderer struct {
							Icon struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
						} `json:"buttonRenderer"`
					} `json:"backButton"`
					Participants []struct {
						LiveChatParticipantRenderer struct {
							AuthorName struct {
								SimpleText string `json:"simpleText"`
							} `json:"authorName"`
							AuthorPhoto struct {
								Thumbnails []struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
							} `json:"authorPhoto"`
							AuthorBadges []struct {
								LiveChatAuthorBadgeRenderer struct {
									Icon struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									Tooltip       string `json:"tooltip"`
									Accessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
								} `json:"liveChatAuthorBadgeRenderer"`
							} `json:"authorBadges"`
						} `json:"liveChatParticipantRenderer"`
					} `json:"participants"`
				} `json:"liveChatParticipantsListRenderer"`
			} `json:"participantsList"`
			PopoutMessage struct {
				MessageRenderer struct {
					Text struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
					Button         struct {
						ButtonRenderer struct {
							Style string `json:"style"`
							Text  struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"text"`
							ServiceEndpoint struct {
								ClickTrackingParams    string `json:"clickTrackingParams"`
								PopoutLiveChatEndpoint struct {
									URL string `json:"url"`
								} `json:"popoutLiveChatEndpoint"`
							} `json:"serviceEndpoint"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"button"`
				} `json:"messageRenderer"`
			} `json:"popoutMessage"`
			ClientMessages struct {
				ReconnectMessage struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
				} `json:"reconnectMessage"`
				UnableToReconnectMessage struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
				} `json:"unableToReconnectMessage"`
				FatalError struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
				} `json:"fatalError"`
				ReconnectedMessage struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
				} `json:"reconnectedMessage"`
				GenericError struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
				} `json:"genericError"`
			} `json:"clientMessages"`
		} `json:"liveChatContinuation"`
	} `json:"continuationContents"`
}
