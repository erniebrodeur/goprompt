Feature: Aggregator fallback behavior
  The aggregator merges multiple segments in parallel, handling failures or timeouts gracefully.

  @all_succeed
  Scenario: All segments succeed quickly
    Given I have segments ["DirSegment", "GitSegment"] that succeed quickly
    When I build the prompt with a 100ms timeout
    Then the result should include "DirSegment"
    And the result should include "GitSegment"

  @one_fail
  Scenario: One segment fails immediately
    Given I have segments ["DirSegment", "FailSegment"] where "FailSegment" fails
    When I build the prompt with a 100ms timeout
    Then the result should include "DirSegment"
    And the result should include "[ERR]"

  @one_timeout
  Scenario: One segment times out
    Given I have segments ["QuickSegment", "SlowSegment"] where "SlowSegment" is very slow
    When I build the prompt with a 50ms timeout
    Then the result should include "QuickSegment"
    And the result should include "[ERR]"

  @all_fail
  Scenario: All segments fail or time out
    Given I have segments ["BadSegment1", "BadSegment2"] that all fail
    When I build the prompt with a 100ms timeout
    Then the result should be "$dir() %"

