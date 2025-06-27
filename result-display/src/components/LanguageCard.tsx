import "../styles/LanguageCard.css";
import { Accordion, Span, VStack, HStack, Text } from "@chakra-ui/react";
import LanguageGraph from "./LanguageGraph";

export interface LanguageData {
  lang_name: string;
  image: string;
  frameworks: string[];
  requests: number;
  duration_in_microseconds: number;
  requests_per_sec: number;
  bytes_transfer_per_sec: number;
  connect_errors: number;
  read_errors: number;
  write_errors: number;
  http_errors: number;
  timeouts: number;
  latency_distribution: LatencyObjects[];
}

export interface LatencyObjects {
  percentile: number;
  latency_in_microseconds: number;
}

export default function LanguageCard({
  lang_name,
  frameworks,
  image,
  requests,
  duration_in_microseconds,
  requests_per_sec,
  bytes_transfer_per_sec,
  connect_errors,
  read_errors,
  write_errors,
  http_errors,
  timeouts,
  latency_distribution,
}: LanguageData) {
  const totalErrors =
    connect_errors + read_errors + write_errors + http_errors + timeouts;

  const errorsAccordian = {
    value: "errors",
    title: `Errors (${totalErrors})`,
  };

  return (
    <>
      <div className="card-container">
        <div className="card-header">
          <img
            className="lang-logo"
            src={image}
            alt="programming language logo"
          />
          <p className="lang-name">{lang_name} </p>
          <p className="framework-list">[{...frameworks}]</p>
        </div>
        <div className="card-content">
          <div className="card-column-left">
            <p className="info-text">
              <strong>Total requests:</strong> {requests}
            </p>
            <p className="info-text">
              <strong>Duration (μs):</strong> {duration_in_microseconds}
            </p>
            <p className="info-text">
              <strong>Requests (/s):</strong> {requests_per_sec}
            </p>
            <p className="info-text">
              <strong>Bytes Moved (/s):</strong>
              {bytes_transfer_per_sec}
            </p>
            <Accordion.Root collapsible>
              <Accordion.Item
                value={errorsAccordian.value}
                style={{ border: "none" }}
              >
                <Accordion.ItemTrigger>
                  <Span flex="1" className="info-text">
                    {errorsAccordian.title}
                  </Span>
                  <Accordion.ItemIndicator />
                </Accordion.ItemTrigger>
                <Accordion.ItemContent>
                  <Accordion.ItemBody>
                    <VStack align="start">
                      <HStack justify="space-between" width="100%">
                        <Text className="info-text">Connection Errors:</Text>
                        <Text fontWeight="bold" className="info-text">
                          {connect_errors}
                        </Text>
                      </HStack>
                      <HStack justify="space-between" width="100%">
                        <Text className="info-text">Read Errors:</Text>
                        <Text fontWeight="bold" className="info-text">
                          {read_errors}
                        </Text>
                      </HStack>
                      <HStack justify="space-between" width="100%">
                        <Text className="info-text">Write Errors:</Text>
                        <Text fontWeight="bold" className="info-text">
                          {write_errors}
                        </Text>
                      </HStack>
                      <HStack justify="space-between" width="100%">
                        <Text className="info-text">HTTP Errors:</Text>
                        <Text fontWeight="bold" className="info-text">
                          {http_errors}
                        </Text>
                      </HStack>
                      <HStack justify="space-between" width="100%">
                        <Text className="info-text">Timeouts:</Text>
                        <Text fontWeight="bold" className="info-text">
                          {timeouts}
                        </Text>
                      </HStack>
                    </VStack>
                  </Accordion.ItemBody>
                </Accordion.ItemContent>
              </Accordion.Item>
            </Accordion.Root>
          </div>

          <LanguageGraph data={latency_distribution} lang_name={lang_name} />
        </div>
      </div>
    </>
  );
}
